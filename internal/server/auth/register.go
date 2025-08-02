package auth

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary      Create a new account
// @Description  Creates a new user account with email and password
// @Tags         auth
// @Accept       json
// @Param        payload  body       database.CreateAccountParams  true  "Registration payload"
// @Success      200      {string}   string                              "Created account ID"
// @Router       /auth/register [post]
func (r *AuthRoutes) RegisterHandler(c *gin.Context) {
	ctx := context.Background()

	var request database.CreateAccountParams
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid register data",
		})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Invalid register data",
		})
		return
	}
	request.Password = string(hashedPassword)

	queries := database.New(r.Db)
	id, err := queries.CreateAccount(ctx, request)
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Account with given email already existed",
		})
		return
	}

	tokenString, err := r.Jwt.NewToken(id.String())
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
		return
	}

	c.String(http.StatusOK, tokenString)
}
