package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents the users table.
type User struct {
	ID              uint           `gorm:"primaryKey"`
	Nama            string         `gorm:"type:varchar(255);not null"`
	Email           string         `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password        string         `gorm:"type:varchar(255);not null"`
	Personalization bool           `gorm:"default:false"`
	Foto            *string        `gorm:"type:varchar(255);default:null"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

// TableName overrides the default table name to "users".
func (User) TableName() string {
	return "users"
}

// BeforeCreate always hashes the password for new records.
func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

// BeforeUpdate hashes the password only when it has been changed.
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if !tx.Statement.Changed("Password") {
		return nil
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashed)
	return nil
}

// ComparePassword returns true when the candidate matches the stored hash.
func (u *User) ComparePassword(candidate string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(candidate))
	return err == nil
}
