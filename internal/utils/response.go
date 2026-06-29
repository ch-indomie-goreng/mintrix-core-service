package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Envelope is the standard JSON response wrapper.
type Envelope struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

// Success writes a successful JSON response.
func Success(c *gin.Context, status int, data interface{}) {
	c.JSON(status, Envelope{Data: data})
}

// Error writes an error JSON response.
func Error(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, Envelope{Error: message})
}

// Convenience helpers.
func OK(c *gin.Context, data interface{})            { Success(c, http.StatusOK, data) }
func Created(c *gin.Context, data interface{})       { Success(c, http.StatusCreated, data) }
func BadRequest(c *gin.Context, msg string)          { Error(c, http.StatusBadRequest, msg) }
func Unauthorized(c *gin.Context, msg string)        { Error(c, http.StatusUnauthorized, msg) }
func NotFound(c *gin.Context, msg string)             { Error(c, http.StatusNotFound, msg) }
func Conflict(c *gin.Context, msg string)             { Error(c, http.StatusConflict, msg) }
