package service

import (
	"time"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/caarlos0/env/v11"
	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	cfg config.JwtConfig
}

func NewJwtService() (s JwtService, err error) {
	cfg, err := env.ParseAs[config.JwtConfig]()

	s.cfg = cfg
	return
}

func (s *JwtService) NewToken(subject string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": subject,
		"exp": time.Now().Add(time.Duration(s.cfg.ExpiredIn) * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	token, err := claims.SignedString(s.cfg.Secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
