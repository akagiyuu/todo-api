package service

import (
	"fmt"
	"time"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/caarlos0/env/v11"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	cfg config.JwtConfig
}

func NewJwtService() *JwtService {
	cfg, _ := env.ParseAs[config.JwtConfig]()

	return &JwtService{cfg}
}

func (s *JwtService) NewToken(subject string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Duration(s.cfg.ExpiredIn) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(s.cfg.Secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtService) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return s.cfg.Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
