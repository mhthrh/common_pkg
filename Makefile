grpc-build:
	protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           pkg/model/**/grpc/**/*.proto

.POHNY: grpc-build