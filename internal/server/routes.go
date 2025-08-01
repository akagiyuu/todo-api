package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "github.com/akagiyuu/todo-backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(s.CorsMiddleware())
	r.Use(s.ErrorMiddleware())

	r.GET("/", s.PingHandler)
	r.POST("/auth/register", s.RegisterHandler)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
