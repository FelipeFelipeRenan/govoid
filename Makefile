.PHONY: run build test bench

run:
	go run cmd/server/main.go

build:
	go build -o bin/voidkv cmd/server/main.go

test:
	go test -v ./...

bench:
	go test -bench=. ./internal/...