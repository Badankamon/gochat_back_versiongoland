package entity

import (
	"time"
)

type Profile struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserID    string    `json:"user_id" gorm:"uniqueIndex;not null"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Gender    string    `json:"gender"` // male, female, other
	Region    string    `json:"region"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
