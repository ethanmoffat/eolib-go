package data

import (
	eolib "github.com/ethanmoffat/eolib-go/pkg/eolib"
	"golang.org/x/text/encoding/charmap"
)

// EncodeNumber encodes a number to a sequence of bytes.
func EncodeNumber(number int) []byte {
	value := number

	d := 0xFE
	if number >= THREE_MAX {
		d = value/THREE_MAX + 1
		value = value % THREE_MAX
	}

	c := 0xFE
	if number >= SHORT_MAX {
		c = value/SHORT_MAX + 1
		value = value % SHORT_MAX
	}

	b := 0xFE
	if number >= CHAR_MAX {
		b = value/CHAR_MAX + 1
		value = value % CHAR_MAX
	}

	a := value + 1

	return []byte{byte(a), byte(b), byte(c), byte(d)}
}

// DecodeNumber decodes a number from a sequence of bytes.
func DecodeNumber(bytes []byte) int {
	result := 0
	length := eolib.Min(len(bytes), 4)

	for i := 0; i < length; i++ {
		value := int(bytes[i])

		if value == 0xFE {
			break
		}

		value--

		if i == 0 {
			result += value
		} else if i == 1 {
			result += CHAR_MAX * value
		} else if i == 2 {
			result += SHORT_MAX * value
		} else if i == 3 {
			result += THREE_MAX * value
		}
	}

	return result
}

// EncodeString encodes a string by inverting the bytes and then reversing them.
func EncodeString(str []byte) []byte {
	inverted := invert(str)
	return eolib.Reverse(inverted)
}

// DecodeString decodes a string by reversing the bytes and then inverting them.
func DecodeString(bytes []byte) []byte {
	reversed := eolib.Reverse(bytes)
	return invert(reversed)
}

func invert(bytes []byte) []byte {
	flippy := len(bytes)%2 == 1

	retBytes := make([]byte, len(bytes))
	copy(retBytes, bytes)

	for i, c := range retBytes {
		retBytes[i] = c

		f := 0

		if flippy {
			f = 0x2E
			if c >= 0x50 {
				f *= -1
			}
		}

		if c >= 0x22 && c <= 0x7E {
			retBytes[i] = 0x9F - c - byte(f)
		}

		flippy = !flippy
	}

	return retBytes
}

// windows1252String converts a sequence of bytes to a string using the Windows-1252 character set
func windows1252String(bytes []byte) string {
	var ret []rune
	for _, b := range bytes {
		next := charmap.Windows1252.DecodeByte(b)
		ret = append(ret, next)
	}
	return string(ret)
}

func windows1252Bytes(str string) []byte {
	ret := make([]byte, 0, len(str))
	for _, r := range str {
		next, _ := charmap.Windows1252.EncodeRune(r)
		ret = append(ret, next)
	}
	return ret
}
