package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
)

type encodeNumberParameters struct {
	a      byte
	b      byte
	c      byte
	d      byte
	number int
}

type encodeStringParameters struct {
	decoded string
	encoded string
}

var encodeNumberTestCases []encodeNumberParameters
var encodeStringTestCases []encodeStringParameters

func TestMain(m *testing.M) {
	encodeNumberTestCases = []encodeNumberParameters{
		{0x01, 0xFE, 0xFE, 0xFE, 0},
		{0x02, 0xFE, 0xFE, 0xFE, 1},
		{0x1D, 0xFE, 0xFE, 0xFE, 28},
		{0x65, 0xFE, 0xFE, 0xFE, 100},
		{0x81, 0xFE, 0xFE, 0xFE, 128},
		{0xFD, 0xFE, 0xFE, 0xFE, 252},
		{0x01, 0x02, 0xFE, 0xFE, 253},
		{0x02, 0x02, 0xFE, 0xFE, 254},
		{0x03, 0x02, 0xFE, 0xFE, 255},
		{0x7E, 0x7F, 0xFE, 0xFE, 32003},
		{0x7F, 0x7F, 0xFE, 0xFE, 32004},
		{0x80, 0x7F, 0xFE, 0xFE, 32005},
		{0xFD, 0xFD, 0xFE, 0xFE, 64008},
		{0x01, 0x01, 0x02, 0xFE, 64009},
		{0x02, 0x01, 0x02, 0xFE, 64010},
		{0xB0, 0x3A, 0x9D, 0xFE, 10_000_000},
		{0xFD, 0xFD, 0xFD, 0xFE, 16_194_276},
		{0x01, 0x01, 0x01, 0x02, 16_194_277},
		{0x02, 0x01, 0x01, 0x02, 16_194_278},
		{0x7E, 0x7F, 0x7F, 0x7F, 2_048_576_039},
		{0x7F, 0x7F, 0x7F, 0x7F, 2_048_576_040},
		{0x80, 0x7F, 0x7F, 0x7F, 2_048_576_041},
		{0xFC, 0xFD, 0xFD, 0xFD, int(4097152079)},
		{0xFD, 0xFD, 0xFD, 0xFD, int(4097152080)},
	}

	encodeStringTestCases = []encodeStringParameters{
		{"Hello, World!", "!;a-^H s^3a:)"},
		{"We're ¼ of the way there, so ¾ is remaining.", "C8_6_6l2h- ,d ¾ ^, sh-h7Y T>V h7Y g0 ¼ :[xhH"},
		{"64² = 4096", ";fAk b ²=i"},
		{"© FÒÖ BÃR BÅZ 2014", "=nAm EÅ] MÃ] ÖÒY ©"},
		{"Öxxö Xööx \"Lëïth Säë\" - \"Ÿ\"", "OŸO D OëäL 7YïëSO UööG öU'Ö"},
		{"Padded with 0xFFÿÿÿÿÿÿÿÿ", "ÿÿÿÿÿÿÿÿ+YUo 7Y6V i:i;lO"},
	}

	os.Exit(m.Run())
}

func TestEncodeNumber(t *testing.T) {
	for _, tc := range encodeNumberTestCases {
		t.Run(fmt.Sprintf("%d should encode to [%d, %d, %d, %d]", tc.number, tc.a, tc.b, tc.c, tc.d),
			func(t *testing.T) {
				actual := EncodeNumber(tc.number)
				assert.Equal(t, []byte{tc.a, tc.b, tc.c, tc.d}, actual)
			})
	}
}

func TestDecodeNumber(t *testing.T) {
	for _, tc := range encodeNumberTestCases {
		t.Run(fmt.Sprintf("[%d, %d, %d, %d] should decode to %d", tc.a, tc.b, tc.c, tc.d, tc.number),
			func(t *testing.T) {
				actual := DecodeNumber([]byte{tc.a, tc.b, tc.c, tc.d})
				assert.Equal(t, tc.number, actual)
			})
	}
}

func TestEncodeString(t *testing.T) {
	for _, tc := range encodeStringTestCases {
		t.Run(fmt.Sprintf("%s should encode to %s", tc.decoded, tc.encoded),
			func(t *testing.T) {
				bytes := toBytes(tc.decoded)
				encoded := EncodeString(bytes)
				assert.Equal(t, toBytes(tc.encoded), encoded)
			})
	}
}

func TestDecodeString(t *testing.T) {
	for _, tc := range encodeStringTestCases {
		t.Run(fmt.Sprintf("%s should decode to %s", tc.encoded, tc.decoded),
			func(t *testing.T) {
				bytes := toBytes(tc.encoded)
				decoded := DecodeString(bytes)
				assert.Equal(t, toBytes(tc.decoded), decoded)
			})
	}
}

func toBytes(input string) (ret []byte) {
	for _, r := range input {
		next, _ := charmap.Windows1252.EncodeRune(r)
		ret = append(ret, next)
	}
	return
}
