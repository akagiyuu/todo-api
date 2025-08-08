package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary      Delete a todo
// @Description  Deletes a specific todo by its UUID.
// @Tags         todo
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id   path      string               true  "Todo UUID v4"
// @Success      200  {string}  string               "OK — todo deleted"
// @Failure      400  {object}  middleware.ApiError "Invalid ID or todo not found"
// @Router       /todo/{id} [delete]
func (r *TodoRoutes) DeleteHandler(c *gin.Context) {
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
	err = queries.DeleteTodo(ctx, database.DeleteTodoParams{
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

	c.Status(http.StatusOK)
}
