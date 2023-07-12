package encrypt

import "errors"

// Interleave interleaves a sequence of bytes. When encrypting EO data, bytes are "woven" into each other.
// Used when encrypting packets and data files.
//
// Example:
//
//	{0, 1, 2, 3, 4, 5} -> {0, 5, 1, 4, 2, 3}
func Interleave(data []byte) (output []byte) {
	output = make([]byte, len(data))

	i, ii := 0, 0

	for ; i < len(data); i += 2 {
		output[i] = data[ii]
		ii++
	}

	i--

	if len(data)%2 != 0 {
		i -= 2
	}

	for ; i >= 0; i -= 2 {
		output[i] = data[ii]
		ii++
	}

	return
}

// Deinterleave deinterleaves a sequence of bytes. This is the reverse of [encrypt.Interleave].
// Used when decrypting packets and data files.
//
// Example:
//
//	{0, 1, 2, 3, 4, 5} -> {0, 2, 4, 5, 3, 1}
func Deinterleave(data []byte) (output []byte) {
	output = make([]byte, len(data))

	i, ii := 0, 0

	for ; i < len(data); i += 2 {
		output[ii] = data[i]
		ii++
	}

	i--

	if len(data)%2 != 0 {
		i -= 2
	}

	for ; i >= 0; i -= 2 {
		output[ii] = data[i]
		ii++
	}

	return
}

// FlipMsb flips the most significant bits of each byte in a sequence of bytes (exccluding values 0 and 128).
// Used when encrypting and decrypting packets.
//
// Example:
//
//	{0, 1, 127, 128, 129, 254, 255} -> {0, 129, 255, 128, 1, 126, 127}
func FlipMsb(data []byte) (output []byte) {
	output = make([]byte, len(data))

	for i, b := range data {
		if b&0x7F != 0 {
			output[i] = data[i] ^ 0x80
		}
	}

	return
}

// SwapMultiples swaps the order of contiguous bytes in a sequence of bytes that are divisible by a given multiple value.
// Used when encrypting and decrypting packets and data files.
//
// Example:
//
//	multiple = 3
//	{10, 21, 27} -> {10, 27, 21}
func SwapMultiples(data []byte, multiple int) (output []byte, err error) {
	if multiple < 0 {
		err = errors.New("multiple must be a positive number")
		return
	} else if multiple == 0 {
		output = data
		return
	}

	output = make([]byte, len(data))
	copy(output, data)

	sequenceLength := 0

	for i, b := range output {
		if i != len(output) && int(b)%multiple == 0 {
			sequenceLength++
		} else {
			if sequenceLength > 1 {
				for ii := 0; ii < sequenceLength/2; ii++ {
					a := i - sequenceLength + ii
					b := i - ii - 1
					output[a], output[b] = output[b], output[a]
				}
			}

			sequenceLength = 0
		}
	}

	return
}
