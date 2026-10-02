package codegen

import (
	"fmt"
	"strconv"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/v3/internal/codegen/types"
	"github.com/ethanmoffat/eolib-go/v3/internal/xml"
)

// caseValue is the value of a switch case: an enum value, or an integer.
type caseValue struct {
	Name        string
	PackageName string
	Value       int
}

func (v caseValue) code() jen.Code {
	switch {
	case len(v.Name) == 0:
		return jen.Lit(v.Value)
	case len(v.PackageName) > 0:
		return jen.Qual(types.PackagePath(v.PackageName), v.Name)
	default:
		return jen.Id(v.Name)
	}
}

func (v caseValue) docLink() string {
	switch {
	case v == codeParameter:
		return v.Name
	case len(v.Name) == 0:
		return strconv.Itoa(v.Value)
	case len(v.PackageName) > 0:
		return fmt.Sprintf("[%s.%s]", v.PackageName, v.Name)
	default:
		return fmt.Sprintf("[%s]", v.Name)
	}
}

// enumPackageName gets the package that qualifies references to an enum, or an empty string for the current package.
func enumPackageName(e *xml.ProtocolEnum, currentPackage string) string {
	if e.Package == currentPackage {
		return ""
	}
	return e.Package
}

func enumCaseValue(e *xml.ProtocolEnum, v xml.ProtocolValue, currentPackage string) caseValue {
	return caseValue{
		Name:        fmt.Sprintf("%s_%s", types.SanitizeTypeName(e.Name), v.Name),
		PackageName: enumPackageName(e, currentPackage),
		Value:       int(v.Value),
	}
}

// getCaseValue gets the value of a switch case on a field of the specified type. Integer values that match an enum
// value are referred to by name.
func getCaseValue(fieldType string, value string, currentPackage string, fullSpec xml.Protocol) (caseValue, error) {
	e, isEnum := fullSpec.IsEnum(fieldType)

	if ordinal, err := strconv.Atoi(value); err == nil {
		if isEnum {
			for _, v := range e.Values {
				if int(v.Value) == ordinal {
					return enumCaseValue(e, v, currentPackage), nil
				}
			}
		}
		return caseValue{Value: ordinal}, nil
	}

	if !isEnum {
		return caseValue{}, fmt.Errorf("type %s in switch is not an enum", fieldType)
	}
	for _, v := range e.Values {
		if v.Name == value {
			return enumCaseValue(e, v, currentPackage), nil
		}
	}
	return caseValue{}, fmt.Errorf("%s is not a value of enum %s", value, e.Name)
}

func switchInterfaceName(si *types.StructInfo, switchField string) string {
	return si.SwitchStructQualifier + snakeCaseToPascalCase(switchField) + "Data"
}

func switchCaseStructName(si *types.StructInfo, switchField string, c xml.ProtocolCase) string {
	if c.Default {
		return switchInterfaceName(si, switchField) + "Default"
	}
	return switchInterfaceName(si, switchField) + snakeCaseToPascalCase(c.Value)
}
