package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-basics/internal/responses"
)

func RoleRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetHeader("X-ROLE")

		if role != "admin" {
			responses.Fail(
				c,
				http.StatusForbidden,
				"Forbidden",
				"You are not allowrd",
			)
			return
		}

		c.Next()
	}
}
