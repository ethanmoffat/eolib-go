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
	output.WriteString(packageName + "\n\n")
	output.WriteString("import \"fmt\"\n\n")

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

		output.WriteString(fmt.Sprintf("// String converts a %s value into its string representation\n", e.Name))
		output.WriteString(fmt.Sprintf("func (e %s) String() (string, error) {\n", e.Name))
		output.WriteString("\tswitch e {\n")
		for _, v := range e.Values {
			output.WriteString(fmt.Sprintf("\tcase %s_%s:\n\t\treturn \"%s\", nil\n", sanitizeTypeName(e.Name), v.Name, v.Name))
		}
		output.WriteString(fmt.Sprintf("\tdefault:\n\t\treturn \"\", fmt.Errorf(\"could not convert value %%d of type %s to string\", e)\n", e.Name))
		output.WriteString("\t}\n}\n\n")
	}

	outFileName := path.Join(outputDir, enumFileName)
	return writeToFile(outFileName, output.String())
}
