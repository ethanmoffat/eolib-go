package types

import (
	"errors"
	"fmt"

	"github.com/ethanmoffat/eolib-go/v3/internal/xml"
)

// StructInfo is a type representing the metadata about a struct that should be rendered as generated code.
// It represents the common properties of either a ProtocolPacket or a ProtocolStruct.
type StructInfo struct {
	Name         string                    // Name is the name of the type. It is not converted from protocol naming convention (snake_case).
	Comment      string                    // Comment is an optional type comment for the struct.
	Instructions []xml.ProtocolInstruction // Instructions is a collection of instructions for the struct.
	PackageName  string                    // PackageName is the containing package name for the struct.

	Family                string // Family is the Packet Family of the struct, if the struct is a packet struct.
	Action                string // Action is the Packet Action of the struct, if the struct is a packet struct.
	SwitchStructQualifier string // SwitchStructQualifier is an additional qualifier prepended to structs used in switch cases in packets.
}

// GetStructInfo generates a [StructInfo] for the specified typeName. typeName can be a structure
// or a packet in the full XML spec.
func GetStructInfo(typeName string, fullSpec xml.Protocol) (si *StructInfo, err error) {
	si = &StructInfo{SwitchStructQualifier: ""}
	err = nil

	if structInfo, ok := fullSpec.IsStruct(typeName); ok {
		si.Name = structInfo.Name
		si.Comment = structInfo.Comment
		si.Instructions = structInfo.Instructions
		si.PackageName = structInfo.Package
	} else if packetInfo, ok := fullSpec.IsPacket(typeName); ok {
		si.Name = packetInfo.GetTypeName()
		si.Comment = packetInfo.Comment
		si.Instructions = packetInfo.Instructions
		si.PackageName = packetInfo.Package
		si.SwitchStructQualifier = packetInfo.Family + packetInfo.Action
		si.Family = packetInfo.Family
		si.Action = packetInfo.Action
	} else {
		si = nil
		err = fmt.Errorf("type %s is not a struct or packet in the spec", typeName)
	}

	return
}

// Nested creates a nested [StructInfo] from the specified 'chunked' [xml.ProtocolInstruction].
// This function returns an error if the instruction is not of the 'chunked' type.
func (si *StructInfo) Nested(chunked *xml.ProtocolInstruction) (*StructInfo, error) {
	if chunked.XMLName.Local != "chunked" {
		return nil, errors.New("expected 'chunked' instruction creating nested StructInfo")
	}

	return &StructInfo{
		Instructions:          chunked.Chunked,
		PackageName:           si.PackageName,
		SwitchStructQualifier: si.SwitchStructQualifier,
	}, nil
}
