package server

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/v3/data"
	"github.com/ethanmoffat/eolib-go/v3/protocol"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net"
)

// BigCoords :: Map coordinates with 2-byte values.
type BigCoords struct {
	byteSize int

	X int
	Y int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BigCoords) ByteSize() int {
	return s.byteSize
}

func (s *BigCoords) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// X : field : short
	if err = writer.AddShort(s.X); err != nil {
		return
	}
	// Y : field : short
	if err = writer.AddShort(s.Y); err != nil {
		return
	}
	return
}

func (s *BigCoords) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// X : field : short
	s.X = reader.GetShort()
	// Y : field : short
	s.Y = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EquipmentChange ::  Player equipment data. Sent when a player's visible equipment changes. Note that these values are graphic IDs.
type EquipmentChange struct {
	byteSize int

	Boots  int
	Armor  int
	Hat    int
	Weapon int
	Shield int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EquipmentChange) ByteSize() int {
	return s.byteSize
}

func (s *EquipmentChange) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Boots : field : short
	if err = writer.AddShort(s.Boots); err != nil {
		return
	}
	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}
	// Hat : field : short
	if err = writer.AddShort(s.Hat); err != nil {
		return
	}
	// Weapon : field : short
	if err = writer.AddShort(s.Weapon); err != nil {
		return
	}
	// Shield : field : short
	if err = writer.AddShort(s.Shield); err != nil {
		return
	}
	return
}

func (s *EquipmentChange) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Boots : field : short
	s.Boots = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// Hat : field : short
	s.Hat = reader.GetShort()
	// Weapon : field : short
	s.Weapon = reader.GetShort()
	// Shield : field : short
	s.Shield = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EquipmentMapInfo ::  Player equipment data. Sent with map information about a nearby character. Note that these values are graphic IDs.
type EquipmentMapInfo struct {
	byteSize int

	Boots int

	Armor int

	Hat    int
	Shield int
	Weapon int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EquipmentMapInfo) ByteSize() int {
	return s.byteSize
}

func (s *EquipmentMapInfo) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Boots : field : short
	if err = writer.AddShort(s.Boots); err != nil {
		return
	}
	// 0 : field : short
	if err = writer.AddShort(0); err != nil {
		return
	}
	// 0 : field : short
	if err = writer.AddShort(0); err != nil {
		return
	}
	// 0 : field : short
	if err = writer.AddShort(0); err != nil {
		return
	}
	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}
	// 0 : field : short
	if err = writer.AddShort(0); err != nil {
		return
	}
	// Hat : field : short
	if err = writer.AddShort(s.Hat); err != nil {
		return
	}
	// Shield : field : short
	if err = writer.AddShort(s.Shield); err != nil {
		return
	}
	// Weapon : field : short
	if err = writer.AddShort(s.Weapon); err != nil {
		return
	}
	return
}

func (s *EquipmentMapInfo) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Boots : field : short
	s.Boots = reader.GetShort()
	// 0 : field : short
	reader.GetShort()
	// 0 : field : short
	reader.GetShort()
	// 0 : field : short
	reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// 0 : field : short
	reader.GetShort()
	// Hat : field : short
	s.Hat = reader.GetShort()
	// Shield : field : short
	s.Shield = reader.GetShort()
	// Weapon : field : short
	s.Weapon = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EquipmentCharacterSelect ::  Player equipment data. Sent with a character in the character selection list. Note that these values are graphic IDs.
type EquipmentCharacterSelect struct {
	byteSize int

	Boots  int
	Armor  int
	Hat    int
	Shield int
	Weapon int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EquipmentCharacterSelect) ByteSize() int {
	return s.byteSize
}

func (s *EquipmentCharacterSelect) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Boots : field : short
	if err = writer.AddShort(s.Boots); err != nil {
		return
	}
	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}
	// Hat : field : short
	if err = writer.AddShort(s.Hat); err != nil {
		return
	}
	// Shield : field : short
	if err = writer.AddShort(s.Shield); err != nil {
		return
	}
	// Weapon : field : short
	if err = writer.AddShort(s.Weapon); err != nil {
		return
	}
	return
}

func (s *EquipmentCharacterSelect) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Boots : field : short
	s.Boots = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// Hat : field : short
	s.Hat = reader.GetShort()
	// Shield : field : short
	s.Shield = reader.GetShort()
	// Weapon : field : short
	s.Weapon = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EquipmentWelcome ::  Player equipment data. Sent upon selecting a character and entering the game. Note that these values are item IDs.
type EquipmentWelcome struct {
	byteSize int

	Boots     int
	Gloves    int
	Accessory int
	Armor     int
	Belt      int
	Necklace  int
	Hat       int
	Shield    int
	Weapon    int
	Ring      []int
	Armlet    []int
	Bracer    []int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EquipmentWelcome) ByteSize() int {
	return s.byteSize
}

func (s *EquipmentWelcome) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Boots : field : short
	if err = writer.AddShort(s.Boots); err != nil {
		return
	}
	// Gloves : field : short
	if err = writer.AddShort(s.Gloves); err != nil {
		return
	}
	// Accessory : field : short
	if err = writer.AddShort(s.Accessory); err != nil {
		return
	}
	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}
	// Belt : field : short
	if err = writer.AddShort(s.Belt); err != nil {
		return
	}
	// Necklace : field : short
	if err = writer.AddShort(s.Necklace); err != nil {
		return
	}
	// Hat : field : short
	if err = writer.AddShort(s.Hat); err != nil {
		return
	}
	// Shield : field : short
	if err = writer.AddShort(s.Shield); err != nil {
		return
	}
	// Weapon : field : short
	if err = writer.AddShort(s.Weapon); err != nil {
		return
	}
	// Ring : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if len(s.Ring) != 2 {
			err = fmt.Errorf("expected Ring with length 2, got %d", len(s.Ring))
			return
		}

		if err = writer.AddShort(s.Ring[ndx]); err != nil {
			return
		}
	}

	// Armlet : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if len(s.Armlet) != 2 {
			err = fmt.Errorf("expected Armlet with length 2, got %d", len(s.Armlet))
			return
		}

		if err = writer.AddShort(s.Armlet[ndx]); err != nil {
			return
		}
	}

	// Bracer : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if len(s.Bracer) != 2 {
			err = fmt.Errorf("expected Bracer with length 2, got %d", len(s.Bracer))
			return
		}

		if err = writer.AddShort(s.Bracer[ndx]); err != nil {
			return
		}
	}

	return
}

func (s *EquipmentWelcome) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Boots : field : short
	s.Boots = reader.GetShort()
	// Gloves : field : short
	s.Gloves = reader.GetShort()
	// Accessory : field : short
	s.Accessory = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// Belt : field : short
	s.Belt = reader.GetShort()
	// Necklace : field : short
	s.Necklace = reader.GetShort()
	// Hat : field : short
	s.Hat = reader.GetShort()
	// Shield : field : short
	s.Shield = reader.GetShort()
	// Weapon : field : short
	s.Weapon = reader.GetShort()
	// Ring : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Ring = append(s.Ring, 0)
		s.Ring[ndx] = reader.GetShort()
	}

	// Armlet : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Armlet = append(s.Armlet, 0)
		s.Armlet[ndx] = reader.GetShort()
	}

	// Bracer : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Bracer = append(s.Bracer, 0)
		s.Bracer[ndx] = reader.GetShort()
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// EquipmentPaperdoll ::  Player equipment data. Sent with information about a player's paperdoll. Note that these values are item IDs.
type EquipmentPaperdoll struct {
	byteSize int

	Boots     int
	Accessory int
	Gloves    int
	Belt      int
	Armor     int
	Necklace  int
	Hat       int
	Shield    int
	Weapon    int
	Ring      []int
	Armlet    []int
	Bracer    []int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EquipmentPaperdoll) ByteSize() int {
	return s.byteSize
}

func (s *EquipmentPaperdoll) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Boots : field : short
	if err = writer.AddShort(s.Boots); err != nil {
		return
	}
	// Accessory : field : short
	if err = writer.AddShort(s.Accessory); err != nil {
		return
	}
	// Gloves : field : short
	if err = writer.AddShort(s.Gloves); err != nil {
		return
	}
	// Belt : field : short
	if err = writer.AddShort(s.Belt); err != nil {
		return
	}
	// Armor : field : short
	if err = writer.AddShort(s.Armor); err != nil {
		return
	}
	// Necklace : field : short
	if err = writer.AddShort(s.Necklace); err != nil {
		return
	}
	// Hat : field : short
	if err = writer.AddShort(s.Hat); err != nil {
		return
	}
	// Shield : field : short
	if err = writer.AddShort(s.Shield); err != nil {
		return
	}
	// Weapon : field : short
	if err = writer.AddShort(s.Weapon); err != nil {
		return
	}
	// Ring : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if len(s.Ring) != 2 {
			err = fmt.Errorf("expected Ring with length 2, got %d", len(s.Ring))
			return
		}

		if err = writer.AddShort(s.Ring[ndx]); err != nil {
			return
		}
	}

	// Armlet : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if len(s.Armlet) != 2 {
			err = fmt.Errorf("expected Armlet with length 2, got %d", len(s.Armlet))
			return
		}

		if err = writer.AddShort(s.Armlet[ndx]); err != nil {
			return
		}
	}

	// Bracer : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if len(s.Bracer) != 2 {
			err = fmt.Errorf("expected Bracer with length 2, got %d", len(s.Bracer))
			return
		}

		if err = writer.AddShort(s.Bracer[ndx]); err != nil {
			return
		}
	}

	return
}

func (s *EquipmentPaperdoll) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Boots : field : short
	s.Boots = reader.GetShort()
	// Accessory : field : short
	s.Accessory = reader.GetShort()
	// Gloves : field : short
	s.Gloves = reader.GetShort()
	// Belt : field : short
	s.Belt = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	// Necklace : field : short
	s.Necklace = reader.GetShort()
	// Hat : field : short
	s.Hat = reader.GetShort()
	// Shield : field : short
	s.Shield = reader.GetShort()
	// Weapon : field : short
	s.Weapon = reader.GetShort()
	// Ring : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Ring = append(s.Ring, 0)
		s.Ring[ndx] = reader.GetShort()
	}

	// Armlet : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Armlet = append(s.Armlet, 0)
		s.Armlet[ndx] = reader.GetShort()
	}

	// Bracer : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Bracer = append(s.Bracer, 0)
		s.Bracer[ndx] = reader.GetShort()
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterMapInfo ::  Information about a nearby character. The official client skips these if they're under 42 bytes in length.
type CharacterMapInfo struct {
	byteSize int

	Name       string
	PlayerId   int
	MapId      int
	Coords     BigCoords
	Direction  protocol.Direction
	ClassId    int
	GuildTag   string
	Level      int
	Gender     protocol.Gender
	HairStyle  int
	HairColor  int
	Skin       int
	MaxHp      int
	Hp         int
	MaxTp      int
	Tp         int
	Equipment  EquipmentMapInfo
	SitState   SitState
	Invisible  bool
	WarpEffect *WarpEffect
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterMapInfo) ByteSize() int {
	return s.byteSize
}

func (s *CharacterMapInfo) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// MapId : field : short
	if err = writer.AddShort(s.MapId); err != nil {
		return
	}
	// Coords : field : BigCoords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	// ClassId : field : char
	if err = writer.AddChar(s.ClassId); err != nil {
		return
	}
	// GuildTag : field : string
	if len(s.GuildTag) != 3 {
		err = fmt.Errorf("expected GuildTag with length 3, got %d", len(s.GuildTag))
		return
	}
	if err = writer.AddFixedString(s.GuildTag, 3); err != nil {
		return
	}
	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}
	// Gender : field : Gender
	if err = writer.AddChar(int(s.Gender)); err != nil {
		return
	}
	// HairStyle : field : char
	if err = writer.AddChar(s.HairStyle); err != nil {
		return
	}
	// HairColor : field : char
	if err = writer.AddChar(s.HairColor); err != nil {
		return
	}
	// Skin : field : char
	if err = writer.AddChar(s.Skin); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}
	// Equipment : field : EquipmentMapInfo
	if err = s.Equipment.Serialize(writer); err != nil {
		return
	}
	// SitState : field : SitState
	if err = writer.AddChar(int(s.SitState)); err != nil {
		return
	}
	// Invisible : field : bool
	if s.Invisible {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	// WarpEffect : field : WarpEffect
	if s.WarpEffect != nil {
		if err = writer.AddChar(int(*s.WarpEffect)); err != nil {
			return
		}
	}
	writer.SanitizeStrings = false
	return
}

func (s *CharacterMapInfo) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// MapId : field : short
	s.MapId = reader.GetShort()
	// Coords : field : BigCoords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// ClassId : field : char
	s.ClassId = reader.GetChar()
	// GuildTag : field : string
	if s.GuildTag, err = reader.GetFixedString(3); err != nil {
		return
	}

	// Level : field : char
	s.Level = reader.GetChar()
	// Gender : field : Gender
	s.Gender = protocol.Gender(reader.GetChar())
	// HairStyle : field : char
	s.HairStyle = reader.GetChar()
	// HairColor : field : char
	s.HairColor = reader.GetChar()
	// Skin : field : char
	s.Skin = reader.GetChar()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// Equipment : field : EquipmentMapInfo
	if err = s.Equipment.Deserialize(reader); err != nil {
		return
	}
	// SitState : field : SitState
	s.SitState = SitState(reader.GetChar())
	// Invisible : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Invisible = true
	} else {
		s.Invisible = false
	}
	// WarpEffect : field : WarpEffect
	if reader.Remaining() > 0 {
		s.WarpEffect = new(WarpEffect)
		*s.WarpEffect = WarpEffect(reader.GetChar())
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcMapInfo :: Information about a nearby NPC.
type NpcMapInfo struct {
	byteSize int

	Index     int
	Id        int
	Coords    protocol.Coords
	Direction protocol.Direction
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcMapInfo) ByteSize() int {
	return s.byteSize
}

func (s *NpcMapInfo) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Index : field : char
	if err = writer.AddChar(s.Index); err != nil {
		return
	}
	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *NpcMapInfo) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Index : field : char
	s.Index = reader.GetChar()
	// Id : field : short
	s.Id = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ItemMapInfo :: Information about a nearby item on the ground.
type ItemMapInfo struct {
	byteSize int

	Uid    int
	Id     int
	Coords protocol.Coords
	Amount int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ItemMapInfo) ByteSize() int {
	return s.byteSize
}

func (s *ItemMapInfo) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Uid : field : short
	if err = writer.AddShort(s.Uid); err != nil {
		return
	}
	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Amount : field : three
	if err = writer.AddThree(s.Amount); err != nil {
		return
	}
	return
}

func (s *ItemMapInfo) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Uid : field : short
	s.Uid = reader.GetShort()
	// Id : field : short
	s.Id = reader.GetShort()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Amount : field : three
	s.Amount = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// AvatarChange :: Information about a nearby player's appearance changing.
type AvatarChange struct {
	byteSize int

	PlayerId       int
	ChangeType     AvatarChangeType
	Sound          bool
	ChangeTypeData ChangeTypeData
}

type ChangeTypeData interface {
	protocol.EoData
}

type ChangeTypeDataEquipment struct {
	byteSize int

	Equipment EquipmentChange
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChangeTypeDataEquipment) ByteSize() int {
	return s.byteSize
}

func (s *ChangeTypeDataEquipment) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Equipment : field : EquipmentChange
	if err = s.Equipment.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *ChangeTypeDataEquipment) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Equipment : field : EquipmentChange
	if err = s.Equipment.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type ChangeTypeDataHair struct {
	byteSize int

	HairStyle int
	HairColor int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChangeTypeDataHair) ByteSize() int {
	return s.byteSize
}

func (s *ChangeTypeDataHair) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// HairStyle : field : char
	if err = writer.AddChar(s.HairStyle); err != nil {
		return
	}
	// HairColor : field : char
	if err = writer.AddChar(s.HairColor); err != nil {
		return
	}
	return
}

func (s *ChangeTypeDataHair) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// HairStyle : field : char
	s.HairStyle = reader.GetChar()
	// HairColor : field : char
	s.HairColor = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

type ChangeTypeDataHairColor struct {
	byteSize int

	HairColor int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ChangeTypeDataHairColor) ByteSize() int {
	return s.byteSize
}

func (s *ChangeTypeDataHairColor) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// HairColor : field : char
	if err = writer.AddChar(s.HairColor); err != nil {
		return
	}
	return
}

func (s *ChangeTypeDataHairColor) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// HairColor : field : char
	s.HairColor = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *AvatarChange) ByteSize() int {
	return s.byteSize
}

func (s *AvatarChange) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// ChangeType : field : AvatarChangeType
	if err = writer.AddChar(int(s.ChangeType)); err != nil {
		return
	}
	// Sound : field : bool
	if s.Sound {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	switch s.ChangeType {
	case AvatarChange_Equipment:
		switch s.ChangeTypeData.(type) {
		case *ChangeTypeDataEquipment:
			if err = s.ChangeTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ChangeType)
			return
		}
	case AvatarChange_Hair:
		switch s.ChangeTypeData.(type) {
		case *ChangeTypeDataHair:
			if err = s.ChangeTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ChangeType)
			return
		}
	case AvatarChange_HairColor:
		switch s.ChangeTypeData.(type) {
		case *ChangeTypeDataHairColor:
			if err = s.ChangeTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.ChangeType)
			return
		}
	}
	return
}

func (s *AvatarChange) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// ChangeType : field : AvatarChangeType
	s.ChangeType = AvatarChangeType(reader.GetChar())
	// Sound : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Sound = true
	} else {
		s.Sound = false
	}
	switch s.ChangeType {
	case AvatarChange_Equipment:
		s.ChangeTypeData = &ChangeTypeDataEquipment{}
		if err = s.ChangeTypeData.Deserialize(reader); err != nil {
			return
		}
	case AvatarChange_Hair:
		s.ChangeTypeData = &ChangeTypeDataHair{}
		if err = s.ChangeTypeData.Deserialize(reader); err != nil {
			return
		}
	case AvatarChange_HairColor:
		s.ChangeTypeData = &ChangeTypeDataHairColor{}
		if err = s.ChangeTypeData.Deserialize(reader); err != nil {
			return
		}
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NearbyInfo :: Information about nearby entities.
type NearbyInfo struct {
	byteSize int

	Characters []CharacterMapInfo
	Npcs       []NpcMapInfo
	Items      []ItemMapInfo
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NearbyInfo) ByteSize() int {
	return s.byteSize
}

func (s *NearbyInfo) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// CharactersCount : length : char
	if err = writer.AddChar(len(s.Characters)); err != nil {
		return
	}
	writer.SanitizeStrings = true
	writer.AddByte(255)
	// Characters : array : CharacterMapInfo
	for ndx := 0; ndx < len(s.Characters); ndx++ {
		if err = s.Characters[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	// Npcs : array : NpcMapInfo
	for ndx := 0; ndx < len(s.Npcs); ndx++ {
		if err = s.Npcs[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.AddByte(255)
	// Items : array : ItemMapInfo
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	writer.SanitizeStrings = false
	return
}

func (s *NearbyInfo) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// CharactersCount : length : char
	charactersCount := reader.GetChar()
	reader.SetIsChunked(true)
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Characters : array : CharacterMapInfo
	for ndx := 0; ndx < charactersCount; ndx++ {
		s.Characters = append(s.Characters, CharacterMapInfo{})
		if err = s.Characters[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	// Npcs : array : NpcMapInfo
	NpcsRemaining := reader.Remaining()
	for ndx := 0; ndx < NpcsRemaining/6; ndx++ {
		s.Npcs = append(s.Npcs, NpcMapInfo{})
		if err = s.Npcs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Items : array : ItemMapInfo
	ItemsRemaining := reader.Remaining()
	for ndx := 0; ndx < ItemsRemaining/9; ndx++ {
		s.Items = append(s.Items, ItemMapInfo{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapFile :: A map file (EMF).
type MapFile struct {
	byteSize int

	Content []byte
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapFile) ByteSize() int {
	return s.byteSize
}

func (s *MapFile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Content : field : blob
	if err = writer.AddBytes(s.Content); err != nil {
		return
	}
	return
}

func (s *MapFile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Content : field : blob
	s.Content = reader.GetBytes(reader.Remaining())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PubFile :: A pub file (EIF, ENF, ECF, ESF).
type PubFile struct {
	byteSize int

	FileId  int
	Content []byte
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PubFile) ByteSize() int {
	return s.byteSize
}

func (s *PubFile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// FileId : field : char
	if err = writer.AddChar(s.FileId); err != nil {
		return
	}
	// Content : field : blob
	if err = writer.AddBytes(s.Content); err != nil {
		return
	}
	return
}

func (s *PubFile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// FileId : field : char
	s.FileId = reader.GetChar()
	// Content : field : blob
	s.Content = reader.GetBytes(reader.Remaining())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// OnlinePlayer :: A player in the online list.
type OnlinePlayer struct {
	byteSize int

	Name     string
	Title    string
	Level    int
	Icon     CharacterIcon
	ClassId  int
	GuildTag string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *OnlinePlayer) ByteSize() int {
	return s.byteSize
}

func (s *OnlinePlayer) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// Title : field : string
	if err = writer.AddString(s.Title); err != nil {
		return
	}
	writer.AddByte(255)
	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}
	// Icon : field : CharacterIcon
	if err = writer.AddChar(int(s.Icon)); err != nil {
		return
	}
	// ClassId : field : char
	if err = writer.AddChar(s.ClassId); err != nil {
		return
	}
	// GuildTag : field : string
	if err = writer.AddString(s.GuildTag); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *OnlinePlayer) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Title : field : string
	if s.Title, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Level : field : char
	s.Level = reader.GetChar()
	// Icon : field : CharacterIcon
	s.Icon = CharacterIcon(reader.GetChar())
	// ClassId : field : char
	s.ClassId = reader.GetChar()
	// GuildTag : field : string
	if s.GuildTag, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersList :: Information about online players.
type PlayersList struct {
	byteSize int

	Players []OnlinePlayer
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersList) ByteSize() int {
	return s.byteSize
}

func (s *PlayersList) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayersCount : length : short
	if err = writer.AddShort(len(s.Players)); err != nil {
		return
	}
	writer.AddByte(255)
	// Players : array : OnlinePlayer
	for ndx := 0; ndx < len(s.Players); ndx++ {
		if err = s.Players[ndx].Serialize(writer); err != nil {
			return
		}
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *PlayersList) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayersCount : length : short
	playersCount := reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Players : array : OnlinePlayer
	for ndx := 0; ndx < playersCount; ndx++ {
		s.Players = append(s.Players, OnlinePlayer{})
		if err = s.Players[ndx].Deserialize(reader); err != nil {
			return
		}
		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayersListFriends ::  Information about online players. Sent in reply to friends list requests.
type PlayersListFriends struct {
	byteSize int

	Players []string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayersListFriends) ByteSize() int {
	return s.byteSize
}

func (s *PlayersListFriends) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayersCount : length : short
	if err = writer.AddShort(len(s.Players)); err != nil {
		return
	}
	writer.AddByte(255)
	// Players : array : string
	for ndx := 0; ndx < len(s.Players); ndx++ {
		if err = writer.AddString(s.Players[ndx]); err != nil {
			return
		}
		writer.AddByte(255)
	}

	writer.SanitizeStrings = false
	return
}

func (s *PlayersListFriends) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayersCount : length : short
	playersCount := reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Players : array : string
	for ndx := 0; ndx < playersCount; ndx++ {
		s.Players = append(s.Players, "")
		if s.Players[ndx], err = reader.GetString(); err != nil {
			return
		}

		if err = reader.NextChunk(); err != nil {
			return
		}
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterSelectionListEntry :: Character selection screen character.
type CharacterSelectionListEntry struct {
	byteSize int

	Name      string
	Id        int
	Level     int
	Gender    protocol.Gender
	HairStyle int
	HairColor int
	Skin      int
	Admin     protocol.AdminLevel
	Equipment EquipmentCharacterSelect
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterSelectionListEntry) ByteSize() int {
	return s.byteSize
}

func (s *CharacterSelectionListEntry) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// Id : field : int
	if err = writer.AddInt(s.Id); err != nil {
		return
	}
	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}
	// Gender : field : Gender
	if err = writer.AddChar(int(s.Gender)); err != nil {
		return
	}
	// HairStyle : field : char
	if err = writer.AddChar(s.HairStyle); err != nil {
		return
	}
	// HairColor : field : char
	if err = writer.AddChar(s.HairColor); err != nil {
		return
	}
	// Skin : field : char
	if err = writer.AddChar(s.Skin); err != nil {
		return
	}
	// Admin : field : AdminLevel
	if err = writer.AddChar(int(s.Admin)); err != nil {
		return
	}
	// Equipment : field : EquipmentCharacterSelect
	if err = s.Equipment.Serialize(writer); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *CharacterSelectionListEntry) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Id : field : int
	s.Id = reader.GetInt()
	// Level : field : char
	s.Level = reader.GetChar()
	// Gender : field : Gender
	s.Gender = protocol.Gender(reader.GetChar())
	// HairStyle : field : char
	s.HairStyle = reader.GetChar()
	// HairColor : field : char
	s.HairColor = reader.GetChar()
	// Skin : field : char
	s.Skin = reader.GetChar()
	// Admin : field : AdminLevel
	s.Admin = protocol.AdminLevel(reader.GetChar())
	// Equipment : field : EquipmentCharacterSelect
	if err = s.Equipment.Deserialize(reader); err != nil {
		return
	}
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ServerSettings :: Settings sent with WELCOME_REPLY packet.
type ServerSettings struct {
	byteSize int

	JailMap                   int
	RescueMap                 int
	RescueCoords              protocol.Coords
	SpyAndLightGuideFloodRate int
	GuardianFloodRate         int
	GameMasterFloodRate       int
	HighGameMasterFloodRate   int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ServerSettings) ByteSize() int {
	return s.byteSize
}

func (s *ServerSettings) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// JailMap : field : short
	if err = writer.AddShort(s.JailMap); err != nil {
		return
	}
	// RescueMap : field : short
	if err = writer.AddShort(s.RescueMap); err != nil {
		return
	}
	// RescueCoords : field : Coords
	if err = s.RescueCoords.Serialize(writer); err != nil {
		return
	}
	// SpyAndLightGuideFloodRate : field : short
	if err = writer.AddShort(s.SpyAndLightGuideFloodRate); err != nil {
		return
	}
	// GuardianFloodRate : field : short
	if err = writer.AddShort(s.GuardianFloodRate); err != nil {
		return
	}
	// GameMasterFloodRate : field : short
	if err = writer.AddShort(s.GameMasterFloodRate); err != nil {
		return
	}
	// HighGameMasterFloodRate : field : short
	if err = writer.AddShort(s.HighGameMasterFloodRate); err != nil {
		return
	}
	return
}

func (s *ServerSettings) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// JailMap : field : short
	s.JailMap = reader.GetShort()
	// RescueMap : field : short
	s.RescueMap = reader.GetShort()
	// RescueCoords : field : Coords
	if err = s.RescueCoords.Deserialize(reader); err != nil {
		return
	}
	// SpyAndLightGuideFloodRate : field : short
	s.SpyAndLightGuideFloodRate = reader.GetShort()
	// GuardianFloodRate : field : short
	s.GuardianFloodRate = reader.GetShort()
	// GameMasterFloodRate : field : short
	s.GameMasterFloodRate = reader.GetShort()
	// HighGameMasterFloodRate : field : short
	s.HighGameMasterFloodRate = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopTradeItem :: An item that a shop can buy or sell.
type ShopTradeItem struct {
	byteSize int

	ItemId       int
	BuyPrice     int
	SellPrice    int
	MaxBuyAmount int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopTradeItem) ByteSize() int {
	return s.byteSize
}

func (s *ShopTradeItem) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// BuyPrice : field : three
	if err = writer.AddThree(s.BuyPrice); err != nil {
		return
	}
	// SellPrice : field : three
	if err = writer.AddThree(s.SellPrice); err != nil {
		return
	}
	// MaxBuyAmount : field : char
	if err = writer.AddChar(s.MaxBuyAmount); err != nil {
		return
	}
	return
}

func (s *ShopTradeItem) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// BuyPrice : field : three
	s.BuyPrice = reader.GetThree()
	// SellPrice : field : three
	s.SellPrice = reader.GetThree()
	// MaxBuyAmount : field : char
	s.MaxBuyAmount = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopCraftItem :: An item that a shop can craft.
type ShopCraftItem struct {
	byteSize int

	ItemId      int
	Ingredients []net.CharItem
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopCraftItem) ByteSize() int {
	return s.byteSize
}

func (s *ShopCraftItem) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// Ingredients : array : CharItem
	for ndx := 0; ndx < 4; ndx++ {
		if len(s.Ingredients) != 4 {
			err = fmt.Errorf("expected Ingredients with length 4, got %d", len(s.Ingredients))
			return
		}

		if err = s.Ingredients[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ShopCraftItem) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// Ingredients : array : CharItem
	for ndx := 0; ndx < 4; ndx++ {
		s.Ingredients = append(s.Ingredients, net.CharItem{})
		if err = s.Ingredients[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ShopSoldItem :: A sold item when selling an item to a shop.
type ShopSoldItem struct {
	byteSize int

	Amount int
	Id     int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *ShopSoldItem) ByteSize() int {
	return s.byteSize
}

func (s *ShopSoldItem) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Amount : field : int
	if err = writer.AddInt(s.Amount); err != nil {
		return
	}
	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	return
}

func (s *ShopSoldItem) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Amount : field : int
	s.Amount = reader.GetInt()
	// Id : field : short
	s.Id = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterBaseStats :: The 6 base character stats.
type CharacterBaseStats struct {
	byteSize int

	Str  int
	Intl int
	Wis  int
	Agi  int
	Con  int
	Cha  int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterBaseStats) ByteSize() int {
	return s.byteSize
}

func (s *CharacterBaseStats) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

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

func (s *CharacterBaseStats) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterBaseStatsWelcome ::  The 6 base character stats. Sent upon selecting a character and entering the game.
type CharacterBaseStatsWelcome struct {
	byteSize int

	Str  int
	Wis  int
	Intl int
	Agi  int
	Con  int
	Cha  int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterBaseStatsWelcome) ByteSize() int {
	return s.byteSize
}

func (s *CharacterBaseStatsWelcome) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Str : field : short
	if err = writer.AddShort(s.Str); err != nil {
		return
	}
	// Wis : field : short
	if err = writer.AddShort(s.Wis); err != nil {
		return
	}
	// Intl : field : short
	if err = writer.AddShort(s.Intl); err != nil {
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

func (s *CharacterBaseStatsWelcome) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Str : field : short
	s.Str = reader.GetShort()
	// Wis : field : short
	s.Wis = reader.GetShort()
	// Intl : field : short
	s.Intl = reader.GetShort()
	// Agi : field : short
	s.Agi = reader.GetShort()
	// Con : field : short
	s.Con = reader.GetShort()
	// Cha : field : short
	s.Cha = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterSecondaryStats :: The 5 secondary character stats.
type CharacterSecondaryStats struct {
	byteSize int

	MinDamage int
	MaxDamage int
	Accuracy  int
	Evade     int
	Armor     int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterSecondaryStats) ByteSize() int {
	return s.byteSize
}

func (s *CharacterSecondaryStats) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

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
	return
}

func (s *CharacterSecondaryStats) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
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
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterSecondaryStatsInfoLookup ::  The 5 secondary character stats. Sent with character info lookups.
type CharacterSecondaryStatsInfoLookup struct {
	byteSize int

	MaxDamage int
	MinDamage int
	Accuracy  int
	Evade     int
	Armor     int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterSecondaryStatsInfoLookup) ByteSize() int {
	return s.byteSize
}

func (s *CharacterSecondaryStatsInfoLookup) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MaxDamage : field : short
	if err = writer.AddShort(s.MaxDamage); err != nil {
		return
	}
	// MinDamage : field : short
	if err = writer.AddShort(s.MinDamage); err != nil {
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
	return
}

func (s *CharacterSecondaryStatsInfoLookup) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// MaxDamage : field : short
	s.MaxDamage = reader.GetShort()
	// MinDamage : field : short
	s.MinDamage = reader.GetShort()
	// Accuracy : field : short
	s.Accuracy = reader.GetShort()
	// Evade : field : short
	s.Evade = reader.GetShort()
	// Armor : field : short
	s.Armor = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterElementalStats :: The 6 elemental character stats.
type CharacterElementalStats struct {
	byteSize int

	Light int
	Dark  int
	Fire  int
	Water int
	Earth int
	Wind  int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterElementalStats) ByteSize() int {
	return s.byteSize
}

func (s *CharacterElementalStats) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Light : field : short
	if err = writer.AddShort(s.Light); err != nil {
		return
	}
	// Dark : field : short
	if err = writer.AddShort(s.Dark); err != nil {
		return
	}
	// Fire : field : short
	if err = writer.AddShort(s.Fire); err != nil {
		return
	}
	// Water : field : short
	if err = writer.AddShort(s.Water); err != nil {
		return
	}
	// Earth : field : short
	if err = writer.AddShort(s.Earth); err != nil {
		return
	}
	// Wind : field : short
	if err = writer.AddShort(s.Wind); err != nil {
		return
	}
	return
}

func (s *CharacterElementalStats) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Light : field : short
	s.Light = reader.GetShort()
	// Dark : field : short
	s.Dark = reader.GetShort()
	// Fire : field : short
	s.Fire = reader.GetShort()
	// Water : field : short
	s.Water = reader.GetShort()
	// Earth : field : short
	s.Earth = reader.GetShort()
	// Wind : field : short
	s.Wind = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterStatsReset ::  Character stats data. Sent when resetting stats and skills at a skill master NPC.
type CharacterStatsReset struct {
	byteSize int

	StatPoints  int
	SkillPoints int
	Hp          int
	MaxHp       int
	Tp          int
	MaxTp       int
	MaxSp       int
	Base        CharacterBaseStats
	Secondary   CharacterSecondaryStats
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterStatsReset) ByteSize() int {
	return s.byteSize
}

func (s *CharacterStatsReset) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// StatPoints : field : short
	if err = writer.AddShort(s.StatPoints); err != nil {
		return
	}
	// SkillPoints : field : short
	if err = writer.AddShort(s.SkillPoints); err != nil {
		return
	}
	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// MaxSp : field : short
	if err = writer.AddShort(s.MaxSp); err != nil {
		return
	}
	// Base : field : CharacterBaseStats
	if err = s.Base.Serialize(writer); err != nil {
		return
	}
	// Secondary : field : CharacterSecondaryStats
	if err = s.Secondary.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *CharacterStatsReset) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// StatPoints : field : short
	s.StatPoints = reader.GetShort()
	// SkillPoints : field : short
	s.SkillPoints = reader.GetShort()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// MaxSp : field : short
	s.MaxSp = reader.GetShort()
	// Base : field : CharacterBaseStats
	if err = s.Base.Deserialize(reader); err != nil {
		return
	}
	// Secondary : field : CharacterSecondaryStats
	if err = s.Secondary.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterStatsWelcome ::  Character stats data. Sent upon selecting a character and entering the game.
type CharacterStatsWelcome struct {
	byteSize int

	Hp          int
	MaxHp       int
	Tp          int
	MaxTp       int
	MaxSp       int
	StatPoints  int
	SkillPoints int
	Karma       int
	Secondary   CharacterSecondaryStats
	Base        CharacterBaseStatsWelcome
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterStatsWelcome) ByteSize() int {
	return s.byteSize
}

func (s *CharacterStatsWelcome) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// MaxSp : field : short
	if err = writer.AddShort(s.MaxSp); err != nil {
		return
	}
	// StatPoints : field : short
	if err = writer.AddShort(s.StatPoints); err != nil {
		return
	}
	// SkillPoints : field : short
	if err = writer.AddShort(s.SkillPoints); err != nil {
		return
	}
	// Karma : field : short
	if err = writer.AddShort(s.Karma); err != nil {
		return
	}
	// Secondary : field : CharacterSecondaryStats
	if err = s.Secondary.Serialize(writer); err != nil {
		return
	}
	// Base : field : CharacterBaseStatsWelcome
	if err = s.Base.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *CharacterStatsWelcome) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// MaxSp : field : short
	s.MaxSp = reader.GetShort()
	// StatPoints : field : short
	s.StatPoints = reader.GetShort()
	// SkillPoints : field : short
	s.SkillPoints = reader.GetShort()
	// Karma : field : short
	s.Karma = reader.GetShort()
	// Secondary : field : CharacterSecondaryStats
	if err = s.Secondary.Deserialize(reader); err != nil {
		return
	}
	// Base : field : CharacterBaseStatsWelcome
	if err = s.Base.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterStatsUpdate ::  Character stats data. Sent when stats are updated.
type CharacterStatsUpdate struct {
	byteSize int

	BaseStats      CharacterBaseStats
	MaxHp          int
	MaxTp          int
	MaxSp          int
	MaxWeight      int
	SecondaryStats CharacterSecondaryStats
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterStatsUpdate) ByteSize() int {
	return s.byteSize
}

func (s *CharacterStatsUpdate) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BaseStats : field : CharacterBaseStats
	if err = s.BaseStats.Serialize(writer); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// MaxSp : field : short
	if err = writer.AddShort(s.MaxSp); err != nil {
		return
	}
	// MaxWeight : field : short
	if err = writer.AddShort(s.MaxWeight); err != nil {
		return
	}
	// SecondaryStats : field : CharacterSecondaryStats
	if err = s.SecondaryStats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *CharacterStatsUpdate) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// BaseStats : field : CharacterBaseStats
	if err = s.BaseStats.Deserialize(reader); err != nil {
		return
	}
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// MaxSp : field : short
	s.MaxSp = reader.GetShort()
	// MaxWeight : field : short
	s.MaxWeight = reader.GetShort()
	// SecondaryStats : field : CharacterSecondaryStats
	if err = s.SecondaryStats.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterStatsInfoLookup ::  Character stats data. Sent with character info lookups.
type CharacterStatsInfoLookup struct {
	byteSize int

	Hp             int
	MaxHp          int
	Tp             int
	MaxTp          int
	BaseStats      CharacterBaseStats
	SecondaryStats CharacterSecondaryStatsInfoLookup
	ElementalStats CharacterElementalStats
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterStatsInfoLookup) ByteSize() int {
	return s.byteSize
}

func (s *CharacterStatsInfoLookup) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// Tp : field : short
	if err = writer.AddShort(s.Tp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// BaseStats : field : CharacterBaseStats
	if err = s.BaseStats.Serialize(writer); err != nil {
		return
	}
	// SecondaryStats : field : CharacterSecondaryStatsInfoLookup
	if err = s.SecondaryStats.Serialize(writer); err != nil {
		return
	}
	// ElementalStats : field : CharacterElementalStats
	if err = s.ElementalStats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *CharacterStatsInfoLookup) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Hp : field : short
	s.Hp = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// Tp : field : short
	s.Tp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// BaseStats : field : CharacterBaseStats
	if err = s.BaseStats.Deserialize(reader); err != nil {
		return
	}
	// SecondaryStats : field : CharacterSecondaryStatsInfoLookup
	if err = s.SecondaryStats.Deserialize(reader); err != nil {
		return
	}
	// ElementalStats : field : CharacterElementalStats
	if err = s.ElementalStats.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterStatsEquipmentChange ::  Character stats data. Sent when an item is equipped or unequipped.
type CharacterStatsEquipmentChange struct {
	byteSize int

	MaxHp          int
	MaxTp          int
	BaseStats      CharacterBaseStats
	SecondaryStats CharacterSecondaryStats
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterStatsEquipmentChange) ByteSize() int {
	return s.byteSize
}

func (s *CharacterStatsEquipmentChange) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// BaseStats : field : CharacterBaseStats
	if err = s.BaseStats.Serialize(writer); err != nil {
		return
	}
	// SecondaryStats : field : CharacterSecondaryStats
	if err = s.SecondaryStats.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *CharacterStatsEquipmentChange) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// BaseStats : field : CharacterBaseStats
	if err = s.BaseStats.Deserialize(reader); err != nil {
		return
	}
	// SecondaryStats : field : CharacterSecondaryStats
	if err = s.SecondaryStats.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// SkillLearn :: A skill that can be learned from a skill master NPC.
type SkillLearn struct {
	byteSize int

	Id                int
	LevelRequirement  int
	ClassRequirement  int
	Cost              int
	SkillRequirements []int
	StatRequirements  CharacterBaseStats
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *SkillLearn) ByteSize() int {
	return s.byteSize
}

func (s *SkillLearn) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// LevelRequirement : field : char
	if err = writer.AddChar(s.LevelRequirement); err != nil {
		return
	}
	// ClassRequirement : field : char
	if err = writer.AddChar(s.ClassRequirement); err != nil {
		return
	}
	// Cost : field : int
	if err = writer.AddInt(s.Cost); err != nil {
		return
	}
	// SkillRequirements : array : short
	for ndx := 0; ndx < 4; ndx++ {
		if len(s.SkillRequirements) != 4 {
			err = fmt.Errorf("expected SkillRequirements with length 4, got %d", len(s.SkillRequirements))
			return
		}

		if err = writer.AddShort(s.SkillRequirements[ndx]); err != nil {
			return
		}
	}

	// StatRequirements : field : CharacterBaseStats
	if err = s.StatRequirements.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *SkillLearn) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Id : field : short
	s.Id = reader.GetShort()
	// LevelRequirement : field : char
	s.LevelRequirement = reader.GetChar()
	// ClassRequirement : field : char
	s.ClassRequirement = reader.GetChar()
	// Cost : field : int
	s.Cost = reader.GetInt()
	// SkillRequirements : array : short
	for ndx := 0; ndx < 4; ndx++ {
		s.SkillRequirements = append(s.SkillRequirements, 0)
		s.SkillRequirements[ndx] = reader.GetShort()
	}

	// StatRequirements : field : CharacterBaseStats
	if err = s.StatRequirements.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// BoardPostListing :: An entry in the list of town board posts.
type BoardPostListing struct {
	byteSize int

	PostId  int
	Author  string
	Subject string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *BoardPostListing) ByteSize() int {
	return s.byteSize
}

func (s *BoardPostListing) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PostId : field : short
	if err = writer.AddShort(s.PostId); err != nil {
		return
	}
	writer.AddByte(255)
	// Author : field : string
	if err = writer.AddString(s.Author); err != nil {
		return
	}
	writer.AddByte(255)
	// Subject : field : string
	if err = writer.AddString(s.Subject); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *BoardPostListing) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PostId : field : short
	s.PostId = reader.GetShort()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Author : field : string
	if s.Author, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Subject : field : string
	if s.Subject, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// CharacterDetails :: Information displayed on the paperdoll and book.
type CharacterDetails struct {
	byteSize int

	Name      string
	Home      string
	Partner   string
	Title     string
	Guild     string
	GuildRank string
	PlayerId  int
	ClassId   int
	Gender    protocol.Gender
	Admin     protocol.AdminLevel
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *CharacterDetails) ByteSize() int {
	return s.byteSize
}

func (s *CharacterDetails) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// Home : field : string
	if err = writer.AddString(s.Home); err != nil {
		return
	}
	writer.AddByte(255)
	// Partner : field : string
	if err = writer.AddString(s.Partner); err != nil {
		return
	}
	writer.AddByte(255)
	// Title : field : string
	if err = writer.AddString(s.Title); err != nil {
		return
	}
	writer.AddByte(255)
	// Guild : field : string
	if err = writer.AddString(s.Guild); err != nil {
		return
	}
	writer.AddByte(255)
	// GuildRank : field : string
	if err = writer.AddString(s.GuildRank); err != nil {
		return
	}
	writer.AddByte(255)
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// ClassId : field : char
	if err = writer.AddChar(s.ClassId); err != nil {
		return
	}
	// Gender : field : Gender
	if err = writer.AddChar(int(s.Gender)); err != nil {
		return
	}
	// Admin : field : AdminLevel
	if err = writer.AddChar(int(s.Admin)); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *CharacterDetails) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Home : field : string
	if s.Home, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Partner : field : string
	if s.Partner, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Title : field : string
	if s.Title, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Guild : field : string
	if s.Guild, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// GuildRank : field : string
	if s.GuildRank, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// ClassId : field : char
	s.ClassId = reader.GetChar()
	// Gender : field : Gender
	s.Gender = protocol.Gender(reader.GetChar())
	// Admin : field : AdminLevel
	s.Admin = protocol.AdminLevel(reader.GetChar())
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyMember :: A member of the player's party.
type PartyMember struct {
	byteSize int

	PlayerId     int
	Leader       bool
	Level        int
	HpPercentage int
	Name         string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyMember) ByteSize() int {
	return s.byteSize
}

func (s *PartyMember) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Leader : field : bool
	if s.Leader {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	return
}

func (s *PartyMember) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Leader : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.Leader = true
	} else {
		s.Leader = false
	}
	// Level : field : char
	s.Level = reader.GetChar()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PartyExpShare :: EXP gain for a member of the player's party.
type PartyExpShare struct {
	byteSize int

	PlayerId   int
	Experience int
	LevelUp    int //  A value greater than 0 is "new level" and indicates the player leveled up.
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PartyExpShare) ByteSize() int {
	return s.byteSize
}

func (s *PartyExpShare) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Experience : field : int
	if err = writer.AddInt(s.Experience); err != nil {
		return
	}
	// LevelUp : field : char
	if err = writer.AddChar(s.LevelUp); err != nil {
		return
	}
	return
}

func (s *PartyExpShare) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Experience : field : int
	s.Experience = reader.GetInt()
	// LevelUp : field : char
	s.LevelUp = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildStaff :: Information about a guild staff member (recruiter or leader).
type GuildStaff struct {
	byteSize int

	Rank int
	Name string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildStaff) ByteSize() int {
	return s.byteSize
}

func (s *GuildStaff) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Rank : field : char
	if err = writer.AddChar(s.Rank); err != nil {
		return
	}
	writer.AddByte(255)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *GuildStaff) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Rank : field : char
	s.Rank = reader.GetChar()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GuildMember :: Information about a guild member.
type GuildMember struct {
	byteSize int

	Rank     int
	Name     string
	RankName string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GuildMember) ByteSize() int {
	return s.byteSize
}

func (s *GuildMember) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Rank : field : char
	if err = writer.AddChar(s.Rank); err != nil {
		return
	}
	writer.AddByte(255)
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// RankName : field : string
	if err = writer.AddString(s.RankName); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *GuildMember) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Rank : field : char
	s.Rank = reader.GetChar()
	if err = reader.NextChunk(); err != nil {
		return
	}
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// RankName : field : string
	if s.RankName, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GroupHealTargetPlayer :: Nearby player hit by a group heal spell.
type GroupHealTargetPlayer struct {
	byteSize int

	PlayerId     int
	HpPercentage int
	Hp           int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GroupHealTargetPlayer) ByteSize() int {
	return s.byteSize
}

func (s *GroupHealTargetPlayer) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// Hp : field : short
	if err = writer.AddShort(s.Hp); err != nil {
		return
	}
	return
}

func (s *GroupHealTargetPlayer) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// Hp : field : short
	s.Hp = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TradeItemData :: Trade window item data.
type TradeItemData struct {
	byteSize int

	PlayerId int
	Items    []net.Item
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TradeItemData) ByteSize() int {
	return s.byteSize
}

func (s *TradeItemData) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Items : array : Item
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *TradeItemData) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Items : array : Item
	ItemsRemaining := reader.Remaining()
	for ndx := 0; ndx < ItemsRemaining/6; ndx++ {
		s.Items = append(s.Items, net.Item{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcKilledData :: Information about an NPC that has been killed.
type NpcKilledData struct {
	byteSize int

	KillerId        int
	KillerDirection protocol.Direction
	NpcIndex        int
	DropIndex       int
	DropId          int
	DropCoords      protocol.Coords
	DropAmount      int
	Damage          int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcKilledData) ByteSize() int {
	return s.byteSize
}

func (s *NpcKilledData) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// KillerId : field : short
	if err = writer.AddShort(s.KillerId); err != nil {
		return
	}
	// KillerDirection : field : Direction
	if err = writer.AddChar(int(s.KillerDirection)); err != nil {
		return
	}
	// NpcIndex : field : short
	if err = writer.AddShort(s.NpcIndex); err != nil {
		return
	}
	// DropIndex : field : short
	if err = writer.AddShort(s.DropIndex); err != nil {
		return
	}
	// DropId : field : short
	if err = writer.AddShort(s.DropId); err != nil {
		return
	}
	// DropCoords : field : Coords
	if err = s.DropCoords.Serialize(writer); err != nil {
		return
	}
	// DropAmount : field : int
	if err = writer.AddInt(s.DropAmount); err != nil {
		return
	}
	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	return
}

func (s *NpcKilledData) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// KillerId : field : short
	s.KillerId = reader.GetShort()
	// KillerDirection : field : Direction
	s.KillerDirection = protocol.Direction(reader.GetChar())
	// NpcIndex : field : short
	s.NpcIndex = reader.GetShort()
	// DropIndex : field : short
	s.DropIndex = reader.GetShort()
	// DropId : field : short
	s.DropId = reader.GetShort()
	// DropCoords : field : Coords
	if err = s.DropCoords.Deserialize(reader); err != nil {
		return
	}
	// DropAmount : field : int
	s.DropAmount = reader.GetInt()
	// Damage : field : three
	s.Damage = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// LevelUpStats :: Level and stat updates.
type LevelUpStats struct {
	byteSize int

	Level       int
	StatPoints  int
	SkillPoints int
	MaxHp       int
	MaxTp       int
	MaxSp       int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *LevelUpStats) ByteSize() int {
	return s.byteSize
}

func (s *LevelUpStats) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Level : field : char
	if err = writer.AddChar(s.Level); err != nil {
		return
	}
	// StatPoints : field : short
	if err = writer.AddShort(s.StatPoints); err != nil {
		return
	}
	// SkillPoints : field : short
	if err = writer.AddShort(s.SkillPoints); err != nil {
		return
	}
	// MaxHp : field : short
	if err = writer.AddShort(s.MaxHp); err != nil {
		return
	}
	// MaxTp : field : short
	if err = writer.AddShort(s.MaxTp); err != nil {
		return
	}
	// MaxSp : field : short
	if err = writer.AddShort(s.MaxSp); err != nil {
		return
	}
	return
}

func (s *LevelUpStats) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Level : field : char
	s.Level = reader.GetChar()
	// StatPoints : field : short
	s.StatPoints = reader.GetShort()
	// SkillPoints : field : short
	s.SkillPoints = reader.GetShort()
	// MaxHp : field : short
	s.MaxHp = reader.GetShort()
	// MaxTp : field : short
	s.MaxTp = reader.GetShort()
	// MaxSp : field : short
	s.MaxSp = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcUpdatePosition :: An NPC walking.
type NpcUpdatePosition struct {
	byteSize int

	NpcIndex  int
	Coords    protocol.Coords
	Direction protocol.Direction
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcUpdatePosition) ByteSize() int {
	return s.byteSize
}

func (s *NpcUpdatePosition) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : char
	if err = writer.AddChar(s.NpcIndex); err != nil {
		return
	}
	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	return
}

func (s *NpcUpdatePosition) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : char
	s.NpcIndex = reader.GetChar()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcUpdateAttack :: An NPC attacking.
type NpcUpdateAttack struct {
	byteSize int

	NpcIndex     int
	Killed       PlayerKilledState
	Direction    protocol.Direction
	PlayerId     int
	Damage       int
	HpPercentage int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcUpdateAttack) ByteSize() int {
	return s.byteSize
}

func (s *NpcUpdateAttack) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : char
	if err = writer.AddChar(s.NpcIndex); err != nil {
		return
	}
	// Killed : field : PlayerKilledState
	if err = writer.AddChar(int(s.Killed)); err != nil {
		return
	}
	// Direction : field : Direction
	if err = writer.AddChar(int(s.Direction)); err != nil {
		return
	}
	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// Damage : field : three
	if err = writer.AddThree(s.Damage); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	return
}

func (s *NpcUpdateAttack) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : char
	s.NpcIndex = reader.GetChar()
	// Killed : field : PlayerKilledState
	s.Killed = PlayerKilledState(reader.GetChar())
	// Direction : field : Direction
	s.Direction = protocol.Direction(reader.GetChar())
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// Damage : field : three
	s.Damage = reader.GetThree()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// NpcUpdateChat :: An NPC talking.
type NpcUpdateChat struct {
	byteSize int

	NpcIndex int
	Message  string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *NpcUpdateChat) ByteSize() int {
	return s.byteSize
}

func (s *NpcUpdateChat) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcIndex : field : char
	if err = writer.AddChar(s.NpcIndex); err != nil {
		return
	}
	// MessageLength : length : char
	if err = writer.AddChar(len(s.Message)); err != nil {
		return
	}
	// Message : field : string
	if err = writer.AddFixedString(s.Message, len(s.Message)); err != nil {
		return
	}
	return
}

func (s *NpcUpdateChat) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// NpcIndex : field : char
	s.NpcIndex = reader.GetChar()
	// MessageLength : length : char
	messageLength := reader.GetChar()
	// Message : field : string
	if s.Message, err = reader.GetFixedString(messageLength); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// QuestProgressEntry :: An entry in the Quest Progress window.
type QuestProgressEntry struct {
	byteSize int

	Name        string
	Description string
	Icon        QuestRequirementIcon
	Progress    int
	Target      int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *QuestProgressEntry) ByteSize() int {
	return s.byteSize
}

func (s *QuestProgressEntry) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// Name : field : string
	if err = writer.AddString(s.Name); err != nil {
		return
	}
	writer.AddByte(255)
	// Description : field : string
	if err = writer.AddString(s.Description); err != nil {
		return
	}
	writer.AddByte(255)
	// Icon : field : QuestRequirementIcon
	if err = writer.AddShort(int(s.Icon)); err != nil {
		return
	}
	// Progress : field : short
	if err = writer.AddShort(s.Progress); err != nil {
		return
	}
	// Target : field : short
	if err = writer.AddShort(s.Target); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *QuestProgressEntry) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// Name : field : string
	if s.Name, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Description : field : string
	if s.Description, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Icon : field : QuestRequirementIcon
	s.Icon = QuestRequirementIcon(reader.GetShort())
	// Progress : field : short
	s.Progress = reader.GetShort()
	// Target : field : short
	s.Target = reader.GetShort()
	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// DialogQuestEntry :: An entry in the quest switcher.
type DialogQuestEntry struct {
	byteSize int

	QuestId   int
	QuestName string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *DialogQuestEntry) ByteSize() int {
	return s.byteSize
}

func (s *DialogQuestEntry) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// QuestId : field : short
	if err = writer.AddShort(s.QuestId); err != nil {
		return
	}
	// QuestName : field : string
	if err = writer.AddString(s.QuestName); err != nil {
		return
	}
	return
}

func (s *DialogQuestEntry) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// QuestId : field : short
	s.QuestId = reader.GetShort()
	// QuestName : field : string
	if s.QuestName, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// DialogEntry :: An entry in a quest dialog.
type DialogEntry struct {
	byteSize int

	EntryType     DialogEntryType
	EntryTypeData EntryTypeData
	Line          string
}

type EntryTypeData interface {
	protocol.EoData
}

type EntryTypeDataLink struct {
	byteSize int

	LinkId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *EntryTypeDataLink) ByteSize() int {
	return s.byteSize
}

func (s *EntryTypeDataLink) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// LinkId : field : short
	if err = writer.AddShort(s.LinkId); err != nil {
		return
	}
	return
}

func (s *EntryTypeDataLink) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// LinkId : field : short
	s.LinkId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *DialogEntry) ByteSize() int {
	return s.byteSize
}

func (s *DialogEntry) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// EntryType : field : DialogEntryType
	if err = writer.AddShort(int(s.EntryType)); err != nil {
		return
	}
	switch s.EntryType {
	case DialogEntry_Link:
		switch s.EntryTypeData.(type) {
		case *EntryTypeDataLink:
			if err = s.EntryTypeData.Serialize(writer); err != nil {
				return
			}
		default:
			err = fmt.Errorf("invalid switch struct type for switch value %d", s.EntryType)
			return
		}
	}
	// Line : field : string
	if err = writer.AddString(s.Line); err != nil {
		return
	}
	return
}

func (s *DialogEntry) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// EntryType : field : DialogEntryType
	s.EntryType = DialogEntryType(reader.GetShort())
	switch s.EntryType {
	case DialogEntry_Link:
		s.EntryTypeData = &EntryTypeDataLink{}
		if err = s.EntryTypeData.Deserialize(reader); err != nil {
			return
		}
	}
	// Line : field : string
	if s.Line, err = reader.GetString(); err != nil {
		return
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapDrainDamageOther :: Another player taking damage from a map HP drain.
type MapDrainDamageOther struct {
	byteSize int

	PlayerId     int
	HpPercentage int
	Damage       int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapDrainDamageOther) ByteSize() int {
	return s.byteSize
}

func (s *MapDrainDamageOther) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// HpPercentage : field : char
	if err = writer.AddChar(s.HpPercentage); err != nil {
		return
	}
	// Damage : field : short
	if err = writer.AddShort(s.Damage); err != nil {
		return
	}
	return
}

func (s *MapDrainDamageOther) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// HpPercentage : field : char
	s.HpPercentage = reader.GetChar()
	// Damage : field : short
	s.Damage = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// GlobalBackfillMessage :: A backfilled global chat message.
type GlobalBackfillMessage struct {
	byteSize int

	PlayerName string
	Message    string
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *GlobalBackfillMessage) ByteSize() int {
	return s.byteSize
}

func (s *GlobalBackfillMessage) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	writer.SanitizeStrings = true
	// PlayerName : field : string
	if err = writer.AddString(s.PlayerName); err != nil {
		return
	}
	writer.AddByte(255)
	// Message : field : string
	if err = writer.AddString(s.Message); err != nil {
		return
	}
	writer.SanitizeStrings = false
	return
}

func (s *GlobalBackfillMessage) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	reader.SetIsChunked(true)
	// PlayerName : field : string
	if s.PlayerName, err = reader.GetString(); err != nil {
		return
	}

	if err = reader.NextChunk(); err != nil {
		return
	}
	// Message : field : string
	if s.Message, err = reader.GetString(); err != nil {
		return
	}

	reader.SetIsChunked(false)
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// PlayerEffect :: An effect playing on a player.
type PlayerEffect struct {
	byteSize int

	PlayerId int
	EffectId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *PlayerEffect) ByteSize() int {
	return s.byteSize
}

func (s *PlayerEffect) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// PlayerId : field : short
	if err = writer.AddShort(s.PlayerId); err != nil {
		return
	}
	// EffectId : field : three
	if err = writer.AddThree(s.EffectId); err != nil {
		return
	}
	return
}

func (s *PlayerEffect) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// PlayerId : field : short
	s.PlayerId = reader.GetShort()
	// EffectId : field : three
	s.EffectId = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// TileEffect :: An effect playing on a tile.
type TileEffect struct {
	byteSize int

	Coords   protocol.Coords
	EffectId int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *TileEffect) ByteSize() int {
	return s.byteSize
}

func (s *TileEffect) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// EffectId : field : short
	if err = writer.AddShort(s.EffectId); err != nil {
		return
	}
	return
}

func (s *TileEffect) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// EffectId : field : short
	s.EffectId = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}
