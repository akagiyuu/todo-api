package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message string `json:"message"`
}

// @Description  Check if server is running
// @Produce      json
// @Router       / [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{
		Message: "pong",
	})
}
