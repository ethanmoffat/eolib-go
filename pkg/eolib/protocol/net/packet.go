package net

import (
	"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol"
)

// Packet represents a packet object in the EO network protocol.
type Packet interface {
	protocol.Serializer

	Family() PacketFamily // Family gets the family of the EO packet.
	Action() PacketAction // Action gets the action of the EO packet.EoData
}
