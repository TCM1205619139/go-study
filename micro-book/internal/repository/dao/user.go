package dao

import (
	"context"
	"time"

	"gorm.io/gorm"
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
	// 存毫秒数
	now := time.Now().UnixMilli()

	user.CreateTime = now
	user.UpdateTime = now

	return ud.db.WithContext(ctx).Create(&user).Error
}

// User 直接对应于数据库表
type UserEntity struct {
	Id       int64  `gorm:"primaryKey,autoIncreatement"`
	Email    string `gorm:"unique"`
	Password string

	CreateTime int64 // 毫秒数 规避时区问题
	UpdateTime int64
}
