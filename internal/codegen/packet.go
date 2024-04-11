package codegen

import (
	"fmt"
	"path"

	"github.com/dave/jennifer/jen"
	"github.com/ethanmoffat/eolib-go/internal/codegen/types"
	"github.com/ethanmoffat/eolib-go/internal/xml"
)

func GeneratePackets(outputDir string, packets []xml.ProtocolPacket, fullSpec xml.Protocol) error {
	if len(packets) == 0 {
		return nil
	}

	packageName, err := getPackageName(outputDir)
	if err != nil {
		return err
	}

	f := jen.NewFile(packageName)
	types.AddImports(f)

	// collect type names to generate packet structs
	var typeNames []string
	f.Var().Id("packetMap").Op("=").Map(jen.Int()).Qual("reflect", "Type").BlockFunc(func(g *jen.Group) {
		// Note that this block is using "BlockFunc"
		// Official docs advices to use "Values" with "DictFunc". However, default sorting is alphabetical, which
		//    creates a nasty git diff of the existing generated code
		for _, p := range packets {
			typeNames = append(typeNames, p.GetTypeName())

			g.Qual(types.PackagePath("net"), "PacketId").Call(
				jen.Qual(types.PackagePath("net"), fmt.Sprintf("PacketFamily_%s", p.Family)),
				jen.Qual(types.PackagePath("net"), fmt.Sprintf("PacketAction_%s", p.Action)),
			).Op(":").Qual("reflect", "TypeOf").Call(
				jen.Id(snakeCaseToCamelCase(p.GetTypeName())).Values(),
			).Op(",")
		}
	})

	f.Comment("PacketFromId creates a typed packet instance from a [net.PacketFamily] and [net.PacketAction].")
	f.Comment("This function calls [PacketFromIntegerId] internally.")

	f.Func().Id("PacketFromId").Params(
		jen.Id("family").Qual(types.PackagePath("net"), "PacketFamily"),
		jen.Id("action").Qual(types.PackagePath("net"), "PacketAction"),
	).Params(
		jen.Qual(types.PackagePath("net"), "Packet"),
		jen.Error(),
	).Block(
		jen.Return(jen.Id("PacketFromIntegerId").Call(
			jen.Qual(types.PackagePath("net"), "PacketId").Call(jen.Id("family"), jen.Id("action")),
		)),
	)

	f.Comment(`// PacketFromIntegerId creates a typed packet instance from a packet's ID. An ID may be converted from a family/action pair via the [net.PacketId] function.
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
//   }`)

	f.Func().Id("PacketFromIntegerId").Params(
		jen.Id("id").Int(), // func declaration: int parameter 'id'
	).Params(
		jen.Qual(types.PackagePath("net"), "Packet"), // func declaration: return types (net.Packet, error)
		jen.Error(),
	).Block(
		// try to get the packet type out of the map (indexed by the id)
		jen.List(jen.Id("packetType"), jen.Id("idOk")).Op(":=").Id("packetMap").Index(jen.Id("id")),
		// check that id is ok, return error otherwise
		jen.If(jen.Op("!").Id("idOk")).Block(
			jen.Return(jen.List(jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit("could not find packet with id %d"), jen.Id("id")))),
		).Line(),
		// type assert that creating the packet type results in an interface that satisfies net.Packet
		jen.List(jen.Id("packetInstance"), jen.Id("typeOk").Op(":=").Qual("reflect", "New").Call(
			jen.Id("packetType"),
		).Dot("Interface").Call().Assert(
			jen.Qual(types.PackagePath("net"), "Packet"),
		)),
		// check that type is ok, return error otherwise
		jen.If(jen.Op("!").Id("typeOk")).Block(
			jen.Return(jen.List(jen.Nil(), jen.Qual("fmt", "Errorf").Call(jen.Lit("could not create packet from id %d"), jen.Id("id")))),
		).Line(),
		// return packetInstance, nil
		jen.Return(jen.Id("packetInstance"), jen.Nil()),
	)

	const packetMapFileName = "packetmap_generated.go"
	if err := writeToFileJen(f, path.Join(outputDir, packetMapFileName)); err != nil {
		return err
	}

	const packetFileName = "packets_generated.go"
	return generateStructsShared(outputDir, packetFileName, typeNames, fullSpec)
}
