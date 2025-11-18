package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func Connect(ctx context.Context) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("DATABASE_URL must be set")
	}
	c, err := pgx.Connect(ctx, url)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	Conn = c
}
