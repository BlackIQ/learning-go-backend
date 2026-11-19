package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-basics/internal/models"
)

func Me(c *gin.Context) {
	userValue, exists := c.Get("user")

	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Unauthorized",
			},
		)
		return
	}

	user, ok := userValue.(*models.User)

	if !ok {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Invalid user context",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		user,
	)
}
