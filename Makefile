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

apply-migrate:
	 sudo atlas schema apply --url "postgres://root:root@localhost:5432/khushidb?sslmode=disable" --to "file://db/migration/20221213125830.sql" --dev-url "docker://postgres/12/khushi"

server:
	go run main.go

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

swag: 
	exec ~/go/bin/swag init

swag-fmt:
	exec ~/go/bin/swag fmt

debug:
	exec ~/go/bin/air

.PHONY: postgres createdb dropdb test sqlc server migrate new-migration apply-migrate swag swag-fmt debug