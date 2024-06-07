package data_test

import (
	"fmt"
	"testing"

	"github.com/ethanmoffat/eolib-go/v3/data"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	reader := createReader([]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06})
	reader.GetByte()
	reader.SetIsChunked(true)

	reader2, err := reader.SliceFromCurrent()
	assert.NoError(t, err)
	assert.Equal(t, 0, reader2.Position())
	assert.Equal(t, 5, reader2.Remaining())
	assert.False(t, reader2.IsChunked())

	reader3, err := reader2.SliceFromIndex(1)
	assert.NoError(t, err)
	assert.Equal(t, 0, reader3.Position())
	assert.Equal(t, 4, reader3.Remaining())
	assert.False(t, reader3.IsChunked())

	reader4, err := reader3.Slice(1, 2)
	assert.NoError(t, err)
	assert.Equal(t, 0, reader4.Position())
	assert.Equal(t, 2, reader4.Remaining())
	assert.False(t, reader4.IsChunked())

	assert.Equal(t, 1, reader.Position())
	assert.Equal(t, 5, reader.Remaining())
	assert.True(t, reader.IsChunked())
}

func TestSliceOverRead(t *testing.T) {
	reader := createReader([]byte{0x01, 0x02, 0x03})

	res, _ := reader.Slice(2, 5)
	assert.Equal(t, 1, res.Remaining())

	res, _ = reader.SliceFromIndex(3)
	assert.Equal(t, 0, res.Remaining())

	res, _ = reader.SliceFromIndex(4)
	assert.Equal(t, 0, res.Remaining())

	res, _ = reader.Slice(4, 12345)
	assert.Equal(t, 0, res.Remaining())
}

func TestSliceNegativeInputs(t *testing.T) {
	reader := data.EoReader{}

	_, err := reader.SliceFromIndex(-1)
	assert.Error(t, err)

	_, err = reader.Slice(0, -1)
	assert.Error(t, err)
}

func TestReaderGetByte(t *testing.T) {
	byteValues := []byte{0x00, 0x01, 0x02, 0x80, 0xFD, 0xFE, 0xFF}
	for _, expected := range byteValues {
		t.Run(fmt.Sprintf("ReadByte_%d", expected), func(t *testing.T) {
			reader := createReader([]byte{expected})
			actual := reader.GetByte()
			assert.Equal(t, expected, actual)
		})
	}
}

func TestReaderOverReadByte(t *testing.T) {
	reader := createReader([]byte{})
	value := reader.GetByte()
	assert.Equal(t, byte(0x00), value)
}

func TestReaderGetBytes(t *testing.T) {
	reader := createReader([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
	assert.Equal(t, []byte{0x01, 0x02, 0x03}, reader.GetBytes(3))
	assert.Equal(t, []byte{0x04, 0x05}, reader.GetBytes(10))
	assert.Equal(t, []byte{}, reader.GetBytes(1))
}

func TestReaderGetChar(t *testing.T) {
	reader := createReader([]byte{0x01, 0x02, 0x80, 0x81, 0xFD, 0xFE, 0xFF})
	assert.Equal(t, 0, reader.GetChar())
	assert.Equal(t, 1, reader.GetChar())
	assert.Equal(t, 127, reader.GetChar())
	assert.Equal(t, 128, reader.GetChar())
	assert.Equal(t, 252, reader.GetChar())
	assert.Equal(t, 0, reader.GetChar())
	assert.Equal(t, 254, reader.GetChar())
}

func TestReaderGetShort(t *testing.T) {
	reader := data.NewEoReader([]byte{0x01, 0xFE, 0x02, 0xFE, 0x80, 0xFE, 0xFD, 0xFE,
		0xFE, 0xFE, 0xFE, 0x80, 0x7F, 0x7F, 0xFD, 0xFD})
	assert.Equal(t, 0, reader.GetShort())
	assert.Equal(t, 1, reader.GetShort())
	assert.Equal(t, 127, reader.GetShort())
	assert.Equal(t, 252, reader.GetShort())
	assert.Equal(t, 0, reader.GetShort())
	assert.Equal(t, 0, reader.GetShort())
	assert.Equal(t, 32004, reader.GetShort())
	assert.Equal(t, 64008, reader.GetShort())
}

func TestReaderGetThree(t *testing.T) {
	reader := data.NewEoReader([]byte{
		0x01, 0xFE, 0xFE, // 0
		0x02, 0xFE, 0xFE, // 1
		0x80, 0xFE, 0xFE, // 127
		0xFD, 0xFE, 0xFE, // 252
		0xFE, 0xFE, 0xFE, // 0
		0xFE, 0x80, 0x81, // 0
		0x7F, 0x7F, 0xFE, // 32004
		0xFD, 0xFD, 0xFE, // 64008
		0xFD, 0xFD, 0xFD}) // 16194276
	assert.Equal(t, 0, reader.GetThree())
	assert.Equal(t, 1, reader.GetThree())
	assert.Equal(t, 127, reader.GetThree())
	assert.Equal(t, 252, reader.GetThree())
	assert.Equal(t, 0, reader.GetThree())
	assert.Equal(t, 0, reader.GetThree())
	assert.Equal(t, 32004, reader.GetThree())
	assert.Equal(t, 64008, reader.GetThree())
	assert.Equal(t, 16194276, reader.GetThree())
}

func TestReaderGetInt(t *testing.T) {
	reader := data.NewEoReader([]byte{
		0x01, 0xFE, 0xFE, 0xFE, // 0
		0x02, 0xFE, 0xFE, 0xFE, // 1
		0x80, 0xFE, 0xFE, 0xFE, // 127
		0xFD, 0xFE, 0xFE, 0xFE, // 252
		0xFE, 0xFE, 0xFE, 0xFE, // 0
		0xFE, 0x80, 0x81, 0x82, // 0
		0x7F, 0x7F, 0xFE, 0xFE, // 32004
		0xFD, 0xFD, 0xFE, 0xFE, // 64008
		0xFD, 0xFD, 0xFD, 0xFE, // 16194276
		0x7F, 0x7F, 0x7F, 0x7F, // 2048576040
		0xFD, 0xFD, 0xFD, 0xFD}) // 4097152080
	assert.Equal(t, 0, reader.GetInt())
	assert.Equal(t, 1, reader.GetInt())
	assert.Equal(t, 127, reader.GetInt())
	assert.Equal(t, 252, reader.GetInt())
	assert.Equal(t, 0, reader.GetInt())
	assert.Equal(t, 0, reader.GetInt())
	assert.Equal(t, 32004, reader.GetInt())
	assert.Equal(t, 64008, reader.GetInt())
	assert.Equal(t, 16194276, reader.GetInt())
	assert.Equal(t, 2_048_576_040, reader.GetInt())
	assert.Equal(t, 4_097_152_080, reader.GetInt())
}

func TestReaderGetString(t *testing.T) {
	const expected = "Hello, World!"
	reader := createReader(toBytes(expected))
	actual, _ := reader.GetString()
	assert.Equal(t, expected, actual)
}

func TestReaderGetFixedString(t *testing.T) {
	const input = "foobar"
	reader := createReader(toBytes(input))
	actual1, _ := reader.GetFixedString(3)
	actual2, _ := reader.GetFixedString(3)
	assert.Equal(t, input[:3], actual1)
	assert.Equal(t, input[3:], actual2)
}

func TestReaderGetPaddedString(t *testing.T) {
	const input = "fooÿbarÿÿÿ"
	reader := createReader(toBytes(input))
	actual1, _ := reader.GetPaddedString(4)
	actual2, _ := reader.GetPaddedString(6)
	assert.Equal(t, "foo", actual1)
	assert.Equal(t, "bar", actual2)
}

func TestReaderGetStringChunked(t *testing.T) {
	const input = "Hello,ÿWorld!"
	reader := createReader(toBytes(input))
	reader.SetIsChunked(true)

	actual1, _ := reader.GetString()
	assert.Equal(t, "Hello,", actual1)

	reader.NextChunk()
	actual2, _ := reader.GetString()
	assert.Equal(t, "World!", actual2)
}

func TestReaderGetNegativeLengthFixedString(t *testing.T) {
	reader := data.EoReader{}
	_, err := reader.GetFixedString(-1)
	assert.ErrorContains(t, err, "negative length")
}

func TestReaderGetNegativeLengthPaddedString(t *testing.T) {
	reader := data.EoReader{}
	_, err := reader.GetPaddedString(-1)
	assert.ErrorContains(t, err, "negative length")
}

func TestReaderGetEncodedString(t *testing.T) {
	const input = "!;a-^H s^3a:)"
	reader := createReader(toBytes(input))
	actual, _ := reader.GetEncodedString()
	assert.Equal(t, "Hello, World!", actual)
}

func TestReaderGetFixedEncodedString(t *testing.T) {
	const input = "^0g[>k"
	reader := createReader(toBytes(input))
	actual1, _ := reader.GetFixedEncodedString(3)
	actual2, _ := reader.GetFixedEncodedString(3)
	assert.Equal(t, "foo", actual1)
	assert.Equal(t, "bar", actual2)
}

func TestReaderGetPaddedEncodedString(t *testing.T) {
	const input = "ÿ0^9ÿÿÿ-l=S>k"
	reader := createReader(toBytes(input))
	actual1, _ := reader.GetPaddedEncodedString(4)
	actual2, _ := reader.GetPaddedEncodedString(6)
	actual3, _ := reader.GetPaddedEncodedString(3)
	assert.Equal(t, "foo", actual1)
	assert.Equal(t, "bar", actual2)
	assert.Equal(t, "baz", actual3)
}

func TestReaderGetEncodedStringChunked(t *testing.T) {
	const input = "E0a3hWÿ!;a-^H"
	reader := createReader(toBytes(input))
	reader.SetIsChunked(true)

	actual1, _ := reader.GetEncodedString()
	assert.Equal(t, "Hello,", actual1)

	reader.NextChunk()
	actual2, _ := reader.GetEncodedString()
	assert.Equal(t, "World!", actual2)
}

func TestReaderGetNegativeLengthFixedEncodedString(t *testing.T) {
	reader := data.EoReader{}
	_, err := reader.GetFixedEncodedString(-1)
	assert.ErrorContains(t, err, "negative length")
}

func TestReaderGetNegativeLengthPaddedEncodedString(t *testing.T) {
	reader := data.EoReader{}
	_, err := reader.GetPaddedEncodedString(-1)
	assert.ErrorContains(t, err, "negative length")
}

func TestReaderRemaining(t *testing.T) {
	bytes := []byte{0x01, 0x03, 0x04, 0xFE, 0x05, 0xFE, 0xFE, 0x06, 0xFE, 0xFE, 0xFE}
	reader := createReader(bytes)

	assert.Equal(t, len(bytes), reader.Remaining())

	reader.GetByte()
	assert.Equal(t, len(bytes)-1, reader.Remaining())

	reader.GetChar()
	assert.Equal(t, len(bytes)-2, reader.Remaining())

	reader.GetShort()
	assert.Equal(t, len(bytes)-4, reader.Remaining())

	reader.GetThree()
	assert.Equal(t, len(bytes)-7, reader.Remaining())

	reader.GetInt()
	assert.Equal(t, len(bytes)-11, reader.Remaining())

	reader.GetChar()
	assert.Equal(t, 0, reader.Remaining())
}

func TestReaderRemainingChunked(t *testing.T) {
	bytes := []byte{
		0x01, 0x03, 0x04,
		0xFF, // chunk delimiter
		0x05, 0xFE, 0xFE, 0x06, 0xFE, 0xFE, 0xFE}

	reader := createReader(bytes)
	reader.SetIsChunked(true)

	assert.Equal(t, 3, reader.Remaining())

	reader.GetChar()
	reader.GetShort()
	assert.Equal(t, 0, reader.Remaining())

	reader.GetChar()
	assert.Equal(t, 0, reader.Remaining())

	reader.NextChunk()
	assert.Equal(t, 7, reader.Remaining())
}

func TestReaderNextChunk(t *testing.T) {
	bytes := []byte{
		0x01, 0x02,
		0xFF, // chunk delimiter
		0x03, 0x04, 0x5,
		0xFF, // chunk delimiter
		0x06}

	reader := createReader(bytes)
	reader.SetIsChunked(true)

	assert.Equal(t, 0, reader.Position())

	reader.NextChunk()
	assert.Equal(t, 3, reader.Position())

	reader.NextChunk()
	assert.Equal(t, 7, reader.Position())

	reader.NextChunk()
	assert.Equal(t, 8, reader.Position())

	reader.NextChunk()
	assert.Equal(t, 8, reader.Position())
}

func TestReaderNextChunkError(t *testing.T) {
	reader := data.EoReader{}
	err := reader.NextChunk()
	assert.ErrorContains(t, err, "not in chunked reading mode")
}

func TestReaderNextChunkWithChunkedModeToggledInBetween(t *testing.T) {
	bytes := []byte{
		0x01, 0x02,
		0xFF, // chunk delimiter
		0x03, 0x04, 0x5,
		0xFF, // chunk delimiter
		0x06}

	reader := createReader(bytes)

	assert.Equal(t, 0, reader.Position())

	reader.SetIsChunked(true)
	reader.NextChunk()
	reader.SetIsChunked(false)
	assert.Equal(t, 3, reader.Position())

	reader.SetIsChunked(true)
	reader.NextChunk()
	reader.SetIsChunked(false)
	assert.Equal(t, 7, reader.Position())

	reader.SetIsChunked(true)
	reader.NextChunk()
	reader.SetIsChunked(false)
	assert.Equal(t, 8, reader.Position())

	reader.SetIsChunked(true)
	reader.NextChunk()
	reader.SetIsChunked(false)
	assert.Equal(t, 8, reader.Position())
}

func TestReaderUnderRead(t *testing.T) {
	// See: https://github.com/Cirras/eo-protocol/blob/master/docs/chunks.md#1-under-read
	reader := createReader([]byte{0x7C, 0x67, 0x61, 0x72, 0x62, 0x61, 0x67, 0x65, 0xFF, 0xCA, 0x31})
	reader.SetIsChunked(true)

	assert.Equal(t, 123, reader.GetChar()) // byte representation: 123 = 0x7C

	reader.NextChunk()
	assert.Equal(t, 12345, reader.GetShort()) // byte representation: 12345 = 0xCA 0x31
}

func TestOverRead(t *testing.T) {
	// See: https://github.com/Cirras/eo-protocol/blob/master/docs/chunks.md#2-over-read
	reader := createReader([]byte{0xFF, 0x7C})
	reader.SetIsChunked(true)

	assert.Equal(t, 0, reader.GetInt())

	reader.NextChunk()
	assert.Equal(t, 123, reader.GetShort())
}

func TestDoubleRead(t *testing.T) {
	// See: https://github.com/Cirras/eo-protocol/blob/master/docs/chunks.md#3-double-read
	reader := createReader([]byte{0xFF, 0x7C, 0xCA, 0x31})

	// Reading all 4 bytes of the input data
	assert.Equal(t, 790222478, reader.GetInt())

	// Activating chunked mode and seeking to the first break byte with nextChunk(), which actually
	// takes our reader position backwards.
	reader.SetIsChunked(true)
	reader.NextChunk()

	assert.Equal(t, 123, reader.GetChar())
	assert.Equal(t, 12345, reader.GetShort())
}

func createReader(inp []byte) *data.EoReader {
	tmp := make([]byte, len(inp)+20)
	for i, b := range inp {
		tmp[i+10] = b
	}

	reader := data.NewEoReader(tmp)
	retReader, _ := reader.Slice(10, len(inp))
	return retReader
}
