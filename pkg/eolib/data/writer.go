package data

import (
	"errors"
	"strconv"
)

// EoWriter encapsulates operations related to writing EO data to a sequence of bytes.
type EoWriter struct {
	data []byte

	// SanitizeStrings gets or sets the sanitization mode for the writer.
	// With sanitization enabled, the writer will switch 0xFF bytes in strings (ÿ) to 0x79 (y)
	// See [Chunked Reading: Sanitization]
	//
	// [Chunked Reading: Sanitization]: https://github.com/Cirras/eo-protocol/blob/master/docs/chunks.md#sanitization
	SanitizeStrings bool
}

// NewEoWriter initializes an empty [data.EoWriter]
func NewEoWriter() *EoWriter {
	return &EoWriter{make([]byte, 16), false}
}

// Write satisifies the io.Writer interface
//
// Write writes len(p) bytes from p to the writer data.
// This implementation unconditionally writes all bytes in the input slice and returns len(p) and a nil error.
func (w *EoWriter) Write(p []byte) (int, error) {
	w.AddBytes(p)
	return len(p), nil
}

// AddInt adds a raw byte to the writer data.
func (w *EoWriter) AddByte(value int) error {
	if value > 0xFF {
		return errors.New("value is larger than maximum raw byte size")
	}

	w.data = append(w.data, byte(value))

	return nil
}

// AddBytes adds the specified bytes to the writer data.
func (w *EoWriter) AddBytes(bytes []byte) error {
	// append for slices is amortized to O(1) and automatically expands capacity of the underlying array as needed
	w.data = append(w.data, bytes...)
	return nil
}

// AddChar adds an encoded 1-byte integer to the writer data.
func (w *EoWriter) AddChar(number int) error {
	if number > CHAR_MAX-1 {
		return errors.New("value is larger than one byte maximum")
	}

	bytes := EncodeNumber(number)
	w.AddBytes(bytes[:1])

	return nil
}

// AddShort adds an encoded 2-byte integer to the writer data.
func (w *EoWriter) AddShort(number int) error {
	if number > SHORT_MAX-1 {
		return errors.New("value is larger than two byte maximum")
	}

	bytes := EncodeNumber(number)
	w.AddBytes(bytes[:2])

	return nil
}

// AddThree adds an encoded 3-byte integer to the writer data.
func (w *EoWriter) AddThree(number int) error {
	if number > THREE_MAX-1 {
		return errors.New("value is larger than three byte maximum")
	}

	bytes := EncodeNumber(number)
	w.AddBytes(bytes[:3])

	return nil
}

// AddInt adds an encoded 4-byte integer to the writer data.
func (w *EoWriter) AddInt(number int) error {
	if number > INT_MAX-1 {
		return errors.New("value is larger than four byte maximum")
	}

	bytes := EncodeNumber(number)
	w.AddBytes(bytes[:4])

	return nil
}

// AddString adds a string to the writer data.
func (w *EoWriter) AddString(str string) error {
	return w.AddBytes(w.sanitize([]byte(str)))
}

// AddFixedString adds a fixed-length string to the writer data.
func (w *EoWriter) AddFixedString(str string, length int) (err error) {
	if err = w.checkLength(str, length, false); err == nil {
		w.AddBytes(w.sanitize([]byte(str)))
	}
	return
}

// AddPaddedString adds a fixed-length string to the writer data add adds trailing padding (0xFF) bytes.
func (w *EoWriter) AddPaddedString(str string, length int) (err error) {
	if err = w.checkLength(str, length, true); err == nil {
		w.AddBytes(w.addPadding(w.sanitize([]byte(str)), length))
	}
	return
}

// AddEncodedString encodes and adds a string to the writer data.
func (w *EoWriter) AddEncodedString(str string) error {
	sanitized := w.sanitize([]byte(str))
	return w.AddBytes(EncodeString(string(sanitized)))
}

// AddFixedEncodedString encodes and adds a fixed-length string to the writer data.
func (w *EoWriter) AddFixedEncodedString(str string, length int) (err error) {
	if err = w.checkLength(str, length, false); err == nil {
		sanitized := w.sanitize([]byte(str))
		w.AddBytes(EncodeString(string(sanitized)))
	}
	return
}

// AddPaddedEncodedString encodes and adds a fixed-length string to the writer data and adds trailing padding (0xFF) bytes.
func (w *EoWriter) AddPaddedEncodedString(str string, length int) (err error) {
	if err = w.checkLength(str, length, true); err == nil {
		sanitized := w.sanitize([]byte(str))
		padded := w.addPadding(sanitized, length)
		w.AddBytes(EncodeString(string(padded)))
	}
	return
}

// Length gets the length of the writer data.
func (w *EoWriter) Length() int {
	return len(w.data)
}

// ToByteArray gets the writer data as a byte array.
func (w *EoWriter) ToByteArray() []byte {
	ret := make([]byte, len(w.data))
	copy(ret, w.data)
	return ret
}

func (w *EoWriter) sanitize(bytes []byte) (output []byte) {
	if !w.SanitizeStrings {
		return bytes
	}

	output = make([]byte, len(bytes))
	copy(output, bytes)

	for i, b := range output {
		if b == 0xFF {
			output[i] = 0x79
		}
	}

	return
}

func (w *EoWriter) addPadding(bytes []byte, length int) (output []byte) {
	output = make([]byte, length)
	copy(output, bytes)

	if length == len(bytes) {
		return
	}

	for i := len(bytes); i < length; i++ {
		output[i] = 0xFF
	}

	return
}

func (w *EoWriter) checkLength(str string, length int, padded bool) error {
	if padded {
		if length < len(str) {
			return errors.New("padded string " + str + " is too large for length " + strconv.Itoa(length))
		}
	} else if length != len(str) {
		return errors.New("String " + str + " does not have expected length " + strconv.Itoa(length))
	}

	return nil
}
