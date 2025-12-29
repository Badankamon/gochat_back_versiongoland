package service

import "context"

type SMSService interface {
	SendVerificationCode(ctx context.Context, phone string, code string) error
}
