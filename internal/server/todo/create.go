package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Description Payload for POST /todo
type CreateRequest struct {
	Title    string
	Content  string
	Priority database.Priority
}

// @Summary      Create a new todo
// @Description  Create a new todo
// @Tags         todo
// @Security     BearerAuth
// @Accept       json
// @Param        payload  body       CreateRequest  true "Todo data"
// @Success      200      {string}   string
// @Router       /todo [post]
func (r *TodoRoutes) CreateHandler(c *gin.Context) {
	ctx := context.Background()

	var request CreateRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid new todo data",
		})
		return
	}

	accountID := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	id, err := queries.CreateTodo(ctx, database.CreateTodoParams{
		AccountID: accountID,
		Title: request.Title,
		Content: request.Content,
		Priority: request.Priority,
	})
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Todo with given title already existed",
		})
		return
	}

	c.String(http.StatusOK, id.String())
}
