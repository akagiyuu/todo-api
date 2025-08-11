package auth

import (
	"context"

	"github.com/akagiyuu/todo-backend/internal/database"
	"github.com/go-fuego/fuego"
	"golang.org/x/crypto/bcrypt"
)

func (rs AuthResource) Register(c fuego.ContextWithBody[database.CreateAccountParams]) (string, error) {
	request, err := c.Body()
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid register data",
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Invalid register data",
		}
	}
	request.Password = string(hashedPassword)

	queries := database.New(rs.db)
	ctx := context.Background()
	id, err := queries.CreateAccount(ctx, request)
	if err != nil {
		return "", fuego.BadRequestError{
			Err:    err,
			Detail: "Account with given email already existed",
		}
	}

	tokenString, err := rs.jwtService.NewToken(id.String())
	if err != nil {
		return "", fuego.InternalServerError{
			Err:    err,
			Detail: "Failed to generate token",
		}
	}

	return tokenString, nil
}
