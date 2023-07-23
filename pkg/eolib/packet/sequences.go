package packet

import (
	"math/rand"

	"github.com/ethanmoffat/eolib-go/pkg/eolib"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
)

// SequenceGetter represents a value sent by the server to update the client's sequence start, also known as the "starting counter ID".
type SequenceGetter interface {
	// Value gets the sequence value.
	Value() int
}

// ZeroSequence represents an implementation of [packet.SequenceGetter] with a value set to '0'.
type ZeroSequence struct {
}

// InitSequence represents the sequence start value sent with the INIT_INIT server packet.
type InitSequence struct {
	value int
	seq1  int
	seq2  int
}

// PingSequence represents the sequence start value sent with the CONNECTION_PLAYER server packet.
type PingSequence struct {
	value int
	seq1  int
	seq2  int
}

// AccountReplySequence represents the sequence start value sent with the ACCOUNT_REPLY server packet.
type AccountReplySequence struct {
	value int
}

// NewZeroSequence creates a new ZeroSequence instance
func NewZeroSequence() *ZeroSequence {
	return &ZeroSequence{}
}

// NewInitSequence creates an instance of [packet.InitSequence] from the values sent with the INIT_INIT server packet.
func NewInitSequence(seq1 int, seq2 int) *InitSequence {
	return &InitSequence{
		value: seq1*7 + seq2 - 13,
		seq1:  seq1,
		seq2:  seq2,
	}
}

// GenerateInitSequence generates an instance of [packet.InitSequence] with a random value in the range 0-1757.
func GenerateInitSequence(rand rand.Rand) *InitSequence {
	value := rand.Intn(1757)
	seq1Max := value + 13/7
	seq1Min := eolib.Max(0, (value-(data.CHAR_MAX-1)+13+6)/7)

	seq1 := rand.Intn(seq1Max-seq1Min) + seq1Min
	seq2 := value - seq1*7 + 13
	return &InitSequence{
		value: value,
		seq1:  seq1,
		seq2:  seq2,
	}
}

// NewPingSequence creates an instance of [packet.PingSequence] from the values sent with the CONNECTION_PLAYER server packet.
func NewPingSequence(seq1 int, seq2 int) *PingSequence {
	return &PingSequence{
		value: seq1 - seq2,
		seq1:  seq1,
		seq2:  seq2,
	}
}

// GeneratePingSequence generates an isntance of [packet.PingSequence] with a random value in the range 0-1757.
func GeneratePingSequence(rand rand.Rand) *PingSequence {
	value := rand.Intn(1757)
	seq1 := value + rand.Intn(data.CHAR_MAX-1)
	seq2 := seq1 - value

	return &PingSequence{
		value: value,
		seq1:  seq1,
		seq2:  seq2,
	}
}

// NewAccountReplySequence creates an instance of [packet.AccountReplySequence] from the value sent with the ACCOUNT_REPLY server packet.
func NewAccountReplySequence(value int) *AccountReplySequence {
	return &AccountReplySequence{value}
}

// GenerateAccountReplySequence generates an instance of [packet.AccountReplySequence] with a random value in th range 0-240.
func GenerateAccountReplySequence(rand rand.Rand) *AccountReplySequence {
	start := rand.Intn(240)
	return &AccountReplySequence{start}
}

// Value gets the start value of the [packet.ZeroSequence]. This value is always zero.
func (ZeroSequence) Value() int {
	return 0
}

// Value gets the start value of the [packet.InitSequence].
func (s InitSequence) Value() int {
	return s.value
}

// Seq1 gets the Seq1 byte sent with the INIT_INIT server packet.
func (s InitSequence) Seq1() int {
	return s.seq1
}

// Seq2 gets the Seq2 byte sent with the INIT_INIT server packet.
func (s InitSequence) Seq2() int {
	return s.seq2
}

// Value gets the start value of the [packet.PingSequence].
func (s PingSequence) Value() int {
	return s.value
}

// Seq1 gets the Seq1 byte sent with the CONNECTION_PLAYER server packet.
func (s PingSequence) Seq1() int {
	return s.seq1
}

// Seq2 gets the Seq2 byte sent with the CONNECTION_PLAYER server packet.
func (s PingSequence) Seq2() int {
	return s.seq2
}

// Value gets the start value of the [packet.AccountReplySequence].
func (a AccountReplySequence) Value() int {
	return a.value
}
