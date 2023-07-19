package codegen

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/ethanmoffat/eolib-go/internal/xml"
)

const structFileName = "structs_generated.go"

func GenerateStructs(outputDir string, structs []xml.ProtocolStruct, fullSpec xml.Protocol) error {
	packageDeclaration, err := getPackageName(outputDir)
	if err != nil {
		return err
	}

	output := strings.Builder{}
	output.WriteString(packageDeclaration + "\n\n")
	output.WriteString("import (\n\t\"github.com/ethanmoffat/eolib-go/pkg/eolib/data\"\n%s)\n\n")

	var imports []importInfo
	for _, s := range structs {
		imports = append(imports, writeStruct(&output, s, fullSpec)...)
	}

	outputText := output.String()

	var matches map[string]bool = make(map[string]bool)
	var importText string
	for _, imp := range imports {
		if _, ok := matches[imp.Package]; !ok && strings.Split(packageDeclaration, " ")[1] != imp.Package {
			importText = importText + fmt.Sprintf("\t%s \"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol%s\"\n", imp.Package, imp.Path)
			matches[imp.Package] = true
		}
	}
	outputText = fmt.Sprintf(outputText, importText)

	outFileName := path.Join(outputDir, structFileName)
	return writeToFile(outFileName, outputText)
}

func writeStruct(output *strings.Builder, s xml.ProtocolStruct, fullSpec xml.Protocol) (importPaths []importInfo) {
	structName := snakeCaseToPascalCase(s.Name)
	writeTypeComment(output, structName, s.Comment)

	// write out fields
	output.WriteString(fmt.Sprintf("type %s struct {\n", structName))
	switches, nextImports := writeStructFields(output, s.Instructions, s.Package, fullSpec)
	importPaths = append(importPaths, nextImports...)
	output.WriteString("}\n\n")

	for _, sw := range switches {
		importPaths = append(importPaths, writeSwitchStructs(output, *sw, s.Package, fullSpec)...)
	}

	// write out serialize method
	output.WriteString(fmt.Sprintf("func (s *%s) Serialize(writer data.EoWriter) (err error) {\n", structName))
	output.WriteString("\toldSanitizeStrings := writer.SanitizeStrings\n")
	output.WriteString("\tdefer func() {writer.SanitizeStrings = oldSanitizeStrings}()\n\n")
	writeSerializeBody(output, s.Instructions, s.Package, fullSpec)
	output.WriteString("\treturn\n")
	output.WriteString("}\n\n")

	// write out deserialize method
	output.WriteString(fmt.Sprintf("func (s *%s) Deserialize(reader data.EoReader) (err error) {\n", structName))
	output.WriteString("\treturn nil\n")
	output.WriteString("}\n\n")

	return
}

func writeStructFields(output *strings.Builder, instructions []xml.ProtocolInstruction, packageName string, fullSpec xml.Protocol) (switches []*xml.ProtocolInstruction, imports []importInfo) {
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
			output.WriteString(fmt.Sprintf("\t%sData %sData", instName, instName))
			switches = append(switches, &instructions[i])
		case "chunked":
			nextSwitches, nextImports := writeStructFields(output, inst.Chunked, packageName, fullSpec)
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

func writeSwitchStructs(output *strings.Builder, switchInst xml.ProtocolInstruction, packageName string, fullSpec xml.Protocol) (imports []importInfo) {
	if switchInst.XMLName.Local != "switch" {
		return
	}

	switchInterfaceName := fmt.Sprintf("%sData", snakeCaseToPascalCase(*switchInst.Field))
	if switchInst.Comment != nil {
		writeTypeComment(output, switchInterfaceName, *switchInst.Comment)
	}

	output.WriteString(fmt.Sprintf("type %s interface {\n\tprotocol.EoData\n}\n\n", switchInterfaceName))

	for _, c := range switchInst.Cases {
		if len(c.Instructions) == 0 {
			continue
		}

		// TODO: handle default (for packets)
		// TODO: handle integer constant cases (for packets)
		caseName := snakeCaseToPascalCase(c.Value)
		caseStructName := fmt.Sprintf("%s%s", switchInterfaceName, caseName)
		writeTypeComment(output, caseStructName, c.Comment)

		output.WriteString(fmt.Sprintf("type %s struct {\n", caseStructName))
		switches, nextImports := writeStructFields(output, c.Instructions, packageName, fullSpec)
		imports = append(imports, nextImports...)
		output.WriteString("}\n\n")

		for _, sw := range switches {
			imports = append(imports, writeSwitchStructs(output, *sw, packageName, fullSpec)...)
		}

		// write out serialize method
		output.WriteString(fmt.Sprintf("func (s *%s) Serialize(writer data.EoWriter) (err error) {\n", caseStructName))
		output.WriteString("\toldSanitizeStrings := writer.SanitizeStrings\n")
		output.WriteString("\tdefer func() {writer.SanitizeStrings = oldSanitizeStrings}()\n\n")
		writeSerializeBody(output, c.Instructions, packageName, fullSpec)
		output.WriteString("\treturn\n")
		output.WriteString("}\n\n")

		// write out deserialize method
		output.WriteString(fmt.Sprintf("func (s *%s) Deserialize(reader data.EoReader) (err error) {\n", caseStructName))
		writeSwitchStructDeserializeBody(output, c.Instructions, packageName)
		output.WriteString("}\n\n")
	}

	return
}

func writeSerializeBody(output *strings.Builder, instructionList []xml.ProtocolInstruction, packageName string, fullSpec xml.Protocol) {
	for _, instruction := range instructionList {
		instructionType := instruction.XMLName.Local

		if instructionType == "chunked" {
			output.WriteString("\twriter.SanitizeStrings = true\n")
			writeSerializeBody(output, instruction.Chunked, packageName, fullSpec)
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
			switchFieldType := ""
			for _, tmpInst := range instructionList {
				if tmpInst.XMLName.Local == "field" && snakeCaseToPascalCase(*tmpInst.Name) == instructionName {
					switchFieldType = sanitizeTypeName(*tmpInst.Type)
				}
			}

			output.WriteString(fmt.Sprintf("\tswitch s.%s {\n", instructionName))

			for _, c := range instruction.Cases {
				if c.Default {
					// TODO: handle default (for packets)
					// output.WriteString(fmt.Sprintf("\tdefault:\n\t\t"))
				} else {
					if _, err := strconv.ParseInt(c.Value, 10, 32); err != nil {
						// case is for an enum value
						output.WriteString(fmt.Sprintf("\tcase %s_%s:\n", switchFieldType, c.Value))
						output.WriteString(fmt.Sprintf("\t\tif err = %sData.Serialize(s.%sData, writer); err != nil {\n", instructionName, instructionName))
						output.WriteString("\t\t\treturn\n\t\t}\n")
					} else {
						// case is for an integer constant
						output.WriteString(fmt.Sprintf("\tcase %s:\n", c.Value))
						// TODO: handle integer constant cases (for packets)
					}
				}
			}

			output.WriteString("\t}\n")
			continue
		}

		typeName, typeSize := getInstructionTypeName(instruction)

		output.WriteString("\t// " + instructionName + " : " + instructionType + " : " + *instruction.Type + "\n")

		if instructionType == "array" {
			var lenExpr string
			if instruction.Length != nil {
				lenExpr = getLengthExpression(*instruction.Length)
			} else {
				lenExpr = fmt.Sprintf("len(s.%s)", instructionName)
			}

			output.WriteString(fmt.Sprintf("\tfor ndx := 0; ndx < %s; ndx++ {\n\t\t", lenExpr))
		}

		switch typeName {
		case "byte":
			writeAddType(output, instructionName, instruction, "Byte", false)
		case "char":
			writeAddType(output, instructionName, instruction, "Char", false)
		case "short":
			writeAddType(output, instructionName, instruction, "Short", false)
		case "three":
			writeAddType(output, instructionName, instruction, "Three", false)
		case "int":
			writeAddType(output, instructionName, instruction, "Int", false)
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
			writeAddType(output, instructionName, instruction, "Bytes", false)
		case "string":
			if instruction.Length != nil {
				if instruction.Padded != nil && *instruction.Padded {
					writeAddStringType(output, instructionName, instruction, "PaddedString")
				} else {
					writeAddStringType(output, instructionName, instruction, "FixedString")
				}
			} else {
				writeAddStringType(output, instructionName, instruction, "String")
			}
		case "encoded_string":
			if instruction.Length != nil {
				if instruction.Padded != nil && *instruction.Padded {
					writeAddStringType(output, instructionName, instruction, "PaddedEncodedString")
				} else {
					writeAddStringType(output, instructionName, instruction, "FixedEncodedString")
				}
			} else {
				writeAddStringType(output, instructionName, instruction, "EncodedString")
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
					writeAddType(output, instructionName, instruction, string(unicode.ToUpper(rune(e.Type[0])))+e.Type[1:], true)
				}
			} else {
				panic("Unable to find type '" + typeName + "' when writing serialization function")
			}
		}

		if instructionType == "array" {
			output.WriteString("\t}\n\n")
		}
	}
}

func writeSwitchStructDeserializeBody(output *strings.Builder, instList []xml.ProtocolInstruction, packageName string) {
	output.WriteString("\treturn nil\n")
}

func getInstructionName(inst xml.ProtocolInstruction) (instName string) {
	if inst.Name != nil {
		instName = snakeCaseToPascalCase(*inst.Name)
	} else if inst.Field != nil {
		instName = snakeCaseToPascalCase(*inst.Field)
	}
	return
}

func getInstructionTypeName(inst xml.ProtocolInstruction) (typeName string, typeSize string) {
	if inst.Type == nil {
		return
	}

	if strings.ContainsRune(*inst.Type, rune(':')) {
		split := strings.Split(*inst.Type, ":")
		typeName, typeSize = split[0], split[1]
	} else {
		typeName = *inst.Type
	}

	return
}

func writeAddType(output *strings.Builder, instName string, inst xml.ProtocolInstruction, methodType string, needsCastToInt bool) {
	if len(instName) == 0 && inst.Content != nil {
		instName = *inst.Content
	} else {
		instName = "s." + instName
	}

	if inst.XMLName.Local == "array" {
		instName = instName + "[ndx]"
	}

	if needsCastToInt {
		instName = "int(" + instName + ")"
	}
	output.WriteString(fmt.Sprintf("\tif err = writer.Add%s(%s); err != nil {\n\t\treturn\n\t}\n\n", methodType, instName))
}

func writeAddStringType(output *strings.Builder, instName string, inst xml.ProtocolInstruction, methodType string) {
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

func getLengthExpression(instLength string) string {
	if _, err := strconv.ParseInt(instLength, 10, 32); err == nil {
		// string length is a numeric constant
		return instLength
	} else {
		// string length is a reference to another field
		return "s." + snakeCaseToPascalCase(instLength)
	}
}
