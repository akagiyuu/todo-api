package auth

import (
	"context"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/akagiyuu/todo-backend/internal/util"
	"github.com/go-fuego/fuego"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rs AuthResource) Login(c fuego.ContextWithBody[LoginRequest]) (string, error) {
	request, err := c.Body()
	if err != nil {
		return "", util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Invalid login data",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	account, err := queries.GetAccountByEmail(ctx, request.Email)
	if err != nil {
		return "", util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Wrong email or password",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		return "", util.ApiError{
			Inner:   err,
			Code:    http.StatusBadRequest,
			Message: "Wrong email or password",
		}
	}

	tokenString, err := rs.jwtService.NewToken(account.ID.String())
	if err != nil {
		return "", util.ApiError{
			Inner:   err,
			Code:    http.StatusInternalServerError,
			Message: "Failed to generate token",
		}
	}

	return tokenString, nil
}
