package grpc

import (
	"clean-arch-hex/internal/cache"
	srv "clean-arch-hex/internal/controller/server"
	"clean-arch-hex/internal/db"
	"clean-arch-hex/pkg/proto"
	"fmt"
	"log"
	"net"
	"net/http"

	gRPC "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	db    db.Database
	cache cache.Cache
	proto.UnimplementedPostServiceServer
}

func NewServer(db db.Database, cache cache.Cache) srv.Server {
	return &Server{
		db:    db,
		cache: cache,
	}
}

// Test implements server.Server.
func (s *Server) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	panic("unimplemented")
}

// Start implements server.Server.
func (s *Server) Start() error {

	// TODO: write intercepter
	rpcServer := gRPC.NewServer()
	reflection.Register(rpcServer)
	proto.RegisterPostServiceServer(rpcServer, s)

	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		return fmt.Errorf("Failed to listen: %w", err)
	}
	defer listener.Close()
	log.Println("GRPC server is running on port 4000.")

	if err := rpcServer.Serve(listener); err != nil {
		return fmt.Errorf("cannot start gRPC server: %w", err)
	}
	return nil
}
