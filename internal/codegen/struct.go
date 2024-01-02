package codegen

import (
	"fmt"
	"path"
	"strconv"
	"strings"
	"unicode"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/internal/xml"
)

func GenerateStructs(outputDir string, structs []xml.ProtocolStruct, fullSpec xml.Protocol) error {
	const structFileName = "structs_generated.go"

	var typeNames []string
	for _, s := range structs {
		typeNames = append(typeNames, s.Name)
	}
	return generateStructsShared(outputDir, structFileName, typeNames, fullSpec)
}

func generateStructsShared(outputDir string, outputFileName string, typeNames []string, fullSpec xml.Protocol) error {
	packageName, err := getPackageName(outputDir)
	if err != nil {
		return err
	}

	f := jen.NewFile(packageName)
	for k, v := range packageAliases {
		f.ImportName(v, k)
	}

	if len(typeNames) > 0 {
		for _, typeName := range typeNames {
			if err := writeStruct(f, typeName, fullSpec); err != nil {
				return err
			}
		}
	}

	outFileName := path.Join(outputDir, outputFileName)
	return writeToFileJen(f, outFileName)
}

func writeStruct(f *jen.File, typeName string, fullSpec xml.Protocol) (err error) {
	var si *structInfo
	if si, err = getStructInfo(typeName, fullSpec); err != nil {
		return err
	}

	err = writeStructShared(f, si, fullSpec)
	return
}

func writeStructShared(f *jen.File, si *structInfo, fullSpec xml.Protocol) (err error) {
	structName := snakeCaseToPascalCase(si.Name)
	writeTypeCommentJen(f, structName, si.Comment)

	// write out fields
	var switches []*xml.ProtocolInstruction
	f.Type().Id(structName).StructFunc(func(g *jen.Group) {
		switches = writeStructFields(g, si, fullSpec)
	}).Line()

	for _, sw := range switches {
		if err = writeSwitchStructs(f, *sw, si, fullSpec); err != nil {
			return
		}
	}

	if len(si.Family) > 0 && len(si.Action) > 0 {
		// write out family/action methods
		f.Func().Params(jen.Id("s").Id(structName)).Id("Family").Params().Qual(packageAliases["net"], "PacketFamily").Block(
			jen.Return(jen.Qual(packageAliases["net"], fmt.Sprintf("PacketFamily_%s", si.Family))),
		).Line()
		f.Func().Params(jen.Id("s").Id(structName)).Id("Action").Params().Qual(packageAliases["net"], "PacketAction").Block(
			jen.Return(jen.Qual(packageAliases["net"], fmt.Sprintf("PacketAction_%s", si.Action))),
		).Line()
	}

	// write out serialize method
	f.Func().Params(jen.Id("s").Op("*").Id(structName)).Id("Serialize").Params(jen.Id("writer").Op("*").Qual(packageAliases["data"], "EoWriter")).Params(jen.Id("err").Id("error")).BlockFunc(func(g *jen.Group) {
		g.Id("oldSanitizeStrings").Op(":=").Id("writer").Dot("SanitizeStrings")
		// defer here uses 'Values' instead of 'Block' so the deferred function is single-line style
		g.Defer().Func().Params().Values(jen.Id("writer").Dot("SanitizeStrings").Op("=").Id("oldSanitizeStrings")).Call().Line()

		// err = writeSerializeBody(g, si, fullSpec)

		g.Line().Return()
	}).Line()

	if err != nil {
		return
	}

	// write out deserialize method
	f.Func().Params(jen.Id("s").Op("*").Id(structName)).Id("Deserialize").Params(jen.Id("reader").Op("*").Qual(packageAliases["data"], "EoReader")).Params(jen.Id("err").Id("error")).BlockFunc(func(g *jen.Group) {
		g.Id("oldIsChunked").Op(":=").Id("reader").Dot("IsChunked").Call()
		// defer here uses 'Values' instead of 'Block' so the deferred function is single-line style
		g.Defer().Func().Params().Values(jen.Id("reader").Dot("SetIsChunked").Call(jen.Id("oldIsChunked"))).Call().Line()

		// err = writeDeserializeBody(g, si, fullSpec)

		g.Line().Return()
	}).Line()

	return
}

func writeStructFields(g *jen.Group, si *structInfo, fullSpec xml.Protocol) (switches []*xml.ProtocolInstruction) {
	for i, inst := range si.Instructions {
		var instName string

		if inst.Name != nil {
			instName = snakeCaseToPascalCase(*inst.Name)
		} else if inst.Field != nil {
			instName = snakeCaseToPascalCase(*inst.Field)
		}

		var fieldTypeInfo struct {
			typeName   string
			nextImport *importInfo
			isPointer  bool
		}
		if inst.Type != nil {
			fieldTypeInfo.typeName, fieldTypeInfo.nextImport = eoTypeToGoType(*inst.Type, si.PackageName, fullSpec)
			if inst.Optional != nil && *inst.Optional {
				switch inst.XMLName.Local {
				// these are the only supported values where the type of the rendered field needs to be modified to a pointer
				// arrays also support the "optional" attribute in the spec but default to nil since they are rendered as slices
				case "field":
					fallthrough
				case "length":
					fieldTypeInfo.isPointer = true
				}
			}
		}

		qualifiedTypeName := func(s *jen.Statement) {
			if fieldTypeInfo.isPointer {
				s.Op("*")
			}

			writeComment := func(ss *jen.Statement) {
				if inst.Comment != nil {
					writeInlineCommentJen(ss, *inst.Comment)
				}
			}

			if fieldTypeInfo.nextImport != nil && fieldTypeInfo.nextImport.Package != si.PackageName {
				s.Qual(fieldTypeInfo.nextImport.Path, fieldTypeInfo.typeName).Do(writeComment)
			} else {
				s.Id(fieldTypeInfo.typeName).Do(writeComment)
			}
		}

		switch inst.XMLName.Local {
		case "field":
			if len(instName) > 0 {
				g.Id(instName).Do(qualifiedTypeName)
			}
		case "array":
			g.Id(instName).Index().Do(qualifiedTypeName)
		case "length":
			g.Id(instName).Do(qualifiedTypeName)
		case "switch":
			g.Id(fmt.Sprintf("%sData", instName)).Id(fmt.Sprintf("%s%sData", si.SwitchStructQualifier, instName))
			switches = append(switches, &si.Instructions[i])
		case "chunked":
			nestedStructInfo := &structInfo{
				PackageName:           si.PackageName,
				Instructions:          inst.Chunked,
				SwitchStructQualifier: si.SwitchStructQualifier,
			}
			switches = append(switches, writeStructFields(g, nestedStructInfo, fullSpec)...)
		case "dummy":
		case "break":
			continue // no data to write
		}
	}

	return
}

func writeSwitchStructs(f *jen.File, switchInst xml.ProtocolInstruction, si *structInfo, fullSpec xml.Protocol) (err error) {
	if switchInst.XMLName.Local != "switch" {
		return
	}

	switchInterfaceName := fmt.Sprintf("%sData", snakeCaseToPascalCase(*switchInst.Field))
	if len(si.SwitchStructQualifier) > 0 {
		switchInterfaceName = si.SwitchStructQualifier + switchInterfaceName
	}

	if switchInst.Comment != nil {
		writeTypeCommentJen(f, switchInterfaceName, *switchInst.Comment)
	}
	f.Type().Id(switchInterfaceName).Interface(jen.Qual(packageAliases["protocol"], "EoData")).Line()

	for _, c := range switchInst.Cases {
		if len(c.Instructions) == 0 {
			continue
		}

		var caseName string
		if c.Default {
			caseName = "Default"
		} else {
			caseName = snakeCaseToPascalCase(c.Value)
		}
		caseStructName := fmt.Sprintf("%s%s", switchInterfaceName, caseName)

		nestedStructInfo := &structInfo{
			Name:         caseStructName,
			Comment:      c.Comment,
			Instructions: c.Instructions,
			PackageName:  si.PackageName,
		}
		err = writeStructShared(f, nestedStructInfo, fullSpec)
		if err != nil {
			return
		}
	}

	return
}

// used to track the 'outer' list of instructions when a <chunked> instruction is encountered
// this allows any nested <switch> instructions to search both the instructions in the <chunked> section at the same level
//
//	as well as the outer instructions in the <struct> or <packet> when determining the type of the switch field
var outerInstructionList []xml.ProtocolInstruction

func writeSerializeBody(output *strings.Builder, instructionList []xml.ProtocolInstruction, switchStructQualifier string, packageName string, fullSpec xml.Protocol) (imports []importInfo, err error) {
	for _, instruction := range instructionList {
		instructionType := instruction.XMLName.Local

		if instructionType == "chunked" {
			output.WriteString("\twriter.SanitizeStrings = true\n")
			oldOuterInstructionList := outerInstructionList
			outerInstructionList = instructionList
			defer func() { outerInstructionList = oldOuterInstructionList }()

			if nextImports, err := writeSerializeBody(output, instruction.Chunked, switchStructQualifier, packageName, fullSpec); err != nil {
				return nil, err
			} else {
				imports = append(imports, nextImports...)
			}

			output.WriteString("\twriter.SanitizeStrings = false\n")
			continue
		}

		if instructionType == "break" {
			output.WriteString("\twriter.AddByte(0xFF)\n")
			continue
		}

		instructionName := getInstructionName(instruction)

		if instructionType == "switch" {
			// get type of Value field
			switchFieldSanitizedType := ""
			switchFieldEnumType := ""
			for _, tmpInst := range append(outerInstructionList, instructionList...) {
				if tmpInst.XMLName.Local == "field" && snakeCaseToPascalCase(*tmpInst.Name) == instructionName {
					switchFieldEnumType = *tmpInst.Type
					switchFieldSanitizedType = sanitizeTypeName(switchFieldEnumType)
					break
				}
			}

			output.WriteString(fmt.Sprintf("\tswitch s.%s {\n", instructionName))

			for _, c := range instruction.Cases {
				if len(c.Instructions) == 0 {
					continue
				}

				var switchDataType string
				if c.Default {
					switchDataType = fmt.Sprintf("%sDataDefault", instructionName)
					output.WriteString("\tdefault:\n")
				} else {
					switchDataType = fmt.Sprintf("%sData%s", instructionName, c.Value)
					if _, err := strconv.ParseInt(c.Value, 10, 32); err != nil {
						// case is for an enum value
						if enumTypeInfo, ok := fullSpec.IsEnum(switchFieldEnumType); !ok {
							return nil, fmt.Errorf("type %s in switch is not an enum", switchFieldEnumType)
						} else {
							packageQualifier := ""
							if enumTypeInfo.Package != packageName {
								packageQualifier = enumTypeInfo.Package + "."
								imports = append(imports, importInfo{enumTypeInfo.Package, enumTypeInfo.PackagePath})
							}
							output.WriteString(fmt.Sprintf("\tcase %s%s_%s:\n", packageQualifier, switchFieldSanitizedType, c.Value))
						}
					} else {
						// case is for an integer constant
						output.WriteString(fmt.Sprintf("\tcase %s:\n", c.Value))
					}
				}

				if len(switchDataType) > 0 {
					output.WriteString(fmt.Sprintf("\t\tswitch s.%sData.(type) {\n", instructionName))
					output.WriteString(fmt.Sprintf("\t\tcase *%s%s:\n\t\t", switchStructQualifier, switchDataType))
				}
				output.WriteString(fmt.Sprintf("\t\t\tif err = s.%sData.Serialize(writer); err != nil {\n", instructionName))
				output.WriteString("\t\t\t\treturn\n\t\t\t}\n")

				if len(switchDataType) > 0 {
					output.WriteString(fmt.Sprintf("\t\tdefault:\n\t\t\terr = fmt.Errorf(\"invalid switch struct type for switch value %%d\", s.%s)\n\t\t\treturn\n\t\t}\n", instructionName))
				}
			}

			output.WriteString("\t}\n")
			continue
		}

		typeName, typeSize := getInstructionTypeName(instruction)

		output.WriteString("\t// " + instructionName + " : " + instructionType + " : " + *instruction.Type + "\n")

		delimited := instruction.Delimited != nil && *instruction.Delimited
		trailingDelimiter := instruction.TrailingDelimiter == nil || *instruction.TrailingDelimiter
		if instructionType == "array" {
			var lenExpr string
			if instruction.Length != nil {
				lenExpr = getLengthExpression(*instruction.Length)
			} else {
				lenExpr = fmt.Sprintf("len(s.%s)", instructionName)
			}

			output.WriteString(fmt.Sprintf("\tfor ndx := 0; ndx < %s; ndx++ {\n\t\t", lenExpr))

			if delimited && !trailingDelimiter {
				output.WriteString("\t\tif ndx > 0 {\n\t\t\twriter.AddByte(0xFF)\n\t\t}\n\n")
			}
		}

		switch typeName {
		case "byte":
			writeAddTypeForSerialize(output, instructionName, instruction, "Byte", false)
		case "char":
			writeAddTypeForSerialize(output, instructionName, instruction, "Char", false)
		case "short":
			writeAddTypeForSerialize(output, instructionName, instruction, "Short", false)
		case "three":
			writeAddTypeForSerialize(output, instructionName, instruction, "Three", false)
		case "int":
			writeAddTypeForSerialize(output, instructionName, instruction, "Int", false)
		case "bool":
			if len(typeSize) > 0 {
				typeName = string(unicode.ToUpper(rune(typeSize[0]))) + typeSize[1:]
			} else {
				typeName = "Char"
			}
			output.WriteString(fmt.Sprintf("\tif s.%s {\n", instructionName))
			output.WriteString(fmt.Sprintf("\t\terr = writer.Add%s(1)\n\t} else {\n\t\terr = writer.Add%s(0)\n\t}\n", typeName, typeName))
			output.WriteString("\tif err != nil {\n\t\treturn\n\t}\n\n")
		case "blob":
			writeAddTypeForSerialize(output, instructionName, instruction, "Bytes", false)
		case "string":
			if instruction.Length != nil && instructionType == "field" {
				if instruction.Padded != nil && *instruction.Padded {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "PaddedString")
				} else {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "FixedString")
				}
			} else {
				writeAddStringTypeForSerialize(output, instructionName, instruction, "String")
			}
		case "encoded_string":
			if instruction.Length != nil && instructionType == "field" {
				if instruction.Padded != nil && *instruction.Padded {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "PaddedEncodedString")
				} else {
					writeAddStringTypeForSerialize(output, instructionName, instruction, "FixedEncodedString")
				}
			} else {
				writeAddStringTypeForSerialize(output, instructionName, instruction, "EncodedString")
			}
		default:
			if _, ok := fullSpec.IsStruct(typeName); ok {
				if instructionType == "array" {
					instructionName = instructionName + "[ndx]"
				}
				output.WriteString(fmt.Sprintf("\tif err = s.%s.Serialize(writer); err != nil {\n\t\treturn\n\t}\n", instructionName))
			} else if e, ok := fullSpec.IsEnum(typeName); ok {
				switch e.Type {
				case "byte":
					fallthrough
				case "char":
					fallthrough
				case "short":
					fallthrough
				case "three":
					fallthrough
				case "int":
					writeAddTypeForSerialize(output, instructionName, instruction, string(unicode.ToUpper(rune(e.Type[0])))+e.Type[1:], true)
				}
			} else {
				panic("Unable to find type '" + typeName + "' when writing serialization function")
			}
		}

		if instructionType == "array" {
			if delimited && trailingDelimiter {
				output.WriteString("\t\twriter.AddByte(0xFF)\n")
			}
			output.WriteString("\t}\n\n")
		}
	}

	return
}

// flag that determines whether a chunked section is active or not
// this is used to determine if the next chunk should be selected in array delimiters and break bytes
var isChunked bool

func writeDeserializeBody(output *strings.Builder, instructionList []xml.ProtocolInstruction, switchStructQualifier string, packageName string, fullSpec xml.Protocol) (imports []importInfo, err error) {
	for _, instruction := range instructionList {
		instructionType := instruction.XMLName.Local

		if instructionType == "chunked" {
			output.WriteString("\treader.SetIsChunked(true)\n")
			oldChunked := isChunked
			isChunked = true
			oldOuterInstructionList := outerInstructionList
			outerInstructionList = instructionList
			defer func() { isChunked = oldChunked; outerInstructionList = oldOuterInstructionList }()

			nextImports, err := writeDeserializeBody(output, instruction.Chunked, switchStructQualifier, packageName, fullSpec)
			if err != nil {
				return nil, err
			}
			imports = append(imports, nextImports...)

			output.WriteString("\treader.SetIsChunked(false)\n\n")
			continue
		}

		if instructionType == "break" {
			if isChunked {
				output.WriteString("\tif err = reader.NextChunk(); err != nil {\n\t\treturn\n\t}\n")
			} else {
				output.WriteString("\tif breakByte := reader.GetByte(); breakByte != 0xFF {\n")
				output.WriteString("\t\treturn fmt.Errorf(\"missing expected break byte\")\n")
				output.WriteString("\t}\n")
			}
			continue
		}

		instructionName := getInstructionName(instruction)

		if instructionType == "switch" {
			// get type of Value field
			switchFieldSanitizedType := ""
			switchFieldEnumType := ""
			for _, tmpInst := range append(outerInstructionList, instructionList...) {
				if tmpInst.XMLName.Local == "field" && snakeCaseToPascalCase(*tmpInst.Name) == instructionName {
					switchFieldEnumType = *tmpInst.Type
					switchFieldSanitizedType = sanitizeTypeName(switchFieldEnumType)
					break
				}
			}

			output.WriteString(fmt.Sprintf("\tswitch s.%s {\n", instructionName))

			for _, c := range instruction.Cases {
				if len(c.Instructions) == 0 {
					continue
				}

				var switchDataType string
				if c.Default {
					switchDataType = fmt.Sprintf("%sDataDefault", instructionName)
					output.WriteString("\tdefault:\n")
				} else {
					switchDataType = fmt.Sprintf("%sData%s", instructionName, c.Value)
					if _, err := strconv.ParseInt(c.Value, 10, 32); err != nil {
						// case is for an enum value
						if enumTypeInfo, ok := fullSpec.IsEnum(switchFieldEnumType); !ok {
							return nil, fmt.Errorf("type %s in switch is not an enum", switchFieldEnumType)
						} else {
							packageQualifier := ""
							if enumTypeInfo.Package != packageName {
								packageQualifier = enumTypeInfo.Package + "."
								imports = append(imports, importInfo{enumTypeInfo.Package, enumTypeInfo.PackagePath})
							}
							output.WriteString(fmt.Sprintf("\tcase %s%s_%s:\n", packageQualifier, switchFieldSanitizedType, c.Value))
						}
					} else {
						// case is for an integer constant
						output.WriteString(fmt.Sprintf("\tcase %s:\n", c.Value))
					}
				}

				output.WriteString(fmt.Sprintf("\t\ts.%sData = &%s%s{}\n", instructionName, switchStructQualifier, switchDataType))
				output.WriteString(fmt.Sprintf("\t\tif err = s.%sData.Deserialize(reader); err != nil {\n", instructionName))
				output.WriteString("\t\t\treturn\n\t\t}\n")
			}

			output.WriteString("\t}\n")

			continue
		}

		typeName, typeSize := getInstructionTypeName(instruction)

		output.WriteString("\t// " + instructionName + " : " + instructionType + " : " + *instruction.Type + "\n")

		var lenExpr string
		if instructionType == "array" {
			if instruction.Length != nil {
				lenExpr = "ndx < " + getLengthExpression(*instruction.Length)
			} else if (instruction.Delimited == nil || !*instruction.Delimited) && isChunked {
				rawLen, err := calculateTypeSize(typeName, fullSpec)
				if err != nil {
					lenExpr = "reader.Remaining() > 0"
				} else {
					lenExpr = "ndx < reader.Remaining() / " + strconv.Itoa(rawLen)
				}
			} else {
				lenExpr = "reader.Remaining() > 0"
			}

			output.WriteString(fmt.Sprintf("\tfor ndx := 0; %s; ndx++ {\n\t\t", lenExpr))
		}

		switch typeName {
		case "byte":
			castType := "int"
			writeGetTypeForDeserialize(output, instructionName, instruction, "Byte", &castType)
		case "char":
			writeGetTypeForDeserialize(output, instructionName, instruction, "Char", nil)
		case "short":
			writeGetTypeForDeserialize(output, instructionName, instruction, "Short", nil)
		case "three":
			writeGetTypeForDeserialize(output, instructionName, instruction, "Three", nil)
		case "int":
			writeGetTypeForDeserialize(output, instructionName, instruction, "Int", nil)
		case "bool":
			if len(typeSize) > 0 {
				typeName = string(unicode.ToUpper(rune(typeSize[0]))) + typeSize[1:]
			} else {
				typeName = "Char"
			}
			output.WriteString(fmt.Sprintf("\tif boolVal := reader.Get%s(); boolVal > 0 {\n", typeName))
			output.WriteString(fmt.Sprintf("\t\ts.%s = true\n\t} else {\n\t\ts.%s = false\n\t}\n", instructionName, instructionName))
		case "blob":
			writeGetTypeForDeserialize(output, instructionName, instruction, "Bytes", nil)
		case "string":
			if instruction.Length != nil && instructionType == "field" {
				if instruction.Padded != nil && *instruction.Padded {
					writeGetStringTypeForDeserialize(output, instructionName, instruction, "PaddedString")
				} else {
					writeGetStringTypeForDeserialize(output, instructionName, instruction, "FixedString")
				}
			} else {
				writeGetStringTypeForDeserialize(output, instructionName, instruction, "String")
			}
		case "encoded_string":
			if instruction.Length != nil && instructionType == "field" {
				if instruction.Padded != nil && *instruction.Padded {
					writeGetStringTypeForDeserialize(output, instructionName, instruction, "PaddedEncodedString")
				} else {
					writeGetStringTypeForDeserialize(output, instructionName, instruction, "FixedEncodedString")
				}
			} else {
				writeGetStringTypeForDeserialize(output, instructionName, instruction, "EncodedString")
			}
		default:
			if structInfo, ok := fullSpec.IsStruct(typeName); ok {
				if instructionType == "array" {
					if packageName != structInfo.Package {
						typeName = structInfo.Package + "." + typeName
						imports = append(imports, importInfo{structInfo.Package, structInfo.PackagePath})
					}

					output.WriteString(fmt.Sprintf("\ts.%s = append(s.%s, %s{})\n", instructionName, instructionName, typeName))
					instructionName = instructionName + "[ndx]"
				}
				output.WriteString(fmt.Sprintf("\tif err = s.%s.Deserialize(reader); err != nil {\n\t\treturn\n\t}\n", instructionName))
			} else if e, ok := fullSpec.IsEnum(typeName); ok {
				switch e.Type {
				case "byte":
					fallthrough
				case "char":
					fallthrough
				case "short":
					fallthrough
				case "three":
					fallthrough
				case "int":
					if e.Package != packageName {
						typeName = fmt.Sprintf("%s.%s", e.Package, typeName)
					}
					writeGetTypeForDeserialize(output, instructionName, instruction, string(unicode.ToUpper(rune(e.Type[0])))+e.Type[1:], &typeName)
				}
				imports = append(imports, importInfo{e.Package, e.PackagePath})
			} else {
				panic("Unable to find type '" + typeName + "' when writing serialization function")
			}
		}

		delimited := instruction.Delimited != nil && *instruction.Delimited
		trailingDelimiter := instruction.TrailingDelimiter == nil || *instruction.TrailingDelimiter
		if instructionType == "array" {
			if delimited && isChunked {
				if !trailingDelimiter {
					if instruction.Length == nil {
						return nil, fmt.Errorf("delimited arrays with trailing-delimiter=false must have a length (array %s)", instructionName)
					}
					output.WriteString(fmt.Sprintf("\t\tif ndx + 1 < %s {\n", getLengthExpression(*instruction.Length)))
				}
				output.WriteString("\t\tif err = reader.NextChunk(); err != nil {\n\t\t\treturn\n\t\t}\n")
				if !trailingDelimiter {
					output.WriteString("\t\t}\n")
				}
			}
			output.WriteString("\t}\n\n")
		}
	}

	return
}

func getInstructionName(inst xml.ProtocolInstruction) (instName string) {
	if inst.Name != nil {
		instName = snakeCaseToPascalCase(*inst.Name)
	} else if inst.Field != nil {
		instName = snakeCaseToPascalCase(*inst.Field)
	}
	return
}

func writeAddTypeForSerialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string, needsCastToInt bool) {
	optional := instruction.Optional != nil && *instruction.Optional

	if len(instructionName) == 0 && instruction.Content != nil {
		instructionName = *instruction.Content
	} else {
		instructionName = "s." + instructionName
	}

	if instruction.XMLName.Local == "array" {
		instructionName = instructionName + "[ndx]"

		// optional arrays that are unset will be nil.
		// The length expression in the loop checks the length of the nil slice, which evaluates to 0.
		// This means that arrays do not need additional dereferencing when optional.
		optional = false
	} else if optional {
		output.WriteString(fmt.Sprintf("\tif %s != nil {\n", instructionName))
		instructionName = "*" + instructionName
	}

	if needsCastToInt {
		instructionName = "int(" + instructionName + ")"
	}

	output.WriteString(fmt.Sprintf("\t\tif err = writer.Add%s(%s); err != nil {\n\t\t\treturn\n\t\t}\n", methodType, instructionName))

	if optional {
		output.WriteString("\t}\n")
	}
}

func writeGetTypeForDeserialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string, castType *string) {
	optional := instruction.Optional != nil && *instruction.Optional

	lengthExpr := ""
	if instruction.XMLName.Local != "array" {
		if instruction.Length != nil {
			lengthExpr = getLengthExpression(*instruction.Length)
		} else if methodType == "Bytes" {
			lengthExpr = "reader.Remaining()"
		}
	} else {
		// optional arrays that are unset will be nil.
		// The length expression in the loop checks the length of the nil slice, which evaluates to 0.
		// This means that arrays do not need additional dereferencing when optional.
		optional = false
	}

	if optional {
		output.WriteString("\tif reader.Remaining() > 0 {\n")
	}

	if len(instructionName) == 0 && instruction.Content != nil {
		output.WriteString(fmt.Sprintf("\treader.Get%s(%s)\n", methodType, lengthExpr))
	} else {
		if instruction.XMLName.Local == "array" {
			output.WriteString(fmt.Sprintf("\t\ts.%s = append(s.%s, 0)\n", instructionName, instructionName))
			instructionName = instructionName + "[ndx]"
		}

		if castType != nil {
			if optional {
				output.WriteString(fmt.Sprintf("\t\ts.%s = new(%s)\n\t\t*s.", instructionName, *castType))
			} else {
				output.WriteString("\t\ts.")
			}

			output.WriteString(fmt.Sprintf("%s = %s(reader.Get%s(%s))\n", instructionName, *castType, methodType, lengthExpr))
		} else {
			if optional {
				output.WriteString(fmt.Sprintf("\t\ts.%s = new(int)\n\t\t*s.", instructionName))
			} else {
				output.WriteString("\t\ts.")
			}

			output.WriteString(fmt.Sprintf("%s = reader.Get%s(%s)\n", instructionName, methodType, lengthExpr))
		}
	}

	if optional {
		output.WriteString("\t}\n")
	}
}

func writeAddStringTypeForSerialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string) {
	optional := instruction.Optional != nil && *instruction.Optional

	if len(instructionName) == 0 && instruction.Content != nil {
		instructionName = `"` + *instruction.Content + `"`
	} else {
		instructionName = "s." + instructionName
	}

	if instruction.XMLName.Local == "array" {
		instructionName = instructionName + "[ndx]"
		optional = false
	} else if instruction.Length != nil {
		instructionName = instructionName + ", " + getLengthExpression(*instruction.Length)
	}

	if optional {
		output.WriteString(fmt.Sprintf("\tif %s != nil {\n", instructionName))
		instructionName = "*" + instructionName
	}

	output.WriteString(fmt.Sprintf("\t\tif err = writer.Add%s(%s); err != nil {\n\t\t\treturn\n\t\t}\n", methodType, instructionName))

	if optional {
		output.WriteString("\t}\n")
	}
}

func writeGetStringTypeForDeserialize(output *strings.Builder, instructionName string, instruction xml.ProtocolInstruction, methodType string) {
	optional := instruction.Optional != nil && *instruction.Optional

	lengthExpr := ""
	if instruction.XMLName.Local != "array" {
		if instruction.Length != nil {
			lengthExpr = getLengthExpression(*instruction.Length)
		}
	} else {
		optional = false
	}

	if optional {
		output.WriteString("\tif reader.Remaining() > 0 {\n")
	}

	if len(instructionName) == 0 && instruction.Content != nil {
		output.WriteString(fmt.Sprintf("\tif _, err = reader.Get%s(%s); err != nil {\n\t\treturn\n\t}\n", methodType, lengthExpr))
	} else {
		if instruction.XMLName.Local == "array" {
			output.WriteString(fmt.Sprintf("\t\ts.%s = append(s.%s, \"\")\n", instructionName, instructionName))
			instructionName = instructionName + "[ndx]"
		}

		if optional {
			output.WriteString(fmt.Sprintf("\t\ts.%s = new(string)\n\t\tif *s.", instructionName))
		} else {
			output.WriteString("\t\tif s.")
		}

		output.WriteString(fmt.Sprintf("%s, err = reader.Get%s(%s); err != nil {\n\t\treturn\n\t}\n\n", instructionName, methodType, lengthExpr))
	}

	if optional {
		output.WriteString("\t}\n")
	}
}

func getLengthExpression(instLength string) string {
	if _, err := strconv.ParseInt(instLength, 10, 32); err == nil {
		// string length is a numeric constant
		return instLength
	} else {
		// string length is a reference to another field
		return "s." + snakeCaseToPascalCase(instLength)
	}
}
