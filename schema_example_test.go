package schema_test

import (
	"no_vcs/me/schema"

	"fmt"

	"github.com/hamba/avro/registry"
)

type Test struct {
	ID   int    `avro:"id"`
	Name string `avro:"name"`
}

func Example_LatestSchemaFromSubject_EncodeAndDecode() {
	// schema registry endpoint
	cli, err := registry.NewClient("http://0.0.0.0:8081")
	if err != nil {
		panic(err)
	}

	// instantiate our library
	sut := schema.NewRegistryClient(cli)

	// the input we want to encode and decode
	in := Test{ID: 1, Name: "test one"}

	// returns a codec, that will use the latest version of the subject asked for, that lives in the schema registry
	// simpler: get me the latest version of a schema called "test"
	codec, err := sut.LatestSchemaFromSubject("test")
	if err != nil {
		panic(err)
	}

	// let us encode using this schema
	b, err := codec.Encode(in)
	if err != nil {
		panic(err)
	}

	// let us decode the previously encoded byte array into a new struct
	out := Test{}
	err = codec.Decode(b, &out)
	if err != nil {
		panic(err)
	}

	fmt.Printf("id: %d, name: %s", out.ID, out.Name)
	// Output: id: 1, name: test one
}
