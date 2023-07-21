package net

import "fmt"

// PacketFamily ::  The type of operation that a packet performs. Part of the unique packet ID.
type PacketFamily int

const (
	PacketFamily_Connection PacketFamily = iota + 1
	PacketFamily_Account
	PacketFamily_Character
	PacketFamily_Login
	PacketFamily_Welcome
	PacketFamily_Walk
	PacketFamily_Face
	PacketFamily_Chair
	PacketFamily_Emote
	PacketFamily_Attack        = 11
	PacketFamily_Spell         = 12
	PacketFamily_Shop          = 13
	PacketFamily_Item          = 14
	PacketFamily_StatSkill     = 16
	PacketFamily_Global        = 17
	PacketFamily_Talk          = 18
	PacketFamily_Warp          = 19
	PacketFamily_Jukebox       = 21
	PacketFamily_Players       = 22
	PacketFamily_Avatar        = 23
	PacketFamily_Party         = 24
	PacketFamily_Refresh       = 25
	PacketFamily_Npc           = 26
	PacketFamily_PlayerRange   = 27
	PacketFamily_NpcRange      = 28
	PacketFamily_Range         = 29
	PacketFamily_Paperdoll     = 30
	PacketFamily_Effect        = 31
	PacketFamily_Trade         = 32
	PacketFamily_Chest         = 33
	PacketFamily_Door          = 34
	PacketFamily_Message       = 35
	PacketFamily_Bank          = 36
	PacketFamily_Locker        = 37
	PacketFamily_Barber        = 38
	PacketFamily_Guild         = 39
	PacketFamily_Music         = 40
	PacketFamily_Sit           = 41
	PacketFamily_Recover       = 42
	PacketFamily_Board         = 43
	PacketFamily_Cast          = 44
	PacketFamily_Arena         = 45
	PacketFamily_Priest        = 46
	PacketFamily_Marriage      = 47
	PacketFamily_AdminInteract = 48
	PacketFamily_Citizen       = 49
	PacketFamily_Quest         = 50
	PacketFamily_Book          = 51
	PacketFamily_Error         = 250
	PacketFamily_Init          = 255
)

// String converts a PacketFamily value into its string representation
func (e PacketFamily) String() (string, error) {
	switch e {
	case PacketFamily_Connection:
		return "Connection", nil
	case PacketFamily_Account:
		return "Account", nil
	case PacketFamily_Character:
		return "Character", nil
	case PacketFamily_Login:
		return "Login", nil
	case PacketFamily_Welcome:
		return "Welcome", nil
	case PacketFamily_Walk:
		return "Walk", nil
	case PacketFamily_Face:
		return "Face", nil
	case PacketFamily_Chair:
		return "Chair", nil
	case PacketFamily_Emote:
		return "Emote", nil
	case PacketFamily_Attack:
		return "Attack", nil
	case PacketFamily_Spell:
		return "Spell", nil
	case PacketFamily_Shop:
		return "Shop", nil
	case PacketFamily_Item:
		return "Item", nil
	case PacketFamily_StatSkill:
		return "StatSkill", nil
	case PacketFamily_Global:
		return "Global", nil
	case PacketFamily_Talk:
		return "Talk", nil
	case PacketFamily_Warp:
		return "Warp", nil
	case PacketFamily_Jukebox:
		return "Jukebox", nil
	case PacketFamily_Players:
		return "Players", nil
	case PacketFamily_Avatar:
		return "Avatar", nil
	case PacketFamily_Party:
		return "Party", nil
	case PacketFamily_Refresh:
		return "Refresh", nil
	case PacketFamily_Npc:
		return "Npc", nil
	case PacketFamily_PlayerRange:
		return "PlayerRange", nil
	case PacketFamily_NpcRange:
		return "NpcRange", nil
	case PacketFamily_Range:
		return "Range", nil
	case PacketFamily_Paperdoll:
		return "Paperdoll", nil
	case PacketFamily_Effect:
		return "Effect", nil
	case PacketFamily_Trade:
		return "Trade", nil
	case PacketFamily_Chest:
		return "Chest", nil
	case PacketFamily_Door:
		return "Door", nil
	case PacketFamily_Message:
		return "Message", nil
	case PacketFamily_Bank:
		return "Bank", nil
	case PacketFamily_Locker:
		return "Locker", nil
	case PacketFamily_Barber:
		return "Barber", nil
	case PacketFamily_Guild:
		return "Guild", nil
	case PacketFamily_Music:
		return "Music", nil
	case PacketFamily_Sit:
		return "Sit", nil
	case PacketFamily_Recover:
		return "Recover", nil
	case PacketFamily_Board:
		return "Board", nil
	case PacketFamily_Cast:
		return "Cast", nil
	case PacketFamily_Arena:
		return "Arena", nil
	case PacketFamily_Priest:
		return "Priest", nil
	case PacketFamily_Marriage:
		return "Marriage", nil
	case PacketFamily_AdminInteract:
		return "AdminInteract", nil
	case PacketFamily_Citizen:
		return "Citizen", nil
	case PacketFamily_Quest:
		return "Quest", nil
	case PacketFamily_Book:
		return "Book", nil
	case PacketFamily_Error:
		return "Error", nil
	case PacketFamily_Init:
		return "Init", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type PacketFamily to string", e)
	}
}

// PacketAction ::  The specific action that a packet performs. Part of the unique packet ID.
type PacketAction int

const (
	PacketAction_Request PacketAction = iota + 1
	PacketAction_Accept
	PacketAction_Reply
	PacketAction_Remove
	PacketAction_Agree
	PacketAction_Create
	PacketAction_Add
	PacketAction_Player
	PacketAction_Take
	PacketAction_Use
	PacketAction_Buy
	PacketAction_Sell
	PacketAction_Open
	PacketAction_Close
	PacketAction_Msg
	PacketAction_Spec
	PacketAction_Admin
	PacketAction_List
	PacketAction_Tell        = 20
	PacketAction_Report      = 21
	PacketAction_Announce    = 22
	PacketAction_Server      = 23
	PacketAction_Drop        = 24
	PacketAction_Junk        = 25
	PacketAction_Obtain      = 26
	PacketAction_Get         = 27
	PacketAction_Kick        = 28
	PacketAction_Rank        = 29
	PacketAction_TargetSelf  = 30
	PacketAction_TargetOther = 31
	PacketAction_TargetGroup = 33
	PacketAction_Dialog      = 34
	PacketAction_Ping        = 240
	PacketAction_Pong        = 241
	PacketAction_Net242      = 242
	PacketAction_Net243      = 243
	PacketAction_Net244      = 244
	PacketAction_Error       = 250
	PacketAction_Init        = 255
)

// String converts a PacketAction value into its string representation
func (e PacketAction) String() (string, error) {
	switch e {
	case PacketAction_Request:
		return "Request", nil
	case PacketAction_Accept:
		return "Accept", nil
	case PacketAction_Reply:
		return "Reply", nil
	case PacketAction_Remove:
		return "Remove", nil
	case PacketAction_Agree:
		return "Agree", nil
	case PacketAction_Create:
		return "Create", nil
	case PacketAction_Add:
		return "Add", nil
	case PacketAction_Player:
		return "Player", nil
	case PacketAction_Take:
		return "Take", nil
	case PacketAction_Use:
		return "Use", nil
	case PacketAction_Buy:
		return "Buy", nil
	case PacketAction_Sell:
		return "Sell", nil
	case PacketAction_Open:
		return "Open", nil
	case PacketAction_Close:
		return "Close", nil
	case PacketAction_Msg:
		return "Msg", nil
	case PacketAction_Spec:
		return "Spec", nil
	case PacketAction_Admin:
		return "Admin", nil
	case PacketAction_List:
		return "List", nil
	case PacketAction_Tell:
		return "Tell", nil
	case PacketAction_Report:
		return "Report", nil
	case PacketAction_Announce:
		return "Announce", nil
	case PacketAction_Server:
		return "Server", nil
	case PacketAction_Drop:
		return "Drop", nil
	case PacketAction_Junk:
		return "Junk", nil
	case PacketAction_Obtain:
		return "Obtain", nil
	case PacketAction_Get:
		return "Get", nil
	case PacketAction_Kick:
		return "Kick", nil
	case PacketAction_Rank:
		return "Rank", nil
	case PacketAction_TargetSelf:
		return "TargetSelf", nil
	case PacketAction_TargetOther:
		return "TargetOther", nil
	case PacketAction_TargetGroup:
		return "TargetGroup", nil
	case PacketAction_Dialog:
		return "Dialog", nil
	case PacketAction_Ping:
		return "Ping", nil
	case PacketAction_Pong:
		return "Pong", nil
	case PacketAction_Net242:
		return "Net242", nil
	case PacketAction_Net243:
		return "Net243", nil
	case PacketAction_Net244:
		return "Net244", nil
	case PacketAction_Error:
		return "Error", nil
	case PacketAction_Init:
		return "Init", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type PacketAction to string", e)
	}
}

// QuestPage :: A page in the Quest menu.
type QuestPage int

const (
	QuestPage_Progress QuestPage = iota + 1
	QuestPage_History
)

// String converts a QuestPage value into its string representation
func (e QuestPage) String() (string, error) {
	switch e {
	case QuestPage_Progress:
		return "Progress", nil
	case QuestPage_History:
		return "History", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type QuestPage to string", e)
	}
}

// PartyRequestType ::  Whether a player is requesting to join a party, or inviting someone to join theirs.
type PartyRequestType int

const (
	PartyRequest_Join PartyRequestType = iota
	PartyRequest_Invite
)

// String converts a PartyRequestType value into its string representation
func (e PartyRequestType) String() (string, error) {
	switch e {
	case PartyRequest_Join:
		return "Join", nil
	case PartyRequest_Invite:
		return "Invite", nil
	default:
		return "", fmt.Errorf("could not convert value %d of type PartyRequestType to string", e)
	}
}
