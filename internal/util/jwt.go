package util

import (
	"fmt"
	"time"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/caarlos0/env/v11"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

type JwtUtil struct {
	cfg config.JwtConfig
}

func NewJwtUtil() *JwtUtil {
	cfg, _ := env.ParseAs[config.JwtConfig]()

	return &JwtUtil{cfg}
}

func (u *JwtUtil) NewToken(subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Duration(u.cfg.ExpiredIn) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(u.cfg.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *JwtUtil) ParseToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return u.cfg.Secret, nil
	})
	if err != nil {
		return uuid.Nil, err
	}
	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	rawID, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	id, err := uuid.Parse(rawID)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
