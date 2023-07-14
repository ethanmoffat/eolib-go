package codegen

import (
	"fmt"
	"path"
	"strings"

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
	output.WriteString("import (\n\t\"github.com/ethanmoffat/eolib-go/pkg/eolib/data\"\n%s\n)\n\n")

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
	output.WriteString(fmt.Sprintf("func (s *%s) Serialize(writer data.EoWriter) error {\n", structName))
	output.WriteString("\treturn nil\n")
	output.WriteString("}\n\n")

	// write out deserialize method
	output.WriteString(fmt.Sprintf("func (s *%s) Deserialize(reader data.EoReader) error {\n", structName))
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
		output.WriteString(fmt.Sprintf("func (s *%s) Serialize(writer data.EoWriter) error {\n", caseStructName))
		output.WriteString("\treturn nil\n")
		output.WriteString("}\n\n")

		// write out deserialize method
		output.WriteString(fmt.Sprintf("func (s *%s) Deserialize(reader data.EoReader) error {\n", caseStructName))
		output.WriteString("\treturn nil\n")
		output.WriteString("}\n\n")
	}

	return
}
