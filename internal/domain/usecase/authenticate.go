package usecase

import (
	"clean-arch-hex/internal/domain/entity"
	"clean-arch-hex/internal/domain/exception"
	"clean-arch-hex/internal/domain/repository"
	"context"
)

type LoginUseCase struct {
	repo repository.UserRepository
}

func NewLoginUseCase(repo repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		repo: repo,
	}
}

func (uc *LoginUseCase) Login(ctx context.Context, email string, password string) error {
	user, err := uc.repo.FindUser(ctx, entity.UserFilter{
		Email:    email,
		Password: password, // TODO: should be encrypted
	})
	if err != nil {
		return err
	}
	if !user.IsActive {
		return exception.ErrAccountSuspended
	}
	// if !security.VerifyPassword(*user.Password, password) {
	// 	return exception.ErrInvalidCredentials
	// }
	return nil
}

func (uc *LoginUseCase) CheckEmail(ctx context.Context, email string) error {
	panic("unimplemented")
}
