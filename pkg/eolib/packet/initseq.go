package packet

import (
	"math/rand"

	"github.com/ethanmoffat/eolib-go/pkg/eolib"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
)

// InitSequence represents the sequence start value sent with the INIT_INIT server packet.
type InitSequence struct {
	value int
	seq1  int
	seq2  int
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
