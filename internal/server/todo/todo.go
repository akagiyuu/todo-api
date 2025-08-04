package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/akagiyuu/todo-backend/internal/database"
)

type TodoRoutes struct {
	db *pgxpool.Pool
}

func RegisterRoutes(g *gin.Engine) {
	r := TodoRoutes{
		db: database.NewPool(),
	}

	g.POST("/todo", r.CreateHandler)
}
