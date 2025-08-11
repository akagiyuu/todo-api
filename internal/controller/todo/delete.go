package todo

import (
	"context"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/middleware"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
)

func (rs TodoResource) Delete(c fuego.ContextNoBody) (any, error) {
	accountID := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "Required UUID v4",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	err = queries.DeleteTodo(ctx, database.DeleteTodoParams{
		ID:        id,
		AccountID: accountID,
	})
	if err != nil {
		return nil, fuego.BadRequestError{
			Err:    err,
			Detail: "No todo with given id",
		}
	}

	return nil, nil
}
