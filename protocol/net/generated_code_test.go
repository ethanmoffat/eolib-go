package net_test

import (
	"testing"

	"github.com/ethanmoffat/eolib-go/v3/data"
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

func TestNamedHardcodedFieldSerializesHardcodedValue(t *testing.T) {
	tests := []struct {
		name          string
		requestString string
	}{
		{"Unset", ""},
		{"DifferentValue", "OLD"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packet := client.CharacterRequestClientPacket{RequestString: tt.requestString}

			writer := data.NewEoWriter()
			require.NoError(t, packet.Serialize(writer))

			assert.Equal(t, []byte{'N', 'E', 'W', 0xFF}, writer.Array())
		})
	}
}
