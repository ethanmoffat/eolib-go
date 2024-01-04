package eotype

type SerializationType int

const (
	Invalid SerializationType = 0
)

const (
	Primitive SerializationType = iota + 0x0100 // flag indicating type is a primitive (supported for bool)
	Byte
	Char
	Short
	Three
	Int
	Bool
)

const (
	Complex SerializationType = iota + 0x0200 // flag indicating type is complex (not supported for bool)
	Bytes
)

const (
	String SerializationType = iota + 0x0400 // flag indicating type is a string type
	PaddedString
	FixedString
)

const (
	EncodedString SerializationType = iota + 0x0800 // flag indicating type is an encoded string type
	PaddedEncodedString
	FixedEncodedString
)

// offsets from String or EncodedString to the other string method types
const (
	_ SerializationType = iota
	Padded
	Fixed
)

func (t SerializationType) String() string {
	switch t {
	case Byte:
		return "Byte"
	case Char:
		return "Char"
	case Short:
		return "Short"
	case Three:
		return "Three"
	case Int:
		return "Int"
	case Bool:
		return "Byte"
	case Bytes:
		return "Bytes"
	case String:
		return "String"
	case PaddedString:
		return "PaddedString"
	case FixedString:
		return "FixedString"
	case EncodedString:
		return "EncodedString"
	case PaddedEncodedString:
		return "PaddedEncodedString"
	case FixedEncodedString:
		return "FixedEncodedString"
	}

	return ""
}

func NewSerializationType(str string) SerializationType {
	switch str {
	case "byte":
		return Byte
	case "char":
		return Char
	case "short":
		return Short
	case "three":
		return Three
	case "int":
		return Int
	case "blob":
		return Bytes
	}

	return Invalid
}
