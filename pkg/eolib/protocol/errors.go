package protocol

// SerializationError represents an error serializing a protocol data structure.
type SerializationError struct {
	message string
}

// Error gets the error message of the [protocol.SerializationError].
func (s SerializationError) Error() string {
	return s.message
}
