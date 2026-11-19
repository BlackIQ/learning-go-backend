package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-basics/internal/database"
	"gin-basics/internal/models"
)

func Me(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Unauthorized",
			},
		)
		return
	}

	db := database.GetDB()

	var user models.User

	resultFind := db.Where("id = ?", userID).First(&user)

	if resultFind.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{
				"message": resultFind.Error.Error(),
			},
		)
		return
	}

	if errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
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
