package rest

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/usecase"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h HTTPServer) GetAllPost(c *fiber.Ctx) error {
	data, found := h.cache.Get(c.Path())
	if !found {
		service := usecase.ReadPost(h.db)
		posts, err := service.GetAll(context.Background(), entity.PostFilter{})
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err})
		}
		h.cache.Set(c.Path(), posts, time.Minute*2)
		return c.JSON(posts)
	}
	return c.JSON(data)
}

func (h HTTPServer) GetPost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err})
	}
	data, found := h.cache.Get(c.Path())
	if !found {
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
		h.cache.Set(c.Path(), post, time.Minute*2)
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
