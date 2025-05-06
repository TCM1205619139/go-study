package sms

import "context"

type SmsService interface {
	Send(
		ctx context.Context,
		template string,
		args []string,
		numbers ...string,
	) error
}
