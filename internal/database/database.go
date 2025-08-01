package database

import (
	"context"
	"log"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

func Init() *pgxpool.Pool {
	cfg, _ := env.ParseAs[config.DatabaseConfig]()

	db, err := pgxpool.New(context.Background(), cfg.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
