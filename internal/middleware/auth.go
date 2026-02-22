package middleware

import (
	"github.com/Maxintelinno/sport-hub-middleware/internal/config"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware returns a JWT middleware with the configured secret
func AuthMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(cfg.JWTSecret),
		// Optional: you can define a custom context key for the user
		// ContextKey: "user",
	})
}

// Optional: Custom claims extraction or additional validation can be added here
