generate:
	 protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

build-docker-image:
	docker build -t grpcalc .

run-docker-image:
	docker run -d -p 8080:8080 grpcalc

run-grpc-ui:
	grpcui --plaintext 127.0.0.1:8080
