package auth

import (
	"context"

	"github.com/akagiyuu/todo-backend/internal/database"
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
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid login data",
		}
	}

	queries := database.New(rs.db)
	ctx := context.Background()
	account, err := queries.GetAccountByEmail(ctx, request.Email)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Wrong email or password",
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(request.Password)); err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Wrong email or password",
		}
	}

	tokenString, err := rs.jwtService.NewToken(account.ID.String())
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Failed to generate token",
		}
	}

	return tokenString, nil
}
