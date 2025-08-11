package todo

import (
	"context"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/middleware"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
)

func (rs TodoResource) Get(c fuego.ContextNoBody) (database.GetTodoRow, error) {
	accountID := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return database.GetTodoRow{}, fuego.BadRequestError{
			Err:    err,
			Detail: "Required UUID v4",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	todo, err := queries.GetTodo(ctx, database.GetTodoParams{
		ID:        id,
		AccountID: accountID,
	})
	if err != nil {
		return database.GetTodoRow{}, fuego.BadRequestError{
			Err:    err,
			Detail: "No todo with given id",
		}
	}

	return todo, nil
}
