package server

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/go-fuego/fuego"
	_ "github.com/joho/godotenv/autoload"

	"github.com/akagiyuu/todo-backend/internal/config"
)

func NewServer() *fuego.Server {
	cfg, _ := env.ParseAs[config.ServerConfig]()
	options := []fuego.ServerOption{
		fuego.WithAddr(fmt.Sprintf(":%d", cfg.Port)),
		fuego.WithGlobalMiddlewares()
	}

	app := fuego.NewServer(options...)
	app.OpenAPI.Description().Info.Title = "Todo API"

	return app
}
