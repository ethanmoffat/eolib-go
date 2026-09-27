package client

import "fmt"

// SitAction :: Whether the player wants to sit or stand.
type SitAction int

const (
	SitAction_Sit SitAction = iota + 1
	SitAction_Stand
)

// String converts a SitAction value into its string representation
func (e SitAction) String() (string, error) {
	switch e {
	case SitAction_Sit:
		return "Sit", nil
	case SitAction_Stand:
		return "Stand", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SitAction to string", e)
	}
}

// GuildInfoType :: The type of guild info being interacted with.
type GuildInfoType int

const (
	GuildInfo_Description GuildInfoType = iota + 1
	GuildInfo_Ranks
	GuildInfo_Bank
)

// String converts a GuildInfoType value into its string representation
func (e GuildInfoType) String() (string, error) {
	switch e {
	case GuildInfo_Description:
		return "Description", nil
	case GuildInfo_Ranks:
		return "Ranks", nil
	case GuildInfo_Bank:
		return "Bank", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type GuildInfoType to string", e)
	}
}

// TrainType :: Whether the player is spending a stat point or a skill point.
type TrainType int

const (
	Train_Stat TrainType = iota + 1
	Train_Skill
)

// String converts a TrainType value into its string representation
func (e TrainType) String() (string, error) {
	switch e {
	case Train_Stat:
		return "Stat", nil
	case Train_Skill:
		return "Skill", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type TrainType to string", e)
	}
}

// DialogReply :: Whether the player has clicked the OK button or a link in a quest dialog.
type DialogReply int

const (
	DialogReply_Ok DialogReply = iota + 1
	DialogReply_Link
)

// String converts a DialogReply value into its string representation
func (e DialogReply) String() (string, error) {
	switch e {
	case DialogReply_Ok:
		return "Ok", nil
	case DialogReply_Link:
		return "Link", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type DialogReply to string", e)
	}
}

// FileType :: Data file type.
type FileType int

const (
	File_Emf FileType = iota + 1
	File_Eif
	File_Enf
	File_Esf
	File_Ecf
)

// String converts a FileType value into its string representation
func (e FileType) String() (string, error) {
	switch e {
	case File_Emf:
		return "Emf", nil
	case File_Eif:
		return "Eif", nil
	case File_Enf:
		return "Enf", nil
	case File_Esf:
		return "Esf", nil
	case File_Ecf:
		return "Ecf", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type FileType to string", e)
	}
}

// StatId :: Base character stat.
type StatId int

const (
	StatId_Str StatId = iota + 1
	StatId_Int
	StatId_Wis
	StatId_Agi
	StatId_Con
	StatId_Cha
)

// String converts a StatId value into its string representation
func (e StatId) String() (string, error) {
	switch e {
	case StatId_Str:
		return "Str", nil
	case StatId_Int:
		return "Int", nil
	case StatId_Wis:
		return "Wis", nil
	case StatId_Agi:
		return "Agi", nil
	case StatId_Con:
		return "Con", nil
	case StatId_Cha:
		return "Cha", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type StatId to string", e)
	}
}

// SpellTargetType :: Target type of a spell cast.
type SpellTargetType int

const (
	SpellTarget_Player SpellTargetType = iota + 1
	SpellTarget_Npc
)

// String converts a SpellTargetType value into its string representation
func (e SpellTargetType) String() (string, error) {
	switch e {
	case SpellTarget_Player:
		return "Player", nil
	case SpellTarget_Npc:
		return "Npc", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SpellTargetType to string", e)
	}
}

// MarriageRequestType :: Request type sent with MARRIAGE_REQUEST packet.
type MarriageRequestType int

const (
	MarriageRequest_MarriageApproval MarriageRequestType = iota + 1
	MarriageRequest_Divorce
)

// String converts a MarriageRequestType value into its string representation
func (e MarriageRequestType) String() (string, error) {
	switch e {
	case MarriageRequest_MarriageApproval:
		return "MarriageApproval", nil
	case MarriageRequest_Divorce:
		return "Divorce", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MarriageRequestType to string", e)
	}
}
