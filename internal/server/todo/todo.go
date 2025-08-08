package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
)

type TodoRoutes struct {
	db *pgxpool.Pool
}

func RegisterRoutes(g *gin.Engine) {
	r := TodoRoutes{
		db: database.NewPool(),
	}

	todo := g.Group("/", middleware.RequireAuthentication())
	{
		todo.POST("/todo", r.CreateHandler)
		todo.GET("/todo/:id", r.GetHandler)
		todo.PATCH("/todo/:id", r.UpdateHandler)
	}
}
