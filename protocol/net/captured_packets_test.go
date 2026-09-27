package net_test

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"unicode"

	"github.com/ethanmoffat/eolib-go/v3/data"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net/client"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Tests for the packets captured from the official client and server, from the eo-captured-packets repository.
//
// Each packet is deserialized from the captured bytes, and its field values are examined via reflection against the
// captured properties. Each packet must then serialize back to the captured bytes.

const (
	capturedPacketsDir = "../../eo-captured-packets"
	modulePath         = "github.com/ethanmoffat/eolib-go/v3"
)

// capturedPacket is the JSON representation of a packet in the eo-captured-packets repository.
type capturedPacket struct {
	Family     string             `json:"family"`
	Action     string             `json:"action"`
	Expected   []byte             `json:"expected"`
	Properties []capturedProperty `json:"properties"`
}

// capturedProperty is the JSON representation of a packet or structure property in the eo-captured-packets repository.
type capturedProperty struct {
	Type     string             `json:"type"`
	Name     string             `json:"name"`
	Value    json.RawMessage    `json:"value"`
	Optional bool               `json:"optional"`
	Children []capturedProperty `json:"children"`
}

func (p capturedProperty) hasValue() bool {
	return len(p.Value) > 0 && string(p.Value) != "null"
}

// capturedPacketFile is a captured packet file discovered in the eo-captured-packets repository.
type capturedPacketFile struct {
	path   string
	side   string
	name   string
	isOrig bool // True for the invalid data sent by the official client and server, which doesn't round-trip.
}

// unserializablePackets are original packets whose data can't be serialized, because it violates the protocol's
// constraints.
var unserializablePackets = map[string]bool{
	// The characters array has more elements than its length field allows.
	"server/original/PlayersAgree": true,
}

func findCapturedPackets(t *testing.T) []capturedPacketFile {
	var result []capturedPacketFile
	for _, side := range []string{"client", "server"} {
		for _, isOrig := range []bool{false, true} {
			dir := filepath.Join(capturedPacketsDir, side)
			if isOrig {
				dir = filepath.Join(dir, "original")
			}

			paths, err := filepath.Glob(filepath.Join(dir, "*.json"))
			require.NoError(t, err)

			for _, path := range paths {
				name := strings.TrimSuffix(filepath.Base(path), ".json")
				if isOrig {
					name = "original/" + name
				}
				result = append(result, capturedPacketFile{path: path, side: side, name: side + "/" + name, isOrig: isOrig})
			}
		}
	}
	return result
}

func TestCapturedPackets(t *testing.T) {
	files := findCapturedPackets(t)
	require.Greater(t, len(files), 100, "Is the eo-captured-packets submodule initialized?")

	for _, file := range files {
		file := file
		t.Run(file.name, func(t *testing.T) {
			captured := loadCapturedPacket(t, file.path)

			packet := newPacket(t, file.side, captured)
			reader := data.NewEoReader(captured.Expected)
			require.NoError(t, packet.Deserialize(reader))

			assertStructProperties(t, reflect.ValueOf(packet).Elem(), captured.Properties, reflect.TypeOf(packet).Elem().Name())

			writer := data.NewEoWriter()
			err := packet.Serialize(writer)
			if unserializablePackets[file.name] {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)

			if !file.isOrig {
				assert.Equal(t, 0, reader.Remaining())
				assert.Equal(t, len(captured.Expected), packet.ByteSize())
				assert.Equal(t, captured.Expected, writer.Array())
				return
			}

			// The original data doesn't conform to the protocol, so it can't be reproduced. Instead, the serialized data
			// must deserialize to the same property values.
			second := newPacket(t, file.side, captured)
			require.NoError(t, second.Deserialize(data.NewEoReader(writer.Array())))
			assertStructProperties(t, reflect.ValueOf(second).Elem(), captured.Properties, reflect.TypeOf(second).Elem().Name())
		})
	}
}

func loadCapturedPacket(t *testing.T, path string) capturedPacket {
	contents, err := os.ReadFile(path)
	require.NoError(t, err)

	var captured capturedPacket
	require.NoError(t, json.Unmarshal(contents, &captured))
	return captured
}

func newPacket(t *testing.T, side string, captured capturedPacket) net.Packet {
	family, err := parseEnum[net.PacketFamily](captured.Family)
	require.NoError(t, err)
	action, err := parseEnum[net.PacketAction](captured.Action)
	require.NoError(t, err)

	var packet net.Packet
	if side == "client" {
		packet, err = client.PacketFromId(family, action)
	} else {
		packet, err = server.PacketFromId(family, action)
	}
	require.NoError(t, err)
	return packet
}

func parseEnum[T interface {
	~int
	String() (string, error)
}](name string) (T, error) {
	for i := 0; i <= 255; i++ {
		if str, err := T(i).String(); err == nil && str == name {
			return T(i), nil
		}
	}
	return 0, fmt.Errorf("unknown %T value: %s", T(0), name)
}

// assertStructProperties asserts that each captured property matches the corresponding field of the structure. Fields
// that are not listed in the captured properties must have their zero value.
func assertStructProperties(t *testing.T, value reflect.Value, properties []capturedProperty, path string) {
	t.Helper()

	listed := map[string]bool{}
	for _, property := range properties {
		fieldName := snakeCaseToPascalCase(property.Name)
		fieldPath := path + "." + fieldName

		field := value.FieldByName(fieldName)
		if !field.IsValid() {
			t.Errorf("%s: unknown property %q", fieldPath, property.Name)
			continue
		}

		listed[fieldName] = true
		assertPropertyValue(t, field, property, fieldPath)
	}

	for i := 0; i < value.NumField(); i++ {
		structField := value.Type().Field(i)
		if structField.IsExported() && !listed[structField.Name] && !value.Field(i).IsZero() {
			t.Errorf("%s.%s: field not in captured properties has non-zero value %+v", path, structField.Name, value.Field(i).Interface())
		}
	}
}

func assertPropertyValue(t *testing.T, field reflect.Value, property capturedProperty, path string) {
	t.Helper()

	switch field.Kind() {
	case reflect.Pointer:
		if property.Optional && !property.hasValue() && property.Children == nil {
			assert.True(t, field.IsNil(), "%s: expected nil optional value, got %+v", path, field)
			return
		}
		if !assert.False(t, field.IsNil(), "%s: expected a value for the optional field", path) {
			return
		}
		assertPropertyValue(t, field.Elem(), property, path)
	case reflect.Interface:
		// Switch case data is stored as a pointer to a case structure in an interface field.
		if !assert.False(t, field.IsNil(), "%s: expected switch case data of type %s", path, property.Type) {
			return
		}
		caseValue := field.Elem().Elem()
		assertTypeName(t, caseValue.Type(), property.Type, path)
		assertStructProperties(t, caseValue, property.Children, path)
	case reflect.Struct:
		assertTypeName(t, field.Type(), property.Type, path)
		assertStructProperties(t, field, property.Children, path)
	case reflect.Slice:
		if field.Type().Elem().Kind() == reflect.Uint8 {
			var expected []byte
			if property.hasValue() {
				require.NoError(t, json.Unmarshal(property.Value, &expected), path)
			}
			assert.Equal(t, len(expected), field.Len(), "%s: length", path)
			if len(expected) > 0 {
				assert.Equal(t, expected, field.Bytes(), path)
			}
			return
		}

		if !assert.Equal(t, len(property.Children), field.Len(), "%s: length", path) {
			return
		}
		elementType := strings.TrimPrefix(property.Type, "[]")
		for i, child := range property.Children {
			if child.Type == "" {
				child.Type = elementType
			}
			assertPropertyValue(t, field.Index(i), child, fmt.Sprintf("%s[%d]", path, i))
		}
	case reflect.Int:
		if strings.Contains(property.Type, "::") {
			assertTypeName(t, field.Type(), property.Type, path)
		}
		var expected int64
		require.NoError(t, json.Unmarshal(property.Value, &expected), path)
		assert.Equal(t, expected, field.Int(), path)
	case reflect.Bool:
		var expected bool
		require.NoError(t, json.Unmarshal(property.Value, &expected), path)
		assert.Equal(t, expected, field.Bool(), path)
	case reflect.String:
		var expected string
		require.NoError(t, json.Unmarshal(property.Value, &expected), path)
		assert.Equal(t, expected, field.String(), path)
	default:
		t.Errorf("%s: unsupported field kind %s", path, field.Kind())
	}
}

// assertTypeName asserts that a type matches a captured property type in the form "protocol/net/server::TypeName".
func assertTypeName(t *testing.T, actual reflect.Type, capturedType string, path string) {
	t.Helper()

	parts := strings.Split(strings.TrimPrefix(capturedType, "[]"), "::")
	if !assert.Len(t, parts, 2, "%s: unexpected captured type %q", path, capturedType) {
		return
	}
	assert.Equal(t, modulePath+"/"+parts[0], actual.PkgPath(), "%s: package", path)
	assert.Equal(t, parts[1], actual.Name(), "%s: type", path)
}

func snakeCaseToPascalCase(input string) string {
	var sb strings.Builder
	for _, part := range strings.Split(input, "_") {
		if part == "" {
			continue
		}
		runes := []rune(part)
		sb.WriteRune(unicode.ToUpper(runes[0]))
		sb.WriteString(string(runes[1:]))
	}
	return sb.String()
}
