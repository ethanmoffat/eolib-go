package codegen

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
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

func writeToFile(outFileName string, output strings.Builder) error {
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
