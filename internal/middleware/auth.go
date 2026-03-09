package middleware

import (
	"github.com/Maxintelinno/sport-hub-middleware/internal/config"
	"github.com/Maxintelinno/sport-hub-middleware/internal/handler"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware returns a JWT middleware with the configured secret
func AuthMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JWTSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handler.JWTClaims)
		},
	})
}

// Optional: Custom claims extraction or additional validation can be added here
