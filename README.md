# fk-daemon-iam

## preCondition 

```shell
$ go get -u \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc \
github.com/envoyproxy/protoc-gen-validate
```

## Build

```sh
go mod vendor
go mod tidy
go build ./cmd/fk-daemon-iam/main
```
