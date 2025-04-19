package dao

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	DuplicateUserEmailError = errors.New("邮箱冲突")
	UserNotFoundError       = gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (ud *UserDAO) Insert(ctx context.Context, user UserEntity) error {
	var u UserEntity
	err := ud.db.WithContext(ctx).Where(&UserEntity{Email: user.Email}).First(&u).Error
	if u.Email == user.Email {
		return DuplicateUserEmailError
	}
	if err.Error() == "record not found" {
		// 存毫秒数
		now := time.Now().UnixMilli()

		user.CreateTime = now
		user.UpdateTime = now

		return ud.db.WithContext(ctx).Create(&user).Error
	}
	return err
}

func (ud *UserDAO) FindByEmail(ctx context.Context, email string) (UserEntity, error) {
	var u UserEntity
	err := ud.db.WithContext(ctx).Where(&UserEntity{Email: email}).First(&u).Error
	return u, err
}

func (ud *UserDAO) FindById(ctx context.Context, id int64) (UserEntity, error) {
	var u UserEntity
	err := ud.db.WithContext(ctx).Where(&UserEntity{Id: id}).First(&u).Error
	return u, err
}

func (ud *UserDAO) UpdateByEmail(ctx context.Context, email string, user UserEntity) (UserEntity, error) {
	var u UserEntity

	now := time.Now().UnixMilli()

	user.UpdateTime = now
	err := ud.db.WithContext(ctx).Where(&UserEntity{Email: email}).Updates(&user).First(&u).Error
	return u, err
}

func (ud *UserDAO) UpdateById(ctx context.Context, id int64, user UserEntity) (UserEntity, error) {
	var u UserEntity

	now := time.Now().UnixMilli()

	user.UpdateTime = now
	err := ud.db.WithContext(ctx).Where(&UserEntity{Id: id}).Updates(&user).First(&u).Error
	return u, err
}

// User 直接对应于数据库表
type UserEntity struct {
	Id          int64  `gorm:"primaryKey,autoIncreatement"`
	Email       string `gorm:"unique"`
	Password    string
	NickName    string
	Birthday    string
	Description string

	CreateTime int64 // 毫秒数 规避时区问题
	UpdateTime int64
}
