package grpc

import (
	"clean-arch-hex/internal/server"
	"fmt"
)

type GRPCServer struct {
}

// Test implements server.Server.
func (GRPCServer) Test() any {
	panic("unimplemented")
}

// Start implements server.Server.
func (GRPCServer) Start() {
	fmt.Println("GRPC server is running...")
}

func New() server.Server {
	return GRPCServer{}
}
