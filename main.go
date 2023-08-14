package main

import (
	grpc "clean-arch-hex/internal/controller/server/gRPC"
)

func main() {
	if err := grpc.NewClient("localhost:4000"); err != nil {
		panic(err)
	}
}
