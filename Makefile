.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

rundb:
	# docker run --name valeradb -v valera:/var/lib/mysql -p 5506:3306 -e MYSQL_ROOT_PASSWORD=dfnheif -d mysql:latest
	docker run --name valeradb -v valera:/var/lib/postgresql/data -p 5506:5432 -e POSTGRES_PASSWORD=dfnheif -d postgres:12.0

psql:
	docker run -it --rm --network bridge postgres:12.0 psql -h 172.17.0.2 -U postgres

stopdb:
	docker stop valeradb

removedb:
	docker stop valeradb
	docker rm valeradb

.DEFAULT_GOAL := build
