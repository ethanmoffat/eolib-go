package codegen

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/v3/internal/codegen/types"
	"github.com/ethanmoffat/eolib-go/v3/internal/xml"
)

type switchFactory struct {
	Name   string
	Owner  string
	Levels []factoryLevel

	DataType  string
	EmptyData bool

	IsDefault  bool
	CodeType   jen.Code
	CaseValues []caseValue
}

type factoryLevel struct {
	CodeField string
	DataField string
	CaseType  string
	Code      caseValue
}

// codeParameter is the code parameter of a default factory, which is the value of the switch field.
var codeParameter = caseValue{Name: "code"}

type switchFactoryBuilder struct {
	si        *types.StructInfo
	fullSpec  xml.Protocol
	owner     string
	prefix    string
	factories []switchFactory
}

func getSwitchFactories(si *types.StructInfo, fullSpec xml.Protocol) ([]switchFactory, error) {
	scope := xml.Flatten(si.Instructions)
	sw := findSwitch(scope)
	if sw == nil {
		return nil, nil
	}

	builder := &switchFactoryBuilder{si: si, fullSpec: fullSpec, owner: snakeCaseToPascalCase(si.Name)}
	builder.prefix = si.SwitchStructQualifier
	if len(builder.prefix) == 0 {
		builder.prefix = builder.owner
	}

	if err := builder.collect(sw, scope, nil, ""); err != nil {
		return nil, fmt.Errorf("type %s: %w", si.Name, err)
	}

	return builder.factories, nil
}

// findSwitch finds the switch in a scope. Validation ensures there is at most one.
func findSwitch(scope []*xml.ProtocolInstruction) *xml.ProtocolInstruction {
	for _, inst := range scope {
		if inst.XMLName.Local == "switch" {
			return inst
		}
	}
	return nil
}

func (b *switchFactoryBuilder) collect(sw *xml.ProtocolInstruction, scope []*xml.ProtocolInstruction, parents []factoryLevel, path string) error {
	fieldName := *sw.Field

	var fieldType string
	for _, inst := range scope {
		if inst.XMLName.Local == "field" && inst.Name != nil && *inst.Name == fieldName {
			fieldType = *inst.Type
			break
		}
	}

	// a switch on an integer field has no enum values, so only its numeric and default cases get factories
	var codeType jen.Code = jen.Int()
	enumType, isEnum := b.fullSpec.IsEnum(fieldType)
	if isEnum {
		codeType = caseValue{Name: enumType.Name, PackageName: enumPackageName(enumType, b.si.PackageName)}.code()
	} else {
		enumType = &xml.ProtocolEnum{}
	}

	// map each case to its value; numeric cases that aren't an enum value are kept separately
	caseByValue := map[int]*xml.ProtocolCase{}
	valueByCase := map[*xml.ProtocolCase]caseValue{}
	var caseValues []caseValue
	var numericCases []*xml.ProtocolCase
	var defaultCase *xml.ProtocolCase
	for i := range sw.Cases {
		c := &sw.Cases[i]
		if c.Default {
			defaultCase = c
			continue
		}

		value, err := getCaseValue(fieldType, c.Value, b.si.PackageName, b.fullSpec)
		if err != nil {
			return err
		}
		if len(value.Name) == 0 {
			numericCases = append(numericCases, c)
		}

		caseByValue[value.Value] = c
		valueByCase[c] = value
		caseValues = append(caseValues, value)
	}

	codeField := snakeCaseToPascalCase(fieldName)
	newLevel := func(c *xml.ProtocolCase, code caseValue) factoryLevel {
		level := factoryLevel{CodeField: codeField, DataField: codeField + "Data", Code: code}
		if c != nil && len(c.Instructions) > 0 {
			level.CaseType = switchCaseStructName(b.si, fieldName, *c)
		}
		return level
	}

	for _, v := range enumType.Values {
		c := caseByValue[int(v.Value)]
		if c == nil && defaultCase != nil {
			// values without a case are handled by the default case
			continue
		}

		level := newLevel(c, enumCaseValue(enumType, v, b.si.PackageName))
		if err := b.addCase(c, append(cloneLevels(parents), level), path+v.Name, fmt.Sprintf("New%sWith%s%s", b.prefix, path, v.Name)); err != nil {
			return err
		}
	}

	for _, c := range numericCases {
		if len(c.Instructions) == 0 {
			continue
		}

		level := newLevel(c, valueByCase[c])
		name := fmt.Sprintf("New%sWith%s", b.prefix, strings.TrimPrefix(level.CaseType, b.si.SwitchStructQualifier))
		if err := b.addCase(c, append(cloneLevels(parents), level), "", name); err != nil {
			return err
		}
	}

	if defaultCase != nil {
		level := newLevel(defaultCase, codeParameter)
		b.addDefault(defaultCase, append(cloneLevels(parents), level), path, codeType, caseValues)
	}

	return nil
}

func (b *switchFactoryBuilder) addCase(c *xml.ProtocolCase, levels []factoryLevel, path string, name string) error {
	if c != nil {
		scope := xml.Flatten(c.Instructions)
		if nested := findSwitch(scope); nested != nil {
			return b.collect(nested, scope, levels, path)
		}
	}

	b.factories = append(b.factories, newSwitchFactory(c, b.owner, levels, name))
	return nil
}

func (b *switchFactoryBuilder) addDefault(c *xml.ProtocolCase, levels []factoryLevel, path string, codeType jen.Code, caseValues []caseValue) {
	factory := newSwitchFactory(c, b.owner, levels, fmt.Sprintf("New%sWith%sDefault", b.prefix, path))
	factory.IsDefault = true
	factory.CodeType = codeType
	factory.CaseValues = caseValues
	b.factories = append(b.factories, factory)
}

func newSwitchFactory(c *xml.ProtocolCase, owner string, levels []factoryLevel, name string) switchFactory {
	factory := switchFactory{Name: name, Owner: owner, Levels: levels}
	if c == nil || len(c.Instructions) == 0 {
		return factory
	}

	if hasExportedMembers(c.Instructions) {
		factory.DataType = levels[len(levels)-1].CaseType
	} else {
		factory.EmptyData = true
	}
	return factory
}

func hasExportedMembers(instructions []xml.ProtocolInstruction) bool {
	for _, inst := range xml.Flatten(instructions) {
		if len(getInstructionMemberName(*inst)) > 0 {
			return true
		}
	}
	return false
}

func cloneLevels(levels []factoryLevel) []factoryLevel {
	return append([]factoryLevel(nil), levels...)
}

func writeSwitchFactories(f *jen.File, si *types.StructInfo, fullSpec xml.Protocol) error {
	factories, err := getSwitchFactories(si, fullSpec)
	if err != nil {
		return err
	}

	for _, factory := range factories {
		writeSwitchFactory(f, factory)
	}

	return nil
}

func writeSwitchFactory(f *jen.File, factory switchFactory) {
	last := factory.Levels[len(factory.Levels)-1]
	if factory.IsDefault {
		f.Commentf("%s creates a new [%s] with %s, for a %s value that is handled by the default case.", factory.Name, factory.Owner, describeFactoryLevels(factory, 0), last.CodeField)
		f.Comment("It returns an error if code is the value of another case.")
	} else {
		f.Commentf("%s creates a new [%s] with %s.", factory.Name, factory.Owner, describeFactoryLevels(factory, 0))
	}
	if len(factory.DataType) > 0 {
		f.Commentf("If data is nil, %s is left nil and serializing the result returns an error.", last.DataField)
	}

	var params []jen.Code
	if factory.IsDefault {
		params = append(params, jen.Id("code").Add(factory.CodeType))
	}
	if len(factory.DataType) > 0 {
		params = append(params, jen.Id("data").Op("*").Id(factory.DataType))
	}

	results := jen.Op("*").Id(factory.Owner)
	returnValue := func(data jen.Code) jen.Code {
		return jen.Return(factoryValue(factory, data))
	}
	if factory.IsDefault {
		results = jen.Params(jen.Op("*").Id(factory.Owner), jen.Error())
		returnValue = func(data jen.Code) jen.Code {
			return jen.Return(factoryValue(factory, data), jen.Nil())
		}
	}

	f.Func().Id(factory.Name).Params(params...).Add(results).BlockFunc(func(g *jen.Group) {
		if factory.IsDefault && len(factory.CaseValues) > 0 {
			var values []jen.Code
			for _, v := range factory.CaseValues {
				values = append(values, v.code())
			}
			g.Switch(jen.Id("code")).Block(
				jen.Case(values...).Block(
					jen.Return(jen.Nil(), jen.Qual("fmt", "Errorf").Call(
						jen.Lit(fmt.Sprintf("%s %%d has its own case and is not handled by the default case", last.CodeField)),
						jen.Id("code"),
					)),
				),
			)
		}

		if len(factory.DataType) > 0 {
			// a nil pointer assigned to the interface data field would make it non-nil (https://go.dev/doc/faq#nil_error)
			g.If(jen.Id("data").Op("==").Nil()).Block(returnValue(nil))
			g.Add(returnValue(jen.Id("data")))
		} else if factory.EmptyData {
			g.Add(returnValue(jen.Op("&").Id(last.CaseType).Values()))
		} else {
			g.Add(returnValue(nil))
		}
	}).Line()
}

func factoryValue(factory switchFactory, leafData jen.Code) jen.Code {
	value := leafData
	for i := len(factory.Levels) - 1; i >= 0; i-- {
		value = factoryLevelValue(factory, i, value)
	}
	return value
}

func factoryLevelValue(factory switchFactory, index int, inner jen.Code) *jen.Statement {
	holderType := factory.Owner
	if index > 0 {
		holderType = factory.Levels[index-1].CaseType
	}

	level := factory.Levels[index]
	values := []jen.Code{jen.Id(level.CodeField).Op(":").Add(level.Code.code())}
	if inner != nil {
		values = append(values, jen.Id(level.DataField).Op(":").Add(inner))
	}
	return jen.Op("&").Id(holderType).Values(values...)
}

func describeFactoryLevels(factory switchFactory, index int) string {
	level := factory.Levels[index]
	description := fmt.Sprintf("%s set to %s", level.CodeField, level.Code.docLink())

	switch {
	case index < len(factory.Levels)-1:
		description += fmt.Sprintf(" and %s set to a new [%s] with %s", level.DataField, level.CaseType, describeFactoryLevels(factory, index+1))
	case len(factory.DataType) > 0:
		description += fmt.Sprintf(" and %s set to data", level.DataField)
	case factory.EmptyData:
		description += fmt.Sprintf(" and %s set to a new [%s]", level.DataField, level.CaseType)
	}

	return description
}
