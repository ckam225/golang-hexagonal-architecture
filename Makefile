.PHONY: install build run clean test generate mock doc


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
