package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Badankamon/gochat_backend/internal/config"
	"github.com/Badankamon/gochat_backend/internal/modules/auth/application/dto"
	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/entity"
	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	"github.com/Badankamon/gochat_backend/internal/shared/crypto"
)

type LoginUseCase struct {
	userRepo         repository.UserRepository
	verificationRepo repository.VerificationRepository
	sessionRepo      repository.SessionRepository
	cfg              *config.Config
}

func NewLoginUseCase(
	userRepo repository.UserRepository,
	verificationRepo repository.VerificationRepository,
	sessionRepo repository.SessionRepository,
	cfg *config.Config,
) *LoginUseCase {
	return &LoginUseCase{
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
		sessionRepo:      sessionRepo,
		cfg:              cfg,
	}
}

func (uc *LoginUseCase) Execute(ctx context.Context, req dto.LoginRequest) (*dto.AuthResponse, error) {
	// 1. Verify Code
	storedCode, err := uc.verificationRepo.GetCode(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if storedCode == "" || storedCode != req.Code {
		return nil, errors.New("invalid verification code")
	}

	// 2. Find User
	user, err := uc.userRepo.FindByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 3. Generate Token
	accessToken, err := crypto.GenerateToken(user.ID, uc.cfg.JWT)
	if err != nil {
		return nil, err
	}
	refreshToken, err := crypto.GenerateToken(user.ID, uc.cfg.JWT)
	if err != nil {
		return nil, err
	}

	// 4. Create Session
	session := &entity.Session{
		UserID:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    "unknown",
		ExpiresAt:    time.Now().Add(time.Duration(uc.cfg.JWT.RefreshExpireDays) * 24 * time.Hour),
	}
	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	// 5. Delete OTP
	_ = uc.verificationRepo.DeleteCode(ctx, req.Phone)

	return &dto.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
