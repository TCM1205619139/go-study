package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"micro-book/internal/domain"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserCache struct {
	redis      redis.Cmdable // 这里不能用指针
	expiration time.Duration
}

func NewUserCache(client redis.Cmdable) *UserCache {
	return &UserCache{
		redis:      client,
		expiration: time.Minute * 15,
	}
}

func (uc *UserCache) GetUser(context context.Context, id int64) (domain.User, bool, error) {
	var user domain.User
	val, err := uc.redis.Get(context, uc.key(id)).Bytes()
	if err == redis.Nil {
		return user, false, nil
	}
	if err != nil {
		return user, false, err
	}
	err = json.Unmarshal(val, &user)
	if err != nil {
		return user, false, err
	}
	return user, true, err
}

func (uc *UserCache) SetUser(context context.Context, u domain.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	id, err := strconv.ParseInt(u.Id, 10, 64)
	if err != nil {
		return err
	}
	key := uc.key(id)
	uc.redis.Set(context, key, val, uc.expiration)
	return nil
}

func (uc *UserCache) key(id int64) string {
	return fmt.Sprintf("user:info:%d", id)
}
