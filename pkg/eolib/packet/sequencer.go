package packet

// PacketSequencer generates packet sequences based on the specified [packet.SequenceGetter]
type PacketSequencer struct {
	sequenceGetter SequenceGetter
	counter        int
}

// NewPacketSequencer creates a new packet sequencer with the specified [SequenceGetter] and a counter of 0.
func NewPacketSequencer(getter SequenceGetter) PacketSequencer {
	return PacketSequencer{getter, 0}
}

// NextSequence returns the next sequence value, updating the sequence counter in the process
//
// Note: this is not a monotonic operation. The sequence counter increases from 0 to 9 before looping back around to 0.
func (p *PacketSequencer) NextSequence() int {
	result := p.sequenceGetter.Value() + p.counter
	p.counter = (p.counter + 1) % 10
	return result
}

// SetSequenceStart sets the sequence start, also known as the "starting counter ID".
//
// Note: this does not reset the sequence counter.
func (p *PacketSequencer) SetSequenceStart(start SequenceGetter) {
	p.sequenceGetter = start
}
