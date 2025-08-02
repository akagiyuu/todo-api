package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/service/jwt"
)

type AuthRoutes struct {
	db  *pgxpool.Pool
	jwtService *jwt.JwtService
}

func RegisterRoutes(g *gin.Engine) {
	r := AuthRoutes{
		db:  database.NewPool(),
		jwtService: jwt.New(),
	}

	g.POST("/auth/register", r.RegisterHandler)
	g.POST("/auth/login", r.LoginHandler)
}
