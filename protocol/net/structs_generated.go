package net

import "github.com/ethanmoffat/eolib-go/v3/data"

// Version :: Client version.
type Version struct {
	byteSize int

	Major int
	Minor int
	Patch int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *Version) ByteSize() int {
	return s.byteSize
}

func (s *Version) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Major : field : char
	if err = writer.AddChar(s.Major); err != nil {
		return
	}
	// Minor : field : char
	if err = writer.AddChar(s.Minor); err != nil {
		return
	}
	// Patch : field : char
	if err = writer.AddChar(s.Patch); err != nil {
		return
	}
	return
}

func (s *Version) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Major : field : char
	s.Major = reader.GetChar()
	// Minor : field : char
	s.Minor = reader.GetChar()
	// Patch : field : char
	s.Patch = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// Weight :: Current carry weight and maximum carry capacity of a player.
type Weight struct {
	byteSize int

	Current int
	Max     int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *Weight) ByteSize() int {
	return s.byteSize
}

func (s *Weight) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Current : field : char
	if err = writer.AddChar(s.Current); err != nil {
		return
	}
	// Max : field : char
	if err = writer.AddChar(s.Max); err != nil {
		return
	}
	return
}

func (s *Weight) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Current : field : char
	s.Current = reader.GetChar()
	// Max : field : char
	s.Max = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// Item :: An item reference with a 4-byte amount.
type Item struct {
	byteSize int

	Id     int
	Amount int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *Item) ByteSize() int {
	return s.byteSize
}

func (s *Item) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// Amount : field : int
	if err = writer.AddInt(s.Amount); err != nil {
		return
	}
	return
}

func (s *Item) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Id : field : short
	s.Id = reader.GetShort()
	// Amount : field : int
	s.Amount = reader.GetInt()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ThreeItem ::  An item reference with a 3-byte amount. Used for shops, lockers, and various item transfers.
type ThreeItem struct {
	byteSize int

	Id     int
	Amount int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ThreeItem) ByteSize() int {
	return s.byteSize
}

func (s *ThreeItem) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// Amount : field : three
	if err = writer.AddThree(s.Amount); err != nil {
		return
	}
	return
}

func (s *ThreeItem) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Id : field : short
	s.Id = reader.GetShort()
	// Amount : field : three
	s.Amount = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharItem ::  An item reference with a 1-byte amount. Used for craft ingredients.
type CharItem struct {
	byteSize int

	Id     int
	Amount int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharItem) ByteSize() int {
	return s.byteSize
}

func (s *CharItem) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// Amount : field : char
	if err = writer.AddChar(s.Amount); err != nil {
		return
	}
	return
}

func (s *CharItem) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Id : field : short
	s.Id = reader.GetShort()
	// Amount : field : char
	s.Amount = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// Spell :: A spell known by the player.
type Spell struct {
	byteSize int

	Id    int
	Level int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *Spell) ByteSize() int {
	return s.byteSize
}

func (s *Spell) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// Level : field : short
	if err = writer.AddShort(s.Level); err != nil {
		return
	}
	return
}

func (s *Spell) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Id : field : short
	s.Id = reader.GetShort()
	// Level : field : short
	s.Level = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}
