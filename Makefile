.PHONY: deps
migrate-create:
	go mod tidy
	go mod vedor

.PHONY: proto-build
proto-build:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative server/server.proto

.PHONY: run-server
run-server:
	go run cmd/server/main.go

.PHONY: run-client
run-client:
	go run cmd/client/main.go