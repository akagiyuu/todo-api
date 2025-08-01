package handler

import "context"

type PingResponse struct {
	Message string `json:"message"`
}

func Ping(ctx context.Context, _ *struct{}) (*struct { Body: PingResponse }, error) {
	response := PingResponse{
		Message: "pong"
	}

	return response, nil
}
