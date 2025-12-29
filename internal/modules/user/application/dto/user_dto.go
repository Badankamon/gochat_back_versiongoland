package dto

import "time"

type UpdateProfileRequest struct {
	Nickname string `json:"nickname" binding:"max=50"`
	Gender   string `json:"gender" binding:"omitempty,oneof=male female other"`
	Region   string `json:"region"`
	Bio      string `json:"bio" binding:"max=200"`
}

type UserProfileResponse struct {
	ID        string    `json:"id"`
	Phone     string    `json:"phone"` // From Auth User
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	Gender    string    `json:"gender"`
	Region    string    `json:"region"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}
