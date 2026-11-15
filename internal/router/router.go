package router

import (
	"github.com/gin-gonic/gin"

	"github.com/BlackIQ/gin-todo/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api")

	todoRouter := apiRouter.Group("/todos")

	todoRouter.GET("", handlers.ListTodos)
	todoRouter.GET("/:id", handlers.GetTodo)
	todoRouter.POST("", handlers.CreateTodo)
	todoRouter.PATCH("/:id", handlers.UpdateTodo)
	todoRouter.DELETE("/:id", handlers.DeleteTodo)

	noteHandler := apiRouter.Group("/notes")

	noteHandler.GET("", handlers.ListNotes)
	noteHandler.GET("/:id", handlers.GetNote)
	noteHandler.POST("", handlers.CreateNote)
	noteHandler.PATCH("/:id", handlers.UpdateNote)
	noteHandler.DELETE("/:id", handlers.DeleteNote)

	return router
}
