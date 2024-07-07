# Building out the dependencies is too cumbersome, so everything is a PHONY target
.PHONY: all server client server-docker client-docker proto format

PROTO_IMPORT_DIR=./proto
PROTO_SERVICE=./proto/dns_service.proto

PROTO_GEN_PATH=./gen/go/proto
OPENAPI_GEN_PATH=./gen/openapiv2

all: proto format server client;

server:
	go build -C server -o ../bin/server -trimpath

client:
	go build -C client -o ../bin/client -trimpath

format:
	gofmt -s -w .
	goimports -w -local yadro .

proto:
	mkdir -p $(PROTO_GEN_PATH)
	mkdir -p $(OPENAPI_GEN_PATH)

	# gRPC stubs
	protoc -I $(PROTO_IMPORT_DIR) \
		--go_out ./gen/go/proto --go_opt paths=source_relative \
		--go-grpc_out ./gen/go/proto --go-grpc_opt paths=source_relative \
		$(PROTO_SERVICE)

	# Gateway
	protoc -I $(PROTO_IMPORT_DIR) --grpc-gateway_out ./gen/go/proto \
		--grpc-gateway_opt paths=source_relative \
		$(PROTO_SERVICE)

	# OpenAPI docs
	protoc -I $(PROTO_IMPORT_DIR) --openapiv2_out ./gen/openapiv2 \
		$(PROTO_SERVICE)
