package pub

import "fmt"

type Element int

const (
	Element_None Element = iota
	Element_Light
	Element_Dark
	Element_Earth
	Element_Wind
	Element_Water
	Element_Fire
)

// String converts a Element value into its string representation
func (e Element) String() (string, error) {
	switch e {
	case Element_None:
		return "None", nil
	case Element_Light:
		return "Light", nil
	case Element_Dark:
		return "Dark", nil
	case Element_Earth:
		return "Earth", nil
	case Element_Wind:
		return "Wind", nil
	case Element_Water:
		return "Water", nil
	case Element_Fire:
		return "Fire", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type Element to string", e)
	}
}

type ItemType int

const (
	Item_General ItemType = iota
	Item_Reserved1
	Item_Currency
	Item_Heal
	Item_Teleport
	Item_Reserved5
	Item_ExpReward
	Item_Reserved7
	Item_Reserved8
	Item_Key
	Item_Weapon
	Item_Shield
	Item_Armor
	Item_Hat
	Item_Boots
	Item_Gloves
	Item_Accessory
	Item_Belt
	Item_Necklace
	Item_Ring
	Item_Armlet
	Item_Bracer
	Item_Alcohol
	Item_EffectPotion
	Item_HairDye
	Item_CureCurse
	Item_Reserved26
	Item_Reserved27
	Item_Reserved28
	Item_Reserved29
)

// String converts a ItemType value into its string representation
func (e ItemType) String() (string, error) {
	switch e {
	case Item_General:
		return "General", nil
	case Item_Reserved1:
		return "Reserved1", nil
	case Item_Currency:
		return "Currency", nil
	case Item_Heal:
		return "Heal", nil
	case Item_Teleport:
		return "Teleport", nil
	case Item_Reserved5:
		return "Reserved5", nil
	case Item_ExpReward:
		return "ExpReward", nil
	case Item_Reserved7:
		return "Reserved7", nil
	case Item_Reserved8:
		return "Reserved8", nil
	case Item_Key:
		return "Key", nil
	case Item_Weapon:
		return "Weapon", nil
	case Item_Shield:
		return "Shield", nil
	case Item_Armor:
		return "Armor", nil
	case Item_Hat:
		return "Hat", nil
	case Item_Boots:
		return "Boots", nil
	case Item_Gloves:
		return "Gloves", nil
	case Item_Accessory:
		return "Accessory", nil
	case Item_Belt:
		return "Belt", nil
	case Item_Necklace:
		return "Necklace", nil
	case Item_Ring:
		return "Ring", nil
	case Item_Armlet:
		return "Armlet", nil
	case Item_Bracer:
		return "Bracer", nil
	case Item_Alcohol:
		return "Alcohol", nil
	case Item_EffectPotion:
		return "EffectPotion", nil
	case Item_HairDye:
		return "HairDye", nil
	case Item_CureCurse:
		return "CureCurse", nil
	case Item_Reserved26:
		return "Reserved26", nil
	case Item_Reserved27:
		return "Reserved27", nil
	case Item_Reserved28:
		return "Reserved28", nil
	case Item_Reserved29:
		return "Reserved29", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type ItemType to string", e)
	}
}

type ItemSubtype int

const (
	ItemSubtype_None ItemSubtype = iota
	ItemSubtype_Ranged
	ItemSubtype_Arrows
	ItemSubtype_Wings
	ItemSubtype_Reserved4
)

// String converts a ItemSubtype value into its string representation
func (e ItemSubtype) String() (string, error) {
	switch e {
	case ItemSubtype_None:
		return "None", nil
	case ItemSubtype_Ranged:
		return "Ranged", nil
	case ItemSubtype_Arrows:
		return "Arrows", nil
	case ItemSubtype_Wings:
		return "Wings", nil
	case ItemSubtype_Reserved4:
		return "Reserved4", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type ItemSubtype to string", e)
	}
}

type ItemSpecial int

const (
	ItemSpecial_Normal ItemSpecial = iota
	ItemSpecial_Rare
	ItemSpecial_Legendary
	ItemSpecial_Unique
	ItemSpecial_Lore
	ItemSpecial_Cursed
)

// String converts a ItemSpecial value into its string representation
func (e ItemSpecial) String() (string, error) {
	switch e {
	case ItemSpecial_Normal:
		return "Normal", nil
	case ItemSpecial_Rare:
		return "Rare", nil
	case ItemSpecial_Legendary:
		return "Legendary", nil
	case ItemSpecial_Unique:
		return "Unique", nil
	case ItemSpecial_Lore:
		return "Lore", nil
	case ItemSpecial_Cursed:
		return "Cursed", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type ItemSpecial to string", e)
	}
}

// ItemSize :: Size of an item in the inventory.
type ItemSize int

const (
	ItemSize_Size1x1 ItemSize = iota
	ItemSize_Size1x2
	ItemSize_Size1x3
	ItemSize_Size1x4
	ItemSize_Size2x1
	ItemSize_Size2x2
	ItemSize_Size2x3
	ItemSize_Size2x4
)

// String converts a ItemSize value into its string representation
func (e ItemSize) String() (string, error) {
	switch e {
	case ItemSize_Size1x1:
		return "Size1x1", nil
	case ItemSize_Size1x2:
		return "Size1x2", nil
	case ItemSize_Size1x3:
		return "Size1x3", nil
	case ItemSize_Size1x4:
		return "Size1x4", nil
	case ItemSize_Size2x1:
		return "Size2x1", nil
	case ItemSize_Size2x2:
		return "Size2x2", nil
	case ItemSize_Size2x3:
		return "Size2x3", nil
	case ItemSize_Size2x4:
		return "Size2x4", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type ItemSize to string", e)
	}
}

type NpcType int

const (
	Npc_Friendly NpcType = iota
	Npc_Passive
	Npc_Aggressive
	Npc_Reserved3
	Npc_Reserved4
	Npc_Reserved5
	Npc_Shop
	Npc_Inn
	Npc_Reserved8
	Npc_Bank
	Npc_Barber
	Npc_Guild
	Npc_Priest
	Npc_Lawyer
	Npc_Trainer
	Npc_Quest
)

// String converts a NpcType value into its string representation
func (e NpcType) String() (string, error) {
	switch e {
	case Npc_Friendly:
		return "Friendly", nil
	case Npc_Passive:
		return "Passive", nil
	case Npc_Aggressive:
		return "Aggressive", nil
	case Npc_Reserved3:
		return "Reserved3", nil
	case Npc_Reserved4:
		return "Reserved4", nil
	case Npc_Reserved5:
		return "Reserved5", nil
	case Npc_Shop:
		return "Shop", nil
	case Npc_Inn:
		return "Inn", nil
	case Npc_Reserved8:
		return "Reserved8", nil
	case Npc_Bank:
		return "Bank", nil
	case Npc_Barber:
		return "Barber", nil
	case Npc_Guild:
		return "Guild", nil
	case Npc_Priest:
		return "Priest", nil
	case Npc_Lawyer:
		return "Lawyer", nil
	case Npc_Trainer:
		return "Trainer", nil
	case Npc_Quest:
		return "Quest", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type NpcType to string", e)
	}
}

type SkillNature int

const (
	SkillNature_Spell SkillNature = iota
	SkillNature_Skill
)

// String converts a SkillNature value into its string representation
func (e SkillNature) String() (string, error) {
	switch e {
	case SkillNature_Spell:
		return "Spell", nil
	case SkillNature_Skill:
		return "Skill", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SkillNature to string", e)
	}
}

type SkillType int

const (
	Skill_Heal SkillType = iota
	Skill_Attack
	Skill_Bard
)

// String converts a SkillType value into its string representation
func (e SkillType) String() (string, error) {
	switch e {
	case Skill_Heal:
		return "Heal", nil
	case Skill_Attack:
		return "Attack", nil
	case Skill_Bard:
		return "Bard", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SkillType to string", e)
	}
}

type SkillTargetRestrict int

const (
	SkillTargetRestrict_Npc SkillTargetRestrict = iota
	SkillTargetRestrict_Friendly
	SkillTargetRestrict_Opponent
)

// String converts a SkillTargetRestrict value into its string representation
func (e SkillTargetRestrict) String() (string, error) {
	switch e {
	case SkillTargetRestrict_Npc:
		return "Npc", nil
	case SkillTargetRestrict_Friendly:
		return "Friendly", nil
	case SkillTargetRestrict_Opponent:
		return "Opponent", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SkillTargetRestrict to string", e)
	}
}

type SkillTargetType int

const (
	SkillTarget_Normal SkillTargetType = iota
	SkillTarget_Self
	SkillTarget_Reserved2
	SkillTarget_Group
)

// String converts a SkillTargetType value into its string representation
func (e SkillTargetType) String() (string, error) {
	switch e {
	case SkillTarget_Normal:
		return "Normal", nil
	case SkillTarget_Self:
		return "Self", nil
	case SkillTarget_Reserved2:
		return "Reserved2", nil
	case SkillTarget_Group:
		return "Group", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SkillTargetType to string", e)
	}
}
