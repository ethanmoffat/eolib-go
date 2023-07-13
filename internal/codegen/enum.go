package codegen

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/ethanmoffat/eolib-go/internal/xml"
)

func GenerateEnums(outputDir string, enums []xml.ProtocolEnum) error {
	outFileName := path.Join(outputDir, "enums_generated.go")

	packageFileName := path.Join(outputDir, "package.go")
	fp, err := os.Open(packageFileName)
	if err != nil {
		return err
	}
	defer fp.Close()

	packageInfo, err := io.ReadAll(fp)
	if err != nil {
		return err
	}

	var packageName string
	packageInfoLines := strings.Split(string(packageInfo), "\n")
	for _, s := range packageInfoLines {
		if strings.HasPrefix(s, "package ") {
			packageName = s
			break
		}
	}

	output := strings.Builder{}
	output.WriteString(packageName + "\n")
	output.WriteString("\n")

	for _, e := range enums {
		if len(e.Comment) > 0 {
			output.WriteString(fmt.Sprintf("// %s :: %s\n", e.Name, sanitizeComment(e.Comment)))
		}

		output.WriteString(fmt.Sprintf("type %s int\n\n", e.Name))
		output.WriteString("const (\n")

		for i, v := range e.Values {
			if i == 0 {
				output.WriteString(fmt.Sprintf("\t%s_%s %s = iota", sanitizeTypeName(e.Name), v.Name, e.Name))
			} else {
				output.WriteString(fmt.Sprintf("\t%s_%s", sanitizeTypeName(e.Name), v.Name))
			}

			if len(v.Comment) > 0 {
				output.WriteString(fmt.Sprintf(" // %s", sanitizeComment(v.Comment)))
			}

			output.WriteString("\n")
		}

		output.WriteString(")\n\n")
	}

	ofp, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer ofp.Close()

	n, err := ofp.Write([]byte(output.String()))
	if err != nil {
		return err
	}
	if n != output.Len() {
		return fmt.Errorf("wrote %d of %d bytes to file %s", n, output.Len(), outFileName)
	}

	return nil
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
