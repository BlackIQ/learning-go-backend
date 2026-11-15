package router

import (
	"github.com/gin-gonic/gin"

	"github.com/BlackIQ/gin-todo/internal/handlers"
	"github.com/BlackIQ/gin-todo/internal/middlewares"
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

	noteRouter := apiRouter.Group("/notes")

	noteRouter.Use(middlewares.RoleRequired())

	noteRouter.GET("", handlers.ListNotes)
	noteRouter.GET("/:id", handlers.GetNote)
	noteRouter.POST("", handlers.CreateNote)
	noteRouter.PATCH("/:id", handlers.UpdateNote)
	noteRouter.DELETE("/:id", handlers.DeleteNote)

	return router
}
