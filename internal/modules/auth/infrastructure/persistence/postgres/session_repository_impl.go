package postgres

import (
	"context"
	"errors"

	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/entity"
	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	"gorm.io/gorm"
)

type SessionRepositoryImpl struct {
	db *gorm.DB
}

func NewSessionRepository(db *gorm.DB) repository.SessionRepository {
	return &SessionRepositoryImpl{db: db}
}

func (r *SessionRepositoryImpl) Create(ctx context.Context, session *entity.Session) error {
	return r.db.WithContext(ctx).Create(session).Error
}

func (r *SessionRepositoryImpl) FindByRefreshToken(ctx context.Context, token string) (*entity.Session, error) {
	var session entity.Session
	if err := r.db.WithContext(ctx).Where("refresh_token = ?", token).First(&session).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.Session{}, "id = ?", id).Error
}

func (r *SessionRepositoryImpl) DeleteByRefreshToken(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Delete(&entity.Session{}, "refresh_token = ?", token).Error
}
