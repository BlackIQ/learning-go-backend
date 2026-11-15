package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/BlackIQ/gin-todo/internal/models"
)

var todos = []models.Todo{}

func ListTodos(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		todos,
	)
}

func GetTodo(c *gin.Context) {
	todoId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			c.JSON(
				http.StatusOK,
				todos[i],
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Todo not found",
		},
	)
}

func CreateTodo(c *gin.Context) {
	var newTodo *models.Todo

	if err := c.ShouldBindBodyWithJSON(&newTodo); err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	newTodo.ID = uuid.New()

	todos = append(todos, *newTodo)

	c.JSON(
		http.StatusCreated,
		newTodo,
	)
}

func UpdateTodo(c *gin.Context) {
	todoId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	var updateTodo *models.Todo

	if err := c.ShouldBindBodyWithJSON(&updateTodo); err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			updateTodo.ID = todoId
			todos[i] = *updateTodo

			c.JSON(
				http.StatusOK,
				todos[i],
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Todo not found",
		},
	)
}

func DeleteTodo(c *gin.Context) {
	todoId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)

			c.Status(
				http.StatusNoContent,
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Todo not found",
		},
	)
}
