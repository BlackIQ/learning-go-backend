package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gin-basics/internal/database"
	"gin-basics/internal/dto"
	"gin-basics/internal/models"
	"gin-basics/internal/utils"
)

func Signup(c *gin.Context) {
	var SignupData dto.SignupRequest

	if errSchema := c.ShouldBindJSON(&SignupData); errSchema != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid Body",
			},
		)
		return
	}

	db := database.GetDB()

	var user models.User

	resultFind := db.Where("email = ?", SignupData.Email).First(&user)

	if resultFind.Error == nil {
		c.AbortWithStatusJSON(
			http.StatusForbidden,
			gin.H{
				"message": "User already exists",
			},
		)
		return
	}

	if !errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": resultFind.Error.Error(),
			},
		)
		return
	}

	hashedPassword, errHash := utils.HashPassword(SignupData.Password)

	if errHash != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": errHash.Error(),
			},
		)
		return
	}

	newUser := models.User{
		ID:           uuid.New(),
		Name:         SignupData.Name,
		Email:        SignupData.Email,
		PaaswordHash: string(hashedPassword),
	}

	resultCreate := db.Create(&newUser)

	if resultCreate.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": resultCreate.Error.Error(),
			},
		)
		return
	}

	token, errToken := utils.GenerateJWT(newUser.ID.String())

	if errToken != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": errToken.Error(),
			},
		)
		return
	}

	response := dto.TokenResponse{
		AccessToken: token,
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}

func Signin(c *gin.Context) {
	var SigninData dto.SigninRequest

	if errSchema := c.ShouldBindJSON(&SigninData); errSchema != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid Body",
			},
		)
		return
	}

	db := database.GetDB()

	var user models.User

	resultFind := db.Where("email = ?", SigninData.Email).First(&user)

	if errors.Is(resultFind.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Wrong user name or password",
			},
		)
		return
	}

	if resultFind.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": resultFind.Error.Error(),
			},
		)
		return
	}

	passwordVerified := utils.VerifyPassword(SigninData.Password, user.PaaswordHash)

	if !passwordVerified {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Wrong username or password",
			},
		)
		return
	}

	token, errToken := utils.GenerateJWT(user.ID.String())

	if errToken != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": errToken.Error(),
			},
		)
		return
	}

	response := dto.TokenResponse{
		AccessToken: token,
	}

	c.JSON(
		http.StatusOK,
		response,
	)
}
