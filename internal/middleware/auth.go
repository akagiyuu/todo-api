package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/akagiyuu/todo-backend/internal/util"
	"github.com/go-fuego/fuego"
)

const (
	authorization         string = "Authorization"
	bearer                string = "Bearer "
	AuthorizationTokenKey string = "token"
)

func RequireAuthentication(next http.Handler) http.Handler {
	jwtUtil := util.NewJwtUtil()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(authorization)
		if authHeader == "" {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Detail: "Missing authorization header",
			})
			return
		}

		tokenString, isBearer := strings.CutPrefix(authHeader, bearer)
		if !isBearer {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Detail: "Missing authorization token",
			})
			return
		}

		token, err := jwtUtil.ParseToken(tokenString)
		if err != nil {
			fuego.SendJSONError(w, nil, fuego.UnauthorizedError{
				Err:    err,
				Detail: "Invalid authorization token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), AuthorizationTokenKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
