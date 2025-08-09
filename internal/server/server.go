package server

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env/v11"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	_ "github.com/joho/godotenv/autoload"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/akagiyuu/todo-backend/internal/controller/auth"
	"github.com/akagiyuu/todo-backend/internal/controller/ping"
	"github.com/akagiyuu/todo-backend/internal/middleware"
)

func openApiHandler(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.BaseLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}

func NewServer() *fuego.Server {
	cfg, _ := env.ParseAs[config.ServerConfig]()

	s := fuego.NewServer(
		fuego.WithAddr(fmt.Sprintf(":%d", cfg.Port)),
		fuego.WithGlobalMiddlewares(middleware.Cors),
		fuego.WithEngineOptions(
			fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
				UIHandler:            openApiHandler,
				DisableDefaultServer: true,
				DisableMessages:      true,
				Info: &openapi3.Info{
					Title:       "Todo API",
					Description: "Todo API",
				},
			}),
		),
	)
	s.Engine.OutputOpenAPISpec().AddServer(&openapi3.Server{
		URL: cfg.Url,
	})

	ping.RegisterRoutes(s)
	auth.RegisterRoutes(s)

	return s
}
