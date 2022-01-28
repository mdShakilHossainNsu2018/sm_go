migrateup:
	migrate -path sql_db/migration -database "postgresql://postgres:example@localhost:5432/sm_db?sslmode=disable" -verbose up

migratedown:
	migrate -path sql_db/migration -database "postgresql://postgres:example@localhost:5432/sm_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown sqlc test
