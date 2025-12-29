package repository

import (
	"context"
	"time"
)

type VerificationRepository interface {
	SaveCode(ctx context.Context, phone string, code string, duration time.Duration) error
	GetCode(ctx context.Context, phone string) (string, error)
	DeleteCode(ctx context.Context, phone string) error
}
