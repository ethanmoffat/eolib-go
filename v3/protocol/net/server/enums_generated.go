package server

import "fmt"

// InitReply :: Reply code sent with INIT_INIT packet.
type InitReply int

const (
	InitReply_OutOfDate InitReply = iota + 1
	InitReply_Ok
	InitReply_Banned // The official client won't display a message until the connection from the server is closed.
	InitReply_WarpMap
	InitReply_FileEmf
	InitReply_FileEif
	InitReply_FileEnf
	InitReply_FileEsf
	InitReply_PlayersList
	InitReply_MapMutation
	InitReply_PlayersListFriends
	InitReply_FileEcf
)

// String converts a InitReply value into its string representation
func (e InitReply) String() (string, error) {
	switch e {
	case InitReply_OutOfDate:
		return "OutOfDate", nil
	case InitReply_Ok:
		return "Ok", nil
	case InitReply_Banned:
		return "Banned", nil
	case InitReply_WarpMap:
		return "WarpMap", nil
	case InitReply_FileEmf:
		return "FileEmf", nil
	case InitReply_FileEif:
		return "FileEif", nil
	case InitReply_FileEnf:
		return "FileEnf", nil
	case InitReply_FileEsf:
		return "FileEsf", nil
	case InitReply_PlayersList:
		return "PlayersList", nil
	case InitReply_MapMutation:
		return "MapMutation", nil
	case InitReply_PlayersListFriends:
		return "PlayersListFriends", nil
	case InitReply_FileEcf:
		return "FileEcf", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type InitReply to string", e)
	}
}

// InitBanType ::  Ban type sent with INIT_INIT packet. The official client treats a value >= 2 as Permanent. Otherwise, it's Temporary.
type InitBanType int

const (
	InitBan_Temporary InitBanType = iota + 1
	InitBan_Permanent
)

// String converts a InitBanType value into its string representation
func (e InitBanType) String() (string, error) {
	switch e {
	case InitBan_Temporary:
		return "Temporary", nil
	case InitBan_Permanent:
		return "Permanent", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type InitBanType to string", e)
	}
}

// CharacterIcon :: Icon displayed in paperdolls, books, and the online list.
type CharacterIcon int

const (
	CharacterIcon_Player   CharacterIcon = iota + 1
	CharacterIcon_Gm                     = 4
	CharacterIcon_Hgm                    = 5
	CharacterIcon_Party                  = 6
	CharacterIcon_GmParty                = 9
	CharacterIcon_HgmParty               = 10
)

// String converts a CharacterIcon value into its string representation
func (e CharacterIcon) String() (string, error) {
	switch e {
	case CharacterIcon_Player:
		return "Player", nil
	case CharacterIcon_Gm:
		return "Gm", nil
	case CharacterIcon_Hgm:
		return "Hgm", nil
	case CharacterIcon_Party:
		return "Party", nil
	case CharacterIcon_GmParty:
		return "GmParty", nil
	case CharacterIcon_HgmParty:
		return "HgmParty", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type CharacterIcon to string", e)
	}
}

// AvatarChangeType :: How a player's appearance is changing.
type AvatarChangeType int

const (
	AvatarChange_Equipment AvatarChangeType = iota + 1
	AvatarChange_Hair
	AvatarChange_HairColor
)

// String converts a AvatarChangeType value into its string representation
func (e AvatarChangeType) String() (string, error) {
	switch e {
	case AvatarChange_Equipment:
		return "Equipment", nil
	case AvatarChange_Hair:
		return "Hair", nil
	case AvatarChange_HairColor:
		return "HairColor", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type AvatarChangeType to string", e)
	}
}

// TalkReply :: Reply code sent with TALK_REPLY packet.
type TalkReply int

const (
	TalkReply_NotFound TalkReply = iota + 1
)

// String converts a TalkReply value into its string representation
func (e TalkReply) String() (string, error) {
	switch e {
	case TalkReply_NotFound:
		return "NotFound", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type TalkReply to string", e)
	}
}

// SitState :: Indicates how a player is sitting (or not sitting).
type SitState int

const (
	SitState_Stand SitState = iota
	SitState_Chair
	SitState_Floor
)

// String converts a SitState value into its string representation
func (e SitState) String() (string, error) {
	switch e {
	case SitState_Stand:
		return "Stand", nil
	case SitState_Chair:
		return "Chair", nil
	case SitState_Floor:
		return "Floor", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SitState to string", e)
	}
}

// MapEffect :: An effect that occurs for all players on a map.
type MapEffect int

const (
	MapEffect_Quake MapEffect = iota + 1
)

// String converts a MapEffect value into its string representation
func (e MapEffect) String() (string, error) {
	switch e {
	case MapEffect_Quake:
		return "Quake", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MapEffect to string", e)
	}
}

// GuildReply :: Reply code sent with GUILD_REPLY packet.
type GuildReply int

const (
	GuildReply_Busy GuildReply = iota + 1
	GuildReply_NotApproved
	GuildReply_AlreadyMember
	GuildReply_NoCandidates
	GuildReply_Exists
	GuildReply_CreateBegin
	GuildReply_CreateAddConfirm
	GuildReply_CreateAdd
	GuildReply_RecruiterOffline
	GuildReply_RecruiterNotHere
	GuildReply_RecruiterWrongGuild
	GuildReply_NotRecruiter
	GuildReply_JoinRequest
	GuildReply_NotPresent
	GuildReply_AccountLow
	GuildReply_Accepted
	GuildReply_NotFound
	GuildReply_Updated
	GuildReply_RanksUpdated
	GuildReply_RemoveLeader
	GuildReply_RemoveNotMember
	GuildReply_Removed
	GuildReply_RankingLeader
	GuildReply_RankingNotMember
)

// String converts a GuildReply value into its string representation
func (e GuildReply) String() (string, error) {
	switch e {
	case GuildReply_Busy:
		return "Busy", nil
	case GuildReply_NotApproved:
		return "NotApproved", nil
	case GuildReply_AlreadyMember:
		return "AlreadyMember", nil
	case GuildReply_NoCandidates:
		return "NoCandidates", nil
	case GuildReply_Exists:
		return "Exists", nil
	case GuildReply_CreateBegin:
		return "CreateBegin", nil
	case GuildReply_CreateAddConfirm:
		return "CreateAddConfirm", nil
	case GuildReply_CreateAdd:
		return "CreateAdd", nil
	case GuildReply_RecruiterOffline:
		return "RecruiterOffline", nil
	case GuildReply_RecruiterNotHere:
		return "RecruiterNotHere", nil
	case GuildReply_RecruiterWrongGuild:
		return "RecruiterWrongGuild", nil
	case GuildReply_NotRecruiter:
		return "NotRecruiter", nil
	case GuildReply_JoinRequest:
		return "JoinRequest", nil
	case GuildReply_NotPresent:
		return "NotPresent", nil
	case GuildReply_AccountLow:
		return "AccountLow", nil
	case GuildReply_Accepted:
		return "Accepted", nil
	case GuildReply_NotFound:
		return "NotFound", nil
	case GuildReply_Updated:
		return "Updated", nil
	case GuildReply_RanksUpdated:
		return "RanksUpdated", nil
	case GuildReply_RemoveLeader:
		return "RemoveLeader", nil
	case GuildReply_RemoveNotMember:
		return "RemoveNotMember", nil
	case GuildReply_Removed:
		return "Removed", nil
	case GuildReply_RankingLeader:
		return "RankingLeader", nil
	case GuildReply_RankingNotMember:
		return "RankingNotMember", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type GuildReply to string", e)
	}
}

// InnUnsubscribeReply ::  Reply code sent with CITIZEN_REMOVE packet. Indicates the result of trying to give up citizenship to a town.
type InnUnsubscribeReply int

const (
	InnUnsubscribeReply_NotCitizen InnUnsubscribeReply = iota
	InnUnsubscribeReply_Unsubscribed
)

// String converts a InnUnsubscribeReply value into its string representation
func (e InnUnsubscribeReply) String() (string, error) {
	switch e {
	case InnUnsubscribeReply_NotCitizen:
		return "NotCitizen", nil
	case InnUnsubscribeReply_Unsubscribed:
		return "Unsubscribed", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type InnUnsubscribeReply to string", e)
	}
}

// CharacterReply :: Reply code sent with CHARACTER_REPLY packet.
type CharacterReply int

const (
	CharacterReply_Exists CharacterReply = iota + 1
	CharacterReply_Full                  //  Only sent in reply to Character_Create packets. Displays the same message as CharacterReply.Full3 in the official client.
	CharacterReply_Full3                 //  Only sent in reply to Character_Request packets. Displays the same message as CharacterReply.Full in the official client.
	CharacterReply_NotApproved
	CharacterReply_Ok
	CharacterReply_Deleted
)

// String converts a CharacterReply value into its string representation
func (e CharacterReply) String() (string, error) {
	switch e {
	case CharacterReply_Exists:
		return "Exists", nil
	case CharacterReply_Full:
		return "Full", nil
	case CharacterReply_Full3:
		return "Full3", nil
	case CharacterReply_NotApproved:
		return "NotApproved", nil
	case CharacterReply_Ok:
		return "Ok", nil
	case CharacterReply_Deleted:
		return "Deleted", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type CharacterReply to string", e)
	}
}

// SkillMasterReply ::  Reply code sent with STATSKILL_REPLY packet. Indicates why an action was unsuccessful.
type SkillMasterReply int

const (
	SkillMasterReply_RemoveItems SkillMasterReply = iota + 1
	SkillMasterReply_WrongClass
)

// String converts a SkillMasterReply value into its string representation
func (e SkillMasterReply) String() (string, error) {
	switch e {
	case SkillMasterReply_RemoveItems:
		return "RemoveItems", nil
	case SkillMasterReply_WrongClass:
		return "WrongClass", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type SkillMasterReply to string", e)
	}
}

// AccountReply :: Reply code sent with ACCOUNT_REPLY packet.
type AccountReply int

const (
	AccountReply_Exists AccountReply = iota + 1
	AccountReply_NotApproved
	AccountReply_Created
	AccountReply_ChangeFailed  = 5
	AccountReply_Changed       = 6
	AccountReply_RequestDenied = 7
)

// String converts a AccountReply value into its string representation
func (e AccountReply) String() (string, error) {
	switch e {
	case AccountReply_Exists:
		return "Exists", nil
	case AccountReply_NotApproved:
		return "NotApproved", nil
	case AccountReply_Created:
		return "Created", nil
	case AccountReply_ChangeFailed:
		return "ChangeFailed", nil
	case AccountReply_Changed:
		return "Changed", nil
	case AccountReply_RequestDenied:
		return "RequestDenied", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type AccountReply to string", e)
	}
}

// LoginReply ::  Reply code sent with LOGIN_REPLY packet. Indicates the result of a login attempt.
type LoginReply int

const (
	LoginReply_WrongUser LoginReply = iota + 1
	LoginReply_WrongUserPassword
	LoginReply_Ok
	LoginReply_Banned // The official client won't display a message until the connection from the server is closed.
	LoginReply_LoggedIn
	LoginReply_Busy // The official client won't display a message until the connection from the server is closed.
)

// String converts a LoginReply value into its string representation
func (e LoginReply) String() (string, error) {
	switch e {
	case LoginReply_WrongUser:
		return "WrongUser", nil
	case LoginReply_WrongUserPassword:
		return "WrongUserPassword", nil
	case LoginReply_Ok:
		return "Ok", nil
	case LoginReply_Banned:
		return "Banned", nil
	case LoginReply_LoggedIn:
		return "LoggedIn", nil
	case LoginReply_Busy:
		return "Busy", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type LoginReply to string", e)
	}
}

// DialogEntryType :: The type of an entry in a quest dialog.
type DialogEntryType int

const (
	DialogEntry_Text DialogEntryType = iota + 1
	DialogEntry_Link
)

// String converts a DialogEntryType value into its string representation
func (e DialogEntryType) String() (string, error) {
	switch e {
	case DialogEntry_Text:
		return "Text", nil
	case DialogEntry_Link:
		return "Link", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type DialogEntryType to string", e)
	}
}

// QuestRequirementIcon :: Icon displayed for each quest in the Quest Progress window.
type QuestRequirementIcon int

const (
	QuestRequirementIcon_Item QuestRequirementIcon = iota + 3
	QuestRequirementIcon_Talk                      = 5
	QuestRequirementIcon_Kill                      = 8
	QuestRequirementIcon_Step                      = 10
)

// String converts a QuestRequirementIcon value into its string representation
func (e QuestRequirementIcon) String() (string, error) {
	switch e {
	case QuestRequirementIcon_Item:
		return "Item", nil
	case QuestRequirementIcon_Talk:
		return "Talk", nil
	case QuestRequirementIcon_Kill:
		return "Kill", nil
	case QuestRequirementIcon_Step:
		return "Step", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type QuestRequirementIcon to string", e)
	}
}

// WarpEffect :: An effect that accompanies a player warp.
type WarpEffect int

const (
	WarpEffect_None   WarpEffect = iota // Does nothing.
	WarpEffect_Scroll                   // Plays the scroll sound effect.
	WarpEffect_Admin                    // Plays the admin warp sound effect and animation.
)

// String converts a WarpEffect value into its string representation
func (e WarpEffect) String() (string, error) {
	switch e {
	case WarpEffect_None:
		return "None", nil
	case WarpEffect_Scroll:
		return "Scroll", nil
	case WarpEffect_Admin:
		return "Admin", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type WarpEffect to string", e)
	}
}

// WarpType ::  Indicates whether a warp is within the current map, or switching to another map.
type WarpType int

const (
	Warp_Local WarpType = iota + 1
	Warp_MapSwitch
)

// String converts a WarpType value into its string representation
func (e WarpType) String() (string, error) {
	switch e {
	case Warp_Local:
		return "Local", nil
	case Warp_MapSwitch:
		return "MapSwitch", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type WarpType to string", e)
	}
}

// WelcomeCode :: Reply code sent with WELCOME_REPLY packet.
type WelcomeCode int

const (
	WelcomeCode_SelectCharacter WelcomeCode = iota + 1
	WelcomeCode_EnterGame
	WelcomeCode_ServerBusy
	WelcomeCode_LoggedIn
)

// String converts a WelcomeCode value into its string representation
func (e WelcomeCode) String() (string, error) {
	switch e {
	case WelcomeCode_SelectCharacter:
		return "SelectCharacter", nil
	case WelcomeCode_EnterGame:
		return "EnterGame", nil
	case WelcomeCode_ServerBusy:
		return "ServerBusy", nil
	case WelcomeCode_LoggedIn:
		return "LoggedIn", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type WelcomeCode to string", e)
	}
}

// LoginMessageCode :: Whether a warning message should be displayed upon entering the game.
type LoginMessageCode int

const (
	LoginMessageCode_No  LoginMessageCode = iota
	LoginMessageCode_Yes                  = 2
)

// String converts a LoginMessageCode value into its string representation
func (e LoginMessageCode) String() (string, error) {
	switch e {
	case LoginMessageCode_No:
		return "No", nil
	case LoginMessageCode_Yes:
		return "Yes", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type LoginMessageCode to string", e)
	}
}

// AdminMessageType :: Type of message sent to admins via the Help menu.
type AdminMessageType int

const (
	AdminMessage_Message AdminMessageType = iota + 1
	AdminMessage_Report
)

// String converts a AdminMessageType value into its string representation
func (e AdminMessageType) String() (string, error) {
	switch e {
	case AdminMessage_Message:
		return "Message", nil
	case AdminMessage_Report:
		return "Report", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type AdminMessageType to string", e)
	}
}

// PlayerKilledState :: Flag to indicate that a player has been killed.
type PlayerKilledState int

const (
	PlayerKilledState_Alive PlayerKilledState = iota + 1
	PlayerKilledState_Killed
)

// String converts a PlayerKilledState value into its string representation
func (e PlayerKilledState) String() (string, error) {
	switch e {
	case PlayerKilledState_Alive:
		return "Alive", nil
	case PlayerKilledState_Killed:
		return "Killed", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type PlayerKilledState to string", e)
	}
}

// NpcKillStealProtectionState :: Flag to indicate whether you are able to attack an NPC.
type NpcKillStealProtectionState int

const (
	NpcKillStealProtectionState_Unprotected NpcKillStealProtectionState = iota + 1
	NpcKillStealProtectionState_Protected
)

// String converts a NpcKillStealProtectionState value into its string representation
func (e NpcKillStealProtectionState) String() (string, error) {
	switch e {
	case NpcKillStealProtectionState_Unprotected:
		return "Unprotected", nil
	case NpcKillStealProtectionState_Protected:
		return "Protected", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type NpcKillStealProtectionState to string", e)
	}
}

// MapDamageType :: Type of damage being caused by the environment.
type MapDamageType int

const (
	MapDamage_TpDrain MapDamageType = iota + 1
	MapDamage_Spikes
)

// String converts a MapDamageType value into its string representation
func (e MapDamageType) String() (string, error) {
	switch e {
	case MapDamage_TpDrain:
		return "TpDrain", nil
	case MapDamage_Spikes:
		return "Spikes", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MapDamageType to string", e)
	}
}

// MarriageReply :: Reply code sent with MARRIAGE_REPLY packet.
type MarriageReply int

const (
	MarriageReply_AlreadyMarried MarriageReply = iota + 1
	MarriageReply_NotMarried
	MarriageReply_Success
	MarriageReply_NotEnoughGold
	MarriageReply_WrongName
	MarriageReply_ServiceBusy
	MarriageReply_DivorceNotification
)

// String converts a MarriageReply value into its string representation
func (e MarriageReply) String() (string, error) {
	switch e {
	case MarriageReply_AlreadyMarried:
		return "AlreadyMarried", nil
	case MarriageReply_NotMarried:
		return "NotMarried", nil
	case MarriageReply_Success:
		return "Success", nil
	case MarriageReply_NotEnoughGold:
		return "NotEnoughGold", nil
	case MarriageReply_WrongName:
		return "WrongName", nil
	case MarriageReply_ServiceBusy:
		return "ServiceBusy", nil
	case MarriageReply_DivorceNotification:
		return "DivorceNotification", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type MarriageReply to string", e)
	}
}

// PriestReply :: Reply code sent with PRIEST_REPLY packet.
type PriestReply int

const (
	PriestReply_NotDressed PriestReply = iota + 1
	PriestReply_LowLevel
	PriestReply_PartnerNotPresent
	PriestReply_PartnerNotDressed
	PriestReply_Busy
	PriestReply_DoYou
	PriestReply_PartnerAlreadyMarried
	PriestReply_NoPermission
)

// String converts a PriestReply value into its string representation
func (e PriestReply) String() (string, error) {
	switch e {
	case PriestReply_NotDressed:
		return "NotDressed", nil
	case PriestReply_LowLevel:
		return "LowLevel", nil
	case PriestReply_PartnerNotPresent:
		return "PartnerNotPresent", nil
	case PriestReply_PartnerNotDressed:
		return "PartnerNotDressed", nil
	case PriestReply_Busy:
		return "Busy", nil
	case PriestReply_DoYou:
		return "DoYou", nil
	case PriestReply_PartnerAlreadyMarried:
		return "PartnerAlreadyMarried", nil
	case PriestReply_NoPermission:
		return "NoPermission", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type PriestReply to string", e)
	}
}

// PartyReplyCode ::  Reply code sent with PARTY_REPLY packet. Indicates why an invite or join request failed.
type PartyReplyCode int

const (
	PartyReplyCode_AlreadyInAnotherParty PartyReplyCode = iota
	PartyReplyCode_AlreadyInYourParty
	PartyReplyCode_PartyIsFull
)

// String converts a PartyReplyCode value into its string representation
func (e PartyReplyCode) String() (string, error) {
	switch e {
	case PartyReplyCode_AlreadyInAnotherParty:
		return "AlreadyInAnotherParty", nil
	case PartyReplyCode_AlreadyInYourParty:
		return "AlreadyInYourParty", nil
	case PartyReplyCode_PartyIsFull:
		return "PartyIsFull", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type PartyReplyCode to string", e)
	}
}
