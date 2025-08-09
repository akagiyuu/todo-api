package ping

import (
	"github.com/go-fuego/fuego"
)

func RegisterRoutes(s *fuego.Server) {
	fuego.Get(s, "/", ping)
}

type pingResponse struct {
	Message string `json:"message"`
}

func ping(c fuego.ContextNoBody) (pingResponse, error) {
	return pingResponse{
		Message: "pong",
	}, nil
}
