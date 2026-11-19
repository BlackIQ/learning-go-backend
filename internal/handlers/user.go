package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Me(c *gin.Context) {
	user, exists := c.Get("user")

	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Unauthorized",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		user,
	)
}
