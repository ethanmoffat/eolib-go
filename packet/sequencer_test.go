package packet_test

import (
	"testing"

	"github.com/ethanmoffat/eolib-go/v3/packet"
	"github.com/stretchr/testify/assert"
)

func TestNextSequence(t *testing.T) {
	start := packet.NewAccountReplySequence(123)
	sequencer := packet.NewPacketSequencer(start)

	// counter should increase 9 times and then wrap around
	for i := 0; i < 10; i++ {
		assert.Equal(t, sequencer.NextSequence(), 123+i)
	}

	// counter should have wrapped around
	assert.Equal(t, sequencer.NextSequence(), 123)
}

func TestSequenceStart(t *testing.T) {
	start := packet.NewAccountReplySequence(100)
	sequencer := packet.NewPacketSequencer(start)

	assert.Equal(t, 100, sequencer.NextSequence())

	start = packet.NewAccountReplySequence(200)
	sequencer.SetSequenceStart(start)

	// when the sequence start is updated, the counter should not reset
	assert.Equal(t, 201, sequencer.NextSequence())
}
