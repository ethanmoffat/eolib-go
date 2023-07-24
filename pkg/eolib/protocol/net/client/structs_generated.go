package client

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	protocol "github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
)

// Ensure fmt import is referenced in generated code
var _ = fmt.Printf

// ByteCoords :: Map coordinates with raw 1-byte values.
type ByteCoords struct {
	X int
	Y int
}

func (s *ByteCoords) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// X : field : byte
	if err = writer.AddByte(s.X); err != nil {
		return
	}

	// Y : field : byte
	if err = writer.AddByte(s.Y); err != nil {
		return
	}

	return
}

func (s *ByteCoords) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// X : field : byte
	s.X = int(reader.GetByte())
	// Y : field : byte
	s.Y = int(reader.GetByte())

	return
}

// WalkAction :: Common data between walk packets.
type WalkAction struct {
	Direction protocol.Direction
	Timestamp int
	Coords    protocol.Coords
}

func (s *WalkAction) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}

	// Timestamp : field : three
	if err = writer.AddThree(s.Timestamp); err != nil {
		return
	}

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *WalkAction) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}

	return
}
