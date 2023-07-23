package encrypt_test

import (
	"fmt"
	"testing"

	"github.com/ethanmoffat/eolib-go/pkg/eolib/data"
	"github.com/ethanmoffat/eolib-go/pkg/eolib/encrypt"
	"github.com/stretchr/testify/assert"
)

func TestServerVerificationHash(t *testing.T) {
	allArgs := []struct {
		in  int
		out int
	}{
		{0, 114000},
		{1, 115191},
		{2, 229432},
		{5, 613210},
		{12345, 266403},
		{100_000, 145554},
		{5_000_000, 339168},
		{11_092_003, 112773},
		{11_092_004, 112655},
		{11_092_005, 112299},
		{11_092_110, 11016},
		{11_092_111, -2787},
		{11_111_111, 103749},
		{12_345_678, -32046},
		{data.THREE_MAX - 1, 105960},
	}

	for _, arg := range allArgs {
		t.Run(fmt.Sprintf("%d should hash to %d", arg.in, arg.out), func(t *testing.T) {
			actual := encrypt.ServerVerificationHash(arg.in)
			assert.Equal(t, arg.out, actual)
		})
	}
}
