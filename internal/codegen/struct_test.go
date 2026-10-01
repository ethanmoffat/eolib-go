package codegen

import (
	"bytes"
	"strings"
	"testing"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/v3/internal/xml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testEnumXml = `<enum name="Code" type="char"><value name="A">1</value><value name="B">2</value><value name="C">3</value></enum>`

func generateStruct(t *testing.T, protocolXml string) string {
	t.Helper()

	var protocol xml.Protocol
	require.NoError(t, xml.Unmarshal([]byte("<protocol>"+protocolXml+"</protocol>"), &protocol))
	require.NoError(t, protocol.Validate())

	f := jen.NewFile("test")
	require.NoError(t, writeStruct(f, "S", protocol))

	var buf bytes.Buffer
	require.NoError(t, f.Render(&buf))
	return buf.String()
}

func TestStructUnnamedInstructionCommentsAddedToTypeDoc(t *testing.T) {
	code := generateStruct(t, `<struct name="S">
		<comment>Summary</comment>
		<!-- Unnamed field -->
		<field type="char">1</field>
		<field name="a" type="char"/>
		<!-- Chunked -->
		<chunked>
			<field name="b" type="char"/>
			<!-- Break -->
			<break/>
			<!-- Dummy -->
			<dummy type="char">0</dummy>
		</chunked>
	</struct>`)

	assert.Contains(t, code, "// S :: Summary.\n//\n// The chunked section after A: Chunked.\n//\n// The break byte after B: Break.\n//\n// The dummy char after B (always 0): Dummy.\ntype S struct")
	assert.NotContains(t, code, "Unnamed field", "comments on unnamed fields document values that aren't visible to consumers")
	assert.Equal(t, 1, strings.Count(code, "Dummy."), "unnamed instruction comments should only be in the type doc")
}

func TestStructUnnamedInstructionCommentWithoutSummaryAddedToTypeDoc(t *testing.T) {
	code := generateStruct(t, `<struct name="S"><field name="a" type="char"/><!-- Dummy --><dummy type="char">0</dummy></struct>`)

	assert.Contains(t, code, "// S :: The dummy char after A (always 0): Dummy.\ntype S struct")
}

func TestStructUnnamedInstructionCommentDescribesInstruction(t *testing.T) {
	code := generateStruct(t, `<struct name="S">
		<!-- Length -->
		<length name="items_count" type="char"/>
		<array name="items" type="char" length="items_count"/>
		<!-- String -->
		<dummy type="string">ABC</dummy>
	</struct>`)

	assert.Contains(t, code, "// S :: The ItemsCount length field: Length.\n//\n// The dummy string after Items (always \"ABC\"): String.\ntype S struct")
}

func TestStructUnnamedFieldCommentIsSkipped(t *testing.T) {
	code := generateStruct(t, `<struct name="S"><field name="a" type="char"/><!-- Unused --><field type="char">0</field></struct>`)

	assert.NotContains(t, code, "Unused")
	assert.NotContains(t, code, "// S ::")
}

func TestStructUnnamedInstructionCommentAfterSwitchDescribesPositionAfterSwitchData(t *testing.T) {
	code := generateStruct(t, `<struct name="S">
		<field name="kind" type="char"/>
		<switch field="kind"><case value="1"><field name="x" type="char"/></case></switch>
		<!-- Dummy -->
		<dummy type="char">0</dummy>
	</struct>`)

	assert.Contains(t, code, "// S :: The dummy char after KindData (always 0): Dummy.\ntype S struct")
}

func TestStructUnnamedInstructionCommentWithoutFieldsHasNoPosition(t *testing.T) {
	code := generateStruct(t, `<struct name="S"><!-- Dummy --><dummy type="short">1</dummy></struct>`)

	assert.Contains(t, code, "// S :: The dummy short (always 1): Dummy.\ntype S struct")
}

func TestStructEmptyCaseCommentDocumentsSwitchField(t *testing.T) {
	code := generateStruct(t, testEnumXml+`<struct name="S">
		<field name="code" type="Code"/>
		<switch field="code">
			<case value="1"><!-- No effect --></case>
			<case value="B"><field name="b" type="char"/></case>
			<case value="3"><!-- No effect --></case>
			<case default="true"><!-- Unknown --></case>
		</switch>
	</struct>`)

	assert.Contains(t, code, "CodeData CodeData // When Code is 1 or 3: No effect. When Code is any other value: Unknown.\n")
	assert.Contains(t, code, "type CodeDataB struct")
	assert.NotContains(t, code, "CodeData1")
	assert.NotContains(t, code, "CodeData3")
	assert.NotContains(t, code, "CodeDataDefault")
	assert.NotContains(t, code, "case 1:")
	assert.NotContains(t, code, "case 3:")
	assert.Equal(t, 1, strings.Count(code, "No effect"), "empty case comments should only be on the switch field")
}

func TestStructSwitchWithoutEmptyCaseCommentsHasNoFieldDoc(t *testing.T) {
	code := generateStruct(t, testEnumXml+`<struct name="S">
		<field name="code" type="Code"/>
		<switch field="code">
			<case value="1"/>
			<case value="B"><!-- case comment --><field name="b" type="char"/></case>
		</switch>
	</struct>`)

	assert.Contains(t, code, "CodeData CodeData\n")
}

func TestStructNamedHardcodedFieldGeneratesDefaultConst(t *testing.T) {
	code := generateStruct(t, `<struct name="S">
		<!-- The server verifies this -->
		<field name="version" type="char">112</field>
		<field name="request" type="string">NEW</field>
	</struct>`)

	assert.Contains(t, code, "\tVersion int    // The server verifies this. A zero value is serialized as [S_DefaultVersion] unless this object was deserialized.\n")
	assert.Contains(t, code, "\tRequest string // A zero value is serialized as [S_DefaultRequest] unless this object was deserialized.\n")
	assert.Contains(t, code, "// S_DefaultVersion :: The server verifies this.\nconst S_DefaultVersion = 112\n")
	assert.Contains(t, code, "// S_DefaultRequest :: The default value of the Request field.\nconst S_DefaultRequest = \"NEW\"\n")
	assert.NotContains(t, code, "always serialized")
}

func TestStructNamedHardcodedFieldZeroValueSerializesDefaultUnlessDeserialized(t *testing.T) {
	code := generateStruct(t, `<struct name="S">
		<field name="version" type="char">112</field>
		<chunked><field name="request" type="string">NEW</field></chunked>
	</struct>`)

	assert.Contains(t, code, "version := s.Version\n\t// byteSize is only non-zero when the object was deserialized\n\tif version == 0 && s.byteSize == 0 {\n\t\tversion = S_DefaultVersion\n\t}\n\tif err = writer.AddChar(version); err != nil {")
	assert.Contains(t, code, "if request == \"\" && s.byteSize == 0 {\n\t\trequest = S_DefaultRequest\n\t}")
	assert.Contains(t, code, "s.Version = reader.GetChar()")
	assert.Contains(t, code, "if s.Request, err = reader.GetString(); err != nil {")
}

func TestStructNamedHardcodedFieldWithZeroDefaultSerializesField(t *testing.T) {
	code := generateStruct(t, `<struct name="S"><field name="version" type="char">0</field></struct>`)

	assert.Contains(t, code, "\tVersion int\n")
	assert.Contains(t, code, "const S_DefaultVersion = 0\n")
	assert.Contains(t, code, "if err = writer.AddChar(s.Version); err != nil {")
	assert.NotContains(t, code, "byteSize == 0")
}

func TestStructNamedHardcodedFieldWithUnsupportedTypeReturnsError(t *testing.T) {
	var protocol xml.Protocol
	require.NoError(t, xml.Unmarshal([]byte(`<protocol>`+testEnumXml+`<struct name="S"><field name="code" type="Code">A</field></struct></protocol>`), &protocol))

	err := writeStruct(jen.NewFile("test"), "S", protocol)

	assert.ErrorContains(t, err, "named hardcoded field S.Code has unsupported type Code")
}

func TestStructOnlyChunkedHasNoTrailingBlankLine(t *testing.T) {
	code := generateStruct(t, `<struct name="S"><chunked><!-- C --><field name="a" type="char"/></chunked></struct>`)

	assert.Contains(t, code, "\tA int // C.\n}")
}
