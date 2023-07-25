package net

import (
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
)

// Packet represents a packet object in the EO network protocol.
type Packet interface {
	protocol.EoData

	Family() PacketFamily // Family gets the family of the EO packet.
	Action() PacketAction // Action gets the action of the EO packet.EoData
}

// PacketId gets an integer representation of the packet ID from a [PacketFamily] and [PacketAction]
func PacketId(family PacketFamily, action PacketAction) int {
	return ((int(action) & 0xFF) << 8) | (int(family) & 0xFF)
}
