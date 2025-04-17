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
	if err != nil {
		return err
	}
	if u.Email == user.Email {
		return DuplicateUserEmailError
	}
	// 存毫秒数
	now := time.Now().UnixMilli()

	user.CreateTime = now
	user.UpdateTime = now

	return ud.db.WithContext(ctx).Create(&user).Error
}

func (ud *UserDAO) Query(ctx context.Context, user UserEntity) error {
	return ud.db.WithContext(ctx).Select(&user).Error
}

func (ud *UserDAO) FindByEmail(ctx context.Context, email string) (UserEntity, error) {
	var u UserEntity
	// err := ud.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	err := ud.db.WithContext(ctx).Where(&UserEntity{Email: email}).First(&u).Error
	return u, err
}

// User 直接对应于数据库表
type UserEntity struct {
	Id       int64  `gorm:"primaryKey,autoIncreatement"`
	Email    string `gorm:"unique"`
	Password string

	CreateTime int64 // 毫秒数 规避时区问题
	UpdateTime int64
}
