package postgres

import (
	"context"
	"errors"

	"github.com/Badankamon/gochat_backend/internal/modules/qr/domain/entity"
	"github.com/Badankamon/gochat_backend/internal/modules/qr/domain/repository"
	"gorm.io/gorm"
)

type TicketRepositoryImpl struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) repository.TicketRepository {
	return &TicketRepositoryImpl{db: db}
}

func (r *TicketRepositoryImpl) Create(ctx context.Context, ticket *entity.QRTicket) error {
	return r.db.WithContext(ctx).Create(ticket).Error
}

func (r *TicketRepositoryImpl) FindByTicket(ctx context.Context, ticketStr string) (*entity.QRTicket, error) {
	var ticket entity.QRTicket
	if err := r.db.WithContext(ctx).Where("ticket = ?", ticketStr).First(&ticket).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &ticket, nil
}
