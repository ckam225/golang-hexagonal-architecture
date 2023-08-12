package rest

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/usecase"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func (h HTTPServer) GetAllPost(c *fiber.Ctx) error {
	ctx := context.Background()
	data, err := h.cache.Get(ctx, c.Path())
	if err != nil || err == redis.Nil {
		service := usecase.ReadPost(h.db)
		posts, err := service.GetAll(context.Background(), entity.PostFilter{})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err})
		}
		h.cache.Set(ctx, c.Path(), posts, time.Minute*2)
		return c.JSON(posts)
	}
	return c.JSON(data)
}

func (h HTTPServer) GetPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}
	ctx := context.Background()
	data, err := h.cache.Get(ctx, c.Path())
	if err != nil || err == redis.Nil {
		service := usecase.ReadPost(h.db)
		post, err := service.Find(context.Background(), entity.PostFilter{
			ID: id,
		})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		if post.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": fmt.Sprintf("Not found: %d", id)})
		}
		h.cache.Set(ctx, c.Path(), post, time.Minute*2)
		return c.JSON(post)
	}
	post, ok := data.(entity.Post)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error": "Data is corrupted"})
	}
	if post.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"error": fmt.Sprintf("Not found: %d", id)})
	}
	return c.JSON(post)
}

func (h HTTPServer) CreatePost(c *fiber.Ctx) error {
	var q entity.Post
	if err := c.BodyParser(&q); err != nil {
		return c.Status(422).JSON(err)
	}
	if q.Content == "" || q.Title == "" {
		log.Println("content or title is empty")
		return c.Status(422).JSON(fiber.ErrUnprocessableEntity)
	}
	ctx := context.Background()
	service := usecase.SavePost(h.db)
	if err := service.Create(ctx, &q); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": err.Error(),
			"code":    500,
		})
	}
	return c.JSON(q)
}
