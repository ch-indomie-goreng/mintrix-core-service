package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"mintrix-backend/internal/auth"
	"mintrix-backend/internal/models"
	"mintrix-backend/internal/repository"

	"gorm.io/gorm"
)

// AuthService defines the business logic contract for authentication.
type AuthService interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error)
	RefreshToken(ctx context.Context, req models.RefreshRequest) (*models.AuthResponse, error)
	GetProfile(ctx context.Context, userID uint) (*models.UserResponse, error)
}

// authService implements AuthService.
type authService struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
	jwtSecret string
	jwtExpiry time.Duration
	refreshExpiry time.Duration
}

// NewAuthService wires the auth service with its dependencies.
func NewAuthService(
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
	jwtSecret string,
	jwtExpiry time.Duration,
	refreshExpiry time.Duration,
) AuthService {
	return &authService{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		jwtSecret: jwtSecret,
		jwtExpiry: jwtExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (s *authService) Register(ctx context.Context, req models.RegisterRequest) (*models.AuthResponse, error) {
	// Check duplicate email.
	if _, err := s.userRepo.FindByEmail(ctx, req.Email); err == nil {
		return nil, fmt.Errorf("email already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.ErrorContext(ctx, "failed to check email", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	user := models.User{
		Nama:     req.Nama,
		Email:    req.Email,
		Password: req.Password, // hashed by BeforeSave hook
	}

	if err := s.userRepo.Create(ctx, &user); err != nil {
		slog.ErrorContext(ctx, "failed to create user", "error", err)
		return nil, fmt.Errorf("failed to create user")
	}

	return s.issueTokenPair(ctx, user.ID)
}

func (s *authService) Login(ctx context.Context, req models.LoginRequest) (*models.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid email or password")
		}
		slog.ErrorContext(ctx, "failed to find user", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	if !user.ComparePassword(req.Password) {
		return nil, fmt.Errorf("invalid email or password")
	}

	return s.issueTokenPair(ctx, user.ID)
}

func (s *authService) RefreshToken(ctx context.Context, req models.RefreshRequest) (*models.AuthResponse, error) {
	rt, err := s.tokenRepo.FindByToken(ctx, req.RefreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("invalid refresh token")
		}
		slog.ErrorContext(ctx, "failed to find refresh token", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	if time.Now().After(rt.ExpiresAt) {
		_ = s.tokenRepo.DeleteByUserID(ctx, rt.UserID)
		return nil, fmt.Errorf("refresh token expired")
	}

	// Rotate: delete the used refresh token and issue a new pair.
	if err := s.tokenRepo.DeleteByUserID(ctx, rt.UserID); err != nil {
		slog.ErrorContext(ctx, "failed to delete old refresh token", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	return s.issueTokenPair(ctx, rt.UserID)
}

func (s *authService) GetProfile(ctx context.Context, userID uint) (*models.UserResponse, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		slog.ErrorContext(ctx, "failed to find user", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	resp := user.ToUserResponse()
	return &resp, nil
}

// issueTokenPair generates an access + refresh token pair and persists the refresh token.
func (s *authService) issueTokenPair(ctx context.Context, userID uint) (*models.AuthResponse, error) {
	accessToken, err := auth.GenerateAccessToken(userID, s.jwtSecret, s.jwtExpiry)
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate access token", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	refreshStr, err := auth.GenerateRefreshToken()
	if err != nil {
		slog.ErrorContext(ctx, "failed to generate refresh token", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	rt := models.RefreshToken{
		UserID:    userID,
		Token:     refreshStr,
		ExpiresAt: time.Now().Add(s.refreshExpiry),
	}

	if err := s.tokenRepo.Create(ctx, &rt); err != nil {
		slog.ErrorContext(ctx, "failed to store refresh token", "error", err)
		return nil, fmt.Errorf("internal error")
	}

	return &models.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshStr,
		ExpiresIn:    int64(s.jwtExpiry.Seconds()),
	}, nil
}
