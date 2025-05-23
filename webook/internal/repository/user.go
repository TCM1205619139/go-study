package repository

import (
	"context"
	"micro-book/internal/domain"
	"micro-book/internal/repository/cache"
	"micro-book/internal/repository/dao"
	"micro-book/pkg"
	"strconv"
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

var (
	UserNotFoundError       = dao.UserNotFoundError
	DuplicateUserEmailError = dao.DuplicateUserEmailError
)

func NewUserRepository(dao *dao.UserDAO, cache *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: cache,
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
		Id:          strconv.FormatInt(user.Id, 10),
		Email:       user.Email,
		Password:    user.Password,
		NickName:    user.NickName,
		Birthday:    user.Birthday,
		Description: user.Description,
	}, nil
}

func (ur *UserRepository) FindById(ctx context.Context, id int64) (domain.User, error) {
	user, ok, err := ur.cache.GetUser(ctx, id)
	if ok {
		return user, nil
	}
	if err != nil {
		return domain.User{}, err
	}

	ue, err := ur.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	user = domain.User{
		Id:          strconv.FormatInt(ue.Id, 10),
		Email:       ue.Email,
		Password:    ue.Password,
		NickName:    ue.NickName,
		Birthday:    ue.Birthday,
		Description: ue.Description,
	}

	go func() {
		if err = ur.cache.SetUser(ctx, user); err != nil {
			// 不用返回错误，做日志监控就可以了
			// return domain.User{}, err
		}
	}()

	return user, nil
}

func (ur *UserRepository) UpdateByEmail(ctx context.Context, email string, user domain.User) (domain.User, error) {
	u, err := ur.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	u, err = ur.dao.UpdateByEmail(ctx, email, dao.UserEntity{
		Email:       u.Email,
		Password:    u.Password,
		Id:          u.Id,
		NickName:    pkg.MaybeString(user.NickName, u.NickName),
		Birthday:    pkg.MaybeString(user.Birthday, u.Birthday),
		Description: pkg.MaybeString(user.Description, u.Description),
	})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Id:          user.Id,
		Email:       u.Email,
		NickName:    u.NickName,
		Birthday:    u.Birthday,
		Description: u.Description,
	}, nil
}

func (ur *UserRepository) UpdateById(ctx context.Context, id int64, user domain.User) (domain.User, error) {
	u, err := ur.dao.FindById(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	u, err = ur.dao.UpdateById(ctx, id, dao.UserEntity{
		Password:    pkg.MaybeString(user.Password, u.Password),
		Email:       pkg.MaybeString(user.Email, u.Email),
		NickName:    pkg.MaybeString(user.NickName, u.NickName),
		Birthday:    pkg.MaybeString(user.Birthday, u.Birthday),
		Description: pkg.MaybeString(user.Description, u.Description),
	})
	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Id:          strconv.FormatInt(id, 10),
		Email:       u.Email,
		NickName:    u.NickName,
		Birthday:    u.Birthday,
		Description: u.Description,
	}, nil
}
