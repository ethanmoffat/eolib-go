package eolib

import "golang.org/x/exp/constraints"

// Min provides a generic method for getting the lesser of two input values.
// In the event that the inputs are equal, the second input will be returned.
func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Max provides a generic method for getting the greater of two input values.
// In the event that the inputs are equal, the second input will be returned.
func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Reverse returns a new slice containing the elements of the input slice in reverse order.
func Reverse[T constraints.Ordered](input []T) []T {
	ret := make([]T, len(input))

	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = input[j], input[i]
	}

	return ret
}

func FindIndex[T constraints.Ordered](input []T, val T, eq func(T, T) bool) int {
	for i, v := range input {
		if eq(v, val) {
			return i
		}
	}
	return -1
}
