package codegen

import (
	"fmt"
	"math"
	"path"
	"strconv"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/internal/codegen/types"
	"github.com/ethanmoffat/eolib-go/internal/xml"
)

func GenerateStructs(outputDir string, structs []xml.ProtocolStruct, fullSpec xml.Protocol) error {
	const structFileName = "structs_generated.go"

	var typeNames []string
	for _, s := range structs {
		typeNames = append(typeNames, s.Name)
	}
	return generateStructsShared(outputDir, structFileName, typeNames, fullSpec)
}

func generateStructsShared(outputDir string, outputFileName string, typeNames []string, fullSpec xml.Protocol) error {
	packageName, err := getPackageName(outputDir)
	if err != nil {
		return err
	}

	f := jen.NewFile(packageName)
	types.AddImports(f)

	if len(typeNames) > 0 {
		for _, typeName := range typeNames {
			if err := writeStruct(f, typeName, fullSpec); err != nil {
				return err
			}
		}
	}

	outFileName := path.Join(outputDir, outputFileName)
	return writeToFileJen(f, outFileName)
}

func writeStruct(f *jen.File, typeName string, fullSpec xml.Protocol) (err error) {
	var si *types.StructInfo
	if si, err = types.GetStructInfo(typeName, fullSpec); err != nil {
		return err
	}

	err = writeStructShared(f, si, fullSpec)
	return
}

func writeStructShared(f *jen.File, si *types.StructInfo, fullSpec xml.Protocol) (err error) {
	structName := snakeCaseToPascalCase(si.Name)
	writeTypeCommentJen(f, structName, si.Comment)

	// write out fields
	var switches []*xml.ProtocolInstruction
	f.Type().Id(structName).StructFunc(func(g *jen.Group) {
		switches = writeStructFields(g, si, fullSpec)
	}).Line()

	for _, sw := range switches {
		if err = writeSwitchStructs(f, *sw, si, fullSpec); err != nil {
			return
		}
	}

	if len(si.Family) > 0 && len(si.Action) > 0 {
		// write out family/action methods
		f.Func().Params(jen.Id("s").Id(structName)).Id("Family").Params().Qual(types.PackagePath("net"), "PacketFamily").Block(
			jen.Return(jen.Qual(types.PackagePath("net"), fmt.Sprintf("PacketFamily_%s", si.Family))),
		).Line()
		f.Func().Params(jen.Id("s").Id(structName)).Id("Action").Params().Qual(types.PackagePath("net"), "PacketAction").Block(
			jen.Return(jen.Qual(types.PackagePath("net"), fmt.Sprintf("PacketAction_%s", si.Action))),
		).Line()
	}

	// write out serialize method
	f.Func().Params(jen.Id("s").Op("*").Id(structName)).Id("Serialize").Params(jen.Id("writer").Op("*").Qual(types.PackagePath("data"), "EoWriter")).Params(jen.Id("err").Id("error")).BlockFunc(func(g *jen.Group) {
		g.Id("oldSanitizeStrings").Op(":=").Id("writer").Dot("SanitizeStrings")
		// defer here uses 'Values' instead of 'Block' so the deferred function is single-line style
		g.Defer().Func().Params().Values(jen.Id("writer").Dot("SanitizeStrings").Op("=").Id("oldSanitizeStrings")).Call().Line()

		err = writeSerializeBody(g, si, fullSpec, nil)

		g.Return()
	}).Line()

	if err != nil {
		return
	}

	// write out deserialize method
	f.Func().Params(jen.Id("s").Op("*").Id(structName)).Id("Deserialize").Params(jen.Id("reader").Op("*").Qual(types.PackagePath("data"), "EoReader")).Params(jen.Id("err").Id("error")).BlockFunc(func(g *jen.Group) {
		g.Id("oldIsChunked").Op(":=").Id("reader").Dot("IsChunked").Call()
		// defer here uses 'Values' instead of 'Block' so the deferred function is single-line style
		g.Defer().Func().Params().Values(jen.Id("reader").Dot("SetIsChunked").Call(jen.Id("oldIsChunked"))).Call().Line()

		err = writeDeserializeBody(g, si, fullSpec, nil)

		g.Line().Return()
	}).Line()

	return
}

func writeStructFields(g *jen.Group, si *types.StructInfo, fullSpec xml.Protocol) (switches []*xml.ProtocolInstruction) {
	isEmpty := true

	for i, inst := range si.Instructions {
		var instName string

		if inst.Name != nil {
			instName = snakeCaseToPascalCase(*inst.Name)
		} else if inst.Field != nil {
			instName = snakeCaseToPascalCase(*inst.Field)
		}

		var fieldTypeInfo struct {
			typeName   string
			nextImport *types.ImportInfo
			isPointer  bool
		}
		if inst.Type != nil {
			fieldTypeInfo.typeName, fieldTypeInfo.nextImport = types.ProtocolSpecTypeToGoType(*inst.Type, si.PackageName, fullSpec)
			if inst.Optional != nil && *inst.Optional {
				switch inst.XMLName.Local {
				// these are the only supported values where the type of the rendered field needs to be modified to a pointer
				// arrays also support the "optional" attribute in the spec but default to nil since they are rendered as slices
				case "field":
					fallthrough
				case "length":
					fieldTypeInfo.isPointer = true
				}
			}
		}

		qualifiedTypeName := func(s *jen.Statement) {
			if fieldTypeInfo.isPointer {
				s.Op("*")
			}

			writeComment := func(ss *jen.Statement) {
				if inst.Comment != nil {
					writeInlineCommentJen(ss, *inst.Comment)
				}
			}

			if fieldTypeInfo.nextImport != nil && fieldTypeInfo.nextImport.Package != si.PackageName {
				s.Qual(fieldTypeInfo.nextImport.Path, fieldTypeInfo.typeName).Do(writeComment)
			} else {
				s.Id(fieldTypeInfo.typeName).Do(writeComment)
			}
		}

		switch inst.XMLName.Local {
		case "field":
			if len(instName) > 0 {
				g.Id(instName).Do(qualifiedTypeName)
			} else {
				g.Line()
			}
			isEmpty = false
		case "array":
			g.Id(instName).Index().Do(qualifiedTypeName)
			isEmpty = false
		case "length":
			g.Id(instName).Do(qualifiedTypeName)
			isEmpty = false
		case "switch":
			g.Id(fmt.Sprintf("%sData", instName)).Id(fmt.Sprintf("%s%sData", si.SwitchStructQualifier, instName))
			switches = append(switches, &si.Instructions[i])
			isEmpty = false
		case "chunked":
			nestedStructInfo, _ := si.Nested(&inst)
			switches = append(switches, writeStructFields(g, nestedStructInfo, fullSpec)...)
		case "dummy":
		case "break":
			continue // no data to write
		}
	}

	if isEmpty {
		g.Line()
	}

	return
}

func writeSwitchStructs(f *jen.File, switchInst xml.ProtocolInstruction, si *types.StructInfo, fullSpec xml.Protocol) (err error) {
	if switchInst.XMLName.Local != "switch" {
		return
	}

	switchInterfaceName := fmt.Sprintf("%sData", snakeCaseToPascalCase(*switchInst.Field))
	if len(si.SwitchStructQualifier) > 0 {
		switchInterfaceName = si.SwitchStructQualifier + switchInterfaceName
	}

	if switchInst.Comment != nil {
		writeTypeCommentJen(f, switchInterfaceName, *switchInst.Comment)
	}
	f.Type().Id(switchInterfaceName).Interface(jen.Qual(types.PackagePath("protocol"), "EoData")).Line()

	for _, c := range switchInst.Cases {
		if len(c.Instructions) == 0 {
			continue
		}

		var caseName string
		if c.Default {
			caseName = "Default"
		} else {
			caseName = snakeCaseToPascalCase(c.Value)
		}
		caseStructName := fmt.Sprintf("%s%s", switchInterfaceName, caseName)

		nestedStructInfo := &types.StructInfo{
			Name:                  caseStructName,
			Comment:               c.Comment,
			Instructions:          c.Instructions,
			PackageName:           si.PackageName,
			SwitchStructQualifier: si.SwitchStructQualifier,
		}
		err = writeStructShared(f, nestedStructInfo, fullSpec)
		if err != nil {
			return
		}
	}

	return
}

func writeSerializeBody(g *jen.Group, si *types.StructInfo, fullSpec xml.Protocol, outerInstructionList []xml.ProtocolInstruction) (err error) {
	for _, instruction := range si.Instructions {
		instructionType := instruction.XMLName.Local
		instructionName := getInstructionName(instruction)

		switch instructionType {
		case "chunked":
			g.Id("writer").Dot("SanitizeStrings").Op("=").True()

			var nestedInfo *types.StructInfo
			if nestedInfo, err = si.Nested(&instruction); err != nil {
				return
			}

			if err = writeSerializeBody(g, nestedInfo, fullSpec, si.Instructions); err != nil {
				return
			}

			g.Id("writer").Dot("SanitizeStrings").Op("=").False()
		case "break":
			g.Id("writer").Dot("AddByte").Call(jen.Lit(0xFF))
		case "switch":
			// get type of Value field
			switchFieldSanitizedType := ""
			switchFieldEnumType := ""
			for _, tmpInst := range append(outerInstructionList, si.Instructions...) {
				if tmpInst.XMLName.Local == "field" && snakeCaseToPascalCase(*tmpInst.Name) == instructionName {
					switchFieldEnumType = *tmpInst.Type
					switchFieldSanitizedType = types.SanitizeTypeName(switchFieldEnumType)
					break
				}
			}

			var switchBlock []jen.Code
			for _, c := range instruction.Cases {
				if len(c.Instructions) == 0 {
					continue
				}

				var switchDataType string
				if c.Default {
					switchDataType = fmt.Sprintf("%sDataDefault", instructionName)
					switchBlock = append(switchBlock, jen.Default())
				} else {
					switchDataType = fmt.Sprintf("%sData%s", instructionName, c.Value)
					if value, err := strconv.ParseInt(c.Value, 10, 32); err != nil {
						// case is for an enum value
						if enumTypeInfo, ok := fullSpec.IsEnum(switchFieldEnumType); !ok {
							return fmt.Errorf("type %s in switch is not an enum", switchFieldEnumType)
						} else {
							packageQualifier := ""
							if enumTypeInfo.Package != si.PackageName {
								packageQualifier = enumTypeInfo.Package
							}
							switchBlock = append(
								switchBlock,
								jen.CaseFunc(func(g *jen.Group) {
									if packageQualifier != "" {
										g.Qual(types.PackagePath(packageQualifier), fmt.Sprintf("%s_%s", switchFieldSanitizedType, c.Value))
									} else {
										g.Id(fmt.Sprintf("%s_%s", switchFieldSanitizedType, c.Value))
									}
								}),
							)
						}
					} else {
						// case is for an integer constant
						switchBlock = append(switchBlock, jen.Case(jen.Lit(int(value))))
					}
				}

				// Serialize call for the case structure
				caseSerialize := jen.If(
					jen.Id("err").Op("=").Id("s").Dot(fmt.Sprintf("%sData", instructionName)).Dot("Serialize").Call(jen.Id("writer")),
					jen.Id("err").Op("!=").Nil(),
				).Block(jen.Return())

				if len(switchDataType) > 0 {
					// The object to serialize needs a type assertion
					// Wrap it in a type assert switch that returns an error if it does not match
					switchBlock = append(
						switchBlock,
						jen.Switch(
							jen.Id("s").Dot(
								fmt.Sprintf("%sData", instructionName),
							).Assert(jen.Id("type")).Block(
								jen.Case(
									jen.Op("*").Id(fmt.Sprintf("%s%s", si.SwitchStructQualifier, switchDataType)),
								).Block(caseSerialize),
								jen.Default().Block(
									jen.Id("err").Op("=").Qual("fmt", "Errorf").Call(
										jen.Lit("invalid switch struct type for switch value %d"),
										jen.Id("s").Dot(instructionName),
									).Line().Return(),
								),
							),
						),
					)
				} else {
					// The object to serialize does not need a type assertion
					switchBlock = append(switchBlock, caseSerialize)
				}
			}

			g.Switch(jen.Id("s").Dot(instructionName)).Block(switchBlock...)
		default:
			typeName, typeSize := types.GetInstructionTypeName(instruction)

			if len(instructionName) == 0 && instruction.Content != nil {
				instructionName = *instruction.Content
			}
			g.Commentf("// %s : %s : %s", instructionName, instructionType, *instruction.Type)

			stringType := types.String

			var serializeCodes []jen.Code
			switch typeName {
			case "byte":
				fallthrough
			case "char":
				fallthrough
			case "short":
				fallthrough
			case "three":
				fallthrough
			case "int":
				fallthrough
			case "blob":
				serializeCodes = getSerializeForInstruction(instruction, types.NewEoType(typeName), false)
			case "bool":
				if len(typeSize) > 0 {
					typeName = string(unicode.ToUpper(rune(typeSize[0]))) + typeSize[1:]
				} else {
					typeName = "Char"
				}
				serializeCodes = []jen.Code{
					jen.If(jen.Id("s").Dot(instructionName)).Block(
						jen.Id("err").Op("=").Id("writer").Dot(fmt.Sprintf("Add%s", typeName)).Call(jen.Lit(1)),
					).Else().Block(
						jen.Id("err").Op("=").Id("writer").Dot(fmt.Sprintf("Add%s", typeName)).Call(jen.Lit(0)),
					).Line(),
					jen.If(jen.Id("err").Op("!=").Nil()).Block(jen.Return()).Line(),
				}
			case "encoded_string":
				stringType = types.EncodedString
				fallthrough
			case "string":
				if instruction.Length != nil && instructionType == "field" {
					if instruction.Padded != nil && *instruction.Padded {
						serializeCodes = getSerializeForInstruction(instruction, stringType+types.Padded, false)
					} else {
						serializeCodes = getSerializeForInstruction(instruction, stringType+types.Fixed, false)
					}
				} else {
					serializeCodes = getSerializeForInstruction(instruction, stringType, false)
				}
			default:
				if _, ok := fullSpec.IsStruct(typeName); ok {
					serializeCodes = []jen.Code{
						jen.If(
							jen.Id("err").Op("=").Id("s").Dot(instructionName).Do(func(s *jen.Statement) {
								if instructionType == "array" {
									s.Index(jen.Id("ndx"))
								}
							}).Dot("Serialize").Call(jen.Id("writer")),
							jen.Id("err").Op("!=").Nil(),
						).Block(jen.Return()),
					}
				} else if e, ok := fullSpec.IsEnum(typeName); ok {
					if t := types.NewEoType(e.Type); t&types.Primitive > 0 {
						serializeCodes = getSerializeForInstruction(instruction, t, true)
					}
				} else {
					err = fmt.Errorf("unable to find type '%s' when writing serialization function (member: %s, type: %s)", typeName, instructionName, instructionType)
					return
				}
			}

			if instructionType == "array" {
				var lenExpr *jen.Statement
				if instruction.Length != nil {
					lenExpr = getLengthExpression(*instruction.Length)
				} else {
					lenExpr = jen.Len(jen.Id("s").Dot(instructionName))
				}

				delimited := instruction.Delimited != nil && *instruction.Delimited
				trailingDelimiter := instruction.TrailingDelimiter == nil || *instruction.TrailingDelimiter

				if delimited {
					addByteCode := jen.Id("writer").Dot("AddByte").Call(jen.Lit(0xFF))
					if !trailingDelimiter {
						delimiterCode := jen.If(
							jen.Id("ndx").Op(">").Lit(0).Block(addByteCode).Line(),
						)
						serializeCodes = append([]jen.Code{delimiterCode}, serializeCodes...)
					} else {
						serializeCodes = append(serializeCodes, addByteCode)
					}
				}

				g.For(
					jen.Id("ndx").Op(":=").Lit(0),
					jen.Id("ndx").Op("<").Add(lenExpr),
					jen.Id("ndx").Op("++"),
				).Block(serializeCodes...).Line()
			} else {
				g.Add(serializeCodes...)
			}
		}
	}

	return
}

func writeDeserializeBody(g *jen.Group, si *types.StructInfo, fullSpec xml.Protocol, outerInstructionList []xml.ProtocolInstruction) (err error) {
	for _, instruction := range si.Instructions {
		instructionType := instruction.XMLName.Local
		instructionName := getInstructionName(instruction)

		switch instructionType {
		case "chunked":
			g.Id("reader").Dot("SetIsChunked").Call(jen.True())

			var nestedInfo *types.StructInfo
			if nestedInfo, err = si.Nested(&instruction); err != nil {
				return
			}

			if err = writeDeserializeBody(g, nestedInfo, fullSpec, si.Instructions); err != nil {
				return
			}

			g.Id("reader").Dot("SetIsChunked").Call(jen.False())
		case "break":
			if instruction.IsChunked {
				g.If(
					jen.Id("err").Op("=").Id("reader").Dot("NextChunk").Call(),
					jen.Id("err").Op("!=").Nil(),
				).Block(jen.Return())
			} else {
				g.If(
					jen.Id("breakByte").Op(":=").Id("reader").Dot("GetByte").Call(),
					jen.Id("breakByte").Op("!=").Lit(0xFF),
				).Block(
					jen.Return(jen.Qual("fmt", "Errorf").Call(jen.Lit("missing expected break byte"))),
				)
			}
		case "switch":
			// get type of Value field
			switchFieldSanitizedType := ""
			switchFieldEnumType := ""
			for _, tmpInst := range append(outerInstructionList, si.Instructions...) {
				if tmpInst.XMLName.Local == "field" && snakeCaseToPascalCase(*tmpInst.Name) == instructionName {
					switchFieldEnumType = *tmpInst.Type
					switchFieldSanitizedType = types.SanitizeTypeName(switchFieldEnumType)
					break
				}
			}

			var switchBlock []jen.Code
			for _, c := range instruction.Cases {
				if len(c.Instructions) == 0 {
					continue
				}

				var switchDataType string
				if c.Default {
					switchDataType = fmt.Sprintf("%sDataDefault", instructionName)
					switchBlock = append(switchBlock, jen.Default())
				} else {
					switchDataType = fmt.Sprintf("%sData%s", instructionName, c.Value)
					if value, err := strconv.ParseInt(c.Value, 10, 32); err != nil {
						// case is for an enum value
						if enumTypeInfo, ok := fullSpec.IsEnum(switchFieldEnumType); !ok {
							return fmt.Errorf("type %s in switch is not an enum", switchFieldEnumType)
						} else {
							packageQualifier := ""
							if enumTypeInfo.Package != si.PackageName {
								packageQualifier = enumTypeInfo.Package
							}
							switchBlock = append(switchBlock, jen.CaseFunc(func(g *jen.Group) {
								if packageQualifier != "" {
									g.Qual(types.PackagePath(packageQualifier), fmt.Sprintf("%s_%s", switchFieldSanitizedType, c.Value))
								} else {
									g.Id(fmt.Sprintf("%s_%s", switchFieldSanitizedType, c.Value))
								}
							}))
						}
					} else {
						// case is for an integer constant
						switchBlock = append(switchBlock, jen.Case(jen.Lit(int(value))))
					}
				}

				// Deserialize call for the case structure
				sDotData := jen.Id("s").Dot(fmt.Sprintf("%sData", instructionName))
				caseDeserialize := sDotData.Clone().Op("=").Op("&").Id(si.SwitchStructQualifier + switchDataType).Block().Line()
				caseDeserialize = caseDeserialize.If(
					jen.Id("err").Op("=").Add(sDotData).Dot("Deserialize").Call(jen.Id("reader")),
					jen.Id("err").Op("!=").Nil(),
				).Block(jen.Return())

				switchBlock = append(switchBlock, caseDeserialize)
			}

			g.Switch(jen.Id("s").Dot(instructionName)).Block(switchBlock...)
		default:
			typeName, typeSize := types.GetInstructionTypeName(instruction)

			if len(instructionName) == 0 && instruction.Content != nil {
				instructionName = *instruction.Content
			}
			g.Commentf("// %s : %s : %s", instructionName, instructionType, *instruction.Type)

			stringType := types.String

			var deserializeCodes []jen.Code
			switch typeName {
			case "byte":
				deserializeCodes = getDeserializeForInstruction(instruction, types.NewEoType(typeName), jen.Id("int"))
			case "char":
				fallthrough
			case "short":
				fallthrough
			case "three":
				fallthrough
			case "int":
				fallthrough
			case "blob":
				deserializeCodes = getDeserializeForInstruction(instruction, types.NewEoType(typeName), nil)
			case "bool":
				if len(typeSize) > 0 {
					typeName = string(unicode.ToUpper(rune(typeSize[0]))) + typeSize[1:]
				} else {
					typeName = "Char"
				}

				deserializeCodes = []jen.Code{
					jen.If(
						jen.Id("boolVal").Op(":=").Id("reader").Dot("Get"+typeName).Call(),
						jen.Id("boolVal").Op(">").Lit(0),
					).Block(
						jen.Id("s").Dot(instructionName).Op("=").True(),
					).Else().Block(
						jen.Id("s").Dot(instructionName).Op("=").False(),
					),
				}
			case "encoded_string":
				stringType = types.EncodedString
				fallthrough
			case "string":
				if instruction.Length != nil && instructionType == "field" {
					if instruction.Padded != nil && *instruction.Padded {
						deserializeCodes = getDeserializeForInstruction(instruction, stringType+types.Padded, nil)
					} else {
						deserializeCodes = getDeserializeForInstruction(instruction, stringType+types.Fixed, nil)
					}
				} else {
					deserializeCodes = getDeserializeForInstruction(instruction, stringType, nil)
				}
			default:
				if s, ok := fullSpec.IsStruct(typeName); ok {
					arrayCode := jen.Null()
					if instructionType == "array" {
						_, tp := types.ProtocolSpecTypeToGoType(s.Name, si.PackageName, fullSpec)
						arrayCode = jen.Id("s").Dot(instructionName).Op("=").Append(
							jen.Id("s").Dot(instructionName),
							jen.Do(func(s *jen.Statement) {
								if tp != nil {
									s.Qual(tp.Path, typeName)
								} else {
									s.Id(typeName)
								}
							}).Block(),
						)
					}

					deserializeCodes = []jen.Code{
						arrayCode,
						jen.If(
							jen.Id("err").Op("=").Id("s").Dot(instructionName).Do(func(s *jen.Statement) {
								if instructionType == "array" {
									s.Index(jen.Id("ndx"))
								}
							}).Dot("Deserialize").Call(jen.Id("reader")),
							jen.Id("err").Op("!=").Nil(),
						).Block(jen.Return()),
					}
				} else if e, ok := fullSpec.IsEnum(typeName); ok {
					if eoType := types.NewEoType(e.Type); eoType&types.Primitive > 0 {
						_, tp := types.ProtocolSpecTypeToGoType(e.Name, si.PackageName, fullSpec)
						deserializeCodes = getDeserializeForInstruction(
							instruction,
							eoType,
							jen.Do(func(s *jen.Statement) {
								if tp != nil {
									s.Qual(tp.Path, e.Name)
								} else {
									s.Id(e.Name)
								}
							}),
						)
					} else {
						err = fmt.Errorf("expected primitive base type for enum %s when writing deserialize function", e.Name)
					}
				} else {
					panic("Unable to find type '" + typeName + "' when writing serialization function")
				}
			}

			if instructionType == "array" {
				delimited := instruction.Delimited != nil && *instruction.Delimited

				var lenExpr *jen.Statement
				if instruction.Length != nil {
					lenExpr = jen.Id("ndx").Op("<").Add(getLengthExpression(*instruction.Length))
				} else if !delimited && instruction.IsChunked {
					if rawLen, err := types.CalculateTypeSize(typeName, fullSpec); err != nil || rawLen == 1 {
						lenExpr = jen.Id("reader").Dot("Remaining").Call().Op(">").Lit(0)
					} else {
						lenExpr = jen.Id("ndx").Op("<").Id("reader").Dot("Remaining").Call().Op("/").Lit(rawLen)
					}
				} else {
					lenExpr = jen.Id("reader").Dot("Remaining").Call().Op(">").Lit(0)
				}

				trailingDelimiter := instruction.TrailingDelimiter == nil || *instruction.TrailingDelimiter

				if delimited && instruction.IsChunked {
					delimiterExpr := jen.If(
						jen.Id("err").Op("=").Id("reader").Dot("NextChunk").Call(),
						jen.Id("err").Op("!=").Nil(),
					).Block(jen.Return())

					if !trailingDelimiter {
						if instruction.Length == nil {
							err = fmt.Errorf("delimited arrays with trailing-delimiter=false must have a length (array %s)", instructionName)
							return
						}

						delimiterExpr = jen.If(
							jen.Id("ndx").Op("+").Lit(1).Op("<").Add(getLengthExpression(*instruction.Length))).Block(delimiterExpr)
					}

					deserializeCodes = append(deserializeCodes, delimiterExpr)
				}

				g.For(
					jen.Id("ndx").Op(":=").Lit(0),
					lenExpr,
					jen.Id("ndx").Op("++"),
				).Block(deserializeCodes...).Line()
			} else {
				g.Add(deserializeCodes...)
			}
		}
	}

	return
}

func getInstructionName(inst xml.ProtocolInstruction) (instName string) {
	if inst.Name != nil {
		instName = snakeCaseToPascalCase(*inst.Name)
	} else if inst.Field != nil {
		instName = snakeCaseToPascalCase(*inst.Field)
	}
	return
}

func getSerializeForInstruction(instruction xml.ProtocolInstruction, methodType types.EoType, needsCastToInt bool) []jen.Code {
	instructionName := getInstructionName(instruction)

	// the method type is a string if it has the eotype_str or eotype_str_encoded flag
	isString := (methodType&types.String) > 0 || (methodType&types.EncodedString) > 0

	var instructionCode, nilCheckCode *jen.Statement
	if len(instructionName) == 0 && instruction.Content != nil {
		if isString {
			instructionCode = jen.Lit(*instruction.Content)
		} else {
			instructionCode = jen.Id(*instruction.Content)
		}
	} else {
		instructionCode = jen.Id("s").Dot(instructionName)
	}

	isArray := false
	optional := instruction.Optional != nil && *instruction.Optional
	if instruction.XMLName.Local == "array" {
		instructionCode = instructionCode.Index(jen.Id("ndx"))

		// optional arrays that are unset will be nil.
		// The length expression in the loop checks the length of the nil slice, which evaluates to 0.
		// This means that arrays do not need additional dereferencing when optional.
		optional = false
		isArray = true
	}

	if optional {
		nilCheckCode = instructionCode.Clone()
		instructionCode = jen.Op("*").Add(instructionCode)
	}

	if needsCastToInt {
		instructionCode = jen.Int().Call(instructionCode)
	}

	if instruction.Offset != nil {
		var op string
		if *instruction.Offset < 0 {
			op = "+"
		} else {
			op = "-"
		}
		instructionCode = instructionCode.Op(op).Lit(int(math.Abs(float64(*instruction.Offset))))
	}

	serializeCode := jen.If(
		jen.Id("err").Op("=").Id("writer").Dot("Add"+methodType.String()).Call(
			instructionCode,
			jen.Do(func(s *jen.Statement) {
				// strings may have a fixed length that needs to be serialized
				if !isArray && isString && instruction.Length != nil {
					s.Add(getLengthExpression(*instruction.Length))
				}
			}),
		),
		jen.Id("err").Op("!=").Nil(),
	).Block(jen.Return())

	return []jen.Code{
		jen.Do(func(s *jen.Statement) {
			if optional {
				s.If(nilCheckCode.Op("!=").Nil()).Block(serializeCode)
			} else {
				s.Add(serializeCode)
			}
		}),
	}
}

func getDeserializeForInstruction(instruction xml.ProtocolInstruction, methodType types.EoType, castType *jen.Statement) []jen.Code {
	instructionName := getInstructionName(instruction)

	// the method type is a string if it has the eotype_str or eotype_str_encoded flag
	isString := (methodType&types.String) > 0 || (methodType&types.EncodedString) > 0

	isArray := false
	optional := instruction.Optional != nil && *instruction.Optional

	lengthExpr := jen.Null()
	if instruction.XMLName.Local != "array" {
		if instruction.Length != nil {
			lengthExpr = getLengthExpression(*instruction.Length)
		} else if methodType == types.Bytes {
			lengthExpr = jen.Id("reader").Dot("Remaining").Call()
		}
	} else {
		// optional arrays that are unset will be nil.
		// The length expression in the loop checks the length of the nil slice, which evaluates to 0.
		// This means that arrays do not need additional dereferencing when optional.
		optional = false
		isArray = true
	}

	readerGetCode := jen.Id("reader").Dot("Get" + methodType.String()).Call(lengthExpr)
	if instruction.Offset != nil {
		var op string
		if *instruction.Offset < 0 {
			op = "-"
		} else {
			op = "+"
		}
		readerGetCode = readerGetCode.Op(op).Lit(int(math.Abs(float64(*instruction.Offset))))
	}

	var retCodes []jen.Code
	var assignRHS, assignLHS *jen.Statement
	hasAssignTarget := false
	if len(instructionName) == 0 && instruction.Content != nil {
		if isString {
			assignRHS = jen.Op("=").Add(readerGetCode)
			assignLHS = jen.Id("_")
		} else {
			assignRHS = jen.Add(readerGetCode)
			assignLHS = jen.Null()
		}
	} else {
		hasAssignTarget = true

		indexCode := jen.Null()
		if isArray {
			// pre-append an item to the array in the struct field
			var defaultCode *jen.Statement
			if isString {
				defaultCode = jen.Lit("")
			} else {
				defaultCode = jen.Lit(0)
			}

			retCodes = append(retCodes, jen.Id("s").Dot(instructionName).Op("=").Append(jen.Id("s").Dot(instructionName), defaultCode))
			indexCode = jen.Index(jen.Id("ndx"))
		}

		if optional {
			// instantiate the optional struct field
			retCodes = append(retCodes, jen.Id("s").Dot(instructionName).Op("=").New(jen.Do(func(s *jen.Statement) {
				if castType != nil {
					s.Add(castType)
				} else if isString {
					s.String()
				} else {
					s.Int()
				}
			})))

			assignLHS = jen.Op("*").Id("s").Dot(instructionName).Add(indexCode)
		} else {
			assignLHS = jen.Id("s").Dot(instructionName).Add(indexCode)
		}

		assignRHS = jen.Op("=").Do(func(s *jen.Statement) {
			if castType != nil {
				s.Add(castType).Call(readerGetCode)
			} else {
				s.Add(readerGetCode)
			}
		})
	}

	var assignBlock *jen.Statement
	if isString {
		assignBlock = jen.If(
			jen.List(assignLHS, jen.Id("err")).Add(assignRHS),
			jen.Id("err").Op("!=").Nil(),
		).Block(jen.Return()).Do(func(s *jen.Statement) {
			// _, err := strconv.ParseInt(*instruction.Length, 10, 32)
			if hasAssignTarget {
				// For compatibility: prior codegen inserted an extra newline after fixed strings that referenced a length field
				s.Line()
			}
		})
	} else {
		assignBlock = assignLHS.Add(assignRHS)
	}

	if optional {
		retCodes = append(retCodes, assignBlock)
		retCodes = []jen.Code{jen.If(jen.Id("reader").Dot("Remaining").Call().Op(">").Lit(0)).Block(retCodes...)}
	} else {
		retCodes = append(retCodes, assignBlock)
	}

	return retCodes
}

func getLengthExpression(instLength string) *jen.Statement {
	if parsed, err := strconv.ParseInt(instLength, 10, 32); err == nil {
		// string length is a numeric constant
		return jen.Lit(int(parsed))
	} else {
		// string length is a reference to another field
		return jen.Id("s").Dot(snakeCaseToPascalCase(instLength))
	}
}
