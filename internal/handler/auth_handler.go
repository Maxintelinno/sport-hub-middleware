package handler

import (
	"net/http"
	"time"

	"github.com/Maxintelinno/sport-hub-middleware/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// JWTClaims defines custom JWT claims
type JWTClaims struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	UserID   string `json:"userid"`
	jwt.RegisteredClaims
}

type AuthRequest struct {
	Username string `json:"username" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	UserID   string `json:"userid" validate:"required"`
}

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{cfg: cfg}
}

// GenerateToken handles token generation
func (h *AuthHandler) GenerateToken(c echo.Context) error {
	req := new(AuthRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "invalid request body"})
	}

	// Create claims
	claims := &JWTClaims{
		Username: req.Username,
		Phone:    req.Phone,
		UserID:   req.UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // 72 hours
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte(h.cfg.JWTSecret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token": t,
	})
}
