package config

import "os"

type Config struct {
	JWTSecret string
}

func LoadConfig() *Config {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-very-secret-key" // Default for development
	}
	return &Config{
		JWTSecret: secret,
	}
}
