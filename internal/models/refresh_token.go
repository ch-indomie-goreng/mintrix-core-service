package models

import (
	"time"

	"gorm.io/gorm"
)

// RefreshToken stores a refresh token linked to a user.
type RefreshToken struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `gorm:"index;not null"`
	Token     string         `gorm:"type:varchar(512);uniqueIndex;not null"`
	ExpiresAt time.Time      `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the default table name to "refresh_tokens".
func (RefreshToken) TableName() string {
	return "refresh_tokens"
}
