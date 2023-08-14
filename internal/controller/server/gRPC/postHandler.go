package grpc

import (
	"clean-arch-hex/internal/controller/server/gRPC/converter"
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/usecase"
	"clean-arch-hex/pkg/proto"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) CreatePost(ctx context.Context, req *proto.PostRequest) (*proto.Empty, error) {
	if req.Content == "" || req.Title == "" {
		log.Println("content or title is empty")
		return nil, fiber.ErrUnprocessableEntity
	}
	service := usecase.SavePost(s.db)
	userId := int(req.GetUserId())
	q := entity.Post{
		Content: req.GetContent(),
		Title:   req.GetTitle(),
		UserId:  &userId,
	}
	if err := service.Create(ctx, &q); err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (s *Server) GetAllPosts(ctx context.Context, req *proto.PostRequest) (*proto.PostResponse, error) {
	// const cacheKey = "/posts"
	// data, err := s.cache.Get(ctx, cacheKey)
	// if err != nil || err == redis.Nil {
	service := usecase.ReadPost(s.db)
	posts, err := service.GetAll(ctx, entity.PostFilter{
		Limit:   int(req.Limit),
		ID:      int(req.Id),
		Title:   req.Title,
		Content: req.Content,
		UserId:  int(req.UserId),
	})
	if err != nil {
		return nil, err
	}
	// s.cache.Set(ctx, cacheKey, posts, time.Minute*2)
	return &proto.PostResponse{
		Page:      0,
		PageCount: 0,
		Total:     0,
		Data:      converter.PostListToProto(posts),
	}, nil
	// }
	// posts, ok := data.([]entity.Post)
	// if !ok {
	// 	return nil, fmt.Errorf("Posts data is corrupted")
	// }
	// return &proto.PostResponse{
	// 	Page:      0,
	// 	PageCount: 0,
	// 	Total:     0,
	// 	Data:      converter.PostListToProto(posts),
	// }, nil

}

func (s *Server) GetPost(ctx context.Context, req *proto.PostRequest) (*proto.Post, error) {
	// cacheKey := fmt.Sprintf("/posts/%d", req.GetId())
	// data, err := s.cache.Get(ctx, cacheKey)
	// if err != nil || err == redis.Nil {
	service := usecase.ReadPost(s.db)
	post, err := service.Find(ctx, entity.PostFilter{
		ID: int(req.GetId()),
	})
	if err != nil {
		return nil, err
	}
	if post.ID == 0 {
		return nil, fmt.Errorf("Not found: %d", req.GetId())
	}
	// s.cache.Set(ctx, cacheKey, post, time.Minute*2)
	return converter.PostToProto(post), nil
	// }
	// post, ok := data.(entity.Post)
	// if !ok {
	// 	return nil, fmt.Errorf("data is corrupted")
	// }
	// if post.ID == 0 {
	// 	return nil, fmt.Errorf("Not found: %d", req.GetId())
	// }
	// return converter.PostToProto(post), nil
}

// bidirection
func (s *Server) GetPostStream(_ *proto.PostRequest, _ proto.PostService_GetPostStreamServer) error {
	panic("not implemented") // TODO: Implement
}
