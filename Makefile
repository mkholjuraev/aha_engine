postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root aha-makefile

dropdb:
	docker exec -it postgres12 dropdb --username=root --owner=root aha-makefile

migrateup: 
	migrate -path db/migration -database "psql://root:admin@localhost:5433/aha-terminal" -verbose up

migratedown:
	migrate -path db/migration -database "psql://root:admin@localhost:5433/aha-terminal" -verbose down

.PHONY: postgres createdb dropdb
