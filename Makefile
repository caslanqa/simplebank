postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/CharlieAlphaQA/simplebank/db/sqlc Store

migrateupaws:
	migrate -path db/migration -database "postgresql://root:4u8eSI9oKDELJY49c74a@simple-bank.ctqe66s8kpyt.eu-west-1.rds.amazonaws.com:5432/simple_bank" -verbose up


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateupaws