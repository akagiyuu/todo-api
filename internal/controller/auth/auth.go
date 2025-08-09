package auth

import (
	"github.com/go-fuego/fuego"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/util"
)

type authResource struct {
	db         *pgxpool.Pool
	jwtService *util.JwtUtil
}

func RegisterRoutes(s *fuego.Server) {
	rs := authResource{
		db:         database.NewPool(),
		jwtService: util.NewJwtUtil(),
	}

	group := fuego.Group(s, "/auth")
	fuego.Post(group, "/register", rs.Register)
	fuego.Post(group, "/login", rs.Login)
}
