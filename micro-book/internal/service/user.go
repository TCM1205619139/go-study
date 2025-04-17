package service

import (
	"context"
	"errors"
	"micro-book/internal/domain"
	"micro-book/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

var (
	InvavidUserOrPasswordError = errors.New("账号或密码不对")
	DuplicateUserEmailError    = errors.New("邮箱已被注册")
)

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) SignupService(ctx context.Context, user domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err == repository.DuplicateUserEmailError {
		return DuplicateUserEmailError
	}
	if err != nil {
		return err
	}
	user.Password = string(hash)
	if err = us.repo.Create(ctx, user); err != nil {
		return DuplicateUserEmailError
	}
	return nil
}

func (us *UserService) SigninService(ctx context.Context, user domain.User) (domain.User, error) {
	u, err := us.repo.FindByEmail(ctx, user.Email)
	if err == repository.UserNotFoundError {
		return domain.User{}, InvavidUserOrPasswordError
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, InvavidUserOrPasswordError
	}
	return u, nil
}
