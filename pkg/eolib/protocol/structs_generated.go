package protocol

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
)

// Ensure fmt import is referenced in generated code
var _ = fmt.Printf

// Coords :: Map coordinates.
type Coords struct {
	X int
	Y int
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

	// X : field : char
	s.X = reader.GetChar()
	// Y : field : char
	s.Y = reader.GetChar()

	return
}
