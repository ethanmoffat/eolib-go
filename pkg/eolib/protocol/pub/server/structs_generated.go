package serverpub

import "github.com/ethanmoffat/eolib-go/pkg/eolib/data"

// DropRecord :: Record of an item an NPC can drop when killed.
type DropRecord struct {
	ItemId    int
	MinAmount int
	MaxAmount int
	Rate      int // Chance (x in 64,000) of the item being dropped.
}

func (s *DropRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// MinAmount : field : three
	if err = writer.AddThree(s.MinAmount); err != nil {
		return
	}
	// MaxAmount : field : three
	if err = writer.AddThree(s.MaxAmount); err != nil {
		return
	}
	// Rate : field : short
	if err = writer.AddShort(s.Rate); err != nil {
		return
	}
	return
}

func (s *DropRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// MinAmount : field : three
	s.MinAmount = reader.GetThree()
	// MaxAmount : field : three
	s.MaxAmount = reader.GetThree()
	// Rate : field : short
	s.Rate = reader.GetShort()

	return
}

// DropNpcRecord :: Record of potential drops from an NPC.
type DropNpcRecord struct {
	NpcId      int
	DropsCount int
	Drops      []DropRecord
}

func (s *DropNpcRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcId : field : short
	if err = writer.AddShort(s.NpcId); err != nil {
		return
	}
	// DropsCount : length : short
	if err = writer.AddShort(s.DropsCount); err != nil {
		return
	}
	// Drops : array : DropRecord
	for ndx := 0; ndx < s.DropsCount; ndx++ {
		if err = s.Drops[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *DropNpcRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NpcId : field : short
	s.NpcId = reader.GetShort()
	// DropsCount : length : short
	s.DropsCount = reader.GetShort()
	// Drops : array : DropRecord
	for ndx := 0; ndx < s.DropsCount; ndx++ {
		s.Drops = append(s.Drops, DropRecord{})
		if err = s.Drops[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// DropFile :: Endless Drop File.
type DropFile struct {
	Npcs []DropNpcRecord
}

func (s *DropFile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// EDF : field : string
	if err = writer.AddFixedString("EDF", 3); err != nil {
		return
	}
	// Npcs : array : DropNpcRecord
	for ndx := 0; ndx < len(s.Npcs); ndx++ {
		if err = s.Npcs[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *DropFile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// EDF : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Npcs : array : DropNpcRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Npcs = append(s.Npcs, DropNpcRecord{})
		if err = s.Npcs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// InnQuestionRecord :: Record of a question and answer that the player must answer to register citizenship with an inn.
type InnQuestionRecord struct {
	QuestionLength int
	Question       string
	AnswerLength   int
	Answer         string
}

func (s *InnQuestionRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// QuestionLength : length : char
	if err = writer.AddChar(s.QuestionLength); err != nil {
		return
	}
	// Question : field : string
	if err = writer.AddFixedString(s.Question, s.QuestionLength); err != nil {
		return
	}
	// AnswerLength : length : char
	if err = writer.AddChar(s.AnswerLength); err != nil {
		return
	}
	// Answer : field : string
	if err = writer.AddFixedString(s.Answer, s.AnswerLength); err != nil {
		return
	}
	return
}

func (s *InnQuestionRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// QuestionLength : length : char
	s.QuestionLength = reader.GetChar()
	// Question : field : string
	if s.Question, err = reader.GetFixedString(s.QuestionLength); err != nil {
		return
	}

	// AnswerLength : length : char
	s.AnswerLength = reader.GetChar()
	// Answer : field : string
	if s.Answer, err = reader.GetFixedString(s.AnswerLength); err != nil {
		return
	}

	return
}

// InnRecord :: Record of Inn data in an Endless Inn File.
type InnRecord struct {
	BehaviorId            int // Behavior ID of the NPC that runs the inn. 0 for default inn.
	NameLength            int
	Name                  string
	SpawnMap              int  // ID of the map the player is sent to after respawning.
	SpawnX                int  // X coordinate of the map the player is sent to after respawning.
	SpawnY                int  // Y coordinate of the map the player is sent to after respawning.
	SleepMap              int  // ID of the map the player is sent to after sleeping at the inn.
	SleepX                int  // X coordinate of the map the player is sent to after sleeping at the inn.
	SleepY                int  // Y coordinate of the map the player is sent to after sleeping at the inn.
	AlternateSpawnEnabled bool //  Flag for an alternate spawn point. If true, the server will use this alternate spawn. map, x, and, y based on some other condition.  In the official server, this is used to respawn new characters on the noob island. until they reach a certain level.
	AlternateSpawnMap     int
	AlternateSpawnX       int
	AlternateSpawnY       int
	Questions             []InnQuestionRecord
}

func (s *InnRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}
	// NameLength : length : char
	if err = writer.AddChar(s.NameLength); err != nil {
		return
	}
	// Name : field : string
	if err = writer.AddFixedString(s.Name, s.NameLength); err != nil {
		return
	}
	// SpawnMap : field : short
	if err = writer.AddShort(s.SpawnMap); err != nil {
		return
	}
	// SpawnX : field : char
	if err = writer.AddChar(s.SpawnX); err != nil {
		return
	}
	// SpawnY : field : char
	if err = writer.AddChar(s.SpawnY); err != nil {
		return
	}
	// SleepMap : field : short
	if err = writer.AddShort(s.SleepMap); err != nil {
		return
	}
	// SleepX : field : char
	if err = writer.AddChar(s.SleepX); err != nil {
		return
	}
	// SleepY : field : char
	if err = writer.AddChar(s.SleepY); err != nil {
		return
	}
	// AlternateSpawnEnabled : field : bool
	if s.AlternateSpawnEnabled {
		err = writer.AddChar(1)
	} else {
		err = writer.AddChar(0)
	}
	if err != nil {
		return
	}

	// AlternateSpawnMap : field : short
	if err = writer.AddShort(s.AlternateSpawnMap); err != nil {
		return
	}
	// AlternateSpawnX : field : char
	if err = writer.AddChar(s.AlternateSpawnX); err != nil {
		return
	}
	// AlternateSpawnY : field : char
	if err = writer.AddChar(s.AlternateSpawnY); err != nil {
		return
	}
	// Questions : array : InnQuestionRecord
	for ndx := 0; ndx < 3; ndx++ {
		if err = s.Questions[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *InnRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	// NameLength : length : char
	s.NameLength = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetFixedString(s.NameLength); err != nil {
		return
	}

	// SpawnMap : field : short
	s.SpawnMap = reader.GetShort()
	// SpawnX : field : char
	s.SpawnX = reader.GetChar()
	// SpawnY : field : char
	s.SpawnY = reader.GetChar()
	// SleepMap : field : short
	s.SleepMap = reader.GetShort()
	// SleepX : field : char
	s.SleepX = reader.GetChar()
	// SleepY : field : char
	s.SleepY = reader.GetChar()
	// AlternateSpawnEnabled : field : bool
	if boolVal := reader.GetChar(); boolVal > 0 {
		s.AlternateSpawnEnabled = true
	} else {
		s.AlternateSpawnEnabled = false
	}
	// AlternateSpawnMap : field : short
	s.AlternateSpawnMap = reader.GetShort()
	// AlternateSpawnX : field : char
	s.AlternateSpawnX = reader.GetChar()
	// AlternateSpawnY : field : char
	s.AlternateSpawnY = reader.GetChar()
	// Questions : array : InnQuestionRecord
	for ndx := 0; ndx < 3; ndx++ {
		s.Questions = append(s.Questions, InnQuestionRecord{})
		if err = s.Questions[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// InnFile :: Endless Inn File.
type InnFile struct {
	Inns []InnRecord
}

func (s *InnFile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// EID : field : string
	if err = writer.AddFixedString("EID", 3); err != nil {
		return
	}
	// Inns : array : InnRecord
	for ndx := 0; ndx < len(s.Inns); ndx++ {
		if err = s.Inns[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *InnFile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// EID : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Inns : array : InnRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Inns = append(s.Inns, InnRecord{})
		if err = s.Inns[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// SkillMasterSkillRecord :: Record of a skill that a Skill Master NPC can teach.
type SkillMasterSkillRecord struct {
	SkillId           int
	LevelRequirement  int // Level required to learn this skill.
	ClassRequirement  int // Class required to learn this skill.
	Price             int
	SkillRequirements []int // IDs of skills that must be learned before a player can learn this skill.
	StrRequirement    int   // Strength required to learn this skill.
	IntRequirement    int   // Intelligence required to learn this skill.
	WisRequirement    int   // Wisdom required to learn this skill.
	AgiRequirement    int   // Agility required to learn this skill.
	ConRequirement    int   // Constitution required to learn this skill.
	ChaRequirement    int   // Charisma required to learn this skill.
}

func (s *SkillMasterSkillRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// SkillId : field : short
	if err = writer.AddShort(s.SkillId); err != nil {
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
	// Price : field : int
	if err = writer.AddInt(s.Price); err != nil {
		return
	}
	// SkillRequirements : array : short
	for ndx := 0; ndx < 4; ndx++ {
		if err = writer.AddShort(s.SkillRequirements[ndx]); err != nil {
			return
		}
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
	return
}

func (s *SkillMasterSkillRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// SkillId : field : short
	s.SkillId = reader.GetShort()
	// LevelRequirement : field : char
	s.LevelRequirement = reader.GetChar()
	// ClassRequirement : field : char
	s.ClassRequirement = reader.GetChar()
	// Price : field : int
	s.Price = reader.GetInt()
	// SkillRequirements : array : short
	for ndx := 0; ndx < 4; ndx++ {
		s.SkillRequirements = append(s.SkillRequirements, 0)
		s.SkillRequirements[ndx] = reader.GetShort()
	}

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

	return
}

// SkillMasterRecord :: Record of Skill Master data in an Endless Skill Master File.
type SkillMasterRecord struct {
	BehaviorId       int // Behavior ID of the Skill Master NPC.
	NameLength       int
	Name             string
	MinLevel         int // Minimum level required to use this Skill Master.
	MaxLevel         int // Maximum level allowed to use this Skill Master.
	ClassRequirement int // Class required to use this Skill Master.
	SkillsCount      int
	Skills           []SkillMasterSkillRecord
}

func (s *SkillMasterRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}
	// NameLength : length : char
	if err = writer.AddChar(s.NameLength); err != nil {
		return
	}
	// Name : field : string
	if err = writer.AddFixedString(s.Name, s.NameLength); err != nil {
		return
	}
	// MinLevel : field : char
	if err = writer.AddChar(s.MinLevel); err != nil {
		return
	}
	// MaxLevel : field : char
	if err = writer.AddChar(s.MaxLevel); err != nil {
		return
	}
	// ClassRequirement : field : char
	if err = writer.AddChar(s.ClassRequirement); err != nil {
		return
	}
	// SkillsCount : length : short
	if err = writer.AddShort(s.SkillsCount); err != nil {
		return
	}
	// Skills : array : SkillMasterSkillRecord
	for ndx := 0; ndx < s.SkillsCount; ndx++ {
		if err = s.Skills[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *SkillMasterRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	// NameLength : length : char
	s.NameLength = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetFixedString(s.NameLength); err != nil {
		return
	}

	// MinLevel : field : char
	s.MinLevel = reader.GetChar()
	// MaxLevel : field : char
	s.MaxLevel = reader.GetChar()
	// ClassRequirement : field : char
	s.ClassRequirement = reader.GetChar()
	// SkillsCount : length : short
	s.SkillsCount = reader.GetShort()
	// Skills : array : SkillMasterSkillRecord
	for ndx := 0; ndx < s.SkillsCount; ndx++ {
		s.Skills = append(s.Skills, SkillMasterSkillRecord{})
		if err = s.Skills[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// SkillMasterFile :: Endless Skill Master File.
type SkillMasterFile struct {
	SkillMasters []SkillMasterRecord
}

func (s *SkillMasterFile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// EMF : field : string
	if err = writer.AddFixedString("EMF", 3); err != nil {
		return
	}
	// SkillMasters : array : SkillMasterRecord
	for ndx := 0; ndx < len(s.SkillMasters); ndx++ {
		if err = s.SkillMasters[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *SkillMasterFile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// EMF : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// SkillMasters : array : SkillMasterRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.SkillMasters = append(s.SkillMasters, SkillMasterRecord{})
		if err = s.SkillMasters[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ShopTradeRecord :: Record of an item that can be bought or sold in a shop.
type ShopTradeRecord struct {
	ItemId    int
	BuyPrice  int // How much it costs to buy the item from the shop.
	SellPrice int // How much the shop will pay for the item.
	MaxAmount int // Max amount of the item that can be bought or sold at one time.
}

func (s *ShopTradeRecord) Serialize(writer *data.EoWriter) (err error) {
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
	// MaxAmount : field : char
	if err = writer.AddChar(s.MaxAmount); err != nil {
		return
	}
	return
}

func (s *ShopTradeRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// BuyPrice : field : three
	s.BuyPrice = reader.GetThree()
	// SellPrice : field : three
	s.SellPrice = reader.GetThree()
	// MaxAmount : field : char
	s.MaxAmount = reader.GetChar()

	return
}

// ShopCraftIngredientRecord :: Record of an ingredient for crafting an item in a shop.
type ShopCraftIngredientRecord struct {
	ItemId int // Item ID of the craft ingredient, or 0 if the ingredient is not present.
	Amount int
}

func (s *ShopCraftIngredientRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// Amount : field : char
	if err = writer.AddChar(s.Amount); err != nil {
		return
	}
	return
}

func (s *ShopCraftIngredientRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// Amount : field : char
	s.Amount = reader.GetChar()

	return
}

// ShopCraftRecord :: Record of an item that can be crafted in a shop.
type ShopCraftRecord struct {
	ItemId      int
	Ingredients []ShopCraftIngredientRecord
}

func (s *ShopCraftRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ItemId : field : short
	if err = writer.AddShort(s.ItemId); err != nil {
		return
	}
	// Ingredients : array : ShopCraftIngredientRecord
	for ndx := 0; ndx < 4; ndx++ {
		if err = s.Ingredients[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ShopCraftRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ItemId : field : short
	s.ItemId = reader.GetShort()
	// Ingredients : array : ShopCraftIngredientRecord
	for ndx := 0; ndx < 4; ndx++ {
		s.Ingredients = append(s.Ingredients, ShopCraftIngredientRecord{})
		if err = s.Ingredients[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ShopRecord :: Record of Shop data in an Endless Shop File.
type ShopRecord struct {
	BehaviorId       int
	NameLength       int
	Name             string
	MinLevel         int // Minimum level required to use this shop.
	MaxLevel         int // Maximum level allowed to use this shop.
	ClassRequirement int // Class required to use this shop.
	TradesCount      int
	CraftsCount      int
	Trades           []ShopTradeRecord
	Crafts           []ShopCraftRecord
}

func (s *ShopRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// BehaviorId : field : short
	if err = writer.AddShort(s.BehaviorId); err != nil {
		return
	}
	// NameLength : length : char
	if err = writer.AddChar(s.NameLength); err != nil {
		return
	}
	// Name : field : string
	if err = writer.AddFixedString(s.Name, s.NameLength); err != nil {
		return
	}
	// MinLevel : field : char
	if err = writer.AddChar(s.MinLevel); err != nil {
		return
	}
	// MaxLevel : field : char
	if err = writer.AddChar(s.MaxLevel); err != nil {
		return
	}
	// ClassRequirement : field : char
	if err = writer.AddChar(s.ClassRequirement); err != nil {
		return
	}
	// TradesCount : length : short
	if err = writer.AddShort(s.TradesCount); err != nil {
		return
	}
	// CraftsCount : length : char
	if err = writer.AddChar(s.CraftsCount); err != nil {
		return
	}
	// Trades : array : ShopTradeRecord
	for ndx := 0; ndx < s.TradesCount; ndx++ {
		if err = s.Trades[ndx].Serialize(writer); err != nil {
			return
		}
	}

	// Crafts : array : ShopCraftRecord
	for ndx := 0; ndx < s.CraftsCount; ndx++ {
		if err = s.Crafts[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ShopRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// BehaviorId : field : short
	s.BehaviorId = reader.GetShort()
	// NameLength : length : char
	s.NameLength = reader.GetChar()
	// Name : field : string
	if s.Name, err = reader.GetFixedString(s.NameLength); err != nil {
		return
	}

	// MinLevel : field : char
	s.MinLevel = reader.GetChar()
	// MaxLevel : field : char
	s.MaxLevel = reader.GetChar()
	// ClassRequirement : field : char
	s.ClassRequirement = reader.GetChar()
	// TradesCount : length : short
	s.TradesCount = reader.GetShort()
	// CraftsCount : length : char
	s.CraftsCount = reader.GetChar()
	// Trades : array : ShopTradeRecord
	for ndx := 0; ndx < s.TradesCount; ndx++ {
		s.Trades = append(s.Trades, ShopTradeRecord{})
		if err = s.Trades[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	// Crafts : array : ShopCraftRecord
	for ndx := 0; ndx < s.CraftsCount; ndx++ {
		s.Crafts = append(s.Crafts, ShopCraftRecord{})
		if err = s.Crafts[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// ShopFile :: Endless Shop File.
type ShopFile struct {
	Shops []ShopRecord
}

func (s *ShopFile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ESF : field : string
	if err = writer.AddFixedString("ESF", 3); err != nil {
		return
	}
	// Shops : array : ShopRecord
	for ndx := 0; ndx < len(s.Shops); ndx++ {
		if err = s.Shops[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *ShopFile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ESF : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Shops : array : ShopRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Shops = append(s.Shops, ShopRecord{})
		if err = s.Shops[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// TalkMessageRecord :: Record of a message that an NPC can say.
type TalkMessageRecord struct {
	MessageLength int
	Message       string
}

func (s *TalkMessageRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// MessageLength : length : char
	if err = writer.AddChar(s.MessageLength); err != nil {
		return
	}
	// Message : field : string
	if err = writer.AddFixedString(s.Message, s.MessageLength); err != nil {
		return
	}
	return
}

func (s *TalkMessageRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// MessageLength : length : char
	s.MessageLength = reader.GetChar()
	// Message : field : string
	if s.Message, err = reader.GetFixedString(s.MessageLength); err != nil {
		return
	}

	return
}

// TalkRecord :: Record of Talk data in an Endless Talk File.
type TalkRecord struct {
	NpcId         int // ID of the NPC that will talk.
	Rate          int // Chance that the NPC will talk (0-100).
	MessagesCount int
	Messages      []TalkMessageRecord
}

func (s *TalkRecord) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// NpcId : field : short
	if err = writer.AddShort(s.NpcId); err != nil {
		return
	}
	// Rate : field : char
	if err = writer.AddChar(s.Rate); err != nil {
		return
	}
	// MessagesCount : length : char
	if err = writer.AddChar(s.MessagesCount); err != nil {
		return
	}
	// Messages : array : TalkMessageRecord
	for ndx := 0; ndx < s.MessagesCount; ndx++ {
		if err = s.Messages[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *TalkRecord) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// NpcId : field : short
	s.NpcId = reader.GetShort()
	// Rate : field : char
	s.Rate = reader.GetChar()
	// MessagesCount : length : char
	s.MessagesCount = reader.GetChar()
	// Messages : array : TalkMessageRecord
	for ndx := 0; ndx < s.MessagesCount; ndx++ {
		s.Messages = append(s.Messages, TalkMessageRecord{})
		if err = s.Messages[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}

// TalkFile :: Endless Talk File.
type TalkFile struct {
	Npcs []TalkRecord
}

func (s *TalkFile) Serialize(writer *data.EoWriter) (err error) {
	oldSanitizeStrings := writer.SanitizeStrings
	defer func() { writer.SanitizeStrings = oldSanitizeStrings }()

	// ETF : field : string
	if err = writer.AddFixedString("ETF", 3); err != nil {
		return
	}
	// Npcs : array : TalkRecord
	for ndx := 0; ndx < len(s.Npcs); ndx++ {
		if err = s.Npcs[ndx].Serialize(writer); err != nil {
			return
		}
	}

	return
}

func (s *TalkFile) Deserialize(reader *data.EoReader) (err error) {
	oldIsChunked := reader.IsChunked()
	defer func() { reader.SetIsChunked(oldIsChunked) }()

	// ETF : field : string
	if _, err = reader.GetFixedString(3); err != nil {
		return
	}
	// Npcs : array : TalkRecord
	for ndx := 0; reader.Remaining() > 0; ndx++ {
		s.Npcs = append(s.Npcs, TalkRecord{})
		if err = s.Npcs[ndx].Deserialize(reader); err != nil {
			return
		}
	}

	return
}
