package server

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
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
func (s *Server) RegisterHandler(c *gin.Context) {
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
		c.Error(&ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Invalid register data",
		})
		return
	}
	request.Password = string(hashedPassword)

	queries := database.New(s.db)
	id, err := queries.CreateAccount(ctx, request)
	if err != nil {
		c.Error(&ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Account with given email already existed",
		})
		return
	}

	tokenString, err := s.jwt.NewToken(id.String())
	if err != nil {
		c.Error(&ApiError{
			Inner:   err,
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
		return
	}

	c.String(http.StatusOK, tokenString)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary      Login
// @Description  Login with email and password
// @Tags         auth
// @Accept       json
// @Param        payload  body       LoginRequest  true
// @Success      200      {string}   string
// @Router       /auth/login [post]
func (s *Server) LoginHandler(c *gin.Context) {
	ctx := context.Background()

	var request LoginRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ApiError{
			Message: "Invalid login data",
		})
		return
	}

	queries := database.New(s.db)
	account, err := queries.GetAccountByEmail(ctx, request.Email)
	if err != nil {
		c.Error(&ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Wrong email or password",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		c.Error(&ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Wrong email or password",
		})
	}

	tokenString, err := s.jwt.NewToken(account.ID.String())
	if err != nil {
		c.Error(&ApiError{
			Inner:   err,
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		})
		return
	}

	c.String(http.StatusOK, tokenString)
}
