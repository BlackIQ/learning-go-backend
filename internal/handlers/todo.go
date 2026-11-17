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
)

func AllTodos(c *gin.Context) {
	db := database.GetDB()

	var todos []models.Todo

	result := db.Find(&todos)

	if result.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Database error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		todos,
	)
}

func OneTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid ID",
			},
		)
		return
	}

	db := database.GetDB()

	var todo models.Todo

	result := db.Where("id = ?", todoID).First(&todo)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "Todo not found",
			},
		)
		return
	}

	if result.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Database error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		todo,
	)
}

func CreateTodo(c *gin.Context) {
	var todoData dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&todoData); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid Body",
			},
		)
		return
	}

	db := database.GetDB()

	newTodo := models.Todo{
		ID:      uuid.New(),
		Title:   todoData.Title,
		Details: todoData.Details,
	}

	result := db.Create(&newTodo)

	if result.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Database error",
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		newTodo,
	)
}

func UpdateTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid ID",
			},
		)
		return
	}

	var updateTodo dto.UpdateTodoRequest

	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid Body",
			},
		)
		return
	}

	db := database.GetDB()

	var todo models.Todo

	result := db.Where("id = ?", todoID).First(&todo)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "Todo not found",
			},
		)
		return
	}

	if result.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Database error",
			},
		)
		return
	}

	updates := map[string]any{}

	if updateTodo.Title != nil {
		updates["Title"] = *updateTodo.Title
	}

	if updateTodo.Details != nil {
		updates["details"] = *updateTodo.Details
	}

	result = db.Model(&todo).Updates(updates)

	if result.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Database error",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		todo,
	)
}

func DeleteTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid ID",
			},
		)
		return
	}

	db := database.GetDB()

	var todo models.Todo

	result := db.Where("id = ?", todoID).First(&todo)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "Todo not found",
			},
		)
		return
	}

	if result.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Database error",
			},
		)
		return
	}

	result = db.Delete(&todo)

	if result.RowsAffected == 0 {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"message": "Todo not found",
			},
		)
		return
	}

	c.Status(
		http.StatusNoContent,
	)
}
