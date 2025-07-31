package database

import (
	"database/sql"
	"github.com/akagiyuu/todo-backend/internal/config"
	"log"

	"github.com/caarlos0/env/v11"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

func Init() *sql.DB {
	cfg, _ := env.ParseAs[config.DatabaseConfig]()

	db, err := sql.Open("pgx", cfg.GetConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	return db
}
