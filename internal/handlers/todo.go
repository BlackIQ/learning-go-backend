package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-basics/internal/dto"
)

func AllTodos(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		[]dto.TodoResponse{},
	)
}

func OneTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid ID",
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"id": todoID,
		},
	)
}

func CreateTodo(c *gin.Context) {
	var todoData dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&todoData); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid Body",
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": todoData,
		},
	)
}

func UpdateTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid ID",
				"error":   err.Error(),
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
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"id":   todoID,
			"data": updateTodo,
		},
	)
}

func DeleteTodo(c *gin.Context) {
	todoID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Invalid ID",
				"error":   err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"id": todoID,
		},
	)
}
