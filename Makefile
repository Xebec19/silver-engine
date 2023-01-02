postgres:
	docker run --name khushi --network khushi -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -dp 5432:5432 postgres:12-alpine

createdb:
	docker exec -it khushi createdb --username=root --owner=root khushidb

dropdb:
	docker exec -it khushi dropdb khushidb

migrate:
	atlas schema inspect -u 'postgres://root:root@localhost:5432/khushidb?sslmode=disable' > schema.hcl

new-migration:
	atlas migrate new --dir "file://db/migration/"

server:
	go run main.go

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb test sqlc server migrate new-migration