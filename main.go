package main

import (
	"net/http"

	"github.com/Maxintelinno/sport-hub-middleware/internal/config"
	"github.com/Maxintelinno/sport-hub-middleware/internal/handler"
	"github.com/Maxintelinno/sport-hub-middleware/internal/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo
	e := echo.New()

	// Load configuration
	cfg := config.LoadConfig()

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	// Handlers
	authHandler := handler.NewAuthHandler(cfg)

	// Routes
	// Public route
	e.GET("/public", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is a public endpoint")
	})

	e.POST("/api/v1/auth/token", authHandler.GenerateToken)

	// Protected group
	api := e.Group("/api")
	api.Use(middleware.AuthMiddleware(cfg))

	// Protected routes
	api.GET("/profile", func(c echo.Context) error {
		// Example of accessing JWT data if needed
		// user := c.Get("user").(*jwt.Token)
		// claims := user.Claims.(jwt.MapClaims)
		// name := claims["name"].(string)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Welcome to your profile!",
		})
	})

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
