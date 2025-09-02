.DEFAULT := build

.PHONY: fmt vet build test coverage initdb

test:
	go test ./...

coverage: test
	go test -cover ./...

fmt: coverage
	go fmt ./...

vet: fmt
	go vet ./...

initdb:
	goose up -dir ./migrations

build: vet
	go build