package database

import (
	"context"
	"log"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

var instance *pgxpool.Pool

func NewPool() *pgxpool.Pool {
	if instance != nil {
		return instance
	}

	cfg, _ := env.ParseAs[config.DatabaseConfig]()

	instance, err := pgxpool.New(context.Background(), cfg.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return instance
}
