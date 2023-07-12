package packet

import (
	"math/rand"
)

// AccountReplySequence represents the sequence start value sent with the ACCOUNT_REPLY server packet.
type AccountReplySequence struct {
	value int
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

// Value gets the start value of the [packet.AccountReplySequence].
func (a AccountReplySequence) Value() int {
	return a.value
}
