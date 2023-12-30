package codegen

import (
	"fmt"
	"path"
	"strings"

	"github.com/ethanmoffat/eolib-go/internal/xml"
)

func GeneratePackets(outputDir string, packets []xml.ProtocolPacket, fullSpec xml.Protocol) error {
	packageDeclaration, err := getPackageStatement(outputDir)
	if err != nil {
		return err
	}

	output := strings.Builder{}
	output.WriteString(packageDeclaration + "\n\n")
	output.WriteString("import (\n\t\"fmt\"\n\t\"reflect\"\n\t\"github.com/ethanmoffat/eolib-go/pkg/eolib/protocol/net\"\n)\n\n")
	output.WriteString("var packetMap = map[int]reflect.Type{\n")

	// collect type names to generate packet structs
	var typeNames []string
	for _, p := range packets {
		typeNames = append(typeNames, p.GetTypeName())

		output.WriteString(fmt.Sprintf("\tnet.PacketId(net.PacketFamily_%s, net.PacketAction_%s): ", p.Family, p.Action))
		output.WriteString(fmt.Sprintf("reflect.TypeOf(%s{}),\n", snakeCaseToCamelCase(p.GetTypeName())))
	}

	output.WriteString("}\n")

	output.WriteString(`
// PacketFromId creates a typed packet instance from a [net.PacketFamily] and [net.PacketAction].
// This function calls [PacketFromIntegerId] internally.
func PacketFromId(family net.PacketFamily, action net.PacketAction) (net.Packet, error) {
	return PacketFromIntegerId(net.PacketId(family, action))
}

// PacketFromIntegerId creates a typed packet instance from a packet's ID. An ID may be converted from a family/action pair via the [net.PacketId] function.
// The returned packet implements the [net.Packet] interface. It may be serialized/deserialized without further conversion, or a type assertion may be made to examine the data. The expected type of the assertion is a pointer to a packet structure.
// The following example does both: an incoming CHAIR_REQUEST packet is deserialized from a reader without converting from the interface type, and the data is examined via a type assertion.
//
//   pkt, _ := client.PacketFromId(net.PacketFamily_Chair, net.PacketAction_Request)
//   if err = pkt.Deserialize(reader); err != nil {
//     // handle the error
//   }
//   switch v := pkt.(type) {
//   case *client.ChairRequestClientPacket:
//      fmt.Println("SitAction=", v.SitAction)
//      switch d := v.SitActionData.(type) {
//      case *client.ChairRequestSitActionDataSit:
//        fmt.Println("Data.Coords=", v.Data.Coords)
//      }
//   default:
//     fmt.Printf("Unknown type: %s\n", reflect.TypeOf(pkt).Elem().Name())
//   }
func PacketFromIntegerId(id int) (net.Packet, error) {
	packetType, idOk := packetMap[id]
	if !idOk {
		return nil, fmt.Errorf("could not find packet with id %d", id)
	}

	packetInstance, typeOk := reflect.New(packetType).Interface().(net.Packet)
	if !typeOk {
		return nil, fmt.Errorf("could not create packet from id %d", id)
	}

	return packetInstance, nil
}
`)

	if len(packets) > 0 {
		const packetMapFileName = "packetmap_generated.go"
		writeToFile(path.Join(outputDir, packetMapFileName), output.String())
	}

	const packetFileName = "packets_generated.go"
	return generateStructsShared(outputDir, packetFileName, typeNames, fullSpec)
}
