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
	"github.com/Badankamon/gochat_backend/internal/shared/utils"
)

type RegisterUseCase struct {
	userRepo         repository.UserRepository
	verificationRepo repository.VerificationRepository
	sessionRepo      repository.SessionRepository
	cfg              *config.Config
}

func NewRegisterUseCase(
	userRepo repository.UserRepository,
	verificationRepo repository.VerificationRepository,
	sessionRepo repository.SessionRepository,
	cfg *config.Config,
) *RegisterUseCase {
	return &RegisterUseCase{
		userRepo:         userRepo,
		verificationRepo: verificationRepo,
		sessionRepo:      sessionRepo,
		cfg:              cfg,
	}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, req dto.RegisterRequest) (*dto.AuthResponse, error) {
	// 1. Verify Code
	storedCode, err := uc.verificationRepo.GetCode(ctx, req.Phone)
	if err != nil {
		return nil, err
	}
	if storedCode == "" || storedCode != req.Code {
		return nil, errors.New("invalid verification code")
	}

	// 2. Check if user already exists
	existingUser, _ := uc.userRepo.FindByPhone(ctx, req.Phone)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// 3. Create User
	// Generate random 8-char GoChatID for now
	goChatID := "gcid_" + utils.GenerateOTP(12) // Simple generation

	newUser := &entity.User{
		Phone:     req.Phone,
		GoChatID:  goChatID,
		Nickname:  "User_" + goChatID[5:], // Default nickname
		Status:    "active",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(ctx, newUser); err != nil {
		return nil, err
	}

	// 4. Generate Tokens
	accessToken, err := crypto.GenerateToken(newUser.ID, uc.cfg.JWT)
	if err != nil {
		return nil, err
	}
	// For simplicity, using same JWT logic for refresh token but longer expiry
	// Real world: specialized refresh token format
	refreshToken, err := crypto.GenerateToken(newUser.ID, uc.cfg.JWT)
	if err != nil {
		return nil, err
	}

	// 5. Create Session
	session := &entity.Session{
		UserID:       newUser.ID,
		RefreshToken: refreshToken,
		UserAgent:    "unknown", // Can be passed from context/req
		ExpiresAt:    time.Now().Add(time.Duration(uc.cfg.JWT.RefreshExpireDays) * 24 * time.Hour),
	}
	if err := uc.sessionRepo.Create(ctx, session); err != nil {
		return nil, err
	}

	// 6. Delete OTP
	_ = uc.verificationRepo.DeleteCode(ctx, req.Phone)

	return &dto.AuthResponse{
		User:         newUser,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
