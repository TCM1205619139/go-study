package sms

import "context"

type SmsService interface {
	Send(
		ctx context.Context,
		// appId string,
		// signature string,
		template string,
		args []string,
		numbers ...string,
	) error
}
