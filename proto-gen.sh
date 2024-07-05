#!/usr/bin/sh

# Generate gRPC stubs
protoc -I ./proto \
    --go_out ./gen/go/ --go_opt paths=source_relative \
    --go-grpc_out ./gen/go/ --go-grpc_opt paths=source_relative \
    ./proto/dns_service.proto

# Generate OpenAPI documentation
protoc -I ./proto --openapiv2_out ./gen/openapiv2 \
    ./proto/dns_service.proto
