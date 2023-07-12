package data

import eolib "github.com/ethanmoffat/eolib-go/pkg/eolib"

// EncodeString encodes a string by inverting the bytes and then reversing them.
func EncodeString(str string) []byte {
	inverted := invert([]byte(str))
	return eolib.Reverse(inverted)
}

// DecodeString decodes a string by reversing the bytes and then inverting them.
func DecodeString(bytes []byte) string {
	reversed := eolib.Reverse(bytes)
	return string(invert(reversed))
}

func invert(bytes []byte) []byte {
	flippy := len(bytes)%2 == 1

	retBytes := make([]byte, len(bytes))

	for i, c := range bytes {
		retBytes[i] = c

		f := 0

		if flippy {
			f = 0x2E
			if c >= 0x50 {
				f *= -1
			}
		}

		if c >= 0x22 && c <= 0x7e {
			retBytes[i] = byte(0x9F - int(c) - f)
		}

		flippy = !flippy
	}

	return retBytes
}
