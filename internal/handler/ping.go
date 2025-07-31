package handler

import (
	"net/http"

	"github.com/akagiyuu/todo-backend/internal/util"
)

type PingResponse struct {
	Message string `json:"message"`
}

// Ping handler
// @Description test if server is ok
// @Produce  json
// @Router / [get]
func Ping(w http.ResponseWriter, r *http.Request) {
	util.WriteResponse(PingResponse{
		Message: "pong",
	}, w)
}
