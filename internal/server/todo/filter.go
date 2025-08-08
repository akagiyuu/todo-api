package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/moznion/go-optional"
)

// @Description Optional query params for filtering todos
type FilterQuery struct {
	Query    optional.Option[string]            `form:"query" swaggertype:"string"`
	Priority optional.Option[database.Priority] `form:"priority" swaggertype:"string"`
	IsDone   optional.Option[bool]              `form:"isDone" swaggertype:"boolean"`
}

// @Summary      Filter todos
// @Description  Retrieves todos for the authenticated user, filtered by optional query parameters.
// @Tags         todo
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        query    query     string                     false  "Search text"
// @Param        priority query     string                     false  "Priority filter"
// @Param        isDone   query     boolean                    false  "Completion status filter"
// @Success      200      {array}   database.Todo              "List of filtered todos"
// @Failure      400      {object}  middleware.ApiError       "Invalid query parameters"
// @Router       /todo [get]
func (r *TodoRoutes) FilterHandler(c *gin.Context) {
	ctx := context.Background()

	var query FilterQuery
	if err := c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid query params",
		})
		return
	}

	accountID := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	todos, err := queries.FilterTodo(ctx, database.FilterTodoParams{
		AccountID: accountID,
		Query:     query.Query,
		Priority:  query.Priority,
		IsDone:    query.IsDone,
	})
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Failed to query todo with given params",
		})
		return
	}

	c.JSON(http.StatusOK, todos)
}
