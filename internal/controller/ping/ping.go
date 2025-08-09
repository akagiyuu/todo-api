package ping

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {
	fuego.Get(s, "/", Ping)
}

type PingResponse struct {
	Message string `json:"message"`
}

func Ping(c fuego.ContextNoBody) (PingResponse, error) {
	return PingResponse{
		Message: "pong",
	}, nil
}
