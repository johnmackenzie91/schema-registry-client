package schema

import (
	"errors"
	"reflect"

	"github.com/hamba/avro"
)

// Codec will be used to encode and decode inputted values
type Codec struct {
	schema avro.Schema
}

// Encode input via the schema field
func (c Codec) Encode(in interface{}) ([]byte, error) {
	return avro.Marshal(c.schema, in)
}

// PointerNotSupplied will be returned when we are expecting a interface{} param to be a pointer
var PointerNotSupplied = errors.New("Codec.Decode expected a pointer value")

// Decode into the value supplied
func (c Codec) Decode(data []byte, in interface{}) error {
	// we can only work on pointers, inform user if pointer not supplied
	if reflect.ValueOf(in).Kind() != reflect.Ptr {
		return PointerNotSupplied
	}
	return avro.Unmarshal(c.schema, data, in)
}
