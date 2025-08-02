package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/akagiyuu/todo-backend/docs"
	"github.com/akagiyuu/todo-backend/internal/server/auth"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/akagiyuu/todo-backend/internal/server/ping"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	g := gin.Default()

	g.Use(middleware.Cors())
	g.Use(middleware.ErrorHandler())

	ping.PingRoutes{}.RegisterRoutes(g)
	auth.AuthRoutes{
		Db:  s.db,
		Jwt: s.jwt,
	}.RegisterRoutes(g)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return g
}
