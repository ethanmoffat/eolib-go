package packet_test

import (
	"math/rand"
	"testing"

	"github.com/ethanmoffat/eolib-go/pkg/eolib/packet"
	"github.com/stretchr/testify/assert"
)

func TestZeroSequence(t *testing.T) {
	start := packet.NewZeroSequence()
	assert.Equal(t, 0, start.Value())
}

func TestInitSequenceFromValue(t *testing.T) {
	start := packet.NewInitSequence(148, 185)

	assert.Equal(t, 1208, start.Value())
	assert.Equal(t, 148, start.Seq1())
	assert.Equal(t, 185, start.Seq2())
}

func TestInitSequenceGenerate(t *testing.T) {
	const seed = 123
	random := rand.New(rand.NewSource(seed))

	start := packet.GenerateInitSequence(*random)

	// note: the expected values differ from eolib-java due to a different pseudo-random implementation
	assert.Equal(t, 1208, start.Value())
	assert.Equal(t, 148, start.Seq1())
	assert.Equal(t, 185, start.Seq2())
}

func TestPingSequenceFromValue(t *testing.T) {
	start := packet.NewPingSequence(1253, 45)

	assert.Equal(t, 1208, start.Value())
	assert.Equal(t, 1253, start.Seq1())
	assert.Equal(t, 45, start.Seq2())
}

func TestPingSequenceGenerate(t *testing.T) {
	const seed = 123
	random := rand.New(rand.NewSource(seed))

	start := packet.GeneratePingSequence(*random)

	// note: the expected values differ from eolib-java due to a different pseudo-random implementation
	assert.Equal(t, 1208, start.Value())
	assert.Equal(t, 1253, start.Seq1())
	assert.Equal(t, 45, start.Seq2())
}

func TestAccountReplySequenceFromValue(t *testing.T) {
	start := packet.NewAccountReplySequence(22)

	assert.Equal(t, 22, start.Value())
}

func TestAccountReplySequenceGenerate(t *testing.T) {
	const seed = 123
	random := rand.New(rand.NewSource(seed))

	start := packet.GenerateAccountReplySequence(*random)

	// note: the expected values differ from eolib-java due to a different pseudo-random implementation
	assert.Equal(t, 155, start.Value())
}
