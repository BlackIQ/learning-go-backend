package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Todo struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title" binding:"required"`
	Details string    `json:"details" binding:"required"`
}

var todos = []Todo{}

func Home(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Home page",
		},
	)
}

func listTodos(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		todos,
	)
}

func getTodo(c *gin.Context) {
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

func createTodo(c *gin.Context) {
	var newTodo Todo

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

	todos = append(todos, newTodo)

	c.JSON(
		http.StatusCreated,
		newTodo,
	)
}

func updateTodo(c *gin.Context) {
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

	var updateTodo Todo

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
			todos[i] = updateTodo

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

func deleteTodo(c *gin.Context) {
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

func main() {
	router := gin.Default()

	router.GET("/", Home)

	router.GET("/todos", listTodos)         // List Todos
	router.GET("/todos/:id", getTodo)       // One Todo
	router.POST("/todos", createTodo)       // Create Todo
	router.PATCH("todos/:id", updateTodo)   // Update Todo
	router.DELETE("/todos/:id", deleteTodo) // Delete Todo

	err := router.Run("0.0.0.0:8080")

	if err != nil {
		fmt.Println("\n-> Error running the app.")
		fmt.Print(err)
	}
}
