package xml

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func unmarshalStruct(t *testing.T, structXml string) ProtocolStruct {
	t.Helper()

	var protocol Protocol
	require.NoError(t, Unmarshal([]byte("<protocol>"+structXml+"</protocol>"), &protocol))
	require.Len(t, protocol.Structs, 1)
	return protocol.Structs[0]
}

func TestUnmarshalCommentBeforeElementAttachesToNextSibling(t *testing.T) {
	s := unmarshalStruct(t, `<struct name="S"><!-- c --><field name="a" type="char"/><field name="b" type="char"/></struct>`)

	require.Len(t, s.Instructions, 2)
	require.NotNil(t, s.Instructions[0].Comment)
	assert.Equal(t, "c", *s.Instructions[0].Comment)
	assert.Nil(t, s.Instructions[1].Comment)
	assert.Empty(t, s.Comment)
}

func TestUnmarshalConsecutiveCommentsJoinedInOrder(t *testing.T) {
	s := unmarshalStruct(t, "<struct name=\"S\"><!-- first\n     line --><!-- second --><field name=\"a\" type=\"char\"/></struct>")

	require.Len(t, s.Instructions, 1)
	require.NotNil(t, s.Instructions[0].Comment)
	assert.Equal(t, "first line\nsecond", *s.Instructions[0].Comment)
}

func TestUnmarshalExplicitCommentElementMergedBeforeXmlComment(t *testing.T) {
	s := unmarshalStruct(t, `<struct name="S"><!-- preceding --><field name="a" type="char"><comment>explicit</comment><!-- trailing --></field></struct>`)

	require.Len(t, s.Instructions, 1)
	require.NotNil(t, s.Instructions[0].Comment)
	assert.Equal(t, "explicit\ntrailing\npreceding", *s.Instructions[0].Comment)
}

func TestUnmarshalExplicitCommentElementWrappedLinesJoined(t *testing.T) {
	s := unmarshalStruct(t, "<struct name=\"S\"><comment>\n  First sentence wrapped\n  across lines.\n\n  Second sentence.\n</comment></struct>")

	assert.Equal(t, "First sentence wrapped across lines. Second sentence.", s.Comment)
}

func TestUnmarshalExplicitCommentElementWrappedLinesJoinedBeforeXmlComment(t *testing.T) {
	s := unmarshalStruct(t, "<struct name=\"S\"><field name=\"a\" type=\"char\"><comment>wrapped\n  explicit</comment><!-- trailing --></field></struct>")

	require.Len(t, s.Instructions, 1)
	require.NotNil(t, s.Instructions[0].Comment)
	assert.Equal(t, "wrapped explicit\ntrailing", *s.Instructions[0].Comment)
}

func TestUnmarshalTrailingCommentAttachesToParent(t *testing.T) {
	var protocol Protocol
	require.NoError(t, Unmarshal([]byte(`<protocol><struct name="S"><field name="a" type="char"/><switch field="a"><case value="1"><!-- no effect --></case></switch></struct></protocol>`), &protocol))

	require.Len(t, protocol.Structs, 1)
	require.Len(t, protocol.Structs[0].Instructions, 2)
	cases := protocol.Structs[0].Instructions[1].Cases
	require.Len(t, cases, 1)
	assert.Equal(t, "no effect", cases[0].Comment)
	assert.Empty(t, cases[0].Instructions)
}

func TestUnmarshalCommentBeforeClosingParentDoesNotAttachAcrossParentBoundary(t *testing.T) {
	var protocol Protocol
	require.NoError(t, Unmarshal([]byte(`<protocol><struct name="S"><field name="a" type="char"/><!-- inner --></struct><struct name="T"><field name="b" type="char"/></struct></protocol>`), &protocol))

	require.Len(t, protocol.Structs, 2)
	assert.Equal(t, "inner", protocol.Structs[0].Comment)
	assert.Empty(t, protocol.Structs[1].Comment)
	assert.Nil(t, protocol.Structs[1].Instructions[0].Comment)
}

func TestUnmarshalCommentBeforeEnumValueAttachesToValue(t *testing.T) {
	var protocol Protocol
	require.NoError(t, Unmarshal([]byte(`<protocol><enum name="E" type="char"><value name="A">0</value><!-- c --><value name="B">1</value></enum></protocol>`), &protocol))

	require.Len(t, protocol.Enums, 1)
	require.Len(t, protocol.Enums[0].Values, 2)
	assert.Empty(t, protocol.Enums[0].Values[0].Comment)
	assert.Equal(t, "c", protocol.Enums[0].Values[1].Comment)
	assert.Equal(t, OrdinalValue(1), protocol.Enums[0].Values[1].Value)
}

func TestUnmarshalFieldWithContentAndCommentContentIsPreserved(t *testing.T) {
	tests := []struct {
		name string
		xml  string
	}{
		{"ExplicitComment", `<struct name="S"><field name="a" type="char"><comment>c</comment>5</field></struct>`},
		{"XmlComment", "<struct name=\"S\"><!-- c -->\n<field name=\"a\" type=\"char\">5</field></struct>"},
		{"XmlCommentInsideField", `<struct name="S"><field name="a" type="char">5<!-- c --></field></struct>`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := unmarshalStruct(t, tt.xml)

			require.Len(t, s.Instructions, 1)
			require.NotNil(t, s.Instructions[0].Content)
			assert.Equal(t, "5", *s.Instructions[0].Content)
			require.NotNil(t, s.Instructions[0].Comment)
			assert.Equal(t, "c", *s.Instructions[0].Comment)
		})
	}
}

func TestUnmarshalCommentsOnSwitchAndChunkedPassValidation(t *testing.T) {
	var protocol Protocol
	require.NoError(t, Unmarshal([]byte(`<protocol><struct name="S"><field name="a" type="char"/><!-- s --><switch field="a"><case value="1"><!-- c --><chunked><field name="b" type="char"/><!-- b --><break/></chunked></case></switch></struct></protocol>`), &protocol))

	assert.NoError(t, protocol.Validate())
}
