package codegen

import (
	"fmt"
	"path"
	"strings"

	"github.com/ethanmoffat/eolib-go/internal/xml"
)

const enumFileName = "enums_generated.go"

func GenerateEnums(outputDir string, enums []xml.ProtocolEnum) error {
	packageName, err := getPackageName(outputDir)
	if err != nil {
		return err
	}

	output := strings.Builder{}
	output.WriteString(packageName + "\n")
	output.WriteString("\n")

	for _, e := range enums {
		writeTypeComment(&output, e.Name, e.Comment)

		output.WriteString(fmt.Sprintf("type %s int\n\n", e.Name))
		output.WriteString("const (\n")

		for i, v := range e.Values {
			if i == 0 {
				output.WriteString(fmt.Sprintf("\t%s_%s %s = iota", sanitizeTypeName(e.Name), v.Name, e.Name))
			} else {
				output.WriteString(fmt.Sprintf("\t%s_%s", sanitizeTypeName(e.Name), v.Name))
			}

			writeInlineComment(&output, v.Comment)

			output.WriteString("\n")
		}

		output.WriteString(")\n\n")
	}

	outFileName := path.Join(outputDir, enumFileName)
	return writeToFile(outFileName, output.String())
}
