package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/BlackIQ/gin-todo/internal/models"
	"github.com/BlackIQ/gin-todo/internal/responses"
)

var todos = []models.Todo{}

func ListTodos(c *gin.Context) {
	responses.OK(
		c,
		http.StatusOK,
		todos,
	)
}

func GetTodo(c *gin.Context) {
	todoId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"BadRequest",
			"Invalid ID provided",
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			responses.OK(
				c,
				http.StatusOK,
				todos[i],
			)
			return
		}
	}

	responses.Fail(
		c,
		http.StatusNotFound,
		"NotFound",
		"Todo is not found",
	)
}

func CreateTodo(c *gin.Context) {
	var newTodo *models.Todo

	if err := c.ShouldBindBodyWithJSON(&newTodo); err != nil {
		responses.Fail(
			c,
			http.StatusConflict,
			"Conflict",
			"Invalid schema",
		)
		return
	}

	newTodo.ID = uuid.New()

	todos = append(todos, *newTodo)

	responses.OK(
		c,
		http.StatusCreated,
		newTodo,
	)
}

func UpdateTodo(c *gin.Context) {
	todoId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"BadRequest",
			"Invalid ID provided",
		)
		return
	}

	var updateTodo *models.Todo

	if err := c.ShouldBindBodyWithJSON(&updateTodo); err != nil {
		responses.Fail(
			c,
			http.StatusConflict,
			"Conflict",
			"Invalid schema",
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			updateTodo.ID = todoId
			todos[i] = *updateTodo

			responses.OK(
				c,
				http.StatusOK,
				notes[i],
			)
			return
		}
	}

	responses.Fail(
		c,
		http.StatusNotFound,
		"NotFound",
		"Todo is not found",
	)
}

func DeleteTodo(c *gin.Context) {
	todoId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"BadRequest",
			"Invalid ID provided",
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)

			responses.OK(
				c,
				http.StatusNoContent,
				nil,
			)
			return
		}
	}

	responses.Fail(
		c,
		http.StatusNotFound,
		"NotFound",
		"Todo is not found",
	)
}
