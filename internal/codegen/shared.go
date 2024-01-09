package codegen

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
)

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

func writeTypeCommentJen(f *jen.File, typeName string, comment string) {
	if comment = sanitizeComment(comment); len(comment) > 0 {
		f.Commentf("// %s :: %s", typeName, comment)
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

func writeToFileJen(f *jen.File, outFileName string) error {
	return f.Save(outFileName)
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
