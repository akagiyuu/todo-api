package server

import (
	"net/http"
	"github.com/akagiyuu/todo-api/internal/handler"
	"github.com/akagiyuu/todo-api/internal/middleware"

	"github.com/go-pkgz/routegroup"
	_ "github.com/akagiyuu/todo-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (s *Server) RegisterRoutes() http.Handler {
	router := routegroup.New(http.NewServeMux())

	router.Use(middleware.Cors)

	router.HandleFunc("GET /", handler.Ping)

	router.Handle("/swagger/", httpSwagger.WrapHandler)

	return router
}
