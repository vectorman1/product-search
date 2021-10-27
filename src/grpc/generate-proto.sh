#!/usr/bin/env sh
mkdir -p generated/product_service
protoc -I . --proto_path=./proto \
            --go_out=generated/product_service \
            --go_opt=paths=source_relative \
            --go-grpc_out=generated/product_service \
            --go-grpc_opt=paths=source_relative \
            product_service.proto
