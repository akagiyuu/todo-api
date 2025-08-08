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

type UpdateRequest struct {
	Title    optional.Option[string]
	Content  optional.Option[string]
	Priority optional.Option[database.Priority]
}

// @Summary      Update a new todo
// @Description  Update a new todo
// @Tags         todo
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Router       /todo [PATCH]
func (r *TodoRoutes) UpdateHandler(c *gin.Context) {
	ctx := context.Background()

	var request UpdateRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid new todo data",
		})
		return
	}

	accountID := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	err := queries.UpdateTodo(ctx, database.UpdateTodoParams{
		AccountID: accountID,
		Title:     request.Title,
		Content:   request.Content,
		Priority:  request.Priority,
	})
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Todo with given title already existed",
		})
		return
	}

	c.Status(http.StatusOK)
}
