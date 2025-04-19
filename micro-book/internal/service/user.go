package service

import (
	"context"
	"errors"
	"micro-book/internal/domain"
	"micro-book/internal/repository"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

var (
	InvalidUserOrPasswordError = errors.New("账号或密码不对")
	DuplicateUserEmailError    = errors.New("邮箱已被注册")
	InvalidUserEmailError      = errors.New("邮箱不存在")
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
		return domain.User{}, InvalidUserOrPasswordError
	}
	if err != nil {
		return domain.User{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password))
	if err != nil {
		return domain.User{}, InvalidUserOrPasswordError
	}
	return u, nil
}

func (us *UserService) EditService(ctx context.Context, id int64, result domain.User) (domain.User, error) {
	u, err := us.repo.UpdateById(ctx, id, result)
	if err == repository.UserNotFoundError {
		return domain.User{}, InvalidUserEmailError
	}
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (us *UserService) ProfileService(ctx context.Context, id string) (domain.User, error) {
	intId, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		return domain.User{}, err
	}
	u, err := us.repo.FindById(ctx, intId)
	if err == repository.UserNotFoundError {
		return domain.User{}, InvalidUserEmailError
	}
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}
