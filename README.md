# dot-conf

The .conf, dynamic configuration platform, is designed to provide developers with a flexible solution for managing configurations outside of their codebases. By decoupling configurations from the application code, developers gain the ability to modify settings during runtime, reducing the need for frequent deployments and enhancing overall flexibility and security. It's an open-source alternative to Consul by Hashicorp, Etcd, etc.

## Documents
1. [HLD](./documents/HLD.md)
2. [LLD](./documents/LLD.md)

## Setup

### Build Backend Service

1. Run `go mod tidy`.
2. Install `protobuf` 
```
For Linux:
sudo apt install -y protobuf-compiler

# For Mac:
brew install protobuf

# Common for both
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
3. If you've updated the proto file, run the below command
```
## backend server
cd backend
protoc --proto_path=../proto --go_out=paths=source_relative:./proto --go-grpc_out=paths=source_relative:./proto ../proto/config.proto

## java client
cd ../java-client
protoc --proto_path=../proto/ --java_out=build/gen ../proto/config.proto
```
