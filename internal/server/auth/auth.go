package auth

import (
	"github.com/akagiyuu/todo-backend/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRoutes struct {
	Db  *pgxpool.Pool
	Jwt *service.JwtService
}

func (r AuthRoutes) RegisterRoutes(g *gin.Engine) {
	g.POST("/auth/register", r.RegisterHandler)
	g.POST("/auth/login", r.LoginHandler)
}
