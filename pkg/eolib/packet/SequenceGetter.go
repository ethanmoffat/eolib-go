package packet

// SequenceGetter represents a value sent by the server to update the client's sequence start, also known as the "starting counter ID".
type SequenceGetter interface {
	// Value gets the sequence value.
	Value() int
}

// ZeroSequence represents an implementation of [packet.SequenceGetter] with a value set to '0'.
type ZeroSequence struct {
}

// NewZeroSequence creates a new ZeroSequence instance
func NewZeroSequence() *ZeroSequence {
	return &ZeroSequence{}
}

// Value gets the start value of the [packet.ZeroSequence]. This value is always zero.
func (ZeroSequence) Value() int {
	return 0
}
