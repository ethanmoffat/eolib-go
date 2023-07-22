package data

import (
	"errors"
	"io"
	"math"

	"github.com/ethanmoffat/eolib-go/pkg/eolib"
)

type chunkProperties struct {
	isChunked  bool
	chunkStart int
	nextBreak  int
}

// EoReader encapsulates operations related to reading EO data from a sequence of bytes.
//
// EoReader features a "chunked" reading mode, which is important for accurate emulation of the official game client.
//
// See [chunked reading] for more information.
//
// [chunked reading]: https://github.com/Cirras/eo-protocol/blob/master/docs/chunks.md
type EoReader struct {
	data      []byte
	pos       int
	chunkInfo chunkProperties
}

// NewEoReader initializes an [data.EoReader] with the data in the specified byte slice.
func NewEoReader(data []byte) *EoReader {
	return &EoReader{data, 0, chunkProperties{false, 0, -1}}
}

// Read satisfies the io.Reader interface.
//
// Read will read up to len(b) bytes into b. It returns the number of bytes read (0 <= n <= len(b)) and any error encountered.
//
// If Read returns n < len(b), it may use all of b as scratch space during the call.
// If some data is available but not len(b) bytes, Read returns what is available instead of waiting for more.
//
// When Read encounters an error or end-of-file condition after successfully reading n > 0 bytes, it returns the number of bytes read.
// In the end-of-file condition, it returns the (non-nil) error and n == 0 from a subsequent call.
func (r *EoReader) Read(b []byte) (n int, err error) {
	if len(b) == 0 {
		return 0, errors.New("input is nil or empty")
	}

	if r.Remaining() <= 0 {
		return 0, io.EOF
	}

	n = 0
	var nextByte byte
	for nextByte = r.GetByte(); n < len(b) && r.Remaining() > 0; n++ {
		b[n] = nextByte
	}

	return
}

// Seek satisfies the io.Seeker interface.
//
// Seek sets the current position of the reader to the specified offset, based on the specified value of whence.
// - io.SeekStart will set the absolute position of the reader relative to the start of the data.
// - io.SeekCurrent will set a relative position of the reader based on the current read position.
// - io.SeekEnd will set an absolute position from the end of the file (negative offset).
//
// Seek returns the new position of the reader and/or any error that occurred while seeking.
func (r *EoReader) Seek(offset int64, whence int) (i int64, err error) {
	if offset > math.MaxInt32 {
		return 0, errors.New("offset is greater than maximum supported value")
	}

	switch whence {
	case io.SeekStart:
		r.pos = int(offset)
	case io.SeekCurrent:
		r.pos += int(offset)
	case io.SeekEnd:
		r.pos = len(r.data) - 1 + int(offset)
	}

	i = int64(r.pos)

	if r.pos > len(r.data) {
		err = errors.New("attempted to read past the end of available data")
	}

	return
}

// GetByte reads a raw byte from the input data.
func (r *EoReader) GetByte() byte {
	return r.readByte()
}

// GetBytes reads an array of raw bytes from the input data.
func (r *EoReader) GetBytes(length int) []byte {
	return r.readBytes(length)
}

// GetChar reads an encoded 1-byte integer from the input data.
func (r *EoReader) GetChar() int {
	return DecodeNumber(r.readBytes(1))
}

// GetShort reads an encoded 2-byte integer from the input data.
func (r *EoReader) GetShort() int {
	return DecodeNumber(r.readBytes(2))
}

// GetThree reads an encoded 3-byte integer from the input data.
func (r *EoReader) GetThree() int {
	return DecodeNumber(r.readBytes(3))
}

// GetInt reads an encoded 4-byte integer from the input data.
func (r *EoReader) GetInt() int {
	return DecodeNumber(r.readBytes(4))
}

// GetString reads an unencoded string from the input data.
func (r *EoReader) GetString() (string, error) {
	return string(r.readBytes(r.Remaining())), nil
}

// GetFixedString reads an unencoded fixed string from the input data.
func (r *EoReader) GetFixedString(length int) (string, error) {
	if length < 0 {
		return "", errors.New("negative length")
	}

	return string(r.readBytes(length)), nil
}

// GetPaddedString reads an unencoded fixed string from the input data and removes trailing padding bytes (0xFF value).
func (r *EoReader) GetPaddedString(length int) (string, error) {
	if length < 0 {
		return "", errors.New("negative length")
	}

	bytes := r.removePadding(r.readBytes(length))

	return string(bytes), nil
}

// GetEncodedString reads and decodes an encoded string from the input data.
func (r *EoReader) GetEncodedString() (string, error) {
	return windows1252String(DecodeString(r.readBytes(r.Remaining()))), nil
}

// GetFixedEncodedString reads and decodes a fixed string from the input data.
func (r *EoReader) GetFixedEncodedString(length int) (string, error) {
	if length < 0 {
		return "", errors.New("negative length")
	}

	return windows1252String(DecodeString(r.readBytes(length))), nil
}

// GetPaddedEncodedString reads and decodes a fixed string from the input data and removes trailing padding bytes (0xFF value).
func (r *EoReader) GetPaddedEncodedString(length int) (string, error) {
	if length < 0 {
		return "", errors.New("negative length")
	}

	decoded := DecodeString(r.readBytes(length))
	bytes := r.removePadding([]byte(decoded))

	return string(bytes), nil
}

// GetChunkedReadingMode gets whether chunked reading is enabled for the reader.
func (r *EoReader) GetChunkedReadingMode() bool {
	return r.chunkInfo.isChunked
}

// SetChunkedReadingMode sets whether chunked reading is enabled for the reader.
// In chunked reading mode:
// - The reader will treat 0xFF bytes as the end of the current chunk.
// - [EoReader.NextChunk] can be called to move to the next chunk.
func (r *EoReader) SetChunkedReadingMode(value bool) {
	r.chunkInfo.isChunked = value
	if r.chunkInfo.nextBreak == -1 {
		r.chunkInfo.nextBreak = r.findNextBreakIndex()
	}
}

// GetRemaining returns the number of bytes remaining in the input data.
//
// If chunked reading mode is enabled, gets the number of bytes remaining in the current chunk.
// Otherwise, gets the total number of bytes remaining in the input data.
func (r *EoReader) Remaining() int {
	if r.chunkInfo.isChunked {
		return r.chunkInfo.nextBreak - r.pos
	} else {
		return len(r.data) - r.pos
	}
}

// NextChunk moves the reader position to the start of the next chunk in the input data.
// An error is returned if the reader is not in chunked reading mode.
func (r *EoReader) NextChunk() error {
	if !r.chunkInfo.isChunked {
		return errors.New("not in chunked reading mode")
	}

	r.pos = r.chunkInfo.nextBreak
	if r.pos < len(r.data) {
		r.pos++
	}

	r.chunkInfo.chunkStart = r.pos
	r.chunkInfo.nextBreak = r.findNextBreakIndex()

	return nil
}

// Position gets the current position in the input data.
func (r *EoReader) Position() int {
	return r.pos
}

// Length gets the length of the input data.
func (r *EoReader) Length() int {
	return len(r.data)
}

func (r *EoReader) readByte() byte {
	if r.Remaining() > 0 {
		defer r.Seek(1, io.SeekCurrent)
		return r.data[r.pos]
	}

	return 0
}

func (r *EoReader) readBytes(length int) []byte {
	length = eolib.Min(length, r.Remaining())

	defer r.Seek(int64(length), io.SeekCurrent)
	return r.data[r.pos : r.pos+length]
}

func (r *EoReader) removePadding(input []byte) []byte {
	for i, b := range input {
		if b == 0xFF {
			return input[:i]
		}
	}

	return input
}

func (r *EoReader) findNextBreakIndex() int {
	var i int
	for i = r.chunkInfo.chunkStart; i < len(r.data); i++ {
		if r.data[i] == 0xFF {
			break
		}
	}
	return i
}
