package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-ROLE")

		if role != "admin" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Unauthorized",
				},
			)
			return
		}

		c.Next()
	}
}
