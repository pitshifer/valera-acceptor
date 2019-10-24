.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

rundb:
	docker run --name valeradb --network valeranet -v valera:/var/lib/postgresql/data -p 5506:5432 -e POSTGRES_PASSWORD=dfnheif -d postgres:12.0

psql:
	docker run -it --rm --network valeranet postgres:12.0 psql -h valeradb -U postgres

stopdb:
	docker stop valeradb

removedb:
	docker stop valeradb
	docker rm valeradb

migrate:
	migrate --path=./migrations -database postgres://postgres:dfnheif@localhost:5506/acceptor?sslmode=disable ${ARGS}

migrateTest:
	migrate --path=./migrations -database postgres://postgres:dfnheif@localhost:5506/acceptor_test?sslmode=disable ${ARGS}

.DEFAULT_GOAL := build
