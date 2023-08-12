package grpc

import (
	"clean-arch-hex/internal/controller/server"
	"fmt"
	"net/http"
)

type GRPCServer struct {
}

// Test implements server.Server.
func (GRPCServer) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	panic("unimplemented")
}

// Start implements server.Server.
func (GRPCServer) Start() error {
	fmt.Println("GRPC server is running...")
	return nil
}

func New() server.Server {
	return GRPCServer{}
}
