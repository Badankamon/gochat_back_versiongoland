package repository

import (
	"context"

	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
}

type SessionRepository interface {
	Create(ctx context.Context, session *entity.Session) error
	FindByRefreshToken(ctx context.Context, token string) (*entity.Session, error)
	Delete(ctx context.Context, id string) error
	DeleteByRefreshToken(ctx context.Context, token string) error
}
