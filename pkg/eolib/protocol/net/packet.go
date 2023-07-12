package net

import "github.com/ethanmoffat/eolib-go/pkg/eolib/data"

// Packet represents a packet object in the EO network protocol.
type Packet interface {
	Family() int                     // Family gets the family of the EO packet.
	Action() int                     // Action gets the action of the EO packet.
	Serialize(writer *data.EoWriter) // Serialize writes the data in the packet to the specified EoWriter.
}
