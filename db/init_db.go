package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func InitDB(ctx context.Context) (*pgx.Conn, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env files %w", err)
	}
	DATABASE_URL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(ctx, DATABASE_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return conn, nil

}
