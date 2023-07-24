package pub

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
)

// Ensure fmt import is referenced in generated code
var _ = fmt.Printf

// EifRecord :: Record of Item data in an Endless Item File.
type EifRecord struct {
	NameLength       int
	Name             string
	GraphicId        int
	Type             ItemType
	Subtype          ItemSubtype
	Special          ItemSpecial
	Hp               int
	Tp               int
	MinDamage        int
	MaxDamage        int
	Accuracy         int
	Evade            int
	Armor            int
	ReturnDamage     int
	Str              int
	Intl             int
	Wis              int
	Agi              int
	Con              int
	Cha              int
	LightResistance  int
	DarkResistance   int
	EarthResistance  int
	AirResistance    int
	WaterResistance  int
	FireResistance   int
	Spec1            int //  Holds one the following values, depending on item type:. scroll_map, doll_graphic, exp_reward, hair_color, effect, key, alcohol_potency.
	Spec2            int //  Holds one the following values, depending on item type:. scroll_x, gender.
	Spec3            int //  Holds one the following values, depending on item type:. scroll_y.
	LevelRequirement int
	ClassRequirement int
	StrRequirement   int
	IntRequirement   int
	WisRequirement   int
	AgiRequirement   int
	ConRequirement   int
	ChaRequirement   int
	Element          Element
	ElementDamage    int
	Weight           int

	Size ItemSize
}

func (s *EifRecord) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NameLength : length : char
	if err = writer.AddChar(s.NameLength); err != nil {
		return
	}

	// Name : field : string
	if err = writer.AddFixedString(s.Name, s.NameLength); err != nil {
		return
	}

	// GraphicId : field : short
	if err = writer.AddShort(s.GraphicId); err != nil {
		return
	}

	// Type : field : ItemType
	if err = writer.AddChar(int(s.Type)); err != nil {
		return
	}

	// Subtype : field : ItemSubtype
	if err = writer.AddChar(int(s.Subtype)); err != nil {
		return
	}

	// Special : field : ItemSpecial
	if err = writer.AddChar(int(s.Special)); err != nil {
		return
	}

	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}

	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}

	// MinDamage : field : short
	if err = writer.AddShort(s.MinDamage); err != nil {
		return
	}

	// MaxDamage : field : short
	if err = writer.AddShort(s.MaxDamage); err != nil {
		return
	}

	// Accuracy : field : short
	if err = writer.AddShort(s.Accuracy); err != nil {
		return
	}

	// Evade : field : short
	if err = writer.AddShort(s.Evade); err != nil {
		return
	}

	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}

	// ReturnDamage : field : char
	if err = writer.AddChar(s.ReturnDamage); err != nil {
		return
	}

	// Str : field : char
	if err = writer.AddChar(s.Str); err != nil {
		return
	}

	// Intl : field : char
	if err = writer.AddChar(s.Intl); err != nil {
		return
	}

	// Wis : field : char
	if err = writer.AddChar(s.Wis); err != nil {
		return
	}

	// Agi : field : char
	if err = writer.AddChar(s.Agi); err != nil {
		return
	}

	// Con : field : char
	if err = writer.AddChar(s.Con); err != nil {
		return
	}

	// Cha : field : char
	if err = writer.AddChar(s.Cha); err != nil {
		return
	}

	// LightResistance : field : char
	if err = writer.AddChar(s.LightResistance); err != nil {
		return
	}

	// DarkResistance : field : char
	if err = writer.AddChar(s.DarkResistance); err != nil {
		return
	}

	// EarthResistance : field : char
	if err = writer.AddChar(s.EarthResistance); err != nil {
		return
	}

	// AirResistance : field : char
	if err = writer.AddChar(s.AirResistance); err != nil {
		return
	}

	// WaterResistance : field : char
	if err = writer.AddChar(s.WaterResistance); err != nil {
		return
	}

	// FireResistance : field : char
	if err = writer.AddChar(s.FireResistance); err != nil {
		return
	}

	// Spec1 : field : three
	if err = writer.AddThree(s.Spec1); err != nil {
		return
	}

	// Spec2 : field : char
	if err = writer.AddChar(s.Spec2); err != nil {
		return
	}

	// Spec3 : field : char
	if err = writer.AddChar(s.Spec3); err != nil {
		return
	}

	// LevelRequirement : field : short
	if err = writer.AddShort(s.LevelRequirement); err != nil {
		return
	}

	// ClassRequirement : field : short
	if err = writer.AddShort(s.ClassRequirement); err != nil {
		return
	}

	// StrRequirement : field : short
	if err = writer.AddShort(s.StrRequirement); err != nil {
		return
	}

	// IntRequirement : field : short
	if err = writer.AddShort(s.IntRequirement); err != nil {
		return
	}

	// WisRequirement : field : short
	if err = writer.AddShort(s.WisRequirement); err != nil {
		return
	}

	// AgiRequirement : field : short
	if err = writer.AddShort(s.AgiRequirement); err != nil {
		return
	}

	// ConRequirement : field : short
	if err = writer.AddShort(s.ConRequirement); err != nil {
		return
	}

	// ChaRequirement : field : short
	if err = writer.AddShort(s.ChaRequirement); err != nil {
		return
	}

	// Element : field : Element
	if err = writer.AddChar(int(s.Element)); err != nil {
		return
	}

	// ElementDamage : field : char
	if err = writer.AddChar(s.ElementDamage); err != nil {
		return
	}

	// Weight : field : char
	if err = writer.AddChar(s.Weight); err != nil {
		return
	}

	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}

	// Size : field : ItemSize
	if err = writer.AddChar(int(s.Size)); err != nil {
		return
	}

	return
}

func (s *EifRecord) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NameLength : length : char
	s.NameLength = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetFixedString(s.NameLength); err != nil {
		return
	}

	// GraphicId : field : short
	s.GraphicId = reader.GetShort()
	// Type : field : ItemType
	s.Type = ItemType(reader.GetChar())
	// Subtype : field : ItemSubtype
	s.Subtype = ItemSubtype(reader.GetChar())
	// Special : field : ItemSpecial
	s.Special = ItemSpecial(reader.GetChar())
	// Hp : field : short
	s.Hp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// MinDamage : field : short
	s.MinDamage = reader.GetShort()
	// MaxDamage : field : short
	s.MaxDamage = reader.GetShort()
	// Accuracy : field : short
	s.Accuracy = reader.GetShort()
	// Evade : field : short
	s.Evade = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// ReturnDamage : field : char
	s.ReturnDamage = reader.GetChar()
	// Str : field : char
	s.Str = reader.GetChar()
	// Intl : field : char
	s.Intl = reader.GetChar()
	// Wis : field : char
	s.Wis = reader.GetChar()
	// Agi : field : char
	s.Agi = reader.GetChar()
	// Con : field : char
	s.Con = reader.GetChar()
	// Cha : field : char
	s.Cha = reader.GetChar()
	// LightResistance : field : char
	s.LightResistance = reader.GetChar()
	// DarkResistance : field : char
	s.DarkResistance = reader.GetChar()
	// EarthResistance : field : char
	s.EarthResistance = reader.GetChar()
	// AirResistance : field : char
	s.AirResistance = reader.GetChar()
	// WaterResistance : field : char
	s.WaterResistance = reader.GetChar()
	// FireResistance : field : char
	s.FireResistance = reader.GetChar()
	// Spec1 : field : three
	s.Spec1 = reader.GetThree()
	// Spec2 : field : char
	s.Spec2 = reader.GetChar()
	// Spec3 : field : char
	s.Spec3 = reader.GetChar()
	// LevelRequirement : field : short
	s.LevelRequirement = reader.GetShort()
	// ClassRequirement : field : short
	s.ClassRequirement = reader.GetShort()
	// StrRequirement : field : short
	s.StrRequirement = reader.GetShort()
	// IntRequirement : field : short
	s.IntRequirement = reader.GetShort()
	// WisRequirement : field : short
	s.WisRequirement = reader.GetShort()
	// AgiRequirement : field : short
	s.AgiRequirement = reader.GetShort()
	// ConRequirement : field : short
	s.ConRequirement = reader.GetShort()
	// ChaRequirement : field : short
	s.ChaRequirement = reader.GetShort()
	// Element : field : Element
	s.Element = Element(reader.GetChar())
	// ElementDamage : field : char
	s.ElementDamage = reader.GetChar()
	// Weight : field : char
	s.Weight = reader.GetChar()
	//  : field : char
	reader.GetChar()
	// Size : field : ItemSize
	s.Size = ItemSize(reader.GetChar())

	return
}

// Eif :: Endless Item File.
type Eif struct {
	Rid             []int
	TotalItemsCount int
	Version         int
	Items           []EifRecord
}

func (s *Eif) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddFixedString("EIF", 3); err != nil {
		return
	}

	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.Rid[ndx]); err != nil {
			return
		}

	}

	// TotalItemsCount : field : short
	if err = writer.AddShort(s.TotalItemsCount); err != nil {
		return
	}

	// Version : field : char
	if err = writer.AddChar(s.Version); err != nil {
		return
	}

	// Items : array : EifRecord
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *Eif) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Rid[ndx] = reader.GetShort()
	}

	// TotalItemsCount : field : short
	s.TotalItemsCount = reader.GetShort()
	// Version : field : char
	s.Version = reader.GetChar()
	// Items : array : EifRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Items = append(s.Items, EifRecord{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// EnfRecord :: Record of NPC data in an Endless NPC File.
type EnfRecord struct {
	NameLength            int
	Name                  string
	GraphicId             int
	Race                  int
	Boss                  bool
	Child                 bool
	Type                  NpcType
	BehaviorId            int
	Hp                    int
	Tp                    int
	MinDamage             int
	MaxDamage             int
	Accuracy              int
	Evade                 int
	Armor                 int
	ReturnDamage          int
	Element               Element
	ElementDamage         int
	ElementWeakness       Element
	ElementWeaknessDamage int
	Level                 int
	Experience            int
}

func (s *EnfRecord) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NameLength : length : char
	if err = writer.AddChar(s.NameLength); err != nil {
		return
	}

	// Name : field : string
	if err = writer.AddFixedString(s.Name, s.NameLength); err != nil {
		return
	}

	// GraphicId : field : short
	if err = writer.AddShort(s.GraphicId); err != nil {
		return
	}

	// Race : field : char
	if err = writer.AddChar(s.Race); err != nil {
		return
	}

	// Boss : field : bool:short
	if s.Boss {
		err = writer.AddShort(1)
	} else {
		err = writer.AddShort(0)
	}
	if err != nil {
		return
	}

	// Child : field : bool:short
	if s.Child {
		err = writer.AddShort(1)
	} else {
		err = writer.AddShort(0)
	}
	if err != nil {
		return
	}

	// Type : field : NpcType
	if err = writer.AddShort(int(s.Type)); err != nil {
		return
	}

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}

	// Hp : field : three
	if err = writer.AddThree(s.Hp); err != nil {
		return
	}

	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}

	// MinDamage : field : short
	if err = writer.AddShort(s.MinDamage); err != nil {
		return
	}

	// MaxDamage : field : short
	if err = writer.AddShort(s.MaxDamage); err != nil {
		return
	}

	// Accuracy : field : short
	if err = writer.AddShort(s.Accuracy); err != nil {
		return
	}

	// Evade : field : short
	if err = writer.AddShort(s.Evade); err != nil {
		return
	}

	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}

	// ReturnDamage : field : char
	if err = writer.AddChar(s.ReturnDamage); err != nil {
		return
	}

	// Element : field : Element:short
	if err = writer.AddChar(int(s.Element)); err != nil {
		return
	}

	// ElementDamage : field : short
	if err = writer.AddShort(s.ElementDamage); err != nil {
		return
	}

	// ElementWeakness : field : Element:short
	if err = writer.AddChar(int(s.ElementWeakness)); err != nil {
		return
	}

	// ElementWeaknessDamage : field : short
	if err = writer.AddShort(s.ElementWeaknessDamage); err != nil {
		return
	}

	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}

	// Experience : field : three
	if err = writer.AddThree(s.Experience); err != nil {
		return
	}

	return
}

func (s *EnfRecord) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NameLength : length : char
	s.NameLength = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetFixedString(s.NameLength); err != nil {
		return
	}

	// GraphicId : field : short
	s.GraphicId = reader.GetShort()
	// Race : field : char
	s.Race = reader.GetChar()
	// Boss : field : bool:short
	if boolVal := reader.GetShort(); boolVal > 0 {
		s.Boss = true
	} else {
		s.Boss = false
	}
	// Child : field : bool:short
	if boolVal := reader.GetShort(); boolVal > 0 {
		s.Child = true
	} else {
		s.Child = false
	}
	// Type : field : NpcType
	s.Type = NpcType(reader.GetShort())
	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	// Hp : field : three
	s.Hp = reader.GetThree()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// MinDamage : field : short
	s.MinDamage = reader.GetShort()
	// MaxDamage : field : short
	s.MaxDamage = reader.GetShort()
	// Accuracy : field : short
	s.Accuracy = reader.GetShort()
	// Evade : field : short
	s.Evade = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// ReturnDamage : field : char
	s.ReturnDamage = reader.GetChar()
	// Element : field : Element:short
	s.Element = Element(reader.GetChar())
	// ElementDamage : field : short
	s.ElementDamage = reader.GetShort()
	// ElementWeakness : field : Element:short
	s.ElementWeakness = Element(reader.GetChar())
	// ElementWeaknessDamage : field : short
	s.ElementWeaknessDamage = reader.GetShort()
	// Level : field : char
	s.Level = reader.GetChar()
	// Experience : field : three
	s.Experience = reader.GetThree()

	return
}

// Enf :: Endless NPC File.
type Enf struct {
	Rid            []int
	TotalNpcsCount int
	Version        int
	Npcs           []EnfRecord
}

func (s *Enf) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddFixedString("ENF", 3); err != nil {
		return
	}

	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.Rid[ndx]); err != nil {
			return
		}

	}

	// TotalNpcsCount : field : short
	if err = writer.AddShort(s.TotalNpcsCount); err != nil {
		return
	}

	// Version : field : char
	if err = writer.AddChar(s.Version); err != nil {
		return
	}

	// Npcs : array : EnfRecord
	for ndx := 0; ndx < len(s.Npcs); ndx++ {
		if err = s.Npcs[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *Enf) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Rid[ndx] = reader.GetShort()
	}

	// TotalNpcsCount : field : short
	s.TotalNpcsCount = reader.GetShort()
	// Version : field : char
	s.Version = reader.GetChar()
	// Npcs : array : EnfRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Npcs = append(s.Npcs, EnfRecord{})
		if err = s.Npcs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// EcfRecord :: Record of Class data in an Endless Class File.
type EcfRecord struct {
	NameLength int
	Name       string
	ParentType int
	StatGroup  int
	Str        int
	Intl       int
	Wis        int
	Agi        int
	Con        int
	Cha        int
}

func (s *EcfRecord) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NameLength : length : char
	if err = writer.AddChar(s.NameLength); err != nil {
		return
	}

	// Name : field : string
	if err = writer.AddFixedString(s.Name, s.NameLength); err != nil {
		return
	}

	// ParentType : field : char
	if err = writer.AddChar(s.ParentType); err != nil {
		return
	}

	// StatGroup : field : char
	if err = writer.AddChar(s.StatGroup); err != nil {
		return
	}

	// Str : field : short
	if err = writer.AddShort(s.Str); err != nil {
		return
	}

	// Intl : field : short
	if err = writer.AddShort(s.Intl); err != nil {
		return
	}

	// Wis : field : short
	if err = writer.AddShort(s.Wis); err != nil {
		return
	}

	// Agi : field : short
	if err = writer.AddShort(s.Agi); err != nil {
		return
	}

	// Con : field : short
	if err = writer.AddShort(s.Con); err != nil {
		return
	}

	// Cha : field : short
	if err = writer.AddShort(s.Cha); err != nil {
		return
	}

	return
}

func (s *EcfRecord) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NameLength : length : char
	s.NameLength = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetFixedString(s.NameLength); err != nil {
		return
	}

	// ParentType : field : char
	s.ParentType = reader.GetChar()
	// StatGroup : field : char
	s.StatGroup = reader.GetChar()
	// Str : field : short
	s.Str = reader.GetShort()
	// Intl : field : short
	s.Intl = reader.GetShort()
	// Wis : field : short
	s.Wis = reader.GetShort()
	// Agi : field : short
	s.Agi = reader.GetShort()
	// Con : field : short
	s.Con = reader.GetShort()
	// Cha : field : short
	s.Cha = reader.GetShort()

	return
}

// Ecf :: Endless Class File.
type Ecf struct {
	Rid               []int
	TotalClassesCount int
	Version           int
	Classes           []EcfRecord
}

func (s *Ecf) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddFixedString("ECF", 3); err != nil {
		return
	}

	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.Rid[ndx]); err != nil {
			return
		}

	}

	// TotalClassesCount : field : short
	if err = writer.AddShort(s.TotalClassesCount); err != nil {
		return
	}

	// Version : field : char
	if err = writer.AddChar(s.Version); err != nil {
		return
	}

	// Classes : array : EcfRecord
	for ndx := 0; ndx < len(s.Classes); ndx++ {
		if err = s.Classes[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *Ecf) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Rid[ndx] = reader.GetShort()
	}

	// TotalClassesCount : field : short
	s.TotalClassesCount = reader.GetShort()
	// Version : field : char
	s.Version = reader.GetChar()
	// Classes : array : EcfRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Classes = append(s.Classes, EcfRecord{})
		if err = s.Classes[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// EsfRecord :: Record of Skill data in an Endless Skill File.
type EsfRecord struct {
	NameLength  int
	ChantLength int
	Name        string
	Chant       string
	IconId      int
	GraphicId   int
	TpCost      int
	SpCost      int
	CastTime    int
	Nature      SkillNature

	Type           SkillType
	Element        Element
	ElementPower   int
	TargetRestrict SkillTargetRestrict
	TargetType     SkillTargetType
	TargetTime     int

	MaxSkillLevel int
	MinDamage     int
	MaxDamage     int
	Accuracy      int
	Evade         int
	Armor         int
	ReturnDamage  int
	HpHeal        int
	TpHeal        int
	SpHeal        int
	Str           int
	Intl          int
	Wis           int
	Agi           int
	Con           int
	Cha           int
}

func (s *EsfRecord) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NameLength : length : char
	if err = writer.AddChar(s.NameLength); err != nil {
		return
	}

	// ChantLength : length : char
	if err = writer.AddChar(s.ChantLength); err != nil {
		return
	}

	// Name : field : string
	if err = writer.AddFixedString(s.Name, s.NameLength); err != nil {
		return
	}

	// Chant : field : string
	if err = writer.AddFixedString(s.Chant, s.ChantLength); err != nil {
		return
	}

	// IconId : field : short
	if err = writer.AddShort(s.IconId); err != nil {
		return
	}

	// GraphicId : field : short
	if err = writer.AddShort(s.GraphicId); err != nil {
		return
	}

	// TpCost : field : short
	if err = writer.AddShort(s.TpCost); err != nil {
		return
	}

	// SpCost : field : short
	if err = writer.AddShort(s.SpCost); err != nil {
		return
	}

	// CastTime : field : char
	if err = writer.AddChar(s.CastTime); err != nil {
		return
	}

	// Nature : field : SkillNature
	if err = writer.AddChar(int(s.Nature)); err != nil {
		return
	}

	//  : field : char
	if err = writer.AddChar(1); err != nil {
		return
	}

	// Type : field : SkillType
	if err = writer.AddThree(int(s.Type)); err != nil {
		return
	}

	// Element : field : Element
	if err = writer.AddChar(int(s.Element)); err != nil {
		return
	}

	// ElementPower : field : short
	if err = writer.AddShort(s.ElementPower); err != nil {
		return
	}

	// TargetRestrict : field : SkillTargetRestrict
	if err = writer.AddChar(int(s.TargetRestrict)); err != nil {
		return
	}

	// TargetType : field : SkillTargetType
	if err = writer.AddChar(int(s.TargetType)); err != nil {
		return
	}

	// TargetTime : field : char
	if err = writer.AddChar(s.TargetTime); err != nil {
		return
	}

	//  : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}

	// MaxSkillLevel : field : short
	if err = writer.AddShort(s.MaxSkillLevel); err != nil {
		return
	}

	// MinDamage : field : short
	if err = writer.AddShort(s.MinDamage); err != nil {
		return
	}

	// MaxDamage : field : short
	if err = writer.AddShort(s.MaxDamage); err != nil {
		return
	}

	// Accuracy : field : short
	if err = writer.AddShort(s.Accuracy); err != nil {
		return
	}

	// Evade : field : short
	if err = writer.AddShort(s.Evade); err != nil {
		return
	}

	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}

	// ReturnDamage : field : char
	if err = writer.AddChar(s.ReturnDamage); err != nil {
		return
	}

	// HpHeal : field : short
	if err = writer.AddShort(s.HpHeal); err != nil {
		return
	}

	// TpHeal : field : short
	if err = writer.AddShort(s.TpHeal); err != nil {
		return
	}

	// SpHeal : field : char
	if err = writer.AddChar(s.SpHeal); err != nil {
		return
	}

	// Str : field : short
	if err = writer.AddShort(s.Str); err != nil {
		return
	}

	// Intl : field : short
	if err = writer.AddShort(s.Intl); err != nil {
		return
	}

	// Wis : field : short
	if err = writer.AddShort(s.Wis); err != nil {
		return
	}

	// Agi : field : short
	if err = writer.AddShort(s.Agi); err != nil {
		return
	}

	// Con : field : short
	if err = writer.AddShort(s.Con); err != nil {
		return
	}

	// Cha : field : short
	if err = writer.AddShort(s.Cha); err != nil {
		return
	}

	return
}

func (s *EsfRecord) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NameLength : length : char
	s.NameLength = reader.GetChar()
	// ChantLength : length : char
	s.ChantLength = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetFixedString(s.NameLength); err != nil {
		return
	}

	// Chant : field : string
	if s.Chant, err = reader.GetFixedString(s.ChantLength); err != nil {
		return
	}

	// IconId : field : short
	s.IconId = reader.GetShort()
	// GraphicId : field : short
	s.GraphicId = reader.GetShort()
	// TpCost : field : short
	s.TpCost = reader.GetShort()
	// SpCost : field : short
	s.SpCost = reader.GetShort()
	// CastTime : field : char
	s.CastTime = reader.GetChar()
	// Nature : field : SkillNature
	s.Nature = SkillNature(reader.GetChar())
	//  : field : char
	reader.GetChar()
	// Type : field : SkillType
	s.Type = SkillType(reader.GetThree())
	// Element : field : Element
	s.Element = Element(reader.GetChar())
	// ElementPower : field : short
	s.ElementPower = reader.GetShort()
	// TargetRestrict : field : SkillTargetRestrict
	s.TargetRestrict = SkillTargetRestrict(reader.GetChar())
	// TargetType : field : SkillTargetType
	s.TargetType = SkillTargetType(reader.GetChar())
	// TargetTime : field : char
	s.TargetTime = reader.GetChar()
	//  : field : char
	reader.GetChar()
	// MaxSkillLevel : field : short
	s.MaxSkillLevel = reader.GetShort()
	// MinDamage : field : short
	s.MinDamage = reader.GetShort()
	// MaxDamage : field : short
	s.MaxDamage = reader.GetShort()
	// Accuracy : field : short
	s.Accuracy = reader.GetShort()
	// Evade : field : short
	s.Evade = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// ReturnDamage : field : char
	s.ReturnDamage = reader.GetChar()
	// HpHeal : field : short
	s.HpHeal = reader.GetShort()
	// TpHeal : field : short
	s.TpHeal = reader.GetShort()
	// SpHeal : field : char
	s.SpHeal = reader.GetChar()
	// Str : field : short
	s.Str = reader.GetShort()
	// Intl : field : short
	s.Intl = reader.GetShort()
	// Wis : field : short
	s.Wis = reader.GetShort()
	// Agi : field : short
	s.Agi = reader.GetShort()
	// Con : field : short
	s.Con = reader.GetShort()
	// Cha : field : short
	s.Cha = reader.GetShort()

	return
}

// Esf :: Endless Skill File.
type Esf struct {
	Rid              []int
	TotalSkillsCount int
	Version          int
	Skills           []EsfRecord
}

func (s *Esf) Serialize(writer data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	//  : field : string
	if err = writer.AddFixedString("ESF", 3); err != nil {
		return
	}

	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if err = writer.AddShort(s.Rid[ndx]); err != nil {
			return
		}

	}

	// TotalSkillsCount : field : short
	if err = writer.AddShort(s.TotalSkillsCount); err != nil {
		return
	}

	// Version : field : char
	if err = writer.AddChar(s.Version); err != nil {
		return
	}

	// Skills : array : EsfRecord
	for ndx := 0; ndx < len(s.Skills); ndx++ {
		if err = s.Skills[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *Esf) Deserialize(reader data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	//  : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Rid[ndx] = reader.GetShort()
	}

	// TotalSkillsCount : field : short
	s.TotalSkillsCount = reader.GetShort()
	// Version : field : char
	s.Version = reader.GetChar()
	// Skills : array : EsfRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Skills = append(s.Skills, EsfRecord{})
		if err = s.Skills[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}
