package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gin-basics/internal/dto"
	"gin-basics/internal/inits"
	"gin-basics/internal/models"
	"gin-basics/internal/responses"
)

func ListTodos(c *gin.Context) {
	var todos []models.Todo

	result := inits.DB.Find(&todos)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	responses.OK(c, http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_ID",
			"Invalid todo ID",
		)
		return
	}

	var todo models.Todo

	result := inits.DB.First(&todo, "id = ?", todoID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		responses.Fail(
			c,
			http.StatusNotFound,
			"TODO_NOT_FOUND",
			"Todo not found",
		)
		return
	}

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	responses.OK(c, http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {
	var request dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_SCHEMA",
			err.Error(),
		)
		return
	}

	todo := models.Todo{
		ID:      uuid.New(),
		Title:   request.Title,
		Details: request.Details,
	}

	result := inits.DB.Create(&todo)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	responses.OK(c, http.StatusCreated, todo)
}

func UpdateTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_ID",
			"Invalid todo ID",
		)
		return
	}

	var request dto.UpdateTodoRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_SCHEMA",
			err.Error(),
		)
		return
	}

	updates := map[string]any{}

	if request.Title != nil {
		updates["title"] = *request.Title
	}

	if request.Details != nil {
		updates["description"] = *request.Details
	}

	if len(updates) == 0 {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"EMPTY_UPDATE",
			"No fields provided for update",
		)
		return
	}

	result := inits.DB.Model(&models.Todo{}).
		Where("id = ?", todoID).
		Updates(updates)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	if result.RowsAffected == 0 {
		responses.Fail(
			c,
			http.StatusNotFound,
			"TODO_NOT_FOUND",
			"Todo not found",
		)
		return
	}

	var updatedTodo models.Todo

	if err := inits.DB.First(&updatedTodo, "id = ?", todoID).Error; err != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			err.Error(),
		)
		return
	}

	responses.OK(c, http.StatusOK, updatedTodo)
}

func DeleteTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_ID",
			"Invalid todo ID",
		)
		return
	}

	result := inits.DB.Delete(&models.Todo{}, "id = ?", todoID)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	if result.RowsAffected == 0 {
		responses.Fail(
			c,
			http.StatusNotFound,
			"TODO_NOT_FOUND",
			"Todo not found",
		)
		return
	}

	c.Status(http.StatusNoContent)
}
