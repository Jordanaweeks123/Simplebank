postgres:
	docker run --name postgres12 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:USLxCiBSwomdWxRLV2TP@simplebank.ch0mww0aavjg.us-east-1.rds.amazonaws.com:5432/simple_bank" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:USLxCiBSwomdWxRLV2TP@simplebank.ch0mww0aavjg.us-east-1.rds.amazonaws.com:5432/simple_bank" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:USLxCiBSwomdWxRLV2TP@simplebank.ch0mww0aavjg.us-east-1.rds.amazonaws.com:5432/simple_bank" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:USLxCiBSwomdWxRLV2TP@simplebank.ch0mww0aavjg.us-east-1.rds.amazonaws.com:5432/simple_bank" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mock_db github.com/techschool/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc test server mock