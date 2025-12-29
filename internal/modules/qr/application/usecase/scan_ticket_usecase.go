package usecase

import (
	"context"
	"errors"
	"time"

	authRepo "github.com/Badankamon/gochat_backend/internal/modules/auth/domain/repository"
	qrRepo "github.com/Badankamon/gochat_backend/internal/modules/qr/domain/repository"
)

type ScanTicketUseCase struct {
	ticketRepo qrRepo.TicketRepository
	userRepo   authRepo.UserRepository
	// groupRepo  groupRepo.GroupRepository // When Group module exists
}

func NewScanTicketUseCase(ticketRepo qrRepo.TicketRepository, userRepo authRepo.UserRepository) *ScanTicketUseCase {
	return &ScanTicketUseCase{
		ticketRepo: ticketRepo,
		userRepo:   userRepo,
	}
}

type ScanResult struct {
	Type   string      `json:"type"`
	Target interface{} `json:"target"`
}

func (uc *ScanTicketUseCase) Execute(ctx context.Context, ticketStr string) (*ScanResult, error) {
	ticket, err := uc.ticketRepo.FindByTicket(ctx, ticketStr)
	if err != nil {
		return nil, err
	}
	if ticket == nil {
		return nil, errors.New("invalid_ticket")
	}

	// Check Expiration
	if ticket.ExpiresAt != nil && time.Now().After(*ticket.ExpiresAt) {
		return nil, errors.New("expired_ticket")
	}

	// Resolve Target
	if ticket.Type == "user" {
		user, err := uc.userRepo.FindByID(ctx, ticket.TargetID)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New("user_not_found")
		}
		// Return public profile info
		return &ScanResult{
			Type: "user",
			Target: map[string]interface{}{
				"id":       user.ID,
				"nickname": user.Nickname,
				"avatar":   user.Avatar,
				"region":   user.Region,
			},
		}, nil
	} else if ticket.Type == "group" {
		// Placeholder for group lookup
		return &ScanResult{
			Type: "group",
			Target: map[string]interface{}{
				"id":     ticket.TargetID,
				"status": "group_module_not_implemented_yet",
			},
		}, nil
	}

	return nil, errors.New("unknown_ticket_type")
}
