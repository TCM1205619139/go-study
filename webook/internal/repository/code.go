package repository

import (
	"context"
	_ "embed"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type CodeRepository struct {
	redis redis.Cmdable
}

func NewCodeRepository(redis redis.Cmdable) *CodeRepository {
	return &CodeRepository{
		redis: redis,
	}
}

//go:embed lua/set_code.lua
var luaSetScript string

func (cr *CodeRepository) Store(
	ctx context.Context,
	biz string,
	phone string,
	code string,
) error {
	status, err := cr.redis.Eval(
		ctx, luaSetScript,
		[]string{cr.key(biz, phone)},
		code,
	).Int()

	if err != nil {
		return err
	}

	switch status {
	case -1:
		return errors.New("系统错误")
	case -2:
		return errors.New("发送太频繁")
	default:
		return nil
	}
}

//go:embed lua/verify_code.lua
var luaVerifyScript string

func (cr *CodeRepository) Verify(
	ctx context.Context,
	biz string,
	phone string,
	code string,
) error {
	status, err := cr.redis.Eval(ctx, luaVerifyScript, []string{cr.key(biz, phone)}, code).Int()
	if err != nil {
		return err
	}
	switch status {
	case -1:
		return errors.New("输入错误次数太多")
	case -2:
		return errors.New("验证码输入错误")
	default:
		return nil
	}
}

func (cr *CodeRepository) Clear(
	ctx context.Context,
	biz string,
	phone string,
	code string,
) error {
	cr.redis.Del(ctx, cr.key(biz, phone))
	return nil
}

func (cr *CodeRepository) key(biz string, phone string) string {
	return fmt.Sprintf("biz:%s;phone:%s", biz, phone)
}
