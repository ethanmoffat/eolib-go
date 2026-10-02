package net_test

import (
	"testing"

	"github.com/ethanmoffat/eolib-go/v3/data"
	"github.com/ethanmoffat/eolib-go/v3/protocol"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net/client"
	"github.com/ethanmoffat/eolib-go/v3/protocol/net/server"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Regression tests for generated code paths that the captured packets don't cover.

func TestOptionalStructFieldRoundTrip(t *testing.T) {
	experience := 1234

	tests := []struct {
		name    string
		levelUp *server.LevelUpStats
	}{
		{"Unset", nil},
		{"Set", &server.LevelUpStats{Level: 2, StatPoints: 3, SkillPoints: 4, MaxHp: 5, MaxTp: 6, MaxSp: 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packet := server.NpcAcceptServerPacket{Experience: &experience, LevelUp: tt.levelUp}

			writer := data.NewEoWriter()
			require.NoError(t, packet.Serialize(writer))

			var actual server.NpcAcceptServerPacket
			require.NoError(t, actual.Deserialize(data.NewEoReader(writer.Array())))

			require.NotNil(t, actual.Experience)
			assert.Equal(t, experience, *actual.Experience)
			if tt.levelUp == nil {
				assert.Nil(t, actual.LevelUp)
			} else {
				require.NotNil(t, actual.LevelUp)
				assert.Equal(t, tt.levelUp.Level, actual.LevelUp.Level)
				assert.Equal(t, tt.levelUp.MaxSp, actual.LevelUp.MaxSp)
			}
		})
	}
}

func TestOptionalStructFieldTruncatedData(t *testing.T) {
	var packet server.NpcAcceptServerPacket
	require.NoError(t, packet.Deserialize(data.NewEoReader([]byte{})))

	assert.Nil(t, packet.Experience)
	assert.Nil(t, packet.LevelUp)
}

func TestNamedHardcodedFieldZeroValueSerializesDefault(t *testing.T) {
	var packet client.CharacterRequestClientPacket

	writer := data.NewEoWriter()
	require.NoError(t, packet.Serialize(writer))

	assert.Equal(t, []byte{'N', 'E', 'W', 0xFF}, writer.Array())
}

func TestNamedHardcodedFieldSetValueIsSerialized(t *testing.T) {
	packet := client.CharacterRequestClientPacket{RequestString: "OLD"}

	writer := data.NewEoWriter()
	require.NoError(t, packet.Serialize(writer))

	assert.Equal(t, []byte{'O', 'L', 'D', 0xFF}, writer.Array())
}

func TestNamedHardcodedFieldDeserializedValueRoundTrips(t *testing.T) {
	tests := []struct {
		name          string
		input         []byte
		requestString string
	}{
		{"DefaultValue", []byte{'N', 'E', 'W', 0xFF}, "NEW"},
		{"DifferentValue", []byte{'O', 'L', 'D', 0xFF}, "OLD"},
		{"ZeroValue", []byte{0xFF}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var packet client.CharacterRequestClientPacket
			require.NoError(t, packet.Deserialize(data.NewEoReader(tt.input)))
			assert.Equal(t, tt.requestString, packet.RequestString)

			writer := data.NewEoWriter()
			require.NoError(t, packet.Serialize(writer))
			assert.Equal(t, tt.input, writer.Array())
		})
	}
}

func TestSwitchFactoryRoundTrip(t *testing.T) {
	tests := []struct {
		name   string
		packet protocol.EoData
		target protocol.EoData
		check  func(t *testing.T, actual protocol.EoData)
	}{
		{
			"LoginReplyOk",
			server.NewLoginReplyWithOk(&server.LoginReplyReplyCodeDataOk{
				Characters: []server.CharacterSelectionListEntry{{Name: "ethan", Id: 1, Level: 2}},
			}),
			&server.LoginReplyServerPacket{},
			func(t *testing.T, actual protocol.EoData) {
				p := actual.(*server.LoginReplyServerPacket)
				assert.Equal(t, server.LoginReply_Ok, p.ReplyCode)
				ok, isOk := p.ReplyCodeData.(*server.LoginReplyReplyCodeDataOk)
				require.True(t, isOk)
				require.Len(t, ok.Characters, 1)
				assert.Equal(t, "ethan", ok.Characters[0].Name)
				assert.Equal(t, 1, ok.Characters[0].Id)
				assert.Equal(t, 2, ok.Characters[0].Level)
			},
		},
		{
			"LoginReplyWrongUser",
			server.NewLoginReplyWithWrongUser(),
			&server.LoginReplyServerPacket{},
			func(t *testing.T, actual protocol.EoData) {
				p := actual.(*server.LoginReplyServerPacket)
				assert.Equal(t, server.LoginReply_WrongUser, p.ReplyCode)
				assert.IsType(t, &server.LoginReplyReplyCodeDataWrongUser{}, p.ReplyCodeData)
			},
		},
		{
			"InitInitBannedTemporary",
			server.NewInitInitWithBannedTemporary(&server.InitInitBanTypeDataTemporary{MinutesRemaining: 30}),
			&server.InitInitServerPacket{},
			func(t *testing.T, actual protocol.EoData) {
				p := actual.(*server.InitInitServerPacket)
				assert.Equal(t, server.InitReply_Banned, p.ReplyCode)
				banned, isBanned := p.ReplyCodeData.(*server.InitInitReplyCodeDataBanned)
				require.True(t, isBanned)
				assert.Equal(t, server.InitBan_Temporary, banned.BanType)
				temporary, isTemporary := banned.BanTypeData.(*server.InitInitBanTypeDataTemporary)
				require.True(t, isTemporary)
				assert.Equal(t, 30, temporary.MinutesRemaining)
			},
		},
		{
			"InitInitBannedPermanent",
			server.NewInitInitWithBannedPermanent(),
			&server.InitInitServerPacket{},
			func(t *testing.T, actual protocol.EoData) {
				p := actual.(*server.InitInitServerPacket)
				assert.Equal(t, server.InitReply_Banned, p.ReplyCode)
				banned, isBanned := p.ReplyCodeData.(*server.InitInitReplyCodeDataBanned)
				require.True(t, isBanned)
				assert.Equal(t, server.InitBan_Permanent, banned.BanType)
				assert.Nil(t, banned.BanTypeData)
			},
		},
		{
			"InitInitBanTypeData0",
			server.NewInitInitWithBanTypeData0(&server.InitInitBanTypeData0{MinutesRemaining: 5}),
			&server.InitInitServerPacket{},
			func(t *testing.T, actual protocol.EoData) {
				p := actual.(*server.InitInitServerPacket)
				assert.Equal(t, server.InitReply_Banned, p.ReplyCode)
				banned, isBanned := p.ReplyCodeData.(*server.InitInitReplyCodeDataBanned)
				require.True(t, isBanned)
				assert.Equal(t, server.InitBanType(0), banned.BanType)
				zero, isZero := banned.BanTypeData.(*server.InitInitBanTypeData0)
				require.True(t, isZero)
				assert.Equal(t, 5, zero.MinutesRemaining)
			},
		},
		{
			"AccountReplyDefault",
			mustNewAccountReplyWithDefault(t, 1000, &server.AccountReplyReplyCodeDataDefault{SequenceStart: 12}),
			&server.AccountReplyServerPacket{},
			func(t *testing.T, actual protocol.EoData) {
				p := actual.(*server.AccountReplyServerPacket)
				assert.Equal(t, server.AccountReply(1000), p.ReplyCode)
				def, isDefault := p.ReplyCodeData.(*server.AccountReplyReplyCodeDataDefault)
				require.True(t, isDefault)
				assert.Equal(t, 12, def.SequenceStart)
			},
		},
		{
			"DialogEntryLink",
			withLine(server.NewDialogEntryWithLink(&server.EntryTypeDataLink{LinkId: 7}), "Hello"),
			&server.DialogEntry{},
			func(t *testing.T, actual protocol.EoData) {
				p := actual.(*server.DialogEntry)
				assert.Equal(t, server.DialogEntry_Link, p.EntryType)
				link, isLink := p.EntryTypeData.(*server.EntryTypeDataLink)
				require.True(t, isLink)
				assert.Equal(t, 7, link.LinkId)
				assert.Equal(t, "Hello", p.Line)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := data.NewEoWriter()
			require.NoError(t, tt.packet.Serialize(writer))

			require.NoError(t, tt.target.Deserialize(data.NewEoReader(writer.Array())))
			tt.check(t, tt.target)
		})
	}
}

func TestSwitchFactoryNilDataFailsSerialization(t *testing.T) {
	tests := []struct {
		name   string
		packet protocol.EoData
	}{
		{"LoginReplyOk", server.NewLoginReplyWithOk(nil)},
		{"InitInitBannedTemporary", server.NewInitInitWithBannedTemporary(nil)},
		{"AccountReplyDefault", mustNewAccountReplyWithDefault(t, 1000, nil)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Error(t, tt.packet.Serialize(data.NewEoWriter()))
		})
	}
}

func TestSwitchFactoryDefaultInvalidArgumentsReturnError(t *testing.T) {
	tests := []struct {
		name     string
		code     server.AccountReply
		data     *server.AccountReplyReplyCodeDataDefault
		expected string
	}{
		{"NamedCase", server.AccountReply_Exists, &server.AccountReplyReplyCodeDataDefault{}, "ReplyCode 1 has its own case and is not handled by the default case"},
		{"NumericCase", 4, &server.AccountReplyReplyCodeDataDefault{}, "ReplyCode 4 has its own case and is not handled by the default case"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packet, err := server.NewAccountReplyWithDefault(tt.code, tt.data)

			assert.Nil(t, packet)
			assert.EqualError(t, err, tt.expected)
		})
	}
}

func mustNewAccountReplyWithDefault(t *testing.T, code server.AccountReply, d *server.AccountReplyReplyCodeDataDefault) *server.AccountReplyServerPacket {
	t.Helper()

	packet, err := server.NewAccountReplyWithDefault(code, d)
	require.NoError(t, err)
	return packet
}

func withLine(entry *server.DialogEntry, line string) *server.DialogEntry {
	entry.Line = line
	return entry
}
