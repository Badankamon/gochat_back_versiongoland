package external

import (
	"context"
	"log"

	"github.com/Badankamon/gochat_backend/internal/modules/auth/application/service"
)

type MockSMSService struct{}

func NewMockSMSService() service.SMSService {
	return &MockSMSService{}
}

func (s *MockSMSService) SendVerificationCode(ctx context.Context, phone string, code string) error {
	log.Printf("[MOCK SMS] Sending code %s to %s", code, phone)
	return nil
}
