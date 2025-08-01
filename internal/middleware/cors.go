package middleware

import (
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/config"
	"github.com/danielgtaylor/huma/v2"

	"github.com/caarlos0/env/v11"
)

func Cors(ctx huma.Context, next func(huma.Context)) {
	cfg, _ := env.ParseAs[config.CorsConfig]()

	ctx.SetHeader("Access-Control-Allow-Origin", cfg.AllowOrigin)
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
	ctx.SetHeader("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
	ctx.SetHeader("Access-Control-Allow-Credentials", "true")

	if ctx.Method() == http.MethodOptions {
		ctx.SetStatus(http.StatusNoContent)
		return
	}

	next(ctx)
}
