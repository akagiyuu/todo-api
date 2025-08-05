package database

import (
	"context"
	"log"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxUUID "github.com/vgarvardt/pgx-google-uuid/v5"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

var instance *pgxpool.Pool

func NewPool() *pgxpool.Pool {
	if instance != nil {
		return instance
	}

	cfg, _ := env.ParseAs[config.DatabaseConfig]()
	pgxConfig, _ := pgxpool.ParseConfig(cfg.GetConnectionString())
	pgxConfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxUUID.Register(conn.TypeMap())
		return nil
	}

	instance, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		log.Fatal(err)
	}

	return instance
}
