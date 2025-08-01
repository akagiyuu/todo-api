package server

import (
	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/caarlos0/env/v11"
)

func (s *Server) CorsMiddleware() gin.HandlerFunc {
	cfg, _ := env.ParseAs[config.CorsConfig]()

	return cors.New(cors.Config{
		AllowOrigins:     []string{cfg.AllowOrigin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	})
}
