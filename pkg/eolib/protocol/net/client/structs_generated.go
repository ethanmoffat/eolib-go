package client

import (
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
)

// ByteCoords :: Map coordinates with raw 1-byte values.
type ByteCoords struct {
	byteSize int

	X int
	Y int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ByteCoords) ByteSize() int {
	return s.byteSize
}

func (s *ByteCoords) Serialize(writer *data.EoWriter) (err error) {
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

func (s *ByteCoords) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// X : field : byte
	s.X = int(reader.GetByte())
	// Y : field : byte
	s.Y = int(reader.GetByte())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// WalkAction :: Common data between walk packets.
type WalkAction struct {
	byteSize int

	Direction protocol.Direction
	Timestamp int
	Coords    protocol.Coords
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *WalkAction) ByteSize() int {
	return s.byteSize
}

func (s *WalkAction) Serialize(writer *data.EoWriter) (err error) {
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

func (s *WalkAction) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// Timestamp : field : three
	s.Timestamp = reader.GetThree()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}
