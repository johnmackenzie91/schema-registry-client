# Schema Registry Client POC

This is a library that attempts to make it easier to encode and decode from avro schemas that live in a schema registry.

This helper library combines the functionality of https://github.com/hamba/avro and https://github.com/hamba/avro/registry, combining them into a single interface.

## Getting started

1. Run `make down build start` - this will bring up a confluentinc/cp-schema-registry image.
2. Ensure all containers are up and ready...
3. Run `make schema-seed` - this will add a test schema to the schema registry
    1. Check that the new schema has been registered with `make schema-list`.
    2. You can checkout the latest version of this schema via `make schema-get`
3. An example of how to use can be found [here](https://github.com/johnmackenzie91/schema-registry-client/blob/master/schema_example_test.go)
    1. To run this example test run `go test -v`
    
  
## To do
1. Add caching for schemas returned from `GetLatestSchema` currently it calls out each time.