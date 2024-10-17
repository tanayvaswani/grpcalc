# grpcalc

A gRPC based calculator server and client.

## Gears

- go v1.23.2
- protobuf (brew install protobuf)
- protoc-gen-go
- protoc-gen-go-grpc
- some knowledge about protocol buffers
- docker

## How to run ?

1. Using Docker

   ```bash
   docker build -t grpcalc .
   ```

   ```bash
   docker run -d -p 8080:8080 grpcalc
   ```

   Alternatively you can use makefile to run commands

   ```bash
   make build-docker-image
   ```

   ```bash
   docker run-docker-image
   ```

2. Local Setup

   ```bash
   go mod tidy
   ```

   ```bash
   go run server/main.go
   ```

   Alternatively you can use makefile to run commands

   To generate proto code

   ```bash
   make generate
   ```

   ```bash
   make run-server
   ```

## How to consume?

- grpcui (brew install grpc)

```bash
grpcui --plaintext 127.0.0.1:8080
```

- grpcurl

```bash
  grpcurl -d '{
    "a": 10,
    "b": 2
  }' 127.0.0.1:8080
```
