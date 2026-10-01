package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeComment(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Empty", "", ""},
		{"WhitespaceOnly", " \n\t ", ""},
		{"AddsPeriod", "A comment", "A comment."},
		{"KeepsPeriod", "A comment.", "A comment."},
		{"KeepsExclamation", "A comment!", "A comment!"},
		{"KeepsQuestion", "A comment?", "A comment?"},
		{"KeepsPunctuationBeforeQuote", `"A quote."`, `"A quote."`},
		{"KeepsPunctuationBeforeParenthesis", "(A note.)", "(A note.)"},
		{"AddsPeriodAfterQuote", `A "quote"`, `A "quote".`},
		{"AddsPeriodAfterParenthesis", "(A note)", "(A note)."},
		{"KeepsColon", "One of the following:", "One of the following:"},
		{"AddsPeriodAfterQuotedColon", `"Label:"`, `"Label:".`},
		{"TrimsLines", "  First \n Second  ", "First. Second."},
		{"SkipsEmptyLines", "First\n\nSecond", "First. Second."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, sanitizeComment(tt.input))
		})
	}
}
