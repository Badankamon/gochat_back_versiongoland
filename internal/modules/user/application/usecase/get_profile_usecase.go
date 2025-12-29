package usecase

import (
	"context"

	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	"github.com/Badankamon/gochat_backend/internal/modules/user/application/dto"
)

type GetProfileUseCase struct {
	userRepo repository.UserRepository
}

func NewGetProfileUseCase(userRepo repository.UserRepository) *GetProfileUseCase {
	return &GetProfileUseCase{
		userRepo: userRepo,
	}
}

func (uc *GetProfileUseCase) Execute(ctx context.Context, userID string) (*dto.UserProfileResponse, error) {
	// 1. Get User
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil // Or specific error
	}

	return &dto.UserProfileResponse{
		ID:        user.ID,
		Phone:     user.Phone,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Gender:    user.Gender,
		Region:    user.Region,
		Bio:       user.Signature, // Map Signature to Bio in DTO if needed, or update DTO
		CreatedAt: user.CreatedAt,
	}, nil
}
