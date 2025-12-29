package usecase

import (
	"context"
	"io"
	"time"

	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	"github.com/Badankamon/gochat_backend/internal/platform/storage"
)

type UploadAvatarUseCase struct {
	userRepo repository.UserRepository
	storage  storage.Service
}

func NewUploadAvatarUseCase(
	userRepo repository.UserRepository,
	storage storage.Service,
) *UploadAvatarUseCase {
	return &UploadAvatarUseCase{
		userRepo: userRepo,
		storage:  storage,
	}
}

func (uc *UploadAvatarUseCase) Execute(ctx context.Context, userID string, file io.Reader, filename string) (string, error) {
	// 1. Upload to storage
	url, err := uc.storage.Upload(ctx, file, filename)
	if err != nil {
		return "", err
	}

	// 2. Update User with new URL
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return "", err
	}

	if user != nil {
		user.Avatar = url
		user.UpdatedAt = time.Now()
		if err := uc.userRepo.Update(ctx, user); err != nil {
			return "", err
		}
	}

	return url, nil
}
