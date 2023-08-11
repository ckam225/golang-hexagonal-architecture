package usecase

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/repository"
	"context"
)

type PostUseCase struct {
	repo repository.PostRepository
}

func NewPostUseCase(repo repository.PostRepository) *PostUseCase {
	return &PostUseCase{
		repo: repo,
	}
}

func (uc *PostUseCase) GetAll(ctx context.Context, f entity.PostFilter) ([]entity.Post, error) {
	return uc.repo.GetPosts(ctx, f)
}
