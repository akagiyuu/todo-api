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
