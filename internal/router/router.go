package router

import (
	"github.com/gin-gonic/gin"

	"gin-basics/internal/handlers"
	"gin-basics/internal/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.Request())

	router.GET("/", handlers.Home)

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
