package data_test

import (
	"testing"

	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	"github.com/stretchr/testify/assert"
)

func TestWriterAddByte(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddByte(0x00)
	assert.Equal(t, []byte{0x00}, writer.ToByteArray())
}

func TestWriterAddBytes(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddBytes([]byte{0x00, 0xFF})
	assert.Equal(t, []byte{0x00, 0xFF}, writer.ToByteArray())
}

func TestWriterAddChar(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddChar(123)
	assert.Equal(t, []byte{0x7C}, writer.ToByteArray())
}

func TestWriterAddShort(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddShort(12345)
	assert.Equal(t, []byte{0xCA, 0x31}, writer.ToByteArray())
}

func TestWriterAddThree(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddThree(10_000_000)
	assert.Equal(t, []byte{0xB0, 0x3A, 0x9D}, writer.ToByteArray())
}

func TestWriterAddInt(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddInt(2_048_576_040)
	assert.Equal(t, []byte{0x7F, 0x7F, 0x7F, 0x7F}, writer.ToByteArray())
}

func TestWriterAddString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddString("foo")
	assert.Equal(t, toBytes("foo"), writer.ToByteArray())
}

func TestWriterAddFixedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddFixedString("bar", 3)
	assert.Equal(t, toBytes("bar"), writer.ToByteArray())
}

func TestWriterAddPaddedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddPaddedString("bar", 6)
	assert.Equal(t, toBytes("barÿÿÿ"), writer.ToByteArray())
}

func TestWriterAddPaddedStringWithPerfectFit(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddPaddedString("bar", 3)
	assert.Equal(t, toBytes("bar"), writer.ToByteArray())
}

func TestWriterAddEncodedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddEncodedString("foo")
	assert.Equal(t, toBytes("^0g"), writer.ToByteArray())
}

func TestWriterAddFixedEncodedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddFixedEncodedString("bar", 3)
	assert.Equal(t, toBytes("[>k"), writer.ToByteArray())
}

func TestWriterAddPaddedEncodedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddPaddedEncodedString("bar", 6)
	assert.Equal(t, toBytes("ÿÿÿ-l="), writer.ToByteArray())
}

func TestWriterAddPaddedEncodedStringWithPerfectFit(t *testing.T) {
	writer := data.NewEoWriter()
	writer.AddPaddedEncodedString("bar", 3)
	assert.Equal(t, toBytes("[>k"), writer.ToByteArray())
}

func TestWriterAddSanitizedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.SanitizeStrings = true

	writer.AddString("aÿz")
	assert.Equal(t, toBytes("ayz"), writer.ToByteArray())
}

func TestWriterAddFixedSanitizedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.SanitizeStrings = true

	writer.AddFixedString("aÿz", 3)
	assert.Equal(t, toBytes("ayz"), writer.ToByteArray())
}

func TestWriterAddPaddedSanitizedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.SanitizeStrings = true

	writer.AddPaddedString("aÿz", 6)
	assert.Equal(t, toBytes("ayzÿÿÿ"), writer.ToByteArray())
}

func TestWriterAddEncodedSanitizedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.SanitizeStrings = true

	writer.AddEncodedString("aÿz")
	assert.Equal(t, toBytes("S&l"), writer.ToByteArray())
}

func TestWriterAddFixedEncodedSanitizedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.SanitizeStrings = true

	writer.AddFixedEncodedString("aÿz", 3)
	assert.Equal(t, toBytes("S&l"), writer.ToByteArray())
}

func TestWriterAddPaddedEncodedSanitizedString(t *testing.T) {
	writer := data.NewEoWriter()
	writer.SanitizeStrings = true

	writer.AddPaddedEncodedString("aÿz", 6)
	assert.Equal(t, toBytes("ÿÿÿ%T>"), writer.ToByteArray())
}

func TestWriterAddNumbersOnBoundary(t *testing.T) {
	writer := data.NewEoWriter()
	var err error

	err = writer.AddByte(255)
	assert.NoError(t, err)

	err = writer.AddChar(data.CHAR_MAX - 1)
	assert.NoError(t, err)

	err = writer.AddShort(data.SHORT_MAX - 1)
	assert.NoError(t, err)

	err = writer.AddThree(data.THREE_MAX - 1)
	assert.NoError(t, err)

	err = writer.AddInt(data.INT_MAX - 1)
	assert.NoError(t, err)
}

func TestWriterAddNumbersExceedingLimit(t *testing.T) {
	writer := data.NewEoWriter()
	var err error

	err = writer.AddByte(256)
	assert.ErrorContains(t, err, "value is larger than")

	err = writer.AddChar(data.CHAR_MAX)
	assert.ErrorContains(t, err, "value is larger than")

	err = writer.AddShort(data.SHORT_MAX)
	assert.ErrorContains(t, err, "value is larger than")

	err = writer.AddThree(data.THREE_MAX)
	assert.ErrorContains(t, err, "value is larger than")

	err = writer.AddInt(data.INT_MAX)
	assert.ErrorContains(t, err, "value is larger than")
}

func TestWriterAddFixedStringWithIncorrectLength(t *testing.T) {
	writer := data.NewEoWriter()
	var err error

	err = writer.AddFixedString("foo", 2)
	assert.ErrorContains(t, err, "does not have expected length")

	err = writer.AddPaddedString("foo", 2)
	assert.ErrorContains(t, err, "too large for length")

	err = writer.AddFixedString("foo", 4)
	assert.ErrorContains(t, err, "does not have expected length")

	err = writer.AddFixedEncodedString("foo", 2)
	assert.ErrorContains(t, err, "does not have expected length")

	err = writer.AddPaddedEncodedString("foo", 2)
	assert.ErrorContains(t, err, "too large for length")

	err = writer.AddFixedEncodedString("foo", 4)
	assert.ErrorContains(t, err, "does not have expected length")
}

func TestWriterLength(t *testing.T) {
	writer := data.NewEoWriter()
	assert.Equal(t, 0, writer.Length())

	writer.AddString("Lorem ipsum dolor sit amet")
	assert.Equal(t, 26, writer.Length())

	for i := 27; i <= 100; i++ {
		writer.AddByte(0xFF)
	}
	assert.Equal(t, 100, writer.Length())
}
