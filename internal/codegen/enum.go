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

		fmt.Fprintf(&output, "type %s int\n\n", e.Name)
		output.WriteString("const (\n")

		expected := 0
		for i, v := range e.Values {
			if i == 0 {
				fmt.Fprintf(&output, "\t%s_%s %s = iota", sanitizeTypeName(e.Name), v.Name, e.Name)
				if v.Value > 0 {
					fmt.Fprintf(&output, " + %d", v.Value)
					expected = int(v.Value)
				}
			} else {
				fmt.Fprintf(&output, "\t%s_%s", sanitizeTypeName(e.Name), v.Name)
				if expected != int(v.Value) {
					fmt.Fprintf(&output, " = %d", v.Value)
				}
			}

			writeInlineComment(&output, v.Comment)

			output.WriteString("\n")
			expected += 1
		}

		output.WriteString(")\n\n")

		fmt.Fprintf(&output, "// String converts a %s value into its string representation\n", e.Name)
		fmt.Fprintf(&output, "func (e %s) String() (string, error) {\n", e.Name)
		output.WriteString("\tswitch e {\n")
		for _, v := range e.Values {
			fmt.Fprintf(&output, "\tcase %s_%s:\n\t\treturn \"%s\", nil\n", sanitizeTypeName(e.Name), v.Name, v.Name)
		}
		fmt.Fprintf(&output, "\tdefault:\n\t\treturn \"\", fmt.Errorf(\"could not convert value %%d of type %s to string\", e)\n", e.Name)
		output.WriteString("\t}\n}\n\n")
	}

	outFileName := path.Join(outputDir, enumFileName)
	return writeToFile(outFileName, output.String())
}
