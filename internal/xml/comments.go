package xml

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"
)

const commentElementName = "comment"

// node is an XML token in a document tree. Only element nodes have children.
type node struct {
	token    xml.Token
	children []*node
}

// Unmarshal parses XML data into v. XML comments (<!-- -->) are rewritten as <comment> elements before decoding, so
// they are deserialized into the Comment fields of the protocol models.
func Unmarshal(data []byte, v any) error {
	root, err := parseNodes(data)
	if err != nil {
		return err
	}

	for _, child := range root.children {
		if _, ok := child.token.(xml.StartElement); ok {
			rewriteCommentsAsElementsInPlace(child)
		}
	}

	var tokens []xml.Token
	for _, child := range root.children {
		tokens = child.appendTokens(tokens)
	}

	return xml.NewTokenDecoder(&tokenSlice{tokens: tokens}).Decode(v)
}

// parseNodes parses XML data into a tree of nodes. The returned root node has no token; its children are the top-level
// nodes of the document.
func parseNodes(data []byte) (*node, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))

	root := &node{}
	stack := []*node{root}
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			return root, nil
		} else if err != nil {
			return nil, err
		}

		parent := stack[len(stack)-1]
		switch token.(type) {
		case xml.StartElement:
			element := &node{token: xml.CopyToken(token)}
			parent.children = append(parent.children, element)
			stack = append(stack, element)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		default:
			parent.children = append(parent.children, &node{token: xml.CopyToken(token)})
		}
	}
}

// rewriteCommentsAsElementsInPlace removes XML comments from the children of the specified element (recursively) and
// attaches them as <comment> elements. A comment attaches to the next sibling element. A comment with no following
// sibling element attaches to its parent. Existing <comment> text is kept first, followed by trailing comments, then
// preceding comments. Existing <comment> text is normalized the same way as XML comments.
func rewriteCommentsAsElementsInPlace(element *node) {
	var pendingComments []string
	children := element.children[:0]
	for _, child := range element.children {
		switch token := child.token.(type) {
		case xml.Comment:
			pendingComments = append(pendingComments, string(token))
			continue
		case xml.StartElement:
			if token.Name.Local == commentElementName {
				child.children = []*node{{token: xml.CharData(normalizeComment(child.text()))}}
			} else {
				rewriteCommentsAsElementsInPlace(child)
				appendComments(child, pendingComments)
				pendingComments = nil
			}
		}

		children = append(children, child)
	}
	element.children = children

	appendComments(element, pendingComments)
}

func appendComments(element *node, comments []string) {
	var normalized []string
	for _, comment := range comments {
		if comment = normalizeComment(comment); len(comment) > 0 {
			normalized = append(normalized, comment)
		}
	}

	if len(normalized) == 0 {
		return
	}

	commentElement := element.findChild(commentElementName)
	if commentElement == nil {
		commentElement = &node{token: xml.StartElement{Name: xml.Name{Local: commentElementName}}}
		element.children = append([]*node{commentElement}, element.children...)
	} else if existing := commentElement.text(); len(existing) > 0 {
		normalized = append([]string{existing}, normalized...)
	}

	commentElement.children = []*node{{token: xml.CharData(strings.Join(normalized, "\n"))}}
}

// normalizeComment trims each line of a comment and joins the non-empty lines with a space. Comment lines are wrapped
// prose, so a sentence spanning multiple lines is kept together.
func normalizeComment(comment string) string {
	var lines []string
	for _, line := range strings.Split(comment, "\n") {
		if line = strings.TrimSpace(line); len(line) > 0 {
			lines = append(lines, line)
		}
	}
	return strings.Join(lines, " ")
}

func (n *node) findChild(name string) *node {
	for _, child := range n.children {
		if start, ok := child.token.(xml.StartElement); ok && start.Name.Local == name {
			return child
		}
	}
	return nil
}

func (n *node) text() string {
	var sb strings.Builder
	for _, child := range n.children {
		if charData, ok := child.token.(xml.CharData); ok {
			sb.Write(charData)
		}
	}
	return sb.String()
}

func (n *node) appendTokens(tokens []xml.Token) []xml.Token {
	tokens = append(tokens, n.token)
	if start, ok := n.token.(xml.StartElement); ok {
		for _, child := range n.children {
			tokens = child.appendTokens(tokens)
		}
		tokens = append(tokens, start.End())
	}
	return tokens
}

// tokenSlice is an xml.TokenReader over a slice of tokens.
type tokenSlice struct {
	tokens []xml.Token
}

func (t *tokenSlice) Token() (xml.Token, error) {
	if len(t.tokens) == 0 {
		return nil, io.EOF
	}

	token := t.tokens[0]
	t.tokens = t.tokens[1:]
	return token, nil
}
