package auth

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/server/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Description Payload for /auth/login: user's email and password.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginHandler godoc
// @Summary      Login
// @Description  Log in using email and password. Returns a raw JWT token string.
// @Tags         auth
// @Accept       json
// @Produce      plain
// @Param        payload  body       LoginRequest  true  "Login credentials"
// @Success      200      {string}   string                   "JWT access token"
// @Failure      400      {object}   middleware.ApiError      "Invalid login data or wrong credentials"
// @Failure      500      {object}   middleware.ApiError      "Internal failure during token generation"
// @Router       /auth/login [post]
func (r *AuthRoutes) LoginHandler(c *gin.Context) {
	ctx := context.Background()

	var request LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, middleware.ApiError{
			Message: "Invalid login data",
		})
		return
	}

	queries := database.New(r.Db)
	account, err := queries.GetAccountByEmail(ctx, request.Email)
	if err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Wrong email or password",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		c.Error(&middleware.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Wrong email or password",
		})
	}

	tokenString, err := r.Jwt.NewToken(account.ID.String())
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
