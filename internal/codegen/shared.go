package codegen

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/ethanmoffat/eolib-go/internal/xml"
)

func getPackageName(outputDir string) (string, error) {
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

func writeTypeComment(output *strings.Builder, typeName string, comment string) {
	if comment = sanitizeComment(comment); len(comment) > 0 {
		output.WriteString(fmt.Sprintf("// %s :: %s\n", typeName, comment))
	}
}

func writeInlineComment(output *strings.Builder, comment string) {
	if comment = sanitizeComment(comment); len(comment) > 0 {
		output.WriteString(fmt.Sprintf(" // %s", comment))
	}
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

		if structMatch, ok := match.(*xml.ProtocolStruct); ok {
			if structMatch.Package != currentPackage {
				goType = structMatch.Package + "." + eoType
				nextImport = &importInfo{structMatch.Package, structMatch.PackagePath}
			} else {
				goType = eoType
			}
		} else if enumMatch, ok := match.(*xml.ProtocolEnum); ok {
			if enumMatch.Package != currentPackage {
				goType = enumMatch.Package + "." + eoType
				nextImport = &importInfo{enumMatch.Package, enumMatch.PackagePath}
			} else {
				goType = eoType
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
		return
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
				if nestedSize, err := getPrimitizeTypeSize(fieldTypeName, fullSpec); err != nil {
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

func getPrimitizeTypeSize(fieldTypeName string, fullSpec xml.Protocol) (int, error) {
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
			return getPrimitizeTypeSize(enumTypeName, fullSpec)
		} else {
			return 0, fmt.Errorf("cannot get fixed size of unrecognized type %s", fieldTypeName)
		}
	}
}
