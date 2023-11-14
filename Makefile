postgres-pwd := $(shell cat secrets/postgres-pwd)

createdb:
	docker exec -it masterclass-db createdb -U postgres --username=postgres --owner=postgres masterclass

dropdb:
	docker exec -it masterclass-db dropdb -U postgres masterclass

migrateup:
	migrate -path db/migration -database "postgresql://postgres:$(postgres-pwd)@localhost:5432/masterclass?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:$(postgres-pwd)@localhost:5432/masterclass?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./... -env=$(env)

server:
	go run main.go -env=$(env)

mock:
	 mockgen -package mockdb -destination db/mock/store.go github.com/leocardhio/masterclass/db/sqlc Store

.PHONY: createdb dropdb migrateup migratedown sqlc test mock