package usecase

import (
	"context"
	"time"

	authEntity "github.com/Badankamon/gochat_backend/internal/modules/auth/domain/entity"
	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	"github.com/Badankamon/gochat_backend/internal/modules/user/application/dto"
)

type UpdateProfileUseCase struct {
	userRepo repository.UserRepository
}

func NewUpdateProfileUseCase(userRepo repository.UserRepository) *UpdateProfileUseCase {
	return &UpdateProfileUseCase{userRepo: userRepo}
}

func (uc *UpdateProfileUseCase) Execute(ctx context.Context, userID string, req dto.UpdateProfileRequest) (*authEntity.User, error) {
	// 1. Find existing user
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	// 2. Update fields
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Gender != "" {
		user.Gender = req.Gender
	}
	if req.Region != "" {
		user.Region = req.Region
	}
	if req.Bio != "" {
		user.Signature = req.Bio // Mapping Bio DTO -> Signature DB
	}
	user.UpdatedAt = time.Now()

	// 3. Save
	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
