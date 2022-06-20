#!/bin/bash
cd ./internal
protoc -I ./proto --grpc-gateway_out ./proto \
    --go_out=./proto --go_opt=paths=source_relative \
    --go-grpc_out=./proto --go-grpc_opt=paths=source_relative \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --validate_out="lang=go:./proto" \
    --validate_opt paths=source_relative proto/iam/iam.proto
cd -
go build ./cmd/fk-daemon-iam
