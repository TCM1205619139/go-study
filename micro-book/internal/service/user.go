package service

import (
	"context"
	"micro-book/internal/domain"
	"micro-book/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) SignupService(ctx context.Context, user domain.User) error {
	return us.repo.Create(ctx, user)
}
