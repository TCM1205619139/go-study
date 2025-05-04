package tencent

import (
	"context"
	"fmt"

	"github.com/ecodeclub/ekit/slice"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

type Service struct {
	appId     *string
	signature *string
	client    *sms.Client
}

func NewService(appId string, signature string, client *sms.Client) *Service {
	return &Service{
		appId:     &appId,
		signature: &signature,
		client:    client,
	}
}

func (service *Service) Send(ctx context.Context, tpl string, args []string, numbers ...string) error {
	req := sms.NewSendSmsRequest()
	req.SmsSdkAppId = service.appId
	req.SignName = service.signature
	req.TemplateId = &tpl
	req.PhoneNumberSet = slice.Map[string, *string](numbers, func(index int, src string) *string {
		return &src
	})
	req.TemplateParamSet = slice.Map[string, *string](args, func(index int, arg string) *string {
		return &arg
	})

	resp, err := service.client.SendSms(req)
	if err != nil {
		return err
	}
	for _, status := range resp.Response.SendStatusSet {
		if status.Code == nil || *(status.Code) != "Ok" {
			return fmt.Errorf("发送短信失败 %s. %s", *status.Code, *status.Message)
		}
	}
	return nil
}
