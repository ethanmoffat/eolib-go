package client

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net"
	"reflect"
)

var packetMap = map[int]reflect.Type{
	net.PacketId(net.PacketFamily_Init, net.PacketAction_Init):            reflect.TypeOf(InitInitClientPacket{}),
	net.PacketId(net.PacketFamily_Connection, net.PacketAction_Accept):    reflect.TypeOf(ConnectionAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Connection, net.PacketAction_Ping):      reflect.TypeOf(ConnectionPingClientPacket{}),
	net.PacketId(net.PacketFamily_Account, net.PacketAction_Request):      reflect.TypeOf(AccountRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Account, net.PacketAction_Create):       reflect.TypeOf(AccountCreateClientPacket{}),
	net.PacketId(net.PacketFamily_Account, net.PacketAction_Agree):        reflect.TypeOf(AccountAgreeClientPacket{}),
	net.PacketId(net.PacketFamily_Character, net.PacketAction_Request):    reflect.TypeOf(CharacterRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Character, net.PacketAction_Create):     reflect.TypeOf(CharacterCreateClientPacket{}),
	net.PacketId(net.PacketFamily_Character, net.PacketAction_Take):       reflect.TypeOf(CharacterTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Character, net.PacketAction_Remove):     reflect.TypeOf(CharacterRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Login, net.PacketAction_Request):        reflect.TypeOf(LoginRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Request):      reflect.TypeOf(WelcomeRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Msg):          reflect.TypeOf(WelcomeMsgClientPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Agree):        reflect.TypeOf(WelcomeAgreeClientPacket{}),
	net.PacketId(net.PacketFamily_AdminInteract, net.PacketAction_Tell):   reflect.TypeOf(AdminInteractTellClientPacket{}),
	net.PacketId(net.PacketFamily_AdminInteract, net.PacketAction_Report): reflect.TypeOf(AdminInteractReportClientPacket{}),
	net.PacketId(net.PacketFamily_Global, net.PacketAction_Remove):        reflect.TypeOf(GlobalRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Global, net.PacketAction_Player):        reflect.TypeOf(GlobalPlayerClientPacket{}),
	net.PacketId(net.PacketFamily_Global, net.PacketAction_Open):          reflect.TypeOf(GlobalOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Global, net.PacketAction_Close):         reflect.TypeOf(GlobalCloseClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Request):         reflect.TypeOf(TalkRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Open):            reflect.TypeOf(TalkOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Msg):             reflect.TypeOf(TalkMsgClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Tell):            reflect.TypeOf(TalkTellClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Report):          reflect.TypeOf(TalkReportClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Player):          reflect.TypeOf(TalkPlayerClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Use):             reflect.TypeOf(TalkUseClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Admin):           reflect.TypeOf(TalkAdminClientPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Announce):        reflect.TypeOf(TalkAnnounceClientPacket{}),
	net.PacketId(net.PacketFamily_Attack, net.PacketAction_Use):           reflect.TypeOf(AttackUseClientPacket{}),
	net.PacketId(net.PacketFamily_Chair, net.PacketAction_Request):        reflect.TypeOf(ChairRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Sit, net.PacketAction_Request):          reflect.TypeOf(SitRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Emote, net.PacketAction_Report):         reflect.TypeOf(EmoteReportClientPacket{}),
	net.PacketId(net.PacketFamily_Face, net.PacketAction_Player):          reflect.TypeOf(FacePlayerClientPacket{}),
	net.PacketId(net.PacketFamily_Walk, net.PacketAction_Admin):           reflect.TypeOf(WalkAdminClientPacket{}),
	net.PacketId(net.PacketFamily_Walk, net.PacketAction_Spec):            reflect.TypeOf(WalkSpecClientPacket{}),
	net.PacketId(net.PacketFamily_Walk, net.PacketAction_Player):          reflect.TypeOf(WalkPlayerClientPacket{}),
	net.PacketId(net.PacketFamily_Bank, net.PacketAction_Open):            reflect.TypeOf(BankOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Bank, net.PacketAction_Add):             reflect.TypeOf(BankAddClientPacket{}),
	net.PacketId(net.PacketFamily_Bank, net.PacketAction_Take):            reflect.TypeOf(BankTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Barber, net.PacketAction_Buy):           reflect.TypeOf(BarberBuyClientPacket{}),
	net.PacketId(net.PacketFamily_Barber, net.PacketAction_Open):          reflect.TypeOf(BarberOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Add):           reflect.TypeOf(LockerAddClientPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Take):          reflect.TypeOf(LockerTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Open):          reflect.TypeOf(LockerOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Buy):           reflect.TypeOf(LockerBuyClientPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Request):      reflect.TypeOf(CitizenRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Accept):       reflect.TypeOf(CitizenAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Reply):        reflect.TypeOf(CitizenReplyClientPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Remove):       reflect.TypeOf(CitizenRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Open):         reflect.TypeOf(CitizenOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Create):          reflect.TypeOf(ShopCreateClientPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Buy):             reflect.TypeOf(ShopBuyClientPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Sell):            reflect.TypeOf(ShopSellClientPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Open):            reflect.TypeOf(ShopOpenClientPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Open):       reflect.TypeOf(StatSkillOpenClientPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Take):       reflect.TypeOf(StatSkillTakeClientPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Remove):     reflect.TypeOf(StatSkillRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Add):        reflect.TypeOf(StatSkillAddClientPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Junk):       reflect.TypeOf(StatSkillJunkClientPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Use):             reflect.TypeOf(ItemUseClientPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Drop):            reflect.TypeOf(ItemDropClientPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Junk):            reflect.TypeOf(ItemJunkClientPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Get):             reflect.TypeOf(ItemGetClientPacket{}),
	net.PacketId(net.PacketFamily_Board, net.PacketAction_Remove):         reflect.TypeOf(BoardRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Board, net.PacketAction_Create):         reflect.TypeOf(BoardCreateClientPacket{}),
	net.PacketId(net.PacketFamily_Board, net.PacketAction_Take):           reflect.TypeOf(BoardTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Board, net.PacketAction_Open):           reflect.TypeOf(BoardOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Open):         reflect.TypeOf(JukeboxOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Msg):          reflect.TypeOf(JukeboxMsgClientPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Use):          reflect.TypeOf(JukeboxUseClientPacket{}),
	net.PacketId(net.PacketFamily_Warp, net.PacketAction_Accept):          reflect.TypeOf(WarpAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Warp, net.PacketAction_Take):            reflect.TypeOf(WarpTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Paperdoll, net.PacketAction_Request):    reflect.TypeOf(PaperdollRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Paperdoll, net.PacketAction_Remove):     reflect.TypeOf(PaperdollRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Paperdoll, net.PacketAction_Add):        reflect.TypeOf(PaperdollAddClientPacket{}),
	net.PacketId(net.PacketFamily_Book, net.PacketAction_Request):         reflect.TypeOf(BookRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Message, net.PacketAction_Ping):         reflect.TypeOf(MessagePingClientPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Accept):       reflect.TypeOf(PlayersAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Request):      reflect.TypeOf(PlayersRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_List):         reflect.TypeOf(PlayersListClientPacket{}),
	net.PacketId(net.PacketFamily_Door, net.PacketAction_Open):            reflect.TypeOf(DoorOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Open):           reflect.TypeOf(ChestOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Add):            reflect.TypeOf(ChestAddClientPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Take):           reflect.TypeOf(ChestTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Refresh, net.PacketAction_Request):      reflect.TypeOf(RefreshRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Range, net.PacketAction_Request):        reflect.TypeOf(RangeRequestClientPacket{}),
	net.PacketId(net.PacketFamily_PlayerRange, net.PacketAction_Request):  reflect.TypeOf(PlayerRangeRequestClientPacket{}),
	net.PacketId(net.PacketFamily_NpcRange, net.PacketAction_Request):     reflect.TypeOf(NpcRangeRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Request):        reflect.TypeOf(PartyRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Accept):         reflect.TypeOf(PartyAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Remove):         reflect.TypeOf(PartyRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Take):           reflect.TypeOf(PartyTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Request):        reflect.TypeOf(GuildRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Accept):         reflect.TypeOf(GuildAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Remove):         reflect.TypeOf(GuildRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Agree):          reflect.TypeOf(GuildAgreeClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Create):         reflect.TypeOf(GuildCreateClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Player):         reflect.TypeOf(GuildPlayerClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Take):           reflect.TypeOf(GuildTakeClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Use):            reflect.TypeOf(GuildUseClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Buy):            reflect.TypeOf(GuildBuyClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Open):           reflect.TypeOf(GuildOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Tell):           reflect.TypeOf(GuildTellClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Report):         reflect.TypeOf(GuildReportClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Junk):           reflect.TypeOf(GuildJunkClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Kick):           reflect.TypeOf(GuildKickClientPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Rank):           reflect.TypeOf(GuildRankClientPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_Request):        reflect.TypeOf(SpellRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_TargetSelf):     reflect.TypeOf(SpellTargetSelfClientPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_TargetOther):    reflect.TypeOf(SpellTargetOtherClientPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_TargetGroup):    reflect.TypeOf(SpellTargetGroupClientPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_Use):            reflect.TypeOf(SpellUseClientPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Request):        reflect.TypeOf(TradeRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Accept):         reflect.TypeOf(TradeAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Remove):         reflect.TypeOf(TradeRemoveClientPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Agree):          reflect.TypeOf(TradeAgreeClientPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Add):            reflect.TypeOf(TradeAddClientPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Close):          reflect.TypeOf(TradeCloseClientPacket{}),
	net.PacketId(net.PacketFamily_Quest, net.PacketAction_Use):            reflect.TypeOf(QuestUseClientPacket{}),
	net.PacketId(net.PacketFamily_Quest, net.PacketAction_Accept):         reflect.TypeOf(QuestAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Quest, net.PacketAction_List):           reflect.TypeOf(QuestListClientPacket{}),
	net.PacketId(net.PacketFamily_Marriage, net.PacketAction_Open):        reflect.TypeOf(MarriageOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Marriage, net.PacketAction_Request):     reflect.TypeOf(MarriageRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Priest, net.PacketAction_Accept):        reflect.TypeOf(PriestAcceptClientPacket{}),
	net.PacketId(net.PacketFamily_Priest, net.PacketAction_Open):          reflect.TypeOf(PriestOpenClientPacket{}),
	net.PacketId(net.PacketFamily_Priest, net.PacketAction_Request):       reflect.TypeOf(PriestRequestClientPacket{}),
	net.PacketId(net.PacketFamily_Priest, net.PacketAction_Use):           reflect.TypeOf(PriestUseClientPacket{}),
}

// PacketFromId creates a typed packet instance from a [net.PacketFamily] and [net.PacketAction].
// This function calls [PacketFromIntegerId] internally.
func PacketFromId(family net.PacketFamily, action net.PacketAction) (net.Packet, error) {
	return PacketFromIntegerId(net.PacketId(family, action))
}

// PacketFromIntegerId creates a typed packet instance from a packet's ID. An ID may be converted from a family/action pair via the [net.PacketId] function.
// The returned packet implements the [net.Packet] interface. It may be serialized/deserialized without further conversion, or a type assertion may be made to examine the data. The expected type of the assertion is a pointer to a packet structure.
// The following example does both: an incoming CHAIR_REQUEST packet is deserialized from a reader without converting from the interface type, and the data is examined via a type assertion.
//
//	pkt, _ := client.PacketFromId(net.PacketFamily_Chair, net.PacketAction_Request)
//	if err = pkt.Deserialize(reader); err != nil {
//	  // handle the error
//	}
//	switch v := pkt.(type) {
//	case *client.ChairRequestClientPacket:
//	   fmt.Println("SitAction=", v.SitAction)
//	   switch d := v.SitActionData.(type) {
//	   case *client.ChairRequestSitActionDataSit:
//	     fmt.Println("Data.Coords=", v.Data.Coords)
//	   }
//	default:
//	  fmt.Printf("Unknown type: %s\n", reflect.TypeOf(pkt).Elem().Name())
//	}
func PacketFromIntegerId(id int) (net.Packet, error) {
	packetType, idOk := packetMap[id]
	if !idOk {
		return nil, fmt.Errorf("could not find packet with id %d", id)
	}

	packetInstance, typeOk := reflect.New(packetType).Interface().(net.Packet)
	if !typeOk {
		return nil, fmt.Errorf("could not create packet from id %d", id)
	}

	return packetInstance, nil
}
