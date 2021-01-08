#!/bin/sh

go mod tidy && go install \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc

for f in $(find proto -name "*.proto")
do
    protoc -I proto -I /usr/local/include \
        --go_out pb \
        --go_opt paths=source_relative \
        --go-grpc_out pb \
        --go-grpc_opt paths=source_relative \
        "${f}"
done