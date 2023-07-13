package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/ethanmoffat/eolib-go/internal/codegen"
	eoxml "github.com/ethanmoffat/eolib-go/internal/xml"
)

var inputDir string
var outputDir string

func main() {
	flag.StringVar(&inputDir, "i", "eo-protocol", "The input directory for eo-protocol files.")
	flag.StringVar(&outputDir, "o", "pkg/eolib/protocol", "The output directory for generated code.")
	flag.Parse()

	if _, err := os.Stat(inputDir); err != nil {
		fmt.Printf("error: input directory %s does not exist\n", inputDir)
		os.Exit(1)
	}

	if _, err := os.Stat(outputDir); err != nil {
		fmt.Printf("error: output directory %s does not exist\n", outputDir)
		os.Exit(1)
	}

	fmt.Printf("Using parameters:\n  inputDir:  %s\n  outputDir: %s\n", inputDir, outputDir)

	protocolFiles := []string{}
	filepath.WalkDir(path.Join(inputDir, "xml"), func(currentPath string, d fs.DirEntry, err error) error {
		if path.Ext(currentPath) == ".xml" {
			relativeDir := strings.ReplaceAll(currentPath, path.Join(inputDir, "xml"), "")
			protocolFiles = append(protocolFiles, strings.ReplaceAll(relativeDir, "/protocol.xml", ""))
		}
		return nil
	})

	for _, file := range protocolFiles {
		fullInputPath := path.Join(inputDir, "xml", file, "protocol.xml")
		fullOutputPath := path.Join(outputDir, file)

		fp, err := os.Open(fullInputPath)
		if err != nil {
			fmt.Printf("error opening file: %v", err)
			os.Exit(1)
		}
		defer fp.Close()

		bytes, err := io.ReadAll(fp)
		if err != nil {
			fmt.Printf("error reading file: %v", err)
			os.Exit(1)
		}

		var protoc eoxml.Protocol
		if err := xml.Unmarshal(bytes, &protoc); err != nil {
			fmt.Printf("error unmarshalling xml: %v", err)
			os.Exit(1)
		}

		if err := protoc.Validate(); err != nil {
			fmt.Printf("error validating unmarshalled xml: %v", err)
			os.Exit(1)
		}

		fmt.Printf("generating code :: %s\n", file)
		fmt.Printf("      %3d enums\n", len(protoc.Enums))
		if err := codegen.GenerateEnums(fullOutputPath, protoc.Enums); err != nil {
			fmt.Printf("      error generating enums: %v", err)
		}

		fmt.Printf("      %3d structs\n", len(protoc.Structs))
		// todo: structs

		fmt.Printf("      %3d packets\n", len(protoc.Packets))
		// todo: packets
	}
}
