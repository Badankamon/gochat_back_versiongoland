package repository

import (
	"context"

	"github.com/Badankamon/gochat_backend/internal/modules/qr/domain/entity"
)

type TicketRepository interface {
	Create(ctx context.Context, ticket *entity.QRTicket) error
	FindByTicket(ctx context.Context, ticketStr string) (*entity.QRTicket, error)
}
