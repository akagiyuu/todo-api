package server

import (
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/handler"
	"github.com/akagiyuu/todo-backend/internal/middleware"
	"github.com/gin-gonic/gin"

	_ "github.com/akagiyuu/todo-backend/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(middleware.Cors())

	r.GET("/", handler.Ping)
	r.POST("/auth/register", handler.Register(s.db))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
