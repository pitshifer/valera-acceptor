.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

rundb:
	docker run --name valeradb -v valera:/var/lib/mysql -p 5506:3306 -e MYSQL_ROOT_PASSWORD=dfnheif -d mysql:latest

stopdb:
	docker stop valeradb

removedb:
	docker stop valeradb
	docker rm valeradb

.DEFAULT_GOAL := build
