package xml

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func validateProtocolXml(t *testing.T, structXml string) error {
	t.Helper()

	var protocol Protocol
	require.NoError(t, xml.Unmarshal([]byte("<protocol>"+structXml+"</protocol>"), &protocol))
	return protocol.Validate()
}

func TestValidateValidStructs(t *testing.T) {
	tests := []struct {
		name string
		xml  string
	}{
		{"UnsizedArrayLast", `<struct name="S"><field name="a" type="char"/><array name="b" type="char"/></struct>`},
		{"UnsizedArrayBeforeBreak", `<struct name="S"><chunked><array name="a" type="char"/><break/><field name="b" type="char"/></chunked></struct>`},
		{"OptionalFieldsLast", `<struct name="S"><field name="a" type="char"/><field name="b" type="char" optional="true"/><field name="c" type="char" optional="true"/></struct>`},
		{"OptionalResetByBreak", `<struct name="S"><chunked><field name="a" type="char" optional="true"/><break/><field name="b" type="char"/></chunked></struct>`},
		{"DelimitedInChunked", `<struct name="S"><chunked><array name="a" type="string" delimited="true"/></chunked></struct>`},
		{"DummyLast", `<struct name="S"><field name="a" type="char"/><dummy type="byte">0</dummy></struct>`},
		{"SwitchWithDefaultLast", `<struct name="S"><field name="a" type="char"/><switch field="a"><case value="1"><field name="b" type="char"/></case><case default="true"/></switch></struct>`},
		{"SwitchNestedInCase", `<struct name="S"><field name="a" type="char"/><switch field="a"><case value="1"><field name="b" type="char"/><switch field="b"><case value="1"/></switch></case><case value="2"><field name="c" type="char"/><switch field="c"><case value="1"/></switch></case></switch></struct>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, validateProtocolXml(t, tt.xml))
		})
	}
}

func TestValidateInvalidStructs(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		expected string
	}{
		{"UnsizedArrayNotLast", `<struct name="S"><array name="a" type="char"/><field name="b" type="char"/></struct>`, "non-delimited arrays without a length must be the final element"},
		{"RequiredAfterOptional", `<struct name="S"><field name="a" type="char" optional="true"/><field name="b" type="char"/></struct>`, "optional fields may not be followed by non-optional fields (<field> b)"},
		{"DelimitedOutsideChunked", `<struct name="S"><array name="a" type="string" delimited="true"/></struct>`, "delimited arrays are only allowed in chunked sections (<array> a)"},
		{"DummyNotLast", `<struct name="S"><dummy type="byte">0</dummy><field name="a" type="char"/></struct>`, "<dummy> elements must not be followed by any other elements"},
		{"DummyInCaseNotLast", `<struct name="S"><field name="a" type="char"/><switch field="a"><case value="1"><dummy type="byte">0</dummy></case></switch><field name="b" type="char"/></struct>`, "<dummy> elements must not be followed by any other elements"},
		{"LengthReferencedTwice", `<struct name="S"><length name="len" type="char"/><field name="a" type="string" length="len"/><field name="b" type="string" length="len"/></struct>`, "length field len must not be referenced by multiple fields (a, b)"},
		{"SwitchUnknownField", `<struct name="S"><switch field="a"><case value="1"/></switch></struct>`, "switch must reference a preceding field (<switch> a)"},
		{"SwitchDefaultNotLast", `<struct name="S"><field name="a" type="char"/><switch field="a"><case default="true"/><case value="1"/></switch></struct>`, "only the last case in a switch on a can be the default case"},
		{"SwitchDuplicateCase", `<struct name="S"><field name="a" type="char"/><switch field="a"><case value="1"/><case value="1"/></switch></struct>`, "duplicate case value 1 in switch on a"},
		{"MultipleSwitchesInTypeScope", `<struct name="S"><field name="a" type="char"/><switch field="a"><case value="1"/></switch><field name="b" type="char"/><switch field="b"><case value="1"/></switch></struct>`, "switch factories don't support multiple switches in one scope (switches on a and b)"},
		{"MultipleSwitchesAcrossChunks", `<struct name="S"><chunked><field name="a" type="char"/><switch field="a"><case value="1"/></switch><break/><field name="b" type="char"/><switch field="b"><case value="1"/></switch></chunked></struct>`, "switch factories don't support multiple switches in one scope (switches on a and b)"},
		{"MultipleSwitchesInCaseScope", `<struct name="S"><field name="a" type="char"/><switch field="a"><case value="1"><field name="x" type="char"/><switch field="x"><case value="1"/></switch><field name="y" type="char"/><switch field="y"><case value="1"/></switch></case></switch></struct>`, "switch factories don't support multiple switches in one scope (switches on x and y)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProtocolXml(t, tt.xml)
			require.Error(t, err)
			assert.Contains(t, err.Error(), "S: validation error: "+tt.expected)
		})
	}
}
