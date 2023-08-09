package net

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
)

// Ensure fmt import is referenced in generated code
var _ = fmt.Printf

// Version :: Client version.
type Version struct {
	Major int
	Minor int
	Patch int
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

	// Major : field : char
	s.Major = reader.GetChar()
	// Minor : field : char
	s.Minor = reader.GetChar()
	// Patch : field : char
	s.Patch = reader.GetChar()

	return
}

// Weight :: Current carry weight and maximum carry capacity of a player.
type Weight struct {
	Current int
	Max     int
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

	// Current : field : char
	s.Current = reader.GetChar()
	// Max : field : char
	s.Max = reader.GetChar()

	return
}

// Item :: An item reference with a 4-byte amount.
type Item struct {
	Id     int
	Amount int
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

	// Id : field : short
	s.Id = reader.GetShort()
	// Amount : field : int
	s.Amount = reader.GetInt()

	return
}

// ThreeItem ::  An item reference with a 3-byte amount. Used for shops, lockers, and various item transfers.
type ThreeItem struct {
	Id     int
	Amount int
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

	// Id : field : short
	s.Id = reader.GetShort()
	// Amount : field : three
	s.Amount = reader.GetThree()

	return
}

// CharItem ::  An item reference with a 1-byte amount. Used for craft ingredients.
type CharItem struct {
	Id     int
	Amount int
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

	// Id : field : short
	s.Id = reader.GetShort()
	// Amount : field : char
	s.Amount = reader.GetChar()

	return
}

// Spell :: A spell known by the player.
type Spell struct {
	Id    int
	Level int
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

	// Id : field : short
	s.Id = reader.GetShort()
	// Level : field : short
	s.Level = reader.GetShort()

	return
}
