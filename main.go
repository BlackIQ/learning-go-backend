package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID      int    `json:"id"`
	Title   string `json:"title" binding:"required"`
	Details string `json:"details" binding:"required"`
}

var todos = []Todo{}

func nextID() (id int) {
	return len(todos) + 1
}

func Home(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "Home page"},
	)
}

func listTodos(c *gin.Context) {
	c.JSON(200, todos)
}

func getTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			409,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			c.JSON(200, todos[i])
			return
		} else {
			continue
		}
	}

	c.JSON(
		404,
		gin.H{
			"error": "Todo not found",
		},
	)
}

func createTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.ShouldBindBodyWithJSON(&newTodo); err != nil {
		c.JSON(
			400,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	newTodo.ID = nextID()

	todos = append(todos, newTodo)

	c.JSON(201, newTodo)
}

func updateTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			409,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	var updateTodo Todo

	if err := c.ShouldBindBodyWithJSON(&updateTodo); err != nil {
		c.JSON(
			400,
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

			c.JSON(200, todos[i])
			return
		} else {
			continue
		}
	}

	c.JSON(
		404,
		gin.H{
			"error": "Todo not found",
		},
	)
}

func deleteTodo(c *gin.Context) {
	todoId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(
			409,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, todo := range todos {
		if todoId == todo.ID {
			todos = append(todos[:i], todos[i+1:]...)

			c.JSON(204, gin.H{})
			return
		} else {
			continue
		}
	}

	c.JSON(
		404,
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
