package repository

import (
	"context"

	"mintrix-backend/internal/models"

	"gorm.io/gorm"
)

// TokenRepository defines the data access contract for refresh tokens.
type TokenRepository interface {
	Create(ctx context.Context, token *models.RefreshToken) error
	FindByToken(ctx context.Context, token string) (*models.RefreshToken, error)
	DeleteByUserID(ctx context.Context, userID uint) error
}

// tokenRepo implements TokenRepository.
type tokenRepo struct {
	db *gorm.DB
}

// NewTokenRepository returns a concrete TokenRepository backed by GORM.
func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepo{db: db}
}

func (r *tokenRepo) Create(ctx context.Context, token *models.RefreshToken) error {
	return r.db.WithContext(ctx).Create(token).Error
}

func (r *tokenRepo) FindByToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	err := r.db.WithContext(ctx).Where("token = ?", token).First(&rt).Error
	if err != nil {
		return nil, err
	}
	return &rt, nil
}

func (r *tokenRepo) DeleteByUserID(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error
}
