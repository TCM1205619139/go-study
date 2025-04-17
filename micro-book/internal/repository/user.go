package repository

import (
	"context"
	"micro-book/internal/domain"
	"micro-book/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

var (
	UserNotFoundError       = dao.UserNotFoundError
	DuplicateUserEmailError = dao.DuplicateUserEmailError
)

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (ur *UserRepository) Create(ctx context.Context, user domain.User) error {
	return ur.dao.Insert(ctx, dao.UserEntity{
		Email:    user.Email,
		Password: user.Password,
	})
}

func (ur *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	user, err := ur.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
