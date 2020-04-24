build:
	docker-compose build --parallel

start:
	docker-compose stop
	docker-compose up -d
	docker-compose ps

down:
	docker-compose down

restart:
	docker-compose down
	@make start
run:
	go run *.go

logs:
	docker-compose logs -f

schema-seed:
	curl -X POST -H "Content-Type: application/vnd.schemaregistry.v1+json" -d @./avro/seeds/test.json http://0.0.0.0:8081/subjects/test/versions

schema-list:
	curl --silent -X GET http://0.0.0.0:8081/subjects/ | jq .

schema-get:
	curl --silent -X GET http://localhost:8081/subjects/test/versions/latest | jq .

logs4:
	docker-compose logs -f | grep '$(container)'