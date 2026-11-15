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

type Note struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
}

var todos = []Todo{}
var notes = []Note{}

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

func listNotes(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		notes,
	)
}

func getNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			c.JSON(
				http.StatusOK,
				notes[i],
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Note not found",
		},
	)
}

func createNote(c *gin.Context) {
	var newNote Note

	if err := c.ShouldBindBodyWithJSON(&newNote); err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	newNote.ID = uuid.New()

	notes = append(notes, newNote)

	c.JSON(
		http.StatusCreated,
		newNote,
	)
}

func updateNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	var updateNote Note

	if err := c.ShouldBindBodyWithJSON(&updateNote); err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			updateNote.ID = noteId
			notes[i] = updateNote

			c.JSON(
				http.StatusOK,
				notes[i],
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Note not found",
		},
	)
}

func deleteNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			notes = append(notes[:i], notes[i+1:]...)

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

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", Home)

	metaRouter := router.Group("/meta")
	{
		metaRouter.GET("health", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})
	}

	apiRouter := router.Group("/api")
	{
		todoRouter := apiRouter.Group("/todos")

		todoRouter.GET("", listTodos)
		todoRouter.GET("/:id", getTodo)
		todoRouter.POST("", createTodo)
		todoRouter.PATCH("/:id", updateTodo)
		todoRouter.DELETE("/:id", deleteTodo)

		noteRouter := apiRouter.Group("/notes")

		noteRouter.GET("", listNotes)
		noteRouter.GET("/:id", getNote)
		noteRouter.POST("", createNote)
		noteRouter.PATCH("/:id", updateNote)
		noteRouter.DELETE("/:id", deleteNote)
	}

	return router
}

func main() {
	router := setupRouter()

	err := router.Run("0.0.0.0:8080")

	if err != nil {
		fmt.Println("\n-> Error running the app.")
		fmt.Print(err)
	}
}
