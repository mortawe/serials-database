.PHONY: build
build:
	go build -o bin/shows ./src/cmd/shows

.PHONY: run
run:
	./bin/shows

.PHONY: tidy
tidy:
	 go mod tidy

.PHONY: vendor
vendor:
	 go mod vendor

.PHONY: prune
prune:
	docker stop shows-db
	docker rm shows-db
	docker volume rm shows-api
