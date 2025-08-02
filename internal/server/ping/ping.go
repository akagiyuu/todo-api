package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingRoutes struct{}

func (r PingRoutes) RegisterRoutes(g *gin.Engine) {
	g.GET("/", r.PingHandler)
}

type PingResponse struct {
	Message string `json:"message"`
}

// @Description  Check if server is running
// @Produce      json
// @Router       / [get]
func (r *PingRoutes) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, PingResponse{
		Message: "pong",
	})
}
