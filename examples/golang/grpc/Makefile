GO := @go

run-server:
	$(GO) run ./server

run-client:
	$(GO) run ./client

proto:
	protoc --go_out=plugins=grpc:. protos/helloworld.proto
