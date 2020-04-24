package schema

import (
	"sync"

	"github.com/hamba/avro"
)

type IRegistryClient interface {
	LatestSchemaFromSubject(subject string) (Codec, error)
}

type LatestSchemaGetter interface {
	GetLatestSchema(subject string) (avro.Schema, error)
}

var _ IRegistryClient = (*RegistryClient)(nil)

type RegistryClient struct {
	schemaRegistry   LatestSchemaGetter
	schemaCache      map[int]avro.Schema
	schemaCacheMutex *sync.RWMutex
}

func (c RegistryClient) LatestSchemaFromSubject(subject string) (Codec, error) {
	schema, err := c.schemaRegistry.GetLatestSchema(subject)

	if err != nil {
		return Codec{}, err
	}

	return Codec{
		schema: schema,
	}, err
}

func NewRegistryClient(client LatestSchemaGetter) *RegistryClient {
	return &RegistryClient{
		schemaRegistry:   client,
		schemaCache:      map[int]avro.Schema{},
		schemaCacheMutex: &sync.RWMutex{},
	}
}
