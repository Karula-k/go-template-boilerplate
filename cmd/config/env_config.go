package config

import (
	"errors"
	"os"

	"github.com/go-template-boilerplate/cmd/models"
)

func EnvConfig() (models.EnvModel, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")
	if databaseUrl == "" {
		return models.EnvModel{}, errors.New("DATABASE_URL is not set")
	}
	if jwtSecret == "" {
		return models.EnvModel{}, errors.New("JWT_SECRET is not set")
	}
	if port == "" {
		return models.EnvModel{}, errors.New("PORT is not set")
	}

	model := models.EnvModel{
		DatabaseUrl: databaseUrl,
		JwtSecret:   jwtSecret,
		Port:        port,
	}

	return model, nil
}
