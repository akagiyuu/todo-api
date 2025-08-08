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

// @Description Payload for PATCH /todo/{id} 
type UpdateRequest struct {
	Title    optional.Option[string]            `json:"title,omitempty" swaggertype:"string"`
	Content  optional.Option[string]            `json:"content,omitempty" swaggertype:"string"`
	Priority optional.Option[database.Priority] `json:"priority,omitempty" swaggertype:"string"`
}

// @Summary      Update a todo
// @Description  Partially updates an existing todo; only supplied fields are changed.
// @Tags         todo
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      string          true  "Todo UUID v4"
// @Param        payload  body      UpdateRequest   true  "Fields to update (optional)"
// @Success      200      {string}  string          "OK — no content returned"
// @Failure      400      {object}  middleware.ApiError  "Invalid payload or update error"
// @Router       /todo/{id} [patch]
func (r *TodoRoutes) UpdateHandler(c *gin.Context) {
	ctx := context.Background()

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Required UUID v4",
		})
		return
	}

	var request UpdateRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid update todo data",
		})
		return
	}

	accountID := c.MustGet(middleware.AuthorizationTokenKey).(uuid.UUID)

	queries := database.New(r.db)
	err = queries.UpdateTodo(ctx, database.UpdateTodoParams{
		ID:        id,
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
