package protocol

import "github.com/ethanmoffat/eolib-go/pkg/eolib/data"

type Serializer interface {
	Serialize(writer data.EoWriter) error
}

type Deserializer interface {
	Deserialize(reader data.EoReader) error
}

type EoData interface {
	Serializer
	Deserializer
}
