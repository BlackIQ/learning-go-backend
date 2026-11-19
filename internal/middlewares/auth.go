package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-basics/internal/database"
	"gin-basics/internal/models"
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

		db := database.GetDB()

		var user models.User

		resultFind := db.Where("id = ?", userID).First(&user)

		if errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "Unauthorized",
				},
			)
			return
		}

		if resultFind.Error != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{
					"message": "Database error",
				},
			)
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
