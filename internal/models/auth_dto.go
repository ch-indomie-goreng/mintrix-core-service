package models

// RegisterRequest is the payload for POST /api/v1/auth/register.
type RegisterRequest struct {
	Nama     string `json:"nama"     binding:"required,min=1,max=255"`
	Email    string `json:"email"    binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=6,max=72"`
}

// LoginRequest is the payload for POST /api/v1/auth/login.
type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,min=1,max=72"`
}

// RefreshRequest is the payload for POST /api/v1/auth/refresh.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// AuthResponse is returned on successful register, login, and refresh.
type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// UserResponse is the public representation of a user (never includes password).
type UserResponse struct {
	ID              uint   `json:"id"`
	Nama            string `json:"nama"`
	Email           string `json:"email"`
	Personalization bool   `json:"personalization"`
	Foto            *string `json:"foto,omitempty"`
}

// ToUserResponse converts a User domain model to a safe public DTO.
func (u *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:              u.ID,
		Nama:            u.Nama,
		Email:           u.Email,
		Personalization: u.Personalization,
		Foto:            u.Foto,
	}
}
