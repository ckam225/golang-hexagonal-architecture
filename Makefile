PROTO_OUT=pkg/proto
PROTO_PATH=internal/controller/server/grpc/proto


install:
	go mod tidy

build:
	go build cmd/main.go

run:
	go run cmd/main.go

clean:
	go clean -cache
	go clean -i
	go clean -testcache
	go clean -modcache

test:
	go test -v -cover ./...

generate:
	go generate

mock:
	mockgen -package=mockdb -destination=./internal/db/mocks/mock.go "clean-arch-hex/internal/db" Database
	# mockgen -package=mockrepo -destination=./internal/domain/repository/mocks/mock.go "clean-arch-hex/internal/domain/repository" PostRepository,UserRepository

proto:
	rm -rf ${PROTO_OUT}/*.proto
	protoc --proto_path=${PROTO_PATH} \
	--go_out=${PROTO_OUT} \
	--go_opt=paths=source_relative \
	--go-grpc_out=${PROTO_OUT} \
	--go-grpc_opt=paths=source_relative \
	${PROTO_PATH}/*.proto

.PHONY: install build run clean test generate mock doc proto