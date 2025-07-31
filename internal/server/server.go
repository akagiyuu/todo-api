package server

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"

	"github.com/akagiyuu/todo-api/internal/config"
	"github.com/akagiyuu/todo-api/internal/database"
)

type Server struct {
	database *sql.DB
}

func NewServer() *http.Server {
	cfg, _ := env.ParseAs[config.ServerConfig]()

	NewServer := &Server{
		database: database.Init(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
