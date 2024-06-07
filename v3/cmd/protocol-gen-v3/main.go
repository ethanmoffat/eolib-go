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

	"github.com/ethanmoffat/eolib-go/v3/internal/codegen"
	eoxml "github.com/ethanmoffat/eolib-go/v3/internal/xml"
)

var inputDir string
var outputDir string

func main() {
	flag.StringVar(&inputDir, "i", "eo-protocol", "The input directory for eo-protocol files.")
	flag.StringVar(&outputDir, "o", "protocol", "The output directory for generated code.")
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

	dirToPackageName := map[string]string{
		"map":        "eomap",
		"net":        "net",
		"net/client": "client",
		"net/server": "server",
		"pub":        "pub",
		"pub/server": "serverpub",
		"":           "protocol",
	}

	var fullSpec eoxml.Protocol  // all XML specs in a single place, for type lookups
	var protocs []eoxml.Protocol // each individual protoc file
	for _, file := range protocolFiles {
		fullInputPath := path.Join(inputDir, "xml", file, "protocol.xml")

		fp, err := os.Open(fullInputPath)
		if err != nil {
			fmt.Printf("error opening file: %v\n", err)
			os.Exit(1)
		}
		defer fp.Close()

		bytes, err := io.ReadAll(fp)
		if err != nil {
			fmt.Printf("error reading file: %v\n", err)
			os.Exit(1)
		}

		var next eoxml.Protocol
		if err := xml.Unmarshal(bytes, &next); err != nil {
			fmt.Printf("error unmarshalling xml: %v\n", err)
			os.Exit(1)
		}

		for i := range next.Enums {
			next.Enums[i].Package = dirToPackageName[strings.Trim(file, string(os.PathSeparator))]
			next.Enums[i].PackagePath = file
		}

		for i := range next.Structs {
			next.Structs[i].Package = dirToPackageName[strings.Trim(file, string(os.PathSeparator))]
			next.Structs[i].PackagePath = file
		}

		for i := range next.Packets {
			next.Packets[i].Package = dirToPackageName[strings.Trim(file, string(os.PathSeparator))]
			next.Packets[i].PackagePath = file
		}

		fullSpec.Enums = append(fullSpec.Enums, next.Enums...)
		fullSpec.Structs = append(fullSpec.Structs, next.Structs...)
		fullSpec.Packets = append(fullSpec.Packets, next.Packets...)

		protocs = append(protocs, next)
	}

	for i, file := range protocolFiles {
		protoc := protocs[i]

		if err := protoc.Validate(); err != nil {
			fmt.Printf("error validating unmarshalled xml: %v\n", err)
			os.Exit(1)
		}

		fullOutputPath := path.Join(outputDir, file)

		fmt.Printf("generating code :: %s\n", file)
		fmt.Printf("      %3d enums\n", len(protoc.Enums))
		if err := codegen.GenerateEnums(fullOutputPath, protoc.Enums); err != nil {
			fmt.Printf("      error generating enums: %v\n", err)
		}

		fmt.Printf("      %3d structs\n", len(protoc.Structs))
		if err := codegen.GenerateStructs(fullOutputPath, protoc.Structs, fullSpec); err != nil {
			fmt.Printf("      error generating structs: %v\n", err)
		}

		fmt.Printf("      %3d packets\n", len(protoc.Packets))
		if err := codegen.GeneratePackets(fullOutputPath, protoc.Packets, fullSpec); err != nil {
			fmt.Printf("      error generating packets: %v\n", err)
		}
	}
}
