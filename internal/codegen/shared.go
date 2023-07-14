package codegen

import (
	"fmt"
	"io"
	"os"
	"path"
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
