package postgres

import (
	"context"
	"errors"

	"github.com/Badankamon/gochat_backend/internal/modules/user/domain/entity"
	"github.com/Badankamon/gochat_backend/internal/modules/user/domain/repository"
	"gorm.io/gorm"
)

type UserProfileRepositoryImpl struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) repository.UserProfileRepository {
	return &UserProfileRepositoryImpl{db: db}
}

func (r *UserProfileRepositoryImpl) Create(ctx context.Context, profile *entity.Profile) error {
	return r.db.WithContext(ctx).Create(profile).Error
}

func (r *UserProfileRepositoryImpl) FindByUserID(ctx context.Context, userID string) (*entity.Profile, error) {
	var profile entity.Profile
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &profile, nil
}

func (r *UserProfileRepositoryImpl) Update(ctx context.Context, profile *entity.Profile) error {
	return r.db.WithContext(ctx).Save(profile).Error
}
