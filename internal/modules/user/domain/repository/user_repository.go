package repository

import (
	"context"

	"github.com/Badankamon/gochat_backend/internal/modules/user/domain/entity"
)

type UserProfileRepository interface {
	Create(ctx context.Context, profile *entity.Profile) error
	FindByUserID(ctx context.Context, userID string) (*entity.Profile, error)
	Update(ctx context.Context, profile *entity.Profile) error
}
