package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/service"
)

type Server struct {
	db  *pgxpool.Pool
	jwt service.JwtService
}

func NewServer() *http.Server {
	cfg, _ := env.ParseAs[config.ServerConfig]()

	NewServer := &Server{
		db:  database.NewPool(),
		jwt: service.NewJwtService(),
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
