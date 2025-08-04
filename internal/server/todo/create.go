package todo

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
)

// @Summary      Create a new todo
// @Description  Create a new todo
// @Tags         todo
// @Accept       json
// @Param        payload  body       database.CreateTodoParams  true "Todo data"
// @Success      200      {string}   string
// @Router       /todo [post]
func (r *TodoRoutes) CreateHandler(c *gin.Context) {
	ctx := context.Background()

	var request database.CreateTodoParams
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid new todo data",
		})
		return
	}

	queries := database.New(r.db)
	id, err := queries.CreateTodo(ctx, request)
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
