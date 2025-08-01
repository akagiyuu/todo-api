package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Inner   error  `json:"-"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func (s *Server) ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err

			switch e := err.(type) {
			case *ApiError:
				c.JSON(e.Code, e)
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "Internal server error",
				})
			}
		}
	}
}
