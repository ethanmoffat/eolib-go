package codegen

import (
	"github.com/ethanmoffat/eolib-go/internal/xml"
)

func GeneratePackets(outputDir string, packets []xml.ProtocolPacket, fullSpec xml.Protocol) error {
	const packetFileName = "packets_generated.go"

	var typeNames []string
	for _, p := range packets {
		typeNames = append(typeNames, p.GetTypeName())
	}
	return generateStructsShared(outputDir, packetFileName, typeNames, fullSpec)
}
