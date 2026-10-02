package codegen

import (
	"testing"

	"github.com/ethanmoffat/eolib-go/v3/internal/codegen/types"
	"github.com/ethanmoffat/eolib-go/v3/internal/xml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testInnerEnumXml = `<enum name="Inner" type="char"><value name="X">1</value><value name="Y">2</value></enum>`

func unmarshalTestProtocol(t *testing.T, protocolXml string) xml.Protocol {
	t.Helper()

	var protocol xml.Protocol
	require.NoError(t, xml.Unmarshal([]byte("<protocol>"+protocolXml+"</protocol>"), &protocol))
	for i := range protocol.Structs {
		protocol.Structs[i].Package = "test"
	}
	for i := range protocol.Packets {
		protocol.Packets[i].Package = "test"
	}

	require.NoError(t, protocol.Validate())
	return protocol
}

func TestSwitchFactoryForEachEnumValue(t *testing.T) {
	code := generateStruct(t, testEnumXml+`<struct name="S">
		<field name="code" type="Code"/>
		<switch field="code">
			<case value="A"><field name="a" type="char"/></case>
			<case value="B"><field type="string">NO</field></case>
			<case value="0"><field name="zero" type="char"/></case>
			<case value="4"/>
		</switch>
	</struct>`)

	tests := []struct {
		name     string
		expected string
	}{
		{"NamedCaseWithMembers", "// NewSWithA creates a new [S] with Code set to [Code_A] and CodeData set to data.\n" +
			"// If data is nil, CodeData is left nil and serializing the result returns an error.\n" +
			"func NewSWithA(data *CodeDataA) *S {\n\tif data == nil {\n\t\treturn &S{Code: Code_A}\n\t}\n\treturn &S{Code: Code_A, CodeData: data}\n}\n"},
		{"NamedCaseWithoutMembers", "// NewSWithB creates a new [S] with Code set to [Code_B] and CodeData set to a new [CodeDataB].\n" +
			"func NewSWithB() *S {\n\treturn &S{Code: Code_B, CodeData: &CodeDataB{}}\n}\n"},
		{"NamedValueWithoutCase", "// NewSWithC creates a new [S] with Code set to [Code_C].\nfunc NewSWithC() *S {\n\treturn &S{Code: Code_C}\n}\n"},
		{"NumericCaseWithMembers", "// NewSWithCodeData0 creates a new [S] with Code set to 0 and CodeData set to data.\n" +
			"// If data is nil, CodeData is left nil and serializing the result returns an error.\n" +
			"func NewSWithCodeData0(data *CodeData0) *S {\n\tif data == nil {\n\t\treturn &S{Code: 0}\n\t}\n\treturn &S{Code: 0, CodeData: data}\n}\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, code, tt.expected)
		})
	}

	assert.NotContains(t, code, "NewSWithCodeData4", "numeric cases without data are skipped")
	assert.NotContains(t, code, "NewSWithDefault")
}

func TestSwitchFactoryForPacketUsesPacketPrefix(t *testing.T) {
	protocol := unmarshalTestProtocol(t, testEnumXml+`<packet family="Login" action="Reply">
		<field name="code" type="Code"/>
		<switch field="code"><case value="A"><field name="a" type="char"/></case></switch>
	</packet>`)

	si, err := types.GetStructInfo("LoginReplyTestPacket", protocol)
	require.NoError(t, err)
	factories, err := getSwitchFactories(si, protocol)
	require.NoError(t, err)

	require.Len(t, factories, 3)
	assert.Equal(t, "NewLoginReplyWithA", factories[0].Name)
	assert.Equal(t, "LoginReplyTestPacket", factories[0].Owner)
	assert.Equal(t, "LoginReplyCodeDataA", factories[0].DataType)
	assert.Equal(t, "NewLoginReplyWithB", factories[1].Name)
	assert.Equal(t, "NewLoginReplyWithC", factories[2].Name)
}

func TestSwitchFactoryNestedSwitchIsFlattened(t *testing.T) {
	code := generateStruct(t, testEnumXml+testInnerEnumXml+`<struct name="S">
		<field name="code" type="Code"/>
		<switch field="code">
			<case value="A">
				<field name="inner" type="Inner"/>
				<field type="char">0</field>
				<switch field="inner">
					<case value="0"><field name="zero" type="char"/></case>
					<case value="X"><field name="x" type="char"/></case>
				</switch>
			</case>
		</switch>
	</struct>`)

	tests := []struct {
		name     string
		expected string
	}{
		{"NamedInnerCase", "// NewSWithAX creates a new [S] with Code set to [Code_A] and CodeData set to a new [CodeDataA] with Inner set to [Inner_X] and InnerData set to data.\n" +
			"// If data is nil, InnerData is left nil and serializing the result returns an error.\n" +
			"func NewSWithAX(data *InnerDataX) *S {\n\tif data == nil {\n\t\treturn &S{Code: Code_A, CodeData: &CodeDataA{Inner: Inner_X}}\n\t}\n\treturn &S{Code: Code_A, CodeData: &CodeDataA{Inner: Inner_X, InnerData: data}}\n}\n"},
		{"InnerValueWithoutCase", "func NewSWithAY() *S {\n\treturn &S{Code: Code_A, CodeData: &CodeDataA{Inner: Inner_Y}}\n}\n"},
		{"NumericInnerCase", "func NewSWithInnerData0(data *InnerData0) *S {\n\tif data == nil {\n\t\treturn &S{Code: Code_A, CodeData: &CodeDataA{Inner: 0}}\n\t}\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Contains(t, code, tt.expected)
		})
	}

	assert.NotContains(t, code, "func NewSWithA(", "a case with a nested switch is replaced by the inner factories")
	assert.Contains(t, code, "func NewSWithB() *S")
}

func TestSwitchFactoryNestedDefaultCaseSetsOuterSwitch(t *testing.T) {
	code := generateStruct(t, testEnumXml+testInnerEnumXml+`<struct name="S">
		<field name="code" type="Code"/>
		<switch field="code">
			<case value="A">
				<field name="inner" type="Inner"/>
				<switch field="inner">
					<case value="X"/>
					<case default="true"><field name="d" type="char"/></case>
				</switch>
			</case>
		</switch>
	</struct>`)

	assert.Contains(t, code, "func NewSWithADefault(code Inner, data *InnerDataDefault) (*S, error) {\n\tswitch code {\n\tcase Inner_X:\n")
	assert.Contains(t, code, "\treturn &S{Code: Code_A, CodeData: &CodeDataA{Inner: code, InnerData: data}}, nil\n")
}

func TestSwitchFactoryDefaultCase(t *testing.T) {
	code := generateStruct(t, testEnumXml+`<struct name="S">
		<field name="code" type="Code"/>
		<switch field="code">
			<case value="0"/>
			<case value="A"><field type="string">NO</field></case>
			<case value="B"><field name="b" type="char"/></case>
			<case default="true"><field name="session_id" type="short"/></case>
		</switch>
	</struct>`)

	assert.Contains(t, code, "// NewSWithDefault creates a new [S] with Code set to code and CodeData set to data, for a Code value that is handled by the default case.\n"+
		"// It returns an error if code is the value of another case.\n"+
		"// If data is nil, CodeData is left nil and serializing the result returns an error.\n"+
		"func NewSWithDefault(code Code, data *CodeDataDefault) (*S, error) {\n"+
		"\tswitch code {\n\tcase 0, Code_A, Code_B:\n"+
		"\t\treturn nil, fmt.Errorf(\"Code %d has its own case and is not handled by the default case\", code)\n\t}\n"+
		"\tif data == nil {\n\t\treturn &S{Code: code}, nil\n\t}\n\treturn &S{Code: code, CodeData: data}, nil\n}\n")
	assert.Contains(t, code, "func NewSWithA() *S")
	assert.Contains(t, code, "func NewSWithB(data *CodeDataB) *S")
	assert.NotContains(t, code, "NewSWithC", "values without a case are handled by the default case")
}

func TestSwitchFactoryDefaultCaseWithoutMembers(t *testing.T) {
	code := generateStruct(t, testEnumXml+`<struct name="S">
		<field name="code" type="Code"/>
		<switch field="code">
			<case value="A"/>
			<case default="true"><field type="string">OK</field></case>
		</switch>
	</struct>`)

	assert.Contains(t, code, "// It returns an error if code is the value of another case.\n"+
		"func NewSWithDefault(code Code) (*S, error) {\n")
	assert.Contains(t, code, "\treturn &S{Code: code, CodeData: &CodeDataDefault{}}, nil\n")
}

func TestSwitchFactoryIntegerSwitchField(t *testing.T) {
	code := generateStruct(t, `<struct name="S">
		<field name="code" type="char"/>
		<switch field="code">
			<case value="1"><field name="x" type="char"/></case>
			<case value="2"/>
			<case default="true"><field name="d" type="char"/></case>
		</switch>
	</struct>`)

	assert.Contains(t, code, "func NewSWithCodeData1(data *CodeData1) *S {\n\tif data == nil {\n\t\treturn &S{Code: 1}\n\t}\n")
	assert.Contains(t, code, "func NewSWithDefault(code int, data *CodeDataDefault) (*S, error) {\n\tswitch code {\n\tcase 1, 2:\n")
}

func TestSwitchFactoryNoSwitchHasNoFactories(t *testing.T) {
	protocol := unmarshalTestProtocol(t, `<struct name="S"><field name="a" type="char"/></struct>`)

	si, err := types.GetStructInfo("S", protocol)
	require.NoError(t, err)
	factories, err := getSwitchFactories(si, protocol)

	require.NoError(t, err)
	assert.Empty(t, factories)
}
