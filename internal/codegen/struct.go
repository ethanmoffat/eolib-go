package codegen

import (
	"fmt"
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

		// err = writeDeserializeBody(g, si, fullSpec)

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
			Name:         caseStructName,
			Comment:      c.Comment,
			Instructions: c.Instructions,
			PackageName:  si.PackageName,
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

			if len(instructionName) == 0 {
				instructionName = "(no_name)"
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
				serializeCodes = getSerializeForInstruction(instruction, types.NewSerializationType(typeName), false)
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
					if t := types.NewSerializationType(e.Type); t&types.Primitive > 0 {
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

/*
// flag that determines whether a chunked section is active or not
// this is used to determine if the next chunk should be selected in array delimiters and break bytes
var isChunked bool

	func writeDeserializeBody(output *strings.Builder, instructionList []xml.ProtocolInstruction, switchStructQualifier string, packageName string, fullSpec xml.Protocol) (imports []importInfo, err error) {
		for _, instruction := range instructionList {
			instructionType := instruction.XMLName.Local

			if instructionType == "chunked" {
				output.WriteString("\treader.SetIsChunked(true)\n")
				oldChunked := isChunked
				isChunked = true
				oldOuterInstructionList := outerInstructionList
				outerInstructionList = instructionList
				defer func() { isChunked = oldChunked; outerInstructionList = oldOuterInstructionList }()

				nextImports, err := writeDeserializeBody(output, instruction.Chunked, switchStructQualifier, packageName, fullSpec)
				if err != nil {
					return nil, err
				}
				imports = append(imports, nextImports...)

				output.WriteString("\treader.SetIsChunked(false)\n\n")
				continue
			}

			if instructionType == "break" {
				if isChunked {
					output.WriteString("\tif err = reader.NextChunk(); err != nil {\n\t\treturn\n\t}\n")
				} else {
					output.WriteString("\tif breakByte := reader.GetByte(); breakByte != 255 {\n")
					output.WriteString("\t\treturn fmt.Errorf(\"missing expected break byte\")\n")
					output.WriteString("\t}\n")
				}
				continue
			}

			instructionName := getInstructionName(instruction)

			if instructionType == "switch" {
				// get type of Value field
				switchFieldSanitizedType := ""
				switchFieldEnumType := ""
				for _, tmpInst := range append(outerInstructionList, instructionList...) {
					if tmpInst.XMLName.Local == "field" && snakeCaseToPascalCase(*tmpInst.Name) == instructionName {
						switchFieldEnumType = *tmpInst.Type
						switchFieldSanitizedType = sanitizeTypeName(switchFieldEnumType)
						break
					}
				}

				output.WriteString(fmt.Sprintf("\tswitch s.%s {\n", instructionName))

				for _, c := range instruction.Cases {
					if len(c.Instructions) == 0 {
						continue
					}

					var switchDataType string
					if c.Default {
						switchDataType = fmt.Sprintf("%sDataDefault", instructionName)
						output.WriteString("\tdefault:\n")
					} else {
						switchDataType = fmt.Sprintf("%sData%s", instructionName, c.Value)
						if _, err := strconv.ParseInt(c.Value, 10, 32); err != nil {
							// case is for an enum value
							if enumTypeInfo, ok := fullSpec.IsEnum(switchFieldEnumType); !ok {
								return nil, fmt.Errorf("type %s in switch is not an enum", switchFieldEnumType)
							} else {
								packageQualifier := ""
								if enumTypeInfo.Package != packageName {
									packageQualifier = enumTypeInfo.Package + "."
									imports = append(imports, importInfo{enumTypeInfo.Package, enumTypeInfo.PackagePath})
								}
								output.WriteString(fmt.Sprintf("\tcase %s%s_%s:\n", packageQualifier, switchFieldSanitizedType, c.Value))
							}
						} else {
							// case is for an integer constant
							output.WriteString(fmt.Sprintf("\tcase %s:\n", c.Value))
						}
					}

					output.WriteString(fmt.Sprintf("\t\ts.%sData = &%s%s{}\n", instructionName, switchStructQualifier, switchDataType))
					output.WriteString(fmt.Sprintf("\t\tif err = s.%sData.Deserialize(reader); err != nil {\n", instructionName))
					output.WriteString("\t\t\treturn\n\t\t}\n")
				}

				output.WriteString("\t}\n")

				continue
			}

			typeName, typeSize := getInstructionTypeName(instruction)

			instructionNameComment := instructionName
			if len(instructionNameComment) == 0 && instruction.Content != nil {
				instructionNameComment = *instruction.Content
			}
			output.WriteString(fmt.Sprintf("\t// %s : %s : %s\n", instructionNameComment, instructionType, *instruction.Type))

			var lenExpr string
			if instructionType == "array" {
				if instruction.Length != nil {
					lenExpr = "ndx < " + getLengthExpression(*instruction.Length)
				} else if (instruction.Delimited == nil || !*instruction.Delimited) && isChunked {
					rawLen, err := calculateTypeSize(typeName, fullSpec)
					if err != nil {
						lenExpr = "reader.Remaining() > 0"
					} else {
						lenExpr = "ndx < reader.Remaining() / " + strconv.Itoa(rawLen)
					}
				} else {
					lenExpr = "reader.Remaining() > 0"
				}

				output.WriteString(fmt.Sprintf("\tfor ndx := 0; %s; ndx++ {\n\t\t", lenExpr))
			}

			switch typeName {
			case "byte":
				castType := "int"
				writeGetTypeForDeserialize(output, instructionName, instruction, "Byte", &castType)
			case "char":
				writeGetTypeForDeserialize(output, instructionName, instruction, "Char", nil)
			case "short":
				writeGetTypeForDeserialize(output, instructionName, instruction, "Short", nil)
			case "three":
				writeGetTypeForDeserialize(output, instructionName, instruction, "Three", nil)
			case "int":
				writeGetTypeForDeserialize(output, instructionName, instruction, "Int", nil)
			case "bool":
				if len(typeSize) > 0 {
					typeName = string(unicode.ToUpper(rune(typeSize[0]))) + typeSize[1:]
				} else {
					typeName = "Char"
				}
				output.WriteString(fmt.Sprintf("\tif boolVal := reader.Get%s(); boolVal > 0 {\n", typeName))
				output.WriteString(fmt.Sprintf("\t\ts.%s = true\n\t} else {\n\t\ts.%s = false\n\t}\n", instructionName, instructionName))
			case "blob":
				writeGetTypeForDeserialize(output, instructionName, instruction, "Bytes", nil)
			case "string":
				if instruction.Length != nil && instructionType == "field" {
					if instruction.Padded != nil && *instruction.Padded {
						writeGetStringTypeForDeserialize(output, instructionName, instruction, "PaddedString")
					} else {
						writeGetStringTypeForDeserialize(output, instructionName, instruction, "FixedString")
					}
				} else {
					writeGetStringTypeForDeserialize(output, instructionName, instruction, "String")
				}
			case "encoded_string":
				if instruction.Length != nil && instructionType == "field" {
					if instruction.Padded != nil && *instruction.Padded {
						writeGetStringTypeForDeserialize(output, instructionName, instruction, "PaddedEncodedString")
					} else {
						writeGetStringTypeForDeserialize(output, instructionName, instruction, "FixedEncodedString")
					}
				} else {
					writeGetStringTypeForDeserialize(output, instructionName, instruction, "EncodedString")
				}
			default:
				if types.StructInfo, ok := fullSpec.IsStruct(typeName); ok {
					if instructionType == "array" {
						if packageName != types.StructInfo.Package {
							typeName = types.StructInfo.Package + "." + typeName
							imports = append(imports, importInfo{types.StructInfo.Package, types.StructInfo.PackagePath})
						}

						output.WriteString(fmt.Sprintf("\ts.%s = append(s.%s, %s{})\n", instructionName, instructionName, typeName))
						instructionName = instructionName + "[ndx]"
					}
					output.WriteString(fmt.Sprintf("\tif err = s.%s.Deserialize(reader); err != nil {\n\t\treturn\n\t}\n", instructionName))
				} else if e, ok := fullSpec.IsEnum(typeName); ok {
					switch e.Type {
					case "byte":
						fallthrough
					case "char":
						fallthrough
					case "short":
						fallthrough
					case "three":
						fallthrough
					case "int":
						if e.Package != packageName {
							typeName = fmt.Sprintf("%s.%s", e.Package, typeName)
						}
						writeGetTypeForDeserialize(output, instructionName, instruction, string(unicode.ToUpper(rune(e.Type[0])))+e.Type[1:], &typeName)
					}
					imports = append(imports, importInfo{e.Package, e.PackagePath})
				} else {
					panic("Unable to find type '" + typeName + "' when writing serialization function")
				}
			}

			delimited := instruction.Delimited != nil && *instruction.Delimited
			trailingDelimiter := instruction.TrailingDelimiter == nil || *instruction.TrailingDelimiter
			if instructionType == "array" {
				if delimited && isChunked {
					if !trailingDelimiter {
						if instruction.Length == nil {
							return nil, fmt.Errorf("delimited arrays with trailing-delimiter=false must have a length (array %s)", instructionName)
						}
						output.WriteString(fmt.Sprintf("\t\tif ndx + 1 < %s {\n", getLengthExpression(*instruction.Length)))
					}
					output.WriteString("\t\tif err = reader.NextChunk(); err != nil {\n\t\t\treturn\n\t\t}\n")
					if !trailingDelimiter {
						output.WriteString("\t\t}\n")
					}
				}
				output.WriteString("\t}\n\n")
			}
		}

		return
	}
*/
var _ = fmt.Printf

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

/*
	func writeGetTypeForDeserialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string, castType *string) {
		optional := instruction.Optional != nil && *instruction.Optional

		lengthExpr := ""
		if instruction.XMLName.Local != "array" {
			if instruction.Length != nil {
				lengthExpr = getLengthExpression(*instruction.Length)
			} else if methodType == "Bytes" {
				lengthExpr = "reader.Remaining()"
			}
		} else {
			// optional arrays that are unset will be nil.
			// The length expression in the loop checks the length of the nil slice, which evaluates to 0.
			// This means that arrays do not need additional dereferencing when optional.
			optional = false
		}

		if optional {
			output.WriteString("\tif reader.Remaining() > 0 {\n")
		}

		if len(instructionName) == 0 && instruction.Content != nil {
			output.WriteString(fmt.Sprintf("\treader.Get%s(%s)\n", methodType, lengthExpr))
		} else {
			if instruction.XMLName.Local == "array" {
				output.WriteString(fmt.Sprintf("\t\ts.%s = append(s.%s, 0)\n", instructionName, instructionName))
				instructionName = instructionName + "[ndx]"
			}

			if castType != nil {
				if optional {
					output.WriteString(fmt.Sprintf("\t\ts.%s = new(%s)\n\t\t*s.", instructionName, *castType))
				} else {
					output.WriteString("\t\ts.")
				}

				output.WriteString(fmt.Sprintf("%s = %s(reader.Get%s(%s))\n", instructionName, *castType, methodType, lengthExpr))
			} else {
				if optional {
					output.WriteString(fmt.Sprintf("\t\ts.%s = new(int)\n\t\t*s.", instructionName))
				} else {
					output.WriteString("\t\ts.")
				}

				output.WriteString(fmt.Sprintf("%s = reader.Get%s(%s)\n", instructionName, methodType, lengthExpr))
			}
		}

		if optional {
			output.WriteString("\t}\n")
		}
	}
*/
var _ = fmt.Printf

/*
	func writeGetStringTypeForDeserialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string) {
		optional := instruction.Optional != nil && *instruction.Optional

		lengthExpr := ""
		if instruction.XMLName.Local != "array" {
			if instruction.Length != nil {
				lengthExpr = getLengthExpression(*instruction.Length)
			}
		} else {
			optional = false
		}

		if optional {
			output.WriteString("\tif reader.Remaining() > 0 {\n")
		}

		if len(instructionName) == 0 && instruction.Content != nil {
			output.WriteString(fmt.Sprintf("\tif _, err = reader.Get%s(%s); err != nil {\n\t\treturn\n\t}\n", methodType, lengthExpr))
		} else {
			if instruction.XMLName.Local == "array" {
				output.WriteString(fmt.Sprintf("\t\ts.%s = append(s.%s, \"\")\n", instructionName, instructionName))
				instructionName = instructionName + "[ndx]"
			}

			if optional {
				output.WriteString(fmt.Sprintf("\t\ts.%s = new(string)\n\t\tif *s.", instructionName))
			} else {
				output.WriteString("\t\tif s.")
			}

			output.WriteString(fmt.Sprintf("%s, err = reader.Get%s(%s); err != nil {\n\t\treturn\n\t}\n\n", instructionName, methodType, lengthExpr))
		}

		if optional {
			output.WriteString("\t}\n")
		}
	}
*/
var _ = fmt.Printf

func getLengthExpression(instLength string) jen.Code {
	if parsed, err := strconv.ParseInt(instLength, 10, 32); err == nil {
		// string length is a numeric constant
		return jen.Lit(int(parsed))
	} else {
		// string length is a reference to another field
		return jen.Id("s").Dot(snakeCaseToPascalCase(instLength))
	}
}
