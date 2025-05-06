package code

import (
	"context"
	"fmt"
	"math/rand"
	"micro-book/internal/repository"
	"micro-book/internal/service/sms/tencent"
)

type CodeService struct {
	repo *repository.CodeRepository
	sms  *tencent.SmsService
}

func NewCodeService(repo *repository.CodeRepository, sms *tencent.SmsService) *CodeService {
	return &CodeService{
		repo: repo,
		sms:  sms,
	}
}

func (cs *CodeService) Send(
	ctx context.Context,
	biz string,
	phone string,
) error {
	code := cs.generateCode()
	err := cs.repo.Store(ctx, biz, phone, code)
	if err != nil {
		return err
	}
	cs.sms.Send(ctx, "", []string{code}, phone)

	return nil
}

func (cs *CodeService) Valid(
	ctx context.Context,
	biz string,
	phnoe string,
	code string,
) (bool, error) {
	err := cs.repo.Verify(ctx, biz, phnoe, code)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cs *CodeService) generateCode() string {
	num := rand.Intn(1e6)

	return fmt.Sprintf("%6d", num)
}
