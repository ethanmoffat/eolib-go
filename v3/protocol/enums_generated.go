package protocol

import "fmt"

// AdminLevel :: The admin level of a player.
type AdminLevel int

const (
	AdminLevel_Player AdminLevel = iota
	AdminLevel_Spy
	AdminLevel_LightGuide
	AdminLevel_Guardian
	AdminLevel_GameMaster
	AdminLevel_HighGameMaster
)

// String converts a AdminLevel value into its string representation
func (e AdminLevel) String() (string, error) {
	switch e {
	case AdminLevel_Player:
		return "Player", nil
	case AdminLevel_Spy:
		return "Spy", nil
	case AdminLevel_LightGuide:
		return "LightGuide", nil
	case AdminLevel_Guardian:
		return "Guardian", nil
	case AdminLevel_GameMaster:
		return "GameMaster", nil
	case AdminLevel_HighGameMaster:
		return "HighGameMaster", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type AdminLevel to string", e)
	}
}

// Direction :: The direction a player or NPC is facing.
type Direction int

const (
	Direction_Down Direction = iota
	Direction_Left
	Direction_Up
	Direction_Right
)

// String converts a Direction value into its string representation
func (e Direction) String() (string, error) {
	switch e {
	case Direction_Down:
		return "Down", nil
	case Direction_Left:
		return "Left", nil
	case Direction_Up:
		return "Up", nil
	case Direction_Right:
		return "Right", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type Direction to string", e)
	}
}

// Emote :: Emote that can be played over a player's head.
type Emote int

const (
	Emote_Happy Emote = iota + 1
	Emote_Depressed
	Emote_Sad
	Emote_Angry
	Emote_Confused
	Emote_Surprised
	Emote_Hearts
	Emote_Moon
	Emote_Suicidal
	Emote_Embarrassed
	Emote_Drunk
	Emote_Trade
	Emote_LevelUp
	Emote_Playful
)

// String converts a Emote value into its string representation
func (e Emote) String() (string, error) {
	switch e {
	case Emote_Happy:
		return "Happy", nil
	case Emote_Depressed:
		return "Depressed", nil
	case Emote_Sad:
		return "Sad", nil
	case Emote_Angry:
		return "Angry", nil
	case Emote_Confused:
		return "Confused", nil
	case Emote_Surprised:
		return "Surprised", nil
	case Emote_Hearts:
		return "Hearts", nil
	case Emote_Moon:
		return "Moon", nil
	case Emote_Suicidal:
		return "Suicidal", nil
	case Emote_Embarrassed:
		return "Embarrassed", nil
	case Emote_Drunk:
		return "Drunk", nil
	case Emote_Trade:
		return "Trade", nil
	case Emote_LevelUp:
		return "LevelUp", nil
	case Emote_Playful:
		return "Playful", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type Emote to string", e)
	}
}

// Gender :: The gender of a player.
type Gender int

const (
	Gender_Female Gender = iota
	Gender_Male
)

// String converts a Gender value into its string representation
func (e Gender) String() (string, error) {
	switch e {
	case Gender_Female:
		return "Female", nil
	case Gender_Male:
		return "Male", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type Gender to string", e)
	}
}
