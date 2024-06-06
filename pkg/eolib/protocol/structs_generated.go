package protocol

import "github.com/ethanmoffat/eolib-go/pkg/eolib/data"

// Coords :: Map coordinates.
type Coords struct {
	byteSize int

	X int
	Y int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *Coords) ByteSize() int {
	return s.byteSize
}

func (s *Coords) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// X : field : char
	if err = writer.AddChar(s.X); err != nil {
		return
	}
	// Y : field : char
	if err = writer.AddChar(s.Y); err != nil {
		return
	}
	return
}

func (s *Coords) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// X : field : char
	s.X = reader.GetChar()
	// Y : field : char
	s.Y = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}
