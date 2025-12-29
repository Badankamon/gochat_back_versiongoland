package redis

import (
	"context"
	"time"

	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	"github.com/redis/go-redis/v9"
)

type VerificationRepositoryImpl struct {
	client *redis.Client
}

func NewVerificationRepository(client *redis.Client) repository.VerificationRepository {
	return &VerificationRepositoryImpl{client: client}
}

func (r *VerificationRepositoryImpl) SaveCode(ctx context.Context, phone string, code string, duration time.Duration) error {
	return r.client.Set(ctx, "otp:"+phone, code, duration).Err()
}

func (r *VerificationRepositoryImpl) GetCode(ctx context.Context, phone string) (string, error) {
	val, err := r.client.Get(ctx, "otp:"+phone).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (r *VerificationRepositoryImpl) DeleteCode(ctx context.Context, phone string) error {
	return r.client.Del(ctx, "otp:"+phone).Err()
}
