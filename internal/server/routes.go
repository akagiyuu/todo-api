package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/handler"
	"github.com/akagiyuu/todo-backend/internal/middleware"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("My API", "1.0.0"))

	api.UseMiddleware(middleware.Cors)

	huma.Get(api, "/", handler.Ping)
	// Register GET /greeting/{name} handler.
	huma.Get(api, "/greeting/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	return mux
}
