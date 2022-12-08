.SILENT:

include .env

dep:
	go mod download

vet:
	go vet -json
	go vet -json ./cmd
	go vet -json ./internal
	go vet -json ./pkg/dto
	go vet -json ./pkg/transport/http/handler
	go vet -json ./models
	go vet -json ./pkg/repository
	go vet -json ./pkg/service

run:
	git pull
	go run cmd/main.go

build:
	go build cmd/main.go

clean:
	go clean

mod:
	go mod download