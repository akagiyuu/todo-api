package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary      Get todo data by ID
// @Description  Retrieves a specific todo by its UUID and validates ownership.
// @Tags         todo
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path     string  true  "Todo UUID (v4)"
// @Success      200  {object} database.Todo  "Todo data returned"
// @Failure      400  {object} middleware.ApiError  "Bad request or not found"
// @Router       /todo/{id} [get]
func (r *TodoRoutes) GetHandler(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Required UUID v4",
		})
		return
	}

	accountID := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	todo, err := queries.GetTodo(ctx, database.GetTodoParams{
		ID:        id,
		AccountID: accountID,
	})
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "No todo with given id",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}
