package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Description HTTP-level error response wrapper.
type ApiError struct {
	Inner   error          `json:"-"`
	Code    int            `json:"-"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func ErrorHandler() gin.HandlerFunc {
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
