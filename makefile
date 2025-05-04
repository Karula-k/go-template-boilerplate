ifneq (,$(wildcard .env))
    include .env
    export
endif

generate:
ifeq ($(name),)
	@echo "use name as params"
else
	migrate create -ext sql -dir ./db/migrations  -seq $(name)
endif

migrateup:
	migrate -database ${DATABASE_URL} -path db/migrations up
migratedown:
	migrate -database ${DATABASE_URL} -path db/migrations down

sqlc:
	sqlc generate

run:
	air
