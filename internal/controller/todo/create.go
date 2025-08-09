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

type CreateRequest struct {
	Title    string            `json:"title"`
	Content  string            `json:"content"`
	Priority database.Priority `json:"priority"`
}

func (rs TodoResource) Create(c fuego.ContextWithBody[CreateRequest]) (string, error) {
	accountID := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	request, err := c.Body()
	if err != nil {
		return "", util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Invalid new todo data",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	id, err := queries.CreateTodo(ctx, database.CreateTodoParams{
		AccountID: accountID,
		Title:     request.Title,
		Content:   request.Content,
		Priority:  request.Priority,
	})
	if err != nil {
		return "", util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Todo with given title already existed",
		}
	}

	return id.String(), nil
}
