package grpc

import (
	"clean-arch-hex/pkg/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type client struct {
}

func NewClient(serverAddr string) error {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("Could not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewPostServiceClient(conn)
	resp, err := client.GetAllPosts(context.Background(), &proto.PostRequest{})
	if err != nil {
		return fmt.Errorf("Could not call service: %v", err)
	}

	fmt.Println(resp)

	return nil
}
