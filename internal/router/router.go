package router

import (
	"github.com/gin-gonic/gin"

	"gin-basics/internal/handlers"
	"gin-basics/internal/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.Request())

	apiRouter := router.Group("/api")

	todoRouter := apiRouter.Group("/todos")

	todoRouter.GET("", handlers.AllTodos)
	todoRouter.GET("/:id", handlers.OneTodo)
	todoRouter.POST("", handlers.CreateTodo)
	todoRouter.PATCH("/:id", handlers.UpdateTodo)
	todoRouter.DELETE("/:id", handlers.DeleteTodo)

	return router
}
