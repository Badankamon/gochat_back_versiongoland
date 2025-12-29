package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	GoChatID  string    `json:"gochat_id" gorm:"unique;not null;column:gochat_id;index"`
	Phone     string    `json:"phone" gorm:"unique;not null;index"`
	Nickname  string    `json:"nickname" gorm:"not null"`
	Avatar    string    `json:"avatar"`
	Gender    string    `json:"gender" gorm:"default:'not_specified'"`
	Region    string    `json:"region"`
	Signature string    `json:"signature" gorm:"column:signature"` // Matches 'signature' in SQL
	QRCode    string    `json:"qr_code" gorm:"column:qr_code"`
	Status    string    `json:"status" gorm:"default:'active'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate hook to generate UUID if not present and handle any defaults not handled by DB
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}
