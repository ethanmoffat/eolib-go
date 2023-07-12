package net

import "github.com/ethanmoffat/eolib-go/pkg/eolib/data"

// Packet represents a packet object in the EO network protocol.
type Packet interface {
	// Family gets the family of the EO packet.
	Family() int
	// Action gets the action of the EO packet.
	Action() int
	// Serialize writes the data in the packet to the specified [data.EoWriter].
	Serialize(writer *data.EoWriter)
}
