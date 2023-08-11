package usecase

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/repository"
	"context"
)

type ReadPostUseCase struct {
	repo repository.PostRepository
}

func ReadPost(repo repository.PostRepository) *ReadPostUseCase {
	return &ReadPostUseCase{
		repo: repo,
	}
}

func (uc *ReadPostUseCase) GetAll(ctx context.Context, f entity.PostFilter) ([]entity.Post, error) {
	return uc.repo.GetPosts(ctx, f)
}

func (uc *ReadPostUseCase) Find(ctx context.Context, f entity.PostFilter) (entity.Post, error) {
	return uc.repo.FindPost(ctx, f)
}
