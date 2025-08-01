package handler

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Create a new account
// @Description  Creates a new user account with email and password.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        payload  body       database.CreateAccountParams  true  "Registration payload"
// @Success      200      {string}   pgtype.UUID                         "Created account ID"
// @Failure      400      {object}   ApiError                            "Bad request, validation or creation failure"
// @Router       /auth/register [post]
func Register(db *pgxpool.Pool) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := context.Background()

		var request database.CreateAccountParams
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, ApiError{
				Message: "Invalid register data",
			})
			return
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, ApiError{
				Message: "Invalid register data",
			})
			return
		}
		request.Password = string(hashedPassword)

		queries := database.New(db)
		id, err := queries.CreateAccount(ctx, request)
		if err != nil {
			c.JSON(http.StatusBadRequest, ApiError{
				Message: "Failed to create account",
			})
			return
		}

		c.JSON(http.StatusOK, id)
	}
}
