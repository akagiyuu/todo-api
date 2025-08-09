package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/middleware"
	"github.com/akagiyuu/todo-backend/internal/util"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
)

func (rs TodoResource) Get(c fuego.ContextNoBody) (database.GetTodoRow, error) {
	accountID := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return database.GetTodoRow{}, util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Required UUID v4",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	todo, err := queries.GetTodo(ctx, database.GetTodoParams{
		ID:        id,
		AccountID: accountID,
	})
	if err != nil {
		return database.GetTodoRow{}, util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "No todo with given id",
		}
	}

	return todo, nil
}
