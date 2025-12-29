package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QRTicket struct {
	Ticket    string     `json:"ticket" gorm:"primaryKey;type:varchar(64)"`
	Type      string     `json:"type" gorm:"not null"` // user | group
	TargetID  string     `json:"target_id" gorm:"not null"`
	Mode      string     `json:"mode" gorm:"not null"` // permanent | temporary
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedBy string     `json:"created_by" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (t *QRTicket) BeforeCreate(tx *gorm.DB) (err error) {
	if t.Ticket == "" {
		t.Ticket = uuid.New().String()
	}
	return
}
