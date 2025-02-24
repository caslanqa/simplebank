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

run-server-local:
	go run main.go

run-server-aws:
	docker run -p 8080:8080 397919817219.dkr.ecr.eu-west-1.amazonaws.com/simplebank:latest

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/CharlieAlphaQA/simplebank/db/sqlc Store

migrateupaws:
	migrate -path db/migration -database "postgresql://root:zwshkJAFs8lozKcydIW0@simple-bank.ctqe66s8kpyt.eu-west-1.rds.amazonaws.com:5432/simple_bank" -verbose up

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateupaws