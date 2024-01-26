DB_URL=postgresql://rangga:mitsuha@localhost:5432/biodata?sslmode=disable

createdb:
	docker exec -it postgres16 createdb --username=rangga --owner=rangga biodata

dev:
	gin run main.go

dropdb:
	docker exec -it postgres16 dropdb biodata

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ranggaAdiPratama/go_biodata/db/sqlc Store

postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=rangga -e POSTGRES_PASSWORD=mitsuha -d postgres

server:
	go run main.go

sqlc:
	docker run --rm -v F:\go\go_biodata:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dev dropdb migrateup migrateup1 migratedown migratedown1 mock postgres server sqlc test

# migrate create -dir db/migration -ext sql -seq add_user_detail
