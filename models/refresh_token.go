package models

import (
	"time"
)

// RefreshToken represents a user's refresh token in SQL
type RefreshToken struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	JTI       string    `gorm:"size:255;not null;unique" json:"jti"`
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
