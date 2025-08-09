package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/middleware"
	"github.com/akagiyuu/todo-backend/internal/util"
	"github.com/go-fuego/fuego"
	"github.com/google/uuid"
	"github.com/moznion/go-optional"
)

type UpdateRequest struct {
	Title    optional.Option[string]            `json:"title,omitempty"`
	Content  optional.Option[string]            `json:"content,omitempty"`
	Priority optional.Option[database.Priority] `json:"priority,omitempty"`
}

func (rs TodoResource) Update(c fuego.ContextWithBody[UpdateRequest]) (any, error) {
	accountID := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	id, err := uuid.Parse(c.PathParam("id"))
	if err != nil {
		return nil, util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Required UUID v4",
		}
	}

	request, err := c.Body()
	if err != nil {
		return nil, util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Invalid update todo data",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	err = queries.UpdateTodo(ctx, database.UpdateTodoParams{
		ID:        id,
		AccountID: accountID,
		Title:     request.Title,
		Content:   request.Content,
		Priority:  request.Priority,
	})
	if err != nil {
		return nil, util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Failed to update todo",
		}
	}

	return nil, nil
}
