package eomap

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/v3/data"
	"github.com/ethanmoffat/eolib-go/v3/protocol"
)

// MapNpc :: NPC spawn EMF entity.
type MapNpc struct {
	byteSize int

	Coords    protocol.Coords
	Id        int
	SpawnType int
	SpawnTime int
	Amount    int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapNpc) ByteSize() int {
	return s.byteSize
}

func (s *MapNpc) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Id : field : short
	if err = writer.AddShort(s.Id); err != nil {
		return
	}
	// SpawnType : field : char
	if err = writer.AddChar(s.SpawnType); err != nil {
		return
	}
	// SpawnTime : field : short
	if err = writer.AddShort(s.SpawnTime); err != nil {
		return
	}
	// Amount : field : char
	if err = writer.AddChar(s.Amount); err != nil {
		return
	}
	return
}

func (s *MapNpc) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Id : field : short
	s.Id = reader.GetShort()
	// SpawnType : field : char
	s.SpawnType = reader.GetChar()
	// SpawnTime : field : short
	s.SpawnTime = reader.GetShort()
	// Amount : field : char
	s.Amount = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapLegacyDoorKey :: Legacy EMF entity used to specify a key on a door.
type MapLegacyDoorKey struct {
	byteSize int

	Coords protocol.Coords
	Key    int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapLegacyDoorKey) ByteSize() int {
	return s.byteSize
}

func (s *MapLegacyDoorKey) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Key : field : short
	if err = writer.AddShort(s.Key); err != nil {
		return
	}
	return
}

func (s *MapLegacyDoorKey) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Key : field : short
	s.Key = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapItem :: Item spawn EMF entity.
type MapItem struct {
	byteSize int

	Coords    protocol.Coords
	Key       int
	ChestSlot int
	ItemId    int
	SpawnTime int
	Amount    int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapItem) ByteSize() int {
	return s.byteSize
}

func (s *MapItem) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// Key : field : short
	if err = writer.AddShort(s.Key); err != nil {
		return
	}
	// ChestSlot : field : char
	if err = writer.AddChar(s.ChestSlot); err != nil {
		return
	}
	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// SpawnTime : field : short
	if err = writer.AddShort(s.SpawnTime); err != nil {
		return
	}
	// Amount : field : three
	if err = writer.AddThree(s.Amount); err != nil {
		return
	}
	return
}

func (s *MapItem) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// Key : field : short
	s.Key = reader.GetShort()
	// ChestSlot : field : char
	s.ChestSlot = reader.GetChar()
	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// SpawnTime : field : short
	s.SpawnTime = reader.GetShort()
	// Amount : field : three
	s.Amount = reader.GetThree()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapWarp :: Warp EMF entity.
type MapWarp struct {
	byteSize int

	DestinationMap    int
	DestinationCoords protocol.Coords
	LevelRequired     int
	Door              int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapWarp) ByteSize() int {
	return s.byteSize
}

func (s *MapWarp) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// DestinationMap : field : short
	if err = writer.AddShort(s.DestinationMap); err != nil {
		return
	}
	// DestinationCoords : field : Coords
	if err = s.DestinationCoords.Serialize(writer); err != nil {
		return
	}
	// LevelRequired : field : char
	if err = writer.AddChar(s.LevelRequired); err != nil {
		return
	}
	// Door : field : short
	if err = writer.AddShort(s.Door); err != nil {
		return
	}
	return
}

func (s *MapWarp) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// DestinationMap : field : short
	s.DestinationMap = reader.GetShort()
	// DestinationCoords : field : Coords
	if err = s.DestinationCoords.Deserialize(reader); err != nil {
		return
	}
	// LevelRequired : field : char
	s.LevelRequired = reader.GetChar()
	// Door : field : short
	s.Door = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapSign :: Sign EMF entity.
type MapSign struct {
	byteSize int

	Coords      protocol.Coords
	StringData  string
	TitleLength int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapSign) ByteSize() int {
	return s.byteSize
}

func (s *MapSign) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Coords : field : Coords
	if err = s.Coords.Serialize(writer); err != nil {
		return
	}
	// StringDataLength : length : short
	if err = writer.AddShort(len(s.StringData) + 1); err != nil {
		return
	}
	// StringData : field : encoded_string
	if err = writer.AddFixedEncodedString(s.StringData, len(s.StringData)); err != nil {
		return
	}
	// TitleLength : field : char
	if err = writer.AddChar(s.TitleLength); err != nil {
		return
	}
	return
}

func (s *MapSign) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Coords : field : Coords
	if err = s.Coords.Deserialize(reader); err != nil {
		return
	}
	// StringDataLength : length : short
	stringDataLength := reader.GetShort() - 1
	// StringData : field : encoded_string
	if s.StringData, err = reader.GetFixedEncodedString(stringDataLength); err != nil {
		return
	}

	// TitleLength : field : char
	s.TitleLength = reader.GetChar()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapTileSpecRowTile :: A single tile in a row of tilespecs.
type MapTileSpecRowTile struct {
	byteSize int

	X        int
	TileSpec MapTileSpec
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapTileSpecRowTile) ByteSize() int {
	return s.byteSize
}

func (s *MapTileSpecRowTile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// X : field : char
	if err = writer.AddChar(s.X); err != nil {
		return
	}
	// TileSpec : field : MapTileSpec
	if err = writer.AddChar(int(s.TileSpec)); err != nil {
		return
	}
	return
}

func (s *MapTileSpecRowTile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// X : field : char
	s.X = reader.GetChar()
	// TileSpec : field : MapTileSpec
	s.TileSpec = MapTileSpec(reader.GetChar())
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapTileSpecRow :: A row of tilespecs.
type MapTileSpecRow struct {
	byteSize int

	Y     int
	Tiles []MapTileSpecRowTile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapTileSpecRow) ByteSize() int {
	return s.byteSize
}

func (s *MapTileSpecRow) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Y : field : char
	if err = writer.AddChar(s.Y); err != nil {
		return
	}
	// TilesCount : length : char
	if err = writer.AddChar(len(s.Tiles)); err != nil {
		return
	}
	// Tiles : array : MapTileSpecRowTile
	for ndx := 0; ndx < len(s.Tiles); ndx++ {
		if err = s.Tiles[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *MapTileSpecRow) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Y : field : char
	s.Y = reader.GetChar()
	// TilesCount : length : char
	tilesCount := reader.GetChar()
	// Tiles : array : MapTileSpecRowTile
	for ndx := 0; ndx < tilesCount; ndx++ {
		s.Tiles = append(s.Tiles, MapTileSpecRowTile{})
		if err = s.Tiles[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapWarpRowTile :: A single tile in a row of warp entities.
type MapWarpRowTile struct {
	byteSize int

	X    int
	Warp MapWarp
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapWarpRowTile) ByteSize() int {
	return s.byteSize
}

func (s *MapWarpRowTile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// X : field : char
	if err = writer.AddChar(s.X); err != nil {
		return
	}
	// Warp : field : MapWarp
	if err = s.Warp.Serialize(writer); err != nil {
		return
	}
	return
}

func (s *MapWarpRowTile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// X : field : char
	s.X = reader.GetChar()
	// Warp : field : MapWarp
	if err = s.Warp.Deserialize(reader); err != nil {
		return
	}
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapWarpRow :: A row of warp entities.
type MapWarpRow struct {
	byteSize int

	Y     int
	Tiles []MapWarpRowTile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapWarpRow) ByteSize() int {
	return s.byteSize
}

func (s *MapWarpRow) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Y : field : char
	if err = writer.AddChar(s.Y); err != nil {
		return
	}
	// TilesCount : length : char
	if err = writer.AddChar(len(s.Tiles)); err != nil {
		return
	}
	// Tiles : array : MapWarpRowTile
	for ndx := 0; ndx < len(s.Tiles); ndx++ {
		if err = s.Tiles[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *MapWarpRow) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Y : field : char
	s.Y = reader.GetChar()
	// TilesCount : length : char
	tilesCount := reader.GetChar()
	// Tiles : array : MapWarpRowTile
	for ndx := 0; ndx < tilesCount; ndx++ {
		s.Tiles = append(s.Tiles, MapWarpRowTile{})
		if err = s.Tiles[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapGraphicRowTile :: A single tile in a row of map graphics.
type MapGraphicRowTile struct {
	byteSize int

	X       int
	Graphic int
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapGraphicRowTile) ByteSize() int {
	return s.byteSize
}

func (s *MapGraphicRowTile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// X : field : char
	if err = writer.AddChar(s.X); err != nil {
		return
	}
	// Graphic : field : short
	if err = writer.AddShort(s.Graphic); err != nil {
		return
	}
	return
}

func (s *MapGraphicRowTile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// X : field : char
	s.X = reader.GetChar()
	// Graphic : field : short
	s.Graphic = reader.GetShort()
	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapGraphicRow :: A row in a layer of map graphics.
type MapGraphicRow struct {
	byteSize int

	Y     int
	Tiles []MapGraphicRowTile
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapGraphicRow) ByteSize() int {
	return s.byteSize
}

func (s *MapGraphicRow) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// Y : field : char
	if err = writer.AddChar(s.Y); err != nil {
		return
	}
	// TilesCount : length : char
	if err = writer.AddChar(len(s.Tiles)); err != nil {
		return
	}
	// Tiles : array : MapGraphicRowTile
	for ndx := 0; ndx < len(s.Tiles); ndx++ {
		if err = s.Tiles[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *MapGraphicRow) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// Y : field : char
	s.Y = reader.GetChar()
	// TilesCount : length : char
	tilesCount := reader.GetChar()
	// Tiles : array : MapGraphicRowTile
	for ndx := 0; ndx < tilesCount; ndx++ {
		s.Tiles = append(s.Tiles, MapGraphicRowTile{})
		if err = s.Tiles[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// MapGraphicLayer :: A layer of map graphics.
type MapGraphicLayer struct {
	byteSize int

	GraphicRows []MapGraphicRow
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *MapGraphicLayer) ByteSize() int {
	return s.byteSize
}

func (s *MapGraphicLayer) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// GraphicRowsCount : length : char
	if err = writer.AddChar(len(s.GraphicRows)); err != nil {
		return
	}
	// GraphicRows : array : MapGraphicRow
	for ndx := 0; ndx < len(s.GraphicRows); ndx++ {
		if err = s.GraphicRows[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *MapGraphicLayer) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// GraphicRowsCount : length : char
	graphicRowsCount := reader.GetChar()
	// GraphicRows : array : MapGraphicRow
	for ndx := 0; ndx < graphicRowsCount; ndx++ {
		s.GraphicRows = append(s.GraphicRows, MapGraphicRow{})
		if err = s.GraphicRows[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}

// Emf :: Endless Map File.
type Emf struct {
	byteSize int

	Rid            []int
	Name           string
	Type           MapType
	TimedEffect    MapTimedEffect
	MusicId        int
	MusicControl   MapMusicControl
	AmbientSoundId int
	Width          int
	Height         int
	FillTile       int
	MapAvailable   bool
	CanScroll      bool
	RelogX         int
	RelogY         int

	Npcs           []MapNpc
	LegacyDoorKeys []MapLegacyDoorKey
	Items          []MapItem
	TileSpecRows   []MapTileSpecRow
	WarpRows       []MapWarpRow
	GraphicLayers  []MapGraphicLayer //  The 9 layers of map graphics. Order is [Ground, Object, Overlay, Down Wall, Right Wall, Roof, Top, Shadow, Overlay2].
	Signs          []MapSign
}

// ByteSize gets the deserialized size of this object. This value is zero for an object that was not deserialized from data.
func (s *Emf) ByteSize() int {
	return s.byteSize
}

func (s *Emf) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// EMF : field : string
	if err = writer.AddFixedString("EMF", 3); err != nil {
		return
	}
	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		if len(s.Rid) != 2 {
			err = fmt.Errorf("expected Rid with length 2, got %d", len(s.Rid))
			return
		}

		if err = writer.AddShort(s.Rid[ndx]); err != nil {
			return
		}
	}

	// Name : field : encoded_string
	if len(s.Name) > 24 {
		err = fmt.Errorf("expected Name with length 24, got %d", len(s.Name))
		return
	}
	if err = writer.AddPaddedEncodedString(s.Name, 24); err != nil {
		return
	}
	// Type : field : MapType
	if err = writer.AddChar(int(s.Type)); err != nil {
		return
	}
	// TimedEffect : field : MapTimedEffect
	if err = writer.AddChar(int(s.TimedEffect)); err != nil {
		return
	}
	// MusicId : field : char
	if err = writer.AddChar(s.MusicId); err != nil {
		return
	}
	// MusicControl : field : MapMusicControl
	if err = writer.AddChar(int(s.MusicControl)); err != nil {
		return
	}
	// AmbientSoundId : field : short
	if err = writer.AddShort(s.AmbientSoundId); err != nil {
		return
	}
	// Width : field : char
	if err = writer.AddChar(s.Width); err != nil {
		return
	}
	// Height : field : char
	if err = writer.AddChar(s.Height); err != nil {
		return
	}
	// FillTile : field : short
	if err = writer.AddShort(s.FillTile); err != nil {
		return
	}
	// MapAvailable : field : bool
	if s.MapAvailable {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	// CanScroll : field : bool
	if s.CanScroll {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	// RelogX : field : char
	if err = writer.AddChar(s.RelogX); err != nil {
		return
	}
	// RelogY : field : char
	if err = writer.AddChar(s.RelogY); err != nil {
		return
	}
	// 0 : field : char
	if err = writer.AddChar(0); err != nil {
		return
	}
	// NpcsCount : length : char
	if err = writer.AddChar(len(s.Npcs)); err != nil {
		return
	}
	// Npcs : array : MapNpc
	for ndx := 0; ndx < len(s.Npcs); ndx++ {
		if err = s.Npcs[ndx].Serialize(writer); err != nil {
			return
		}
	}

	// LegacyDoorKeysCount : length : char
	if err = writer.AddChar(len(s.LegacyDoorKeys)); err != nil {
		return
	}
	// LegacyDoorKeys : array : MapLegacyDoorKey
	for ndx := 0; ndx < len(s.LegacyDoorKeys); ndx++ {
		if err = s.LegacyDoorKeys[ndx].Serialize(writer); err != nil {
			return
		}
	}

	// ItemsCount : length : char
	if err = writer.AddChar(len(s.Items)); err != nil {
		return
	}
	// Items : array : MapItem
	for ndx := 0; ndx < len(s.Items); ndx++ {
		if err = s.Items[ndx].Serialize(writer); err != nil {
			return
		}
	}

	// TileSpecRowsCount : length : char
	if err = writer.AddChar(len(s.TileSpecRows)); err != nil {
		return
	}
	// TileSpecRows : array : MapTileSpecRow
	for ndx := 0; ndx < len(s.TileSpecRows); ndx++ {
		if err = s.TileSpecRows[ndx].Serialize(writer); err != nil {
			return
		}
	}

	// WarpRowsCount : length : char
	if err = writer.AddChar(len(s.WarpRows)); err != nil {
		return
	}
	// WarpRows : array : MapWarpRow
	for ndx := 0; ndx < len(s.WarpRows); ndx++ {
		if err = s.WarpRows[ndx].Serialize(writer); err != nil {
			return
		}
	}

	// GraphicLayers : array : MapGraphicLayer
	for ndx := 0; ndx < 9; ndx++ {
		if len(s.GraphicLayers) != 9 {
			err = fmt.Errorf("expected GraphicLayers with length 9, got %d", len(s.GraphicLayers))
			return
		}

		if err = s.GraphicLayers[ndx].Serialize(writer); err != nil {
			return
		}
	}

	// SignsCount : length : char
	if err = writer.AddChar(len(s.Signs)); err != nil {
		return
	}
	// Signs : array : MapSign
	for ndx := 0; ndx < len(s.Signs); ndx++ {
		if err = s.Signs[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *Emf) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	readerStartPosition := reader.Position()
	// EMF : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Rid : array : short
	for ndx := 0; ndx < 2; ndx++ {
		s.Rid = append(s.Rid, 0)
		s.Rid[ndx] = reader.GetShort()
	}

	// Name : field : encoded_string
	if s.Name, err = reader.GetPaddedEncodedString(24); err != nil {
		return
	}

	// Type : field : MapType
	s.Type = MapType(reader.GetChar())
	// TimedEffect : field : MapTimedEffect
	s.TimedEffect = MapTimedEffect(reader.GetChar())
	// MusicId : field : char
	s.MusicId = reader.GetChar()
	// MusicControl : field : MapMusicControl
	s.MusicControl = MapMusicControl(reader.GetChar())
	// AmbientSoundId : field : short
	s.AmbientSoundId = reader.GetShort()
	// Width : field : char
	s.Width = reader.GetChar()
	// Height : field : char
	s.Height = reader.GetChar()
	// FillTile : field : short
	s.FillTile = reader.GetShort()
	// MapAvailable : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.MapAvailable = true
	} else {
		s.MapAvailable = false
	}
	// CanScroll : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.CanScroll = true
	} else {
		s.CanScroll = false
	}
	// RelogX : field : char
	s.RelogX = reader.GetChar()
	// RelogY : field : char
	s.RelogY = reader.GetChar()
	// 0 : field : char
	reader.GetChar()
	// NpcsCount : length : char
	npcsCount := reader.GetChar()
	// Npcs : array : MapNpc
	for ndx := 0; ndx < npcsCount; ndx++ {
		s.Npcs = append(s.Npcs, MapNpc{})
		if err = s.Npcs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	// LegacyDoorKeysCount : length : char
	legacyDoorKeysCount := reader.GetChar()
	// LegacyDoorKeys : array : MapLegacyDoorKey
	for ndx := 0; ndx < legacyDoorKeysCount; ndx++ {
		s.LegacyDoorKeys = append(s.LegacyDoorKeys, MapLegacyDoorKey{})
		if err = s.LegacyDoorKeys[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	// ItemsCount : length : char
	itemsCount := reader.GetChar()
	// Items : array : MapItem
	for ndx := 0; ndx < itemsCount; ndx++ {
		s.Items = append(s.Items, MapItem{})
		if err = s.Items[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	// TileSpecRowsCount : length : char
	tileSpecRowsCount := reader.GetChar()
	// TileSpecRows : array : MapTileSpecRow
	for ndx := 0; ndx < tileSpecRowsCount; ndx++ {
		s.TileSpecRows = append(s.TileSpecRows, MapTileSpecRow{})
		if err = s.TileSpecRows[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	// WarpRowsCount : length : char
	warpRowsCount := reader.GetChar()
	// WarpRows : array : MapWarpRow
	for ndx := 0; ndx < warpRowsCount; ndx++ {
		s.WarpRows = append(s.WarpRows, MapWarpRow{})
		if err = s.WarpRows[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	// GraphicLayers : array : MapGraphicLayer
	for ndx := 0; ndx < 9; ndx++ {
		s.GraphicLayers = append(s.GraphicLayers, MapGraphicLayer{})
		if err = s.GraphicLayers[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	// SignsCount : length : char
	signsCount := reader.GetChar()
	// Signs : array : MapSign
	for ndx := 0; ndx < signsCount; ndx++ {
		s.Signs = append(s.Signs, MapSign{})
		if err = s.Signs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	s.byteSize = reader.Position() - readerStartPosition

	return
}
