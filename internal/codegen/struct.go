package codegen

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"unicode"

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
	packageDeclaration, err := getPackageName(outputDir)
	if err != nil {
		return err
	}

	output := strings.Builder{}
	output.WriteString(packageDeclaration + "\n\n")
	var outputText string
	if len(typeNames) > 0 {
		output.WriteString("import (\n\t\"fmt\"\n\t\"github.com/ethanmoffat/eolib-go/pkg/eolib/data\"\n<REPLACE>)\n\n// Ensure fmt import is referenced in generated code\nvar _ = fmt.Printf\n\n")

		var imports []importInfo
		for _, typeName := range typeNames {
			if nextImports, err := writeStruct(&output, typeName, fullSpec); err != nil {
				return err
			} else {
				imports = append(imports, nextImports...)
			}
		}

		outputText = output.String()

		var matches map[string]bool = make(map[string]bool)
		var importText string
		for _, imp := range imports {
			if _, ok := matches[imp.Package]; !ok && strings.Split(packageDeclaration, " ")[1] != imp.Package {
				importText = importText + fmt.Sprintf("\t%s \"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol%s\"\n", imp.Package, imp.Path)
				matches[imp.Package] = true
			}
		}
		outputText = strings.ReplaceAll(outputText, "<REPLACE>", importText)
	} else {
		outputText = output.String()
	}

	outFileName := path.Join(outputDir, outputFileName)
	return writeToFile(outFileName, outputText)
}

func writeStruct(output *strings.Builder, typeName string, fullSpec xml.Protocol) (importPaths []importInfo, err error) {
	var name string
	var comment string
	var instructions []xml.ProtocolInstruction
	var packageName string

	switchStructQualifier := ""
	if structInfo, ok := fullSpec.IsStruct(typeName); ok {
		name = structInfo.Name
		comment = structInfo.Comment
		instructions = structInfo.Instructions
		packageName = structInfo.Package
	} else if packetInfo, ok := fullSpec.IsPacket(typeName); ok {
		name = packetInfo.GetTypeName()
		comment = packetInfo.Comment
		instructions = packetInfo.Instructions
		packageName = packetInfo.Package
		switchStructQualifier = packetInfo.Family + packetInfo.Action
	} else {
		return nil, fmt.Errorf("type %s is not a struct or packet in the spec", typeName)
	}

	structName := snakeCaseToPascalCase(name)
	writeTypeComment(output, structName, comment)

	// write out fields
	output.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	switches, nextImports := writeStructFields(output, instructions, switchStructQualifier, packageName, fullSpec)
	importPaths = append(importPaths, nextImports...)
	output.WriteString("}\n\n")

	for _, sw := range switches {
		if nextImports, err = writeSwitchStructs(output, *sw, switchStructQualifier, packageName, fullSpec); err != nil {
			return nil, err
		}
		importPaths = append(importPaths, nextImports...)
	}

	// write out serialize method
	output.WriteString(fmt.Sprintf("func (s *%s) Serialize(writer data.EoWriter) (err error) {\n", structName))
	output.WriteString("\toldSanitizeStrings := writer.SanitizeStrings\n")
	output.WriteString("\tdefer func() {writer.SanitizeStrings = oldSanitizeStrings}()\n\n")
	if nextImports, err = writeSerializeBody(output, instructions, switchStructQualifier, packageName, fullSpec); err != nil {
		return nil, err
	}
	importPaths = append(importPaths, nextImports...)
	output.WriteString("\treturn\n")
	output.WriteString("}\n\n")

	// write out deserialize method
	output.WriteString(fmt.Sprintf("func (s *%s) Deserialize(reader data.EoReader) (err error) {\n", structName))
	output.WriteString("\toldChunkedReadingMode := reader.GetChunkedReadingMode()\n")
	output.WriteString("\tdefer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()\n\n")
	if nextImports, err = writeDeserializeBody(output, instructions, switchStructQualifier, packageName, fullSpec); err != nil {
		return nil, err
	}
	importPaths = append(importPaths, nextImports...)
	output.WriteString("\n\treturn\n}\n\n")

	return
}

func writeStructFields(output *strings.Builder, instructions []xml.ProtocolInstruction, switchStructQualifier string, packageName string, fullSpec xml.Protocol) (switches []*xml.ProtocolInstruction, imports []importInfo) {
	for i, inst := range instructions {
		var instName string

		if inst.Name != nil {
			instName = snakeCaseToPascalCase(*inst.Name)
		} else if inst.Field != nil {
			instName = snakeCaseToPascalCase(*inst.Field)
		}

		var typeName string
		if inst.Type != nil {
			var nextImport *importInfo
			if typeName, nextImport = eoTypeToGoType(*inst.Type, packageName, fullSpec); nextImport != nil {
				imports = append(imports, *nextImport)
			}
		}

		switch inst.XMLName.Local {
		case "field":
			if len(instName) > 0 {
				output.WriteString(fmt.Sprintf("\t%s %s", instName, typeName))
			}
		case "array":
			output.WriteString(fmt.Sprintf("\t%s []%s", instName, typeName))
		case "length":
			output.WriteString(fmt.Sprintf("\t%s %s", instName, typeName))
		case "switch":
			output.WriteString(fmt.Sprintf("\t%sData %s%sData", instName, switchStructQualifier, instName))
			switches = append(switches, &instructions[i])
		case "chunked":
			nextSwitches, nextImports := writeStructFields(output, inst.Chunked, switchStructQualifier, packageName, fullSpec)
			switches = append(switches, nextSwitches...)
			imports = append(imports, nextImports...)
		case "dummy":
		case "break":
			continue // no data to write
		}

		if inst.Comment != nil {
			writeInlineComment(output, *inst.Comment)
		}

		output.WriteString("\n")
	}

	return
}

func writeSwitchStructs(output *strings.Builder, switchInst xml.ProtocolInstruction, switchStructQualifier string, packageName string, fullSpec xml.Protocol) (imports []importInfo, err error) {
	if switchInst.XMLName.Local != "switch" {
		return
	}

	switchInterfaceName := fmt.Sprintf("%sData", snakeCaseToPascalCase(*switchInst.Field))
	if len(switchStructQualifier) > 0 {
		switchInterfaceName = switchStructQualifier + switchInterfaceName
	}

	if switchInst.Comment != nil {
		writeTypeComment(output, switchInterfaceName, *switchInst.Comment)
	}

	output.WriteString(fmt.Sprintf("type %s interface {\n\tprotocol.EoData\n}\n\n", switchInterfaceName))

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

		writeTypeComment(output, caseStructName, c.Comment)

		output.WriteString(fmt.Sprintf("type %s struct {\n", caseStructName))
		switches, nextImports := writeStructFields(output, c.Instructions, switchStructQualifier, packageName, fullSpec)
		imports = append(imports, nextImports...)
		output.WriteString("}\n\n")

		for _, sw := range switches {
			if nextImports, err = writeSwitchStructs(output, *sw, switchStructQualifier, packageName, fullSpec); err != nil {
				return nil, err
			}
			imports = append(imports, nextImports...)
		}

		// write out serialize method
		output.WriteString(fmt.Sprintf("func (s *%s) Serialize(writer data.EoWriter) (err error) {\n", caseStructName))
		output.WriteString("\toldSanitizeStrings := writer.SanitizeStrings\n")
		output.WriteString("\tdefer func() {writer.SanitizeStrings = oldSanitizeStrings}()\n\n")
		if nextImports, err = writeSerializeBody(output, c.Instructions, switchStructQualifier, packageName, fullSpec); err != nil {
			return nil, err
		}
		imports = append(imports, nextImports...)
		output.WriteString("\treturn\n")
		output.WriteString("}\n\n")

		// write out deserialize method
		output.WriteString(fmt.Sprintf("func (s *%s) Deserialize(reader data.EoReader) (err error) {\n", caseStructName))
		output.WriteString("\toldChunkedReadingMode := reader.GetChunkedReadingMode()\n")
		output.WriteString("\tdefer func() { reader.SetChunkedReadingMode(oldChunkedReadingMode) }()\n\n")
		if nextImports, err = writeDeserializeBody(output, c.Instructions, switchStructQualifier, packageName, fullSpec); err != nil {
			return nil, err
		}
		imports = append(imports, nextImports...)
		output.WriteString("\n\treturn\n}\n\n")
	}

	return
}

// used to track the 'outer' list of instructions when a <chunked> instruction is encountered
// this allows any nested <switch> instructions to search both the instructions in the <chunked> section at the same level
//
//	as well as the outer instructions in the <struct> or <packet> when determining the type of the switch field
var outerInstructionList []xml.ProtocolInstruction

func writeSerializeBody(output *strings.Builder, instructionList []xml.ProtocolInstruction, switchStructQualifier string, packageName string, fullSpec xml.Protocol) (imports []importInfo, err error) {
	for _, instruction := range instructionList {
		instructionType := instruction.XMLName.Local

		if instructionType == "chunked" {
			output.WriteString("\twriter.SanitizeStrings = true\n")
			oldOuterInstructionList := outerInstructionList
			outerInstructionList = instructionList
			defer func() { outerInstructionList = oldOuterInstructionList }()

			if nextImports, err := writeSerializeBody(output, instruction.Chunked, switchStructQualifier, packageName, fullSpec); err != nil {
				return nil, err
			} else {
				imports = append(imports, nextImports...)
			}

			output.WriteString("\twriter.SanitizeStrings = false\n")
			continue
		}

		if instructionType == "break" {
			output.WriteString("\twriter.AddByte(0xFF)\n")
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

				if len(switchDataType) > 0 {
					output.WriteString(fmt.Sprintf("\t\tswitch s.%sData.(type) {\n", instructionName))
					output.WriteString(fmt.Sprintf("\t\tcase *%s%s:\n\t\t", switchStructQualifier, switchDataType))
				}
				output.WriteString(fmt.Sprintf("\t\t\tif err = s.%sData.Serialize(writer); err != nil {\n", instructionName))
				output.WriteString("\t\t\t\treturn\n\t\t\t}\n")

				if len(switchDataType) > 0 {
					output.WriteString(fmt.Sprintf("\t\tdefault:\n\t\t\terr = fmt.Errorf(\"invalid switch struct type for switch value %%d\", s.%s)\n\t\t\treturn\n\t\t}\n", instructionName))
				}
			}

			output.WriteString("\t}\n")
			continue
		}

		typeName, typeSize := getInstructionTypeName(instruction)

		output.WriteString("\t// " + instructionName + " : " + instructionType + " : " + *instruction.Type + "\n")

		delimited := instruction.Delimited != nil && *instruction.Delimited
		trailingDelimiter := instruction.TrailingDelimiter == nil || *instruction.TrailingDelimiter
		if instructionType == "array" {
			var lenExpr string
			if instruction.Length != nil {
				lenExpr = getLengthExpression(*instruction.Length)
			} else {
				lenExpr = fmt.Sprintf("len(s.%s)", instructionName)
			}

			output.WriteString(fmt.Sprintf("\tfor ndx := 0; ndx < %s; ndx++ {\n\t\t", lenExpr))

			if delimited && !trailingDelimiter {
				output.WriteString("\t\tif ndx > 0 {\n\t\t\twriter.AddByte(0xFF)\n\t\t}\n\n")
			}
		}

		switch typeName {
		case "byte":
			writeAddTypeForSerialize(output, instructionName, instruction, "Byte", false)
		case "char":
			writeAddTypeForSerialize(output, instructionName, instruction, "Char", false)
		case "short":
			writeAddTypeForSerialize(output, instructionName, instruction, "Short", false)
		case "three":
			writeAddTypeForSerialize(output, instructionName, instruction, "Three", false)
		case "int":
			writeAddTypeForSerialize(output, instructionName, instruction, "Int", false)
		case "bool":
			if len(typeSize) > 0 {
				typeName = string(unicode.ToUpper(rune(typeSize[0]))) + typeSize[1:]
			} else {
				typeName = "Char"
			}
			output.WriteString(fmt.Sprintf("\tif s.%s {\n", instructionName))
			output.WriteString(fmt.Sprintf("\t\terr = writer.Add%s(1)\n\t} else {\n\t\terr = writer.Add%s(0)\n\t}\n", typeName, typeName))
			output.WriteString("\tif err != nil {\n\t\treturn\n\t}\n\n")
		case "blob":
			writeAddTypeForSerialize(output, instructionName, instruction, "Bytes", false)
		case "string":
			if instruction.Length != nil && instructionType == "field" {
				if instruction.Padded != nil && *instruction.Padded {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "PaddedString")
				} else {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "FixedString")
				}
			} else {
				writeAddStringTypeForSerialize(output, instructionName, instruction, "String")
			}
		case "encoded_string":
			if instruction.Length != nil && instructionType == "field" {
				if instruction.Padded != nil && *instruction.Padded {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "PaddedEncodedString")
				} else {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "FixedEncodedString")
				}
			} else {
				writeAddStringTypeForSerialize(output, instructionName, instruction, "EncodedString")
			}
		default:
			if _, ok := fullSpec.IsStruct(typeName); ok {
				if instructionType == "array" {
					instructionName = instructionName + "[ndx]"
				}
				output.WriteString(fmt.Sprintf("\tif err = s.%s.Serialize(writer); err != nil {\n\t\treturn\n\t}\n", instructionName))
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
					writeAddTypeForSerialize(output, instructionName, instruction, string(unicode.ToUpper(rune(e.Type[0])))+e.Type[1:], true)
				}
			} else {
				panic("Unable to find type '" + typeName + "' when writing serialization function")
			}
		}

		if instructionType == "array" {
			if delimited && trailingDelimiter {
				output.WriteString("\t\twriter.AddByte(0xFF)\n")
			}
			output.WriteString("\t}\n\n")
		}
	}

	return
}

// flag that determines whether a chunked section is active or not
// this is used to determine if the next chunk should be selected in array delimiters and break bytes
var isChunked bool

func writeDeserializeBody(output *strings.Builder, instructionList []xml.ProtocolInstruction, switchStructQualifier string, packageName string, fullSpec xml.Protocol) (imports []importInfo, err error) {
	for _, instruction := range instructionList {
		instructionType := instruction.XMLName.Local

		if instructionType == "chunked" {
			output.WriteString("\treader.SetChunkedReadingMode(true)\n")
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

			output.WriteString("\treader.SetChunkedReadingMode(false)\n\n")
			continue
		}

		if instructionType == "break" {
			if isChunked {
				output.WriteString("\tif err = reader.NextChunk(); err != nil {\n\t\treturn\n\t}\n")
			} else {
				output.WriteString("\tif breakByte := reader.GetByte(); breakByte != 0xFF {\n")
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

		output.WriteString("\t// " + instructionName + " : " + instructionType + " : " + *instruction.Type + "\n")

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
			if structInfo, ok := fullSpec.IsStruct(typeName); ok {
				if instructionType == "array" {
					if packageName != structInfo.Package {
						typeName = structInfo.Package + "." + typeName
						imports = append(imports, importInfo{structInfo.Package, structInfo.PackagePath})
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

func getInstructionName(inst xml.ProtocolInstruction) (instName string) {
	if inst.Name != nil {
		instName = snakeCaseToPascalCase(*inst.Name)
	} else if inst.Field != nil {
		instName = snakeCaseToPascalCase(*inst.Field)
	}
	return
}

func writeAddTypeForSerialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string, needsCastToInt bool) {
	if len(instructionName) == 0 && instruction.Content != nil {
		instructionName = *instruction.Content
	} else {
		instructionName = "s." + instructionName
	}

	if instruction.XMLName.Local == "array" {
		instructionName = instructionName + "[ndx]"
	}

	if needsCastToInt {
		instructionName = "int(" + instructionName + ")"
	}
	output.WriteString(fmt.Sprintf("\tif err = writer.Add%s(%s); err != nil {\n\t\treturn\n\t}\n\n", methodType, instructionName))
}

func writeGetTypeForDeserialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string, castType *string) {
	lengthExpr := ""
	if instruction.XMLName.Local != "array" {
		if instruction.Length != nil {
			lengthExpr = getLengthExpression(*instruction.Length)
		} else if methodType == "Bytes" {
			lengthExpr = "reader.Remaining()"
		}
	}

	if len(instructionName) == 0 && instruction.Content != nil {
		output.WriteString(fmt.Sprintf("\treader.Get%s(%s)\n", methodType, lengthExpr))
	} else {
		if instruction.XMLName.Local == "array" {
			instructionName = instructionName + "[ndx]"
		}

		if castType != nil {
			output.WriteString(fmt.Sprintf("\ts.%s = %s(reader.Get%s(%s))\n", instructionName, *castType, methodType, lengthExpr))
		} else {
			output.WriteString(fmt.Sprintf("\ts.%s = reader.Get%s(%s)\n", instructionName, methodType, lengthExpr))
		}
	}
}

func writeAddStringTypeForSerialize(output *strings.Builder, instName string, inst xml.ProtocolInstruction, methodType string) {
	if len(instName) == 0 && inst.Content != nil {
		instName = `"` + *inst.Content + `"`
	} else {
		instName = "s." + instName
	}

	if inst.XMLName.Local == "array" {
		instName = instName + "[ndx]"
	} else if inst.Length != nil {
		instName = instName + ", " + getLengthExpression(*inst.Length)
	}

	output.WriteString(fmt.Sprintf("\tif err = writer.Add%s(%s); err != nil {\n\t\treturn\n\t}\n\n", methodType, instName))
}

func writeGetStringTypeForDeserialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string) {
	lengthExpr := ""
	if instruction.XMLName.Local != "array" && instruction.Length != nil {
		lengthExpr = getLengthExpression(*instruction.Length)
	}

	if len(instructionName) == 0 && instruction.Content != nil {
		output.WriteString(fmt.Sprintf("\tif _, err = reader.Get%s(%s); err != nil {\n\t\treturn\n\t}\n", methodType, lengthExpr))
	} else {
		if instruction.XMLName.Local == "array" {
			instructionName = instructionName + "[ndx]"
		}

		output.WriteString(fmt.Sprintf("\tif s.%s, err = reader.Get%s(%s); err != nil {\n\t\treturn\n\t}\n\n", instructionName, methodType, lengthExpr))
	}
}

func getLengthExpression(instLength string) string {
	if _, err := strconv.ParseInt(instLength, 10, 32); err == nil {
		// string length is a numeric constant
		return instLength
	} else {
		// string length is a reference to another field
		return "s." + snakeCaseToPascalCase(instLength)
	}
}
