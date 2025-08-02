package middleware

import (
	"net/http"
	"strings"

	"github.com/akagiyuu/todo-backend/internal/service/jwt"
	"github.com/gin-gonic/gin"
)

const (
	authorization         string = "Authorization"
	bearer                string = "Bearer: "
	AuthorizationTokenKey string = "token"
)

func RequireAuthentication() gin.HandlerFunc {
	jwtService := jwt.New()

	return func(c *gin.Context) {
		authHeader := c.GetHeader(authorization)
		if authHeader == "" {
			c.Error(&ApiError{
				Code:    http.StatusUnauthorized,
				Message: "Missing authorization header",
			})
			return
		}

		tokenString, isBearer := strings.CutPrefix(authHeader, bearer)
		if !isBearer {
			c.Error(&ApiError{
				Code:    http.StatusUnauthorized,
				Message: "Missing authorization token",
			})
			return
		}

		token, err := jwtService.ParseToken(tokenString)
		if err != nil {
			c.Error(&ApiError{
				Inner:   err,
				Code:    http.StatusForbidden,
				Message: "Invalid authorization token",
			})
			return
		}

		c.Set(AuthorizationTokenKey, token)
	}
}
