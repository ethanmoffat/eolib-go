package codegen

import (
	"fmt"
	"path"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/internal/codegen/types"
	"github.com/ethanmoffat/eolib-go/internal/xml"
)

const enumFileName = "enums_generated.go"

func GenerateEnums(outputDir string, enums []xml.ProtocolEnum) error {
	packageName, err := getPackageName(outputDir)
	if err != nil {
		return err
	}

	f := jen.NewFile(packageName)

	for _, e := range enums {
		writeTypeCommentJen(f, e.Name, e.Comment)

		f.Type().Id(e.Name).Int()

		defsList := make([]jen.Code, len(e.Values))
		expected := 0
		for i, v := range e.Values {
			var s *jen.Statement
			if i == 0 {
				s = jen.Id(fmt.Sprintf("%s_%s", types.SanitizeTypeName(e.Name), v.Name)).Qual("", e.Name).Op("=").Iota()

				if v.Value > 0 {
					expected = int(v.Value)
					s.Op("+").Lit(expected)
				}
			} else {
				s = jen.Id(fmt.Sprintf("%s_%s", types.SanitizeTypeName(e.Name), v.Name))
				actual := int(v.Value)
				if expected != actual {
					s.Op("=").Lit(actual)
				}
			}

			writeInlineCommentJen(s, v.Comment)
			expected += 1

			defsList[i] = s
		}
		f.Const().Defs(defsList...)

		caseList := make([]jen.Code, len(e.Values)+1)
		for ndx, v := range e.Values {
			caseList[ndx] = jen.Case(jen.Id(fmt.Sprintf("%s_%s", types.SanitizeTypeName(e.Name), v.Name))).Block(jen.Return(jen.Lit(v.Name), jen.Nil()))
		}
		caseList[len(e.Values)] = jen.Default().Block().Return(
			jen.Lit(""), jen.Qual("fmt", "Errorf").Call(jen.Lit(fmt.Sprintf("could not convert value %%d of type %s to string", e.Name)), jen.Id("e")),
		)

		f.Commentf("String converts a %s value into its string representation", e.Name)
		f.Func().Params(
			// enum receiver
			jen.Id("e").Id(e.Name),
		).Id("String").Params(
		// empty parameter list
		).Params(jen.String(), jen.Error()).Block(
			jen.Switch(jen.Id("e").Block(caseList...)),
		)
	}

	outFileName := path.Join(outputDir, enumFileName)
	return writeToFileJen(f, outFileName)
}
