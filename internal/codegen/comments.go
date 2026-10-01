package codegen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/v3/internal/codegen/types"
	"github.com/ethanmoffat/eolib-go/v3/internal/xml"
)

// sanitizeComment converts each line of a comment to a sentence and joins them with a space.
func sanitizeComment(comment string) string {
	var sentences []string
	for _, line := range strings.Split(comment, "\n") {
		if line = strings.TrimSpace(line); len(line) > 0 {
			if !hasTerminalPunctuation(line) {
				line += "."
			}
			sentences = append(sentences, line)
		}
	}

	return strings.Join(sentences, " ")
}

// hasTerminalPunctuation checks whether a line ends a sentence or introduces the text that follows it,
// ignoring closing quotes and parentheses after the punctuation.
func hasTerminalPunctuation(line string) bool {
	trimmed := strings.TrimRight(line, "\"')")
	return strings.HasSuffix(trimmed, ".") || strings.HasSuffix(trimmed, "!") ||
		strings.HasSuffix(trimmed, "?") || (len(trimmed) == len(line) && strings.HasSuffix(line, ":"))
}

func writeTypeCommentJen(f *jen.File, typeName string, comment string, notes ...string) {
	var paragraphs []string
	for _, c := range append([]string{comment}, notes...) {
		if c = sanitizeComment(c); len(c) > 0 {
			paragraphs = append(paragraphs, c)
		}
	}

	if len(paragraphs) > 0 {
		f.Commentf("// %s :: %s", typeName, strings.Join(paragraphs, "\n//\n// "))
	}
}

func writeInlineCommentJen(c jen.Code, comments ...string) {
	var sentences []string
	for _, comment := range comments {
		if comment = sanitizeComment(comment); len(comment) > 0 {
			sentences = append(sentences, comment)
		}
	}

	if len(sentences) == 0 {
		return
	}

	comment := strings.Join(sentences, " ")
	switch v := c.(type) {
	case *jen.Statement:
		v.Comment(comment)
	case *jen.Group:
		v.Comment(comment)
	}
}

func getInstructionNotes(instructions []xml.ProtocolInstruction) (notes []string) {
	flattened := flattenChunkedInstructions(instructions)
	for i, inst := range flattened {
		if getInstructionMemberName(inst) != "" || inst.XMLName.Local == "switch" || inst.XMLName.Local == "field" || inst.Comment == nil {
			continue
		}

		if comment := strings.TrimSpace(*inst.Comment); len(comment) > 0 {
			notes = append(notes, fmt.Sprintf("%s: %s", describeInstruction(flattened, i), comment))
		}
	}
	return
}

func flattenChunkedInstructions(instructions []xml.ProtocolInstruction) (flattened []xml.ProtocolInstruction) {
	for _, inst := range instructions {
		flattened = append(flattened, inst)
		if inst.XMLName.Local == "chunked" {
			flattened = append(flattened, flattenChunkedInstructions(inst.Chunked)...)
		}
	}
	return
}

func getInstructionMemberName(inst xml.ProtocolInstruction) string {
	switch inst.XMLName.Local {
	case "field", "array":
		if inst.Name != nil {
			return snakeCaseToPascalCase(*inst.Name)
		}
	case "switch":
		if inst.Field != nil {
			return snakeCaseToPascalCase(*inst.Field) + "Data"
		}
	}
	return ""
}

func describeInstruction(instructions []xml.ProtocolInstruction, index int) string {
	inst := instructions[index]

	var subject string
	switch inst.XMLName.Local {
	case "dummy":
		subject = fmt.Sprintf("The dummy %s", *inst.Type)
	case "length":
		return fmt.Sprintf("The %s length field", snakeCaseToPascalCase(*inst.Name))
	case "break":
		subject = "The break byte"
	case "chunked":
		subject = "The chunked section"
	default:
		subject = fmt.Sprintf("The %s", inst.XMLName.Local)
	}

	if position := describeInstructionPosition(instructions, index); len(position) > 0 {
		subject += " " + position
	}

	if inst.Content != nil && inst.XMLName.Local == "dummy" {
		if value := strings.TrimSpace(*inst.Content); len(value) > 0 {
			if typeName, _ := types.GetInstructionTypeName(inst); typeName == "string" || typeName == "encoded_string" {
				value = strconv.Quote(value)
			}
			subject += fmt.Sprintf(" (always %s)", value)
		}
	}

	return subject
}

func describeInstructionPosition(instructions []xml.ProtocolInstruction, index int) string {
	for i := index - 1; i >= 0; i-- {
		if name := getInstructionMemberName(instructions[i]); len(name) > 0 {
			return "after " + name
		}
	}

	for i := index + 1; i < len(instructions); i++ {
		if name := getInstructionMemberName(instructions[i]); len(name) > 0 {
			return "before " + name
		}
	}

	return ""
}

func getEmptyCaseComments(switchInst xml.ProtocolInstruction, fieldName string) []string {
	var comments []string
	valuesByComment := map[string][]string{}
	for _, c := range switchInst.Cases {
		if len(c.Instructions) > 0 || len(strings.TrimSpace(c.Comment)) == 0 {
			continue
		}

		value := c.Value
		if c.Default {
			value = "any other value"
		}

		if _, ok := valuesByComment[c.Comment]; !ok {
			comments = append(comments, c.Comment)
		}
		valuesByComment[c.Comment] = append(valuesByComment[c.Comment], value)
	}

	lines := make([]string, len(comments))
	for i, comment := range comments {
		values := valuesByComment[comment]
		joined := values[len(values)-1]
		if len(values) > 1 {
			joined = fmt.Sprintf("%s or %s", strings.Join(values[:len(values)-1], ", "), joined)
		}
		lines[i] = fmt.Sprintf("When %s is %s: %s", fieldName, joined, comment)
	}
	return lines
}
