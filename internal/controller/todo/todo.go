package todo

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	"github.com/gorilla/schema"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/middleware"
)

type TodoResource struct {
	db      *pgxpool.Pool
	decoder *schema.Decoder
}

func RegisterRoutes(s *fuego.Server) {
	rs := TodoResource{
		db:      database.NewPool(),
		decoder: schema.NewDecoder(),
	}

	group := fuego.Group(s, "/todo",
		option.Middleware(middleware.RequireAuthentication),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)

	fuego.Post(group, "/", rs.Create)
	fuego.Get(group, "/", rs.Filter)
	fuego.Get(group, "/{id}", rs.Get)
	fuego.Patch(group, "/{id}", rs.Update)
	fuego.Delete(group, "/{id}", rs.Delete)
}
