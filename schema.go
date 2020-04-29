package schema

import (
	"sync"

	"github.com/hamba/avro"
)

// LatestSchemaGetter implemented github.com/hamba/avro/registry interface
type LatestSchemaGetter interface {
	GetLatestSchema(subject string) (avro.Schema, error)
}

// RegistryClient is the main struct of this package
type RegistryClient struct {
	schemaRegistry   LatestSchemaGetter
	schemaCache      map[string]*Codec
	schemaCacheMutex *sync.RWMutex
}

// NewRegistryClient instantiates a new RegistryClient
func NewRegistryClient(client LatestSchemaGetter) *RegistryClient {
	return &RegistryClient{
		schemaRegistry:   client,
		schemaCache:      map[string]*Codec{},
		schemaCacheMutex: &sync.RWMutex{},
	}
}

// LatestSchemaFromSubject wraps schemaRegistry.LatestSchemaFromSubject, adds cache
func (c RegistryClient) LatestSchemaFromSubject(subject string) (*Codec, error) {

	// check cache for previously requested codec
	if codec := c.getFromSchemaCache(subject); codec != nil {
		return codec, nil
	}

	schema, err := c.schemaRegistry.GetLatestSchema(subject)

	if err != nil {
		return nil, err
	}

	c.setFromSchemaCache(subject, schema)

	return c.schemaCache[subject], nil
}

func (c *RegistryClient) getFromSchemaCache(subject string) *Codec {
	c.schemaCacheMutex.RLock()
	defer c.schemaCacheMutex.RUnlock()
	v, ok := c.schemaCache[subject]
	if !ok {
		return nil
	}
	return v
}

func (c *RegistryClient) setFromSchemaCache(subject string, schema avro.Schema) {
	c.schemaCacheMutex.Lock()
	c.schemaCache[subject] = &Codec{
		schema: schema,
	}
	c.schemaCacheMutex.Unlock()
}
