package schema

import (
	"github.com/hamba/avro"
)

type Codec struct {
	schema avro.Schema
}

func (c Codec) Encode(in interface{}) ([]byte, error) {
	return avro.Marshal(c.schema, in)
}

func (c Codec) Decode(data []byte, in interface{}) error {
	// TODO check is pointer
	return avro.Unmarshal(c.schema, data, in)
}
