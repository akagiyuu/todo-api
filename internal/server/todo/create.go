package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateRequest defines the expected payload for creating a new todo.
// @Description Payload for POST /todo
type CreateRequest struct {
	Title    string            `json:"title" example:"Buy groceries"`
	Content  string            `json:"content" example:"Milk, bread, and eggs"`
	Priority database.Priority `json:"priority" example:"high"`
}

// @Summary      Create a new todo
// @Description  Creates a new todo item for the authenticated user.
// @Tags         todo
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        payload  body       CreateRequest         true  "Todo data"
// @Success      200      {string}   string                "ID of the created todo"
// @Failure      400      {object}   middleware.ApiError  "Invalid request or duplicate title"
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

	c.String(http.StatusOK, id.String())
}
