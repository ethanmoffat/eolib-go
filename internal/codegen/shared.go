package codegen

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/internal/xml"
)

// packageAliases is a map of package short names to package paths. For use with Jennifer.
var packageAliases = map[string]string{
	"data":     "github.com/ethanmoffat/eolib-go/pkg/eolib/data",
	"net":      "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/net",
	"protocol": "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol",
	"pub":      "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/pub",
}

func getPackageStatement(outputDir string) (string, error) {
	packageFileName := path.Join(outputDir, "package.go")
	fp, err := os.Open(packageFileName)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	packageInfo, err := io.ReadAll(fp)
	if err != nil {
		return "", err
	}

	packageInfoLines := strings.Split(string(packageInfo), "\n")
	for _, s := range packageInfoLines {
		if strings.HasPrefix(s, "package ") {
			return s, nil
		}
	}

	return "", fmt.Errorf("package name not found in %s", outputDir)
}

func getPackageName(outputDir string) (packageDeclaration string, err error) {
	if packageDeclaration, err = getPackageStatement(outputDir); err != nil {
		return
	}

	split := strings.Split(packageDeclaration, " ")
	if len(split) < 2 {
		packageDeclaration = ""
		err = fmt.Errorf("unable to determine package name from package declaration")
		return
	}

	packageDeclaration = split[1]
	return
}

func sanitizeComment(comment string) string {
	split := strings.Split(comment, "\n")

	for i := range split {
		split[i] = strings.TrimSpace(split[i])

		if len(split[i]) > 0 && !strings.HasSuffix(split[i], ".") {
			split[i] += "."
		}
	}

	return strings.Join(split, " ")
}

func sanitizeTypeName(typeName string) string {
	if strings.HasSuffix(typeName, "Type") {
		return typeName[:len(typeName)-4]
	}
	return typeName
}

func writeTypeCommentJen(f *jen.File, typeName string, comment string) {
	if comment = sanitizeComment(comment); len(comment) > 0 {
		f.Commentf("// %s :: %s", typeName, comment)
	}
}

func writeTypeComment(output *strings.Builder, typeName string, comment string) {
	if comment = sanitizeComment(comment); len(comment) > 0 {
		output.WriteString(fmt.Sprintf("// %s :: %s\n", typeName, comment))
	}
}

func writeInlineCommentJen(c jen.Code, comment string) {
	if comment = sanitizeComment(comment); len(comment) > 0 {
		switch v := c.(type) {
		case *jen.Statement:
			v.Comment(comment)
		case *jen.Group:
			v.Comment(comment)
		}
	}
}

func writeInlineComment(output *strings.Builder, comment string) {
	if comment = sanitizeComment(comment); len(comment) > 0 {
		output.WriteString(fmt.Sprintf(" // %s", comment))
	}
}

func writeToFileJen(f *jen.File, outFileName string) error {
	return f.Save(outFileName)
}

func writeToFile(outFileName string, outputText string) error {
	ofp, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer ofp.Close()

	n, err := ofp.Write([]byte(outputText))
	if err != nil {
		return err
	}
	if n != len(outputText) {
		return fmt.Errorf("wrote %d of %d bytes to file %s", n, len(outputText), outFileName)
	}

	return nil
}

type importInfo struct {
	Package string
	Path    string
}

func eoTypeToGoType(eoType string, currentPackage string, fullSpec xml.Protocol) (goType string, nextImport *importInfo) {
	if strings.ContainsRune(eoType, rune(':')) {
		eoType = strings.Split(eoType, ":")[0]
	}

	switch eoType {
	case "byte":
		fallthrough
	case "char":
		fallthrough
	case "short":
		fallthrough
	case "three":
		fallthrough
	case "int":
		return "int", nil
	case "bool":
		return "bool", nil
	case "blob":
		return "[]byte", nil
	case "string":
		fallthrough
	case "encoded_string":
		return "string", nil
	default:
		match := fullSpec.FindType(eoType)
		goType = eoType

		if structMatch, ok := match.(*xml.ProtocolStruct); ok && structMatch.Package != currentPackage {
			nextImport = &importInfo{structMatch.Package, structMatch.PackagePath}
		} else if enumMatch, ok := match.(*xml.ProtocolEnum); ok && enumMatch.Package != currentPackage {
			nextImport = &importInfo{enumMatch.Package, enumMatch.PackagePath}
		}

		if nextImport != nil {
			if val, ok := packageAliases[nextImport.Package]; ok {
				nextImport.Path = val
			}
		}

		return
	}
}

func snakeCaseToCamelCase(input string) string {
	if len(input) == 0 {
		return input
	}

	out := strings.Clone(input)
	for ndx := range out {
		if ndx < len(out) && out[ndx] == '_' {
			tmp := out[:ndx]

			if ndx+1 < len(out) {
				if ndx != 0 {
					tmp = tmp + string(unicode.ToUpper(rune(out[ndx+1])))
				} else {
					tmp = tmp + string(out[ndx+1])
				}
			}

			if ndx+2 < len(out) {
				tmp = tmp + out[ndx+2:]
			}

			out = tmp
		}
	}
	return out
}

func snakeCaseToPascalCase(input string) string {
	if len(input) == 0 {
		return input
	}

	camelCase := snakeCaseToCamelCase(input)
	firstRune := []rune(camelCase)[0]
	return string(unicode.ToUpper(firstRune)) + camelCase[1:]
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

func calculateTypeSize(typeName string, fullSpec xml.Protocol) (res int, err error) {
	var structInfo *xml.ProtocolStruct
	var isStruct bool
	if structInfo, isStruct = fullSpec.IsStruct(typeName); !isStruct {
		return getPrimitiveTypeSize(typeName, fullSpec)
	}

	var flattenedInstList []xml.ProtocolInstruction
	for _, instruction := range (*structInfo).Instructions {
		if instruction.XMLName.Local == "chunked" {
			flattenedInstList = append(flattenedInstList, instruction.Chunked...)
		} else {
			flattenedInstList = append(flattenedInstList, instruction)
		}
	}

	for _, instruction := range flattenedInstList {
		switch instruction.XMLName.Local {
		case "field":
			fieldTypeName, fieldTypeSize := getInstructionTypeName(instruction)
			if fieldTypeSize != "" {
				fieldTypeName = fieldTypeSize
			}

			if instruction.Length != nil {
				if length, err := strconv.ParseInt(*instruction.Length, 10, 32); err == nil {
					// length is a numeric constant
					res += int(length)
				} else {
					return 0, fmt.Errorf("instruction length %s must be a fixed size for %s (%s)", *instruction.Length, *instruction.Name, instruction.XMLName.Local)
				}
			} else {
				if nestedSize, err := getPrimitiveTypeSize(fieldTypeName, fullSpec); err != nil {
					return 0, err
				} else {
					res += nestedSize
				}
			}
		case "break":
			res += 1
		case "array":
		case "dummy":
		}
	}

	return
}

func getPrimitiveTypeSize(fieldTypeName string, fullSpec xml.Protocol) (int, error) {
	switch fieldTypeName {
	case "byte":
		fallthrough
	case "char":
		return 1, nil
	case "short":
		return 2, nil
	case "three":
		return 3, nil
	case "int":
		return 4, nil
	case "bool":
		return 1, nil
	case "blob":
		fallthrough
	case "string":
		fallthrough
	case "encoded_string":
		return 0, fmt.Errorf("cannot get size of %s without fixed length", fieldTypeName)
	default:
		if _, isStruct := fullSpec.IsStruct(fieldTypeName); isStruct {
			return calculateTypeSize(fieldTypeName, fullSpec)
		} else if e, isEnum := fullSpec.IsEnum(fieldTypeName); isEnum {
			enumTypeName := sanitizeTypeName(e.Type)
			return getPrimitiveTypeSize(enumTypeName, fullSpec)
		} else {
			return 0, fmt.Errorf("cannot get fixed size of unrecognized type %s", fieldTypeName)
		}
	}
}

// structInfo is a type representing the metadata about a struct that should be rendered as generated code.
// It represents the common properties of either a ProtocolPacket or a ProtocolStruct.
type structInfo struct {
	Name         string                    // Name is the name of the type. It is not converted from protocol naming convention (snake_case).
	Comment      string                    // Comment is an optional type comment for the struct.
	Instructions []xml.ProtocolInstruction // Instructions is a collection of instructions for the struct.
	PackageName  string                    // PackageName is the containing package name for the struct.

	Family                string // Family is the Packet Family of the struct, if the struct is a packet struct.
	Action                string // Action is the Packet Action of the struct, if the struct is a packet struct.
	SwitchStructQualifier string // SwitchStructQualifier is an additional qualifier prepended to structs used in switch cases in packets.
}

func getStructInfo(typeName string, fullSpec xml.Protocol) (si *structInfo, err error) {
	si = &structInfo{SwitchStructQualifier: ""}
	err = nil

	if structInfo, ok := fullSpec.IsStruct(typeName); ok {
		si.Name = structInfo.Name
		si.Comment = structInfo.Comment
		si.Instructions = structInfo.Instructions
		si.PackageName = structInfo.Package
	} else if packetInfo, ok := fullSpec.IsPacket(typeName); ok {
		si.Name = packetInfo.GetTypeName()
		si.Comment = packetInfo.Comment
		si.Instructions = packetInfo.Instructions
		si.PackageName = packetInfo.Package
		si.SwitchStructQualifier = packetInfo.Family + packetInfo.Action
		si.Family = packetInfo.Family
		si.Action = packetInfo.Action
	} else {
		si = nil
		err = fmt.Errorf("type %s is not a struct or packet in the spec", typeName)
	}

	return
}
