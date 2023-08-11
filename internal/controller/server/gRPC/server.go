package grpc

import (
	"clean-arch-hex/internal/controller/server"
	"fmt"
)

type GRPCServer struct {
}

// Start implements server.Server.
func (GRPCServer) Start() error {
	fmt.Println("GRPC server is running...")
	return nil
}

func New() server.Server {
	return GRPCServer{}
}
