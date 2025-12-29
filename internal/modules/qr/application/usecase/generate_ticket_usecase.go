package usecase

import (
	"context"
	"time"

	"github.com/Badankamon/gochat_backend/internal/config"
	"github.com/Badankamon/gochat_backend/internal/modules/qr/domain/entity"
	"github.com/Badankamon/gochat_backend/internal/modules/qr/domain/repository"
)

type GenerateTicketUseCase struct {
	repo repository.TicketRepository
	cfg  *config.Config
}

func NewGenerateTicketUseCase(repo repository.TicketRepository, cfg *config.Config) *GenerateTicketUseCase {
	return &GenerateTicketUseCase{repo: repo, cfg: cfg}
}

func (uc *GenerateTicketUseCase) Execute(ctx context.Context, userID, qrType, targetID, mode string) (*entity.QRTicket, error) {
	ticket := &entity.QRTicket{
		Type:      qrType,
		TargetID:  targetID,
		Mode:      mode,
		CreatedBy: userID,
	}

	if mode == "temporary" {
		// Temporary QRs expire in 7 days (or user configurable)
		exp := time.Now().Add(7 * 24 * time.Hour)
		ticket.ExpiresAt = &exp
	}

	if err := uc.repo.Create(ctx, ticket); err != nil {
		return nil, err
	}

	return ticket, nil
}
