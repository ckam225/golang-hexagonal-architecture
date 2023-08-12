package usecase

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/repository"
	"context"
)

type SavePostUseCase struct {
	repo repository.PostRepository
}

func SavePost(repo repository.PostRepository) *ReadPostUseCase {
	return &ReadPostUseCase{
		repo: repo,
	}
}

func (uc *ReadPostUseCase) Create(ctx context.Context, p *entity.Post) error {
	return uc.repo.CreatePost(ctx, p)
}
