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
	sqlc generate;
	mockgen -package mockdb -destination db/mock/mock.go github.com/dados-id/dados-be/db/sqlc Querier;

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

seeds:
	go run scripts/faculty_seeding/faculty_seeding.go;
	go run scripts/school_seeding/school_seeding.go;
	go run scripts/user_seeding/user_seeding.go;
	go run scripts/course_seeding/course_seeding.go;
	go run scripts/tag_seeding/tag_seeding.go;

	go run scripts/school_rating_seeding/school_rating_seeding.go;
	go run scripts/professor_seeding/professor_seeding.go;
	go run scripts/professor_rating_seeding/professor_rating_seeding.go;

mock:
	mockgen -package mockdb -destination db/mock/mock.go github.com/dados-id/dados-be/db/sqlc Querier

test:
	go test -v -cover ./...

refresh_db:
	docker exec -it dados-container dropdb dados;
	docker exec -it dados-container createdb --username=root --owner=root dados;
	migrate -path db/migration -database "$(DB_URL)" -verbose up;

.PHONY: postgres createdb dropdb db_docs db_schema server sqlc migrateup migratedown seeds mock test refresh_db
