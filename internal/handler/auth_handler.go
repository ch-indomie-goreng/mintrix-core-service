package handler

import (
	"log/slog"

	"mintrix-backend/internal/models"
	"mintrix-backend/internal/service"
	"mintrix-backend/internal/utils"

	"github.com/gin-gonic/gin"
)

// AuthHandler exposes authentication HTTP endpoints.
type AuthHandler struct {
	svc service.AuthService
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// Register handles POST /api/v1/auth/register.
// @Summary Register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "Registration payload"
// @Success 201 {object} utils.Envelope{data=models.AuthResponse}
// @Failure 400 {object} utils.Envelope
// @Failure 409 {object} utils.Envelope
// @Router /api/v1/auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	resp, err := h.svc.Register(c.Request.Context(), req)
	if err != nil {
		slog.ErrorContext(c.Request.Context(), "register failed", "error", err)
		utils.Conflict(c, err.Error())
		return
	}

	utils.Created(c, resp)
}

// Login handles POST /api/v1/auth/login.
// @Summary Authenticate a user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login payload"
// @Success 200 {object} utils.Envelope{data=models.AuthResponse}
// @Failure 400 {object} utils.Envelope
// @Failure 401 {object} utils.Envelope
// @Router /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	resp, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		slog.ErrorContext(c.Request.Context(), "login failed", "error", err)
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.OK(c, resp)
}

// RefreshToken handles POST /api/v1/auth/refresh.
// @Summary Refresh an access token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RefreshRequest true "Refresh token payload"
// @Success 200 {object} utils.Envelope{data=models.AuthResponse}
// @Failure 400 {object} utils.Envelope
// @Failure 401 {object} utils.Envelope
// @Router /api/v1/auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req models.RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	resp, err := h.svc.RefreshToken(c.Request.Context(), req)
	if err != nil {
		slog.ErrorContext(c.Request.Context(), "refresh failed", "error", err)
		utils.Unauthorized(c, err.Error())
		return
	}

	utils.OK(c, resp)
}

// Me handles GET /api/v1/auth/me (requires authentication).
// @Summary Get current user profile
// @Tags Auth
// @Produce json
// @Security BearerAuth
// @Success 200 {object} utils.Envelope{data=models.UserResponse}
// @Failure 401 {object} utils.Envelope
// @Router /api/v1/auth/me [get]
func (h *AuthHandler) Me(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		utils.Unauthorized(c, "missing user context")
		return
	}

	resp, err := h.svc.GetProfile(c.Request.Context(), userID.(uint))
	if err != nil {
		slog.ErrorContext(c.Request.Context(), "get profile failed", "error", err)
		utils.NotFound(c, err.Error())
		return
	}

	utils.OK(c, resp)
}
