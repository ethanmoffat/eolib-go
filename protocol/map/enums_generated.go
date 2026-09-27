package eomap

import "fmt"

type MapType int

const (
	Map_Normal MapType = iota
	Map_Pk             = 3
)

// String converts a MapType value into its string representation
func (e MapType) String() (string, error) {
	switch e {
	case Map_Normal:
		return "Normal", nil
	case Map_Pk:
		return "Pk", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MapType to string", e)
	}
}

// MapTimedEffect :: A timed effect that can occur on a map.
type MapTimedEffect int

const (
	MapTimedEffect_None MapTimedEffect = iota
	MapTimedEffect_HpDrain
	MapTimedEffect_TpDrain
	MapTimedEffect_Quake1
	MapTimedEffect_Quake2
	MapTimedEffect_Quake3
	MapTimedEffect_Quake4
)

// String converts a MapTimedEffect value into its string representation
func (e MapTimedEffect) String() (string, error) {
	switch e {
	case MapTimedEffect_None:
		return "None", nil
	case MapTimedEffect_HpDrain:
		return "HpDrain", nil
	case MapTimedEffect_TpDrain:
		return "TpDrain", nil
	case MapTimedEffect_Quake1:
		return "Quake1", nil
	case MapTimedEffect_Quake2:
		return "Quake2", nil
	case MapTimedEffect_Quake3:
		return "Quake3", nil
	case MapTimedEffect_Quake4:
		return "Quake4", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MapTimedEffect to string", e)
	}
}

// MapMusicControl :: How background music should be played on a map.
type MapMusicControl int

const (
	MapMusicControl_InterruptIfDifferentPlayOnce MapMusicControl = iota
	MapMusicControl_InterruptPlayOnce
	MapMusicControl_FinishPlayOnce
	MapMusicControl_InterruptIfDifferentPlayRepeat
	MapMusicControl_InterruptPlayRepeat
	MapMusicControl_FinishPlayRepeat
	MapMusicControl_InterruptPlayNothing
)

// String converts a MapMusicControl value into its string representation
func (e MapMusicControl) String() (string, error) {
	switch e {
	case MapMusicControl_InterruptIfDifferentPlayOnce:
		return "InterruptIfDifferentPlayOnce", nil
	case MapMusicControl_InterruptPlayOnce:
		return "InterruptPlayOnce", nil
	case MapMusicControl_FinishPlayOnce:
		return "FinishPlayOnce", nil
	case MapMusicControl_InterruptIfDifferentPlayRepeat:
		return "InterruptIfDifferentPlayRepeat", nil
	case MapMusicControl_InterruptPlayRepeat:
		return "InterruptPlayRepeat", nil
	case MapMusicControl_FinishPlayRepeat:
		return "FinishPlayRepeat", nil
	case MapMusicControl_InterruptPlayNothing:
		return "InterruptPlayNothing", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MapMusicControl to string", e)
	}
}

// MapTileSpec :: The type of a tile on a map.
type MapTileSpec int

const (
	MapTileSpec_Wall MapTileSpec = iota
	MapTileSpec_ChairDown
	MapTileSpec_ChairLeft
	MapTileSpec_ChairRight
	MapTileSpec_ChairUp
	MapTileSpec_ChairDownRight
	MapTileSpec_ChairUpLeft
	MapTileSpec_ChairAll
	MapTileSpec_Reserved8
	MapTileSpec_Chest
	MapTileSpec_Reserved10
	MapTileSpec_Reserved11
	MapTileSpec_Reserved12
	MapTileSpec_Reserved13
	MapTileSpec_Reserved14
	MapTileSpec_Reserved15
	MapTileSpec_BankVault
	MapTileSpec_NpcBoundary
	MapTileSpec_Edge
	MapTileSpec_FakeWall
	MapTileSpec_Board1
	MapTileSpec_Board2
	MapTileSpec_Board3
	MapTileSpec_Board4
	MapTileSpec_Board5
	MapTileSpec_Board6
	MapTileSpec_Board7
	MapTileSpec_Board8
	MapTileSpec_Jukebox
	MapTileSpec_Jump
	MapTileSpec_Water
	MapTileSpec_Reserved31
	MapTileSpec_Arena
	MapTileSpec_AmbientSource
	MapTileSpec_TimedSpikes
	MapTileSpec_Spikes
	MapTileSpec_HiddenSpikes
)

// String converts a MapTileSpec value into its string representation
func (e MapTileSpec) String() (string, error) {
	switch e {
	case MapTileSpec_Wall:
		return "Wall", nil
	case MapTileSpec_ChairDown:
		return "ChairDown", nil
	case MapTileSpec_ChairLeft:
		return "ChairLeft", nil
	case MapTileSpec_ChairRight:
		return "ChairRight", nil
	case MapTileSpec_ChairUp:
		return "ChairUp", nil
	case MapTileSpec_ChairDownRight:
		return "ChairDownRight", nil
	case MapTileSpec_ChairUpLeft:
		return "ChairUpLeft", nil
	case MapTileSpec_ChairAll:
		return "ChairAll", nil
	case MapTileSpec_Reserved8:
		return "Reserved8", nil
	case MapTileSpec_Chest:
		return "Chest", nil
	case MapTileSpec_Reserved10:
		return "Reserved10", nil
	case MapTileSpec_Reserved11:
		return "Reserved11", nil
	case MapTileSpec_Reserved12:
		return "Reserved12", nil
	case MapTileSpec_Reserved13:
		return "Reserved13", nil
	case MapTileSpec_Reserved14:
		return "Reserved14", nil
	case MapTileSpec_Reserved15:
		return "Reserved15", nil
	case MapTileSpec_BankVault:
		return "BankVault", nil
	case MapTileSpec_NpcBoundary:
		return "NpcBoundary", nil
	case MapTileSpec_Edge:
		return "Edge", nil
	case MapTileSpec_FakeWall:
		return "FakeWall", nil
	case MapTileSpec_Board1:
		return "Board1", nil
	case MapTileSpec_Board2:
		return "Board2", nil
	case MapTileSpec_Board3:
		return "Board3", nil
	case MapTileSpec_Board4:
		return "Board4", nil
	case MapTileSpec_Board5:
		return "Board5", nil
	case MapTileSpec_Board6:
		return "Board6", nil
	case MapTileSpec_Board7:
		return "Board7", nil
	case MapTileSpec_Board8:
		return "Board8", nil
	case MapTileSpec_Jukebox:
		return "Jukebox", nil
	case MapTileSpec_Jump:
		return "Jump", nil
	case MapTileSpec_Water:
		return "Water", nil
	case MapTileSpec_Reserved31:
		return "Reserved31", nil
	case MapTileSpec_Arena:
		return "Arena", nil
	case MapTileSpec_AmbientSource:
		return "AmbientSource", nil
	case MapTileSpec_TimedSpikes:
		return "TimedSpikes", nil
	case MapTileSpec_Spikes:
		return "Spikes", nil
	case MapTileSpec_HiddenSpikes:
		return "HiddenSpikes", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MapTileSpec to string", e)
	}
}
