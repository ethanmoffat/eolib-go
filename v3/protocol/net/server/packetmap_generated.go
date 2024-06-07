package server

import (
	"fmt"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net"
	"reflect"
)

var packetMap = map[int]reflect.Type{
	net.PacketId(net.PacketFamily_Init, net.PacketAction_Init):            reflect.TypeOf(InitInitServerPacket{}),
	net.PacketId(net.PacketFamily_Warp, net.PacketAction_Player):          reflect.TypeOf(WarpPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Ping):         reflect.TypeOf(WelcomePingServerPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Pong):         reflect.TypeOf(WelcomePongServerPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Net242):       reflect.TypeOf(WelcomeNet242ServerPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Net243):       reflect.TypeOf(WelcomeNet243ServerPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_List):         reflect.TypeOf(PlayersListServerPacket{}),
	net.PacketId(net.PacketFamily_Warp, net.PacketAction_Create):          reflect.TypeOf(WarpCreateServerPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Reply):        reflect.TypeOf(PlayersReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Net244):       reflect.TypeOf(WelcomeNet244ServerPacket{}),
	net.PacketId(net.PacketFamily_Connection, net.PacketAction_Player):    reflect.TypeOf(ConnectionPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Account, net.PacketAction_Reply):        reflect.TypeOf(AccountReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Character, net.PacketAction_Reply):      reflect.TypeOf(CharacterReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Character, net.PacketAction_Player):     reflect.TypeOf(CharacterPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Login, net.PacketAction_Reply):          reflect.TypeOf(LoginReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Welcome, net.PacketAction_Reply):        reflect.TypeOf(WelcomeReplyServerPacket{}),
	net.PacketId(net.PacketFamily_AdminInteract, net.PacketAction_Reply):  reflect.TypeOf(AdminInteractReplyServerPacket{}),
	net.PacketId(net.PacketFamily_AdminInteract, net.PacketAction_Remove): reflect.TypeOf(AdminInteractRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_AdminInteract, net.PacketAction_Agree):  reflect.TypeOf(AdminInteractAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_AdminInteract, net.PacketAction_List):   reflect.TypeOf(AdminInteractListServerPacket{}),
	net.PacketId(net.PacketFamily_AdminInteract, net.PacketAction_Tell):   reflect.TypeOf(AdminInteractTellServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Request):         reflect.TypeOf(TalkRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Open):            reflect.TypeOf(TalkOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Msg):             reflect.TypeOf(TalkMsgServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Tell):            reflect.TypeOf(TalkTellServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Player):          reflect.TypeOf(TalkPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Reply):           reflect.TypeOf(TalkReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Admin):           reflect.TypeOf(TalkAdminServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Announce):        reflect.TypeOf(TalkAnnounceServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Server):          reflect.TypeOf(TalkServerServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_List):            reflect.TypeOf(TalkListServerPacket{}),
	net.PacketId(net.PacketFamily_Message, net.PacketAction_Open):         reflect.TypeOf(MessageOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Message, net.PacketAction_Close):        reflect.TypeOf(MessageCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Message, net.PacketAction_Accept):       reflect.TypeOf(MessageAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_Talk, net.PacketAction_Spec):            reflect.TypeOf(TalkSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Attack, net.PacketAction_Player):        reflect.TypeOf(AttackPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Attack, net.PacketAction_Error):         reflect.TypeOf(AttackErrorServerPacket{}),
	net.PacketId(net.PacketFamily_Avatar, net.PacketAction_Reply):         reflect.TypeOf(AvatarReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Chair, net.PacketAction_Player):         reflect.TypeOf(ChairPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Chair, net.PacketAction_Reply):          reflect.TypeOf(ChairReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Chair, net.PacketAction_Close):          reflect.TypeOf(ChairCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Chair, net.PacketAction_Remove):         reflect.TypeOf(ChairRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Sit, net.PacketAction_Player):           reflect.TypeOf(SitPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Sit, net.PacketAction_Close):            reflect.TypeOf(SitCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Sit, net.PacketAction_Remove):           reflect.TypeOf(SitRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Sit, net.PacketAction_Reply):            reflect.TypeOf(SitReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Emote, net.PacketAction_Player):         reflect.TypeOf(EmotePlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Effect, net.PacketAction_Player):        reflect.TypeOf(EffectPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Face, net.PacketAction_Player):          reflect.TypeOf(FacePlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Avatar, net.PacketAction_Remove):        reflect.TypeOf(AvatarRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Agree):        reflect.TypeOf(PlayersAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Remove):       reflect.TypeOf(PlayersRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Range, net.PacketAction_Reply):          reflect.TypeOf(RangeReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Npc, net.PacketAction_Agree):            reflect.TypeOf(NpcAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Walk, net.PacketAction_Player):          reflect.TypeOf(WalkPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Walk, net.PacketAction_Reply):           reflect.TypeOf(WalkReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Walk, net.PacketAction_Close):           reflect.TypeOf(WalkCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Walk, net.PacketAction_Open):            reflect.TypeOf(WalkOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Bank, net.PacketAction_Open):            reflect.TypeOf(BankOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Bank, net.PacketAction_Reply):           reflect.TypeOf(BankReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Barber, net.PacketAction_Agree):         reflect.TypeOf(BarberAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Barber, net.PacketAction_Open):          reflect.TypeOf(BarberOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Reply):         reflect.TypeOf(LockerReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Get):           reflect.TypeOf(LockerGetServerPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Open):          reflect.TypeOf(LockerOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Buy):           reflect.TypeOf(LockerBuyServerPacket{}),
	net.PacketId(net.PacketFamily_Locker, net.PacketAction_Spec):          reflect.TypeOf(LockerSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Reply):        reflect.TypeOf(CitizenReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Remove):       reflect.TypeOf(CitizenRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Open):         reflect.TypeOf(CitizenOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Request):      reflect.TypeOf(CitizenRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Citizen, net.PacketAction_Accept):       reflect.TypeOf(CitizenAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Create):          reflect.TypeOf(ShopCreateServerPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Buy):             reflect.TypeOf(ShopBuyServerPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Sell):            reflect.TypeOf(ShopSellServerPacket{}),
	net.PacketId(net.PacketFamily_Shop, net.PacketAction_Open):            reflect.TypeOf(ShopOpenServerPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Open):       reflect.TypeOf(StatSkillOpenServerPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Reply):      reflect.TypeOf(StatSkillReplyServerPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Take):       reflect.TypeOf(StatSkillTakeServerPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Remove):     reflect.TypeOf(StatSkillRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Player):     reflect.TypeOf(StatSkillPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Accept):     reflect.TypeOf(StatSkillAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_StatSkill, net.PacketAction_Junk):       reflect.TypeOf(StatSkillJunkServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Reply):           reflect.TypeOf(ItemReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Drop):            reflect.TypeOf(ItemDropServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Add):             reflect.TypeOf(ItemAddServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Remove):          reflect.TypeOf(ItemRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Junk):            reflect.TypeOf(ItemJunkServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Get):             reflect.TypeOf(ItemGetServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Obtain):          reflect.TypeOf(ItemObtainServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Kick):            reflect.TypeOf(ItemKickServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Agree):           reflect.TypeOf(ItemAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Spec):            reflect.TypeOf(ItemSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Board, net.PacketAction_Player):         reflect.TypeOf(BoardPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Board, net.PacketAction_Open):           reflect.TypeOf(BoardOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Agree):        reflect.TypeOf(JukeboxAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Reply):        reflect.TypeOf(JukeboxReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Open):         reflect.TypeOf(JukeboxOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Msg):          reflect.TypeOf(JukeboxMsgServerPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Player):       reflect.TypeOf(JukeboxPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Jukebox, net.PacketAction_Use):          reflect.TypeOf(JukeboxUseServerPacket{}),
	net.PacketId(net.PacketFamily_Warp, net.PacketAction_Request):         reflect.TypeOf(WarpRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Warp, net.PacketAction_Agree):           reflect.TypeOf(WarpAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Paperdoll, net.PacketAction_Reply):      reflect.TypeOf(PaperdollReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Paperdoll, net.PacketAction_Ping):       reflect.TypeOf(PaperdollPingServerPacket{}),
	net.PacketId(net.PacketFamily_Paperdoll, net.PacketAction_Remove):     reflect.TypeOf(PaperdollRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Paperdoll, net.PacketAction_Agree):      reflect.TypeOf(PaperdollAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Avatar, net.PacketAction_Agree):         reflect.TypeOf(AvatarAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Book, net.PacketAction_Reply):           reflect.TypeOf(BookReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Message, net.PacketAction_Pong):         reflect.TypeOf(MessagePongServerPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Ping):         reflect.TypeOf(PlayersPingServerPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Pong):         reflect.TypeOf(PlayersPongServerPacket{}),
	net.PacketId(net.PacketFamily_Players, net.PacketAction_Net242):       reflect.TypeOf(PlayersNet242ServerPacket{}),
	net.PacketId(net.PacketFamily_Door, net.PacketAction_Open):            reflect.TypeOf(DoorOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Door, net.PacketAction_Close):           reflect.TypeOf(DoorCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Open):           reflect.TypeOf(ChestOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Reply):          reflect.TypeOf(ChestReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Get):            reflect.TypeOf(ChestGetServerPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Agree):          reflect.TypeOf(ChestAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Spec):           reflect.TypeOf(ChestSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Chest, net.PacketAction_Close):          reflect.TypeOf(ChestCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Refresh, net.PacketAction_Reply):        reflect.TypeOf(RefreshReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Request):        reflect.TypeOf(PartyRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Reply):          reflect.TypeOf(PartyReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Create):         reflect.TypeOf(PartyCreateServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Add):            reflect.TypeOf(PartyAddServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Remove):         reflect.TypeOf(PartyRemoveServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Close):          reflect.TypeOf(PartyCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_List):           reflect.TypeOf(PartyListServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_Agree):          reflect.TypeOf(PartyAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Party, net.PacketAction_TargetGroup):    reflect.TypeOf(PartyTargetGroupServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Reply):          reflect.TypeOf(GuildReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Request):        reflect.TypeOf(GuildRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Create):         reflect.TypeOf(GuildCreateServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Take):           reflect.TypeOf(GuildTakeServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Rank):           reflect.TypeOf(GuildRankServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Sell):           reflect.TypeOf(GuildSellServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Buy):            reflect.TypeOf(GuildBuyServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Open):           reflect.TypeOf(GuildOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Tell):           reflect.TypeOf(GuildTellServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Report):         reflect.TypeOf(GuildReportServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Agree):          reflect.TypeOf(GuildAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Accept):         reflect.TypeOf(GuildAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_Guild, net.PacketAction_Kick):           reflect.TypeOf(GuildKickServerPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_Request):        reflect.TypeOf(SpellRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_TargetSelf):     reflect.TypeOf(SpellTargetSelfServerPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_Player):         reflect.TypeOf(SpellPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_Error):          reflect.TypeOf(SpellErrorServerPacket{}),
	net.PacketId(net.PacketFamily_Avatar, net.PacketAction_Admin):         reflect.TypeOf(AvatarAdminServerPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_TargetGroup):    reflect.TypeOf(SpellTargetGroupServerPacket{}),
	net.PacketId(net.PacketFamily_Spell, net.PacketAction_TargetOther):    reflect.TypeOf(SpellTargetOtherServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Request):        reflect.TypeOf(TradeRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Open):           reflect.TypeOf(TradeOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Reply):          reflect.TypeOf(TradeReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Admin):          reflect.TypeOf(TradeAdminServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Use):            reflect.TypeOf(TradeUseServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Spec):           reflect.TypeOf(TradeSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Agree):          reflect.TypeOf(TradeAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Trade, net.PacketAction_Close):          reflect.TypeOf(TradeCloseServerPacket{}),
	net.PacketId(net.PacketFamily_Npc, net.PacketAction_Reply):            reflect.TypeOf(NpcReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Cast, net.PacketAction_Reply):           reflect.TypeOf(CastReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Npc, net.PacketAction_Spec):             reflect.TypeOf(NpcSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Npc, net.PacketAction_Accept):           reflect.TypeOf(NpcAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_Cast, net.PacketAction_Spec):            reflect.TypeOf(CastSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Cast, net.PacketAction_Accept):          reflect.TypeOf(CastAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_Npc, net.PacketAction_Junk):             reflect.TypeOf(NpcJunkServerPacket{}),
	net.PacketId(net.PacketFamily_Npc, net.PacketAction_Player):           reflect.TypeOf(NpcPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Npc, net.PacketAction_Dialog):           reflect.TypeOf(NpcDialogServerPacket{}),
	net.PacketId(net.PacketFamily_Quest, net.PacketAction_Report):         reflect.TypeOf(QuestReportServerPacket{}),
	net.PacketId(net.PacketFamily_Quest, net.PacketAction_Dialog):         reflect.TypeOf(QuestDialogServerPacket{}),
	net.PacketId(net.PacketFamily_Quest, net.PacketAction_List):           reflect.TypeOf(QuestListServerPacket{}),
	net.PacketId(net.PacketFamily_Item, net.PacketAction_Accept):          reflect.TypeOf(ItemAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_Arena, net.PacketAction_Drop):           reflect.TypeOf(ArenaDropServerPacket{}),
	net.PacketId(net.PacketFamily_Arena, net.PacketAction_Use):            reflect.TypeOf(ArenaUseServerPacket{}),
	net.PacketId(net.PacketFamily_Arena, net.PacketAction_Spec):           reflect.TypeOf(ArenaSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Arena, net.PacketAction_Accept):         reflect.TypeOf(ArenaAcceptServerPacket{}),
	net.PacketId(net.PacketFamily_Marriage, net.PacketAction_Open):        reflect.TypeOf(MarriageOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Marriage, net.PacketAction_Reply):       reflect.TypeOf(MarriageReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Priest, net.PacketAction_Open):          reflect.TypeOf(PriestOpenServerPacket{}),
	net.PacketId(net.PacketFamily_Priest, net.PacketAction_Reply):         reflect.TypeOf(PriestReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Priest, net.PacketAction_Request):       reflect.TypeOf(PriestRequestServerPacket{}),
	net.PacketId(net.PacketFamily_Recover, net.PacketAction_Player):       reflect.TypeOf(RecoverPlayerServerPacket{}),
	net.PacketId(net.PacketFamily_Recover, net.PacketAction_Agree):        reflect.TypeOf(RecoverAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Recover, net.PacketAction_List):         reflect.TypeOf(RecoverListServerPacket{}),
	net.PacketId(net.PacketFamily_Recover, net.PacketAction_Reply):        reflect.TypeOf(RecoverReplyServerPacket{}),
	net.PacketId(net.PacketFamily_Recover, net.PacketAction_TargetGroup):  reflect.TypeOf(RecoverTargetGroupServerPacket{}),
	net.PacketId(net.PacketFamily_Effect, net.PacketAction_Use):           reflect.TypeOf(EffectUseServerPacket{}),
	net.PacketId(net.PacketFamily_Effect, net.PacketAction_Agree):         reflect.TypeOf(EffectAgreeServerPacket{}),
	net.PacketId(net.PacketFamily_Effect, net.PacketAction_TargetOther):   reflect.TypeOf(EffectTargetOtherServerPacket{}),
	net.PacketId(net.PacketFamily_Effect, net.PacketAction_Report):        reflect.TypeOf(EffectReportServerPacket{}),
	net.PacketId(net.PacketFamily_Effect, net.PacketAction_Spec):          reflect.TypeOf(EffectSpecServerPacket{}),
	net.PacketId(net.PacketFamily_Effect, net.PacketAction_Admin):         reflect.TypeOf(EffectAdminServerPacket{}),
	net.PacketId(net.PacketFamily_Music, net.PacketAction_Player):         reflect.TypeOf(MusicPlayerServerPacket{}),
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
