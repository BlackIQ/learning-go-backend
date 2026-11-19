package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"gin-basics/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Missing authorization header",
				},
			)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)

		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Invalid authorization header",
				},
			)
			return
		}

		userID, err := utils.VerifyJWT(parts[1])

		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Invalid or expired token",
				},
			)
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
