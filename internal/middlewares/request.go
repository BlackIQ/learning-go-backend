package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")

		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set("RequestID", requestID)

		c.Next()
	}
}
