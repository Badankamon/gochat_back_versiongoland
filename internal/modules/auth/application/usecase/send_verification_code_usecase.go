package usecase

import (
	"context"
	"time"

	"github.com/Badankamon/gochat_backend/internal/modules/auth/application/service"
	"github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	"github.com/Badankamon/gochat_backend/internal/shared/utils"
)

type SendVerificationCodeUseCase struct {
	verificationRepo repository.VerificationRepository
	smsService       service.SMSService
}

func NewSendVerificationCodeUseCase(
	verificationRepo repository.VerificationRepository,
	smsService service.SMSService,
) *SendVerificationCodeUseCase {
	return &SendVerificationCodeUseCase{
		verificationRepo: verificationRepo,
		smsService:       smsService,
	}
}

func (uc *SendVerificationCodeUseCase) Execute(ctx context.Context, phone string) error {
	// Generate 6 digit OTP
	code := utils.GenerateOTP(6)

	// Save to Redis (5 minutes expiration)
	if err := uc.verificationRepo.SaveCode(ctx, phone, code, 5*time.Minute); err != nil {
		return err
	}

	// Send SMS
	return uc.smsService.SendVerificationCode(ctx, phone, code)
}
