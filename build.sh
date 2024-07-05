#!/usr/bin/bash

set -xeuo pipefail

# gRPC stubs
protoc -I ./proto \
    --go_out ./gen/go/ --go_opt paths=source_relative \
    --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative \
    proto/dns_service.proto

# OpenAPI docs
protoc -I ./proto --openapiv2_out ./gen/openapiv2 proto/dns_service.proto


go build -o server.out server/main.go
go build -o client.out client/main.go