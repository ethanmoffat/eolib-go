package packet

import (
	"math/rand"

	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
)

// PingSequence represents the sequence start value sent with the CONNECTION_PLAYER server packet.
type PingSequence struct {
	value int
	seq1  int
	seq2  int
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
