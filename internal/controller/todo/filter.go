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

type FilterQuery struct {
	Query    optional.Option[string]            `form:"query"`
	Priority optional.Option[database.Priority] `form:"priority"`
	IsDone   optional.Option[bool]              `form:"isDone"`
}

func (rs TodoResource) Filter(c fuego.ContextNoBody) ([]database.FilterTodoRow, error) {
	accountID := c.Value(middleware.AuthorizationTokenKey).(uuid.UUID)

	var filter FilterQuery
	err := rs.decoder.Decode(&filter, c.QueryParams())
	if err != nil {
		return nil, util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Invalid query params",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	todos, err := queries.FilterTodo(ctx, database.FilterTodoParams{
		AccountID: accountID,
		Query:     filter.Query,
		Priority:  filter.Priority,
		IsDone:    filter.IsDone,
	})
	if err != nil {
		return nil, util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Failed to query todo with given params",
		}
	}

	return todos, nil
}
