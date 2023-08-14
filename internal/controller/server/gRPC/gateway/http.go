package gateway

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/pkg/proto"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

type GrpcHttpGateway struct {
	postClient proto.PostServiceClient
	app        *fiber.App
}

func NewGrpcHttpGateway(grpcClient *grpc.ClientConn) *GrpcHttpGateway {
	app := &GrpcHttpGateway{
		app:        fiber.New(),
		postClient: proto.NewPostServiceClient(grpcClient),
	}
	return app
}

func (s *GrpcHttpGateway) Start(addr string) error {
	s.app.Get("/posts", s.GetAllPosts)
	s.app.Post("/posts", s.CreatePost)
	s.app.Get("/posts/:id", s.GetPost)
	return s.app.Listen(addr)
}

func (s *GrpcHttpGateway) GetAllPosts(c *fiber.Ctx) error {
	ctx := context.Background()
	res, err := s.postClient.GetAllPosts(ctx, &proto.PostRequest{})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (s *GrpcHttpGateway) GetPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}
	ctx := context.Background()
	post, err := s.postClient.GetPost(ctx, &proto.PostRequest{
		Id: int32(id),
	})
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if post == nil {
		return c.Status(404).JSON(fiber.Map{"error": fmt.Sprintf("Not found: %d", id)})
	}
	return c.JSON(post)
}

func (s *GrpcHttpGateway) CreatePost(c *fiber.Ctx) error {
	var q entity.Post
	if err := c.BodyParser(&q); err != nil {
		return c.Status(422).JSON(err)
	}
	if q.Content == "" || q.Title == "" {
		log.Println("content or title is empty")
		return c.Status(422).JSON(fiber.ErrUnprocessableEntity)
	}
	ctx := context.Background()
	_, err := s.postClient.CreatePost(ctx, &proto.PostRequest{
		Title:   q.Title,
		Content: q.Title,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"code":    500,
		})
	}
	return c.SendStatus(204)
}
