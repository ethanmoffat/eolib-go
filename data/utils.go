package data

import "golang.org/x/text/encoding/charmap"

// StringFromBytes converts a sequence of bytes to a string using the Windows-1252 character set
func StringFromBytes(bytes []byte) string {
	var ret []rune
	for _, b := range bytes {
		next := charmap.Windows1252.DecodeByte(b)
		ret = append(ret, next)
	}
	return string(ret)
}

// BytesFromString converts a string to a sequence of bytes using the Windows-1252 character set
func BytesFromString(str string) []byte {
	ret := make([]byte, 0, len(str))
	for _, r := range str {
		next, _ := charmap.Windows1252.EncodeRune(r)
		ret = append(ret, next)
	}
	return ret
}
