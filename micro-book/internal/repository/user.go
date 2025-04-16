package repository

import (
	"context"
	"micro-book/internal/domain"
	"micro-book/internal/repository/dao"
)

type UserRepository struct {
	dao *dao.UserDAO
}

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
