.PHONY: test
test:
	go test -v -cover -race ./...

.PHONY: build
build:
	go build -o bin/apiserver cmd/main.go

.PHONY: run
run:
	go mod tidy
	go run cmd/main.go

.PHONY: docker-build
docker-build:
	docker build -t vsr-apiserver .

.PHONY: docker-run
docker-run:
	docker run vsr-apiserver
