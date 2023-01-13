DB_URL=postgresql://root:secret@localhost:5432/dados?sslmode=disable

db_docs:
	dbdocs build doc/database.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/database.dbml

postgres:
	docker run --name dados-container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14.0-alpine

createdb:
	docker exec -it dados-container createdb --username=root --owner=root dados

dropdb:
	docker exec -it dados-container dropdb dados

server:
	go run main.go

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

.PHONY: postgres createdb dropdb db_docs db_schema server sqlc migrateup migratedown
