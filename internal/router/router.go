package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"gin-basics/internal/handlers"
	"gin-basics/internal/middlewares"
	"gin-basics/internal/repositories"
	"gin-basics/internal/services"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	todoRepository := repositories.NewTodoRepository(db)
	todoServices := services.NewTodoService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoServices)

	noteRepository := repositories.NewNoteRepository(db)
	noteServices := services.NewNoteService(noteRepository)
	noteHandler := handlers.NewNoteHandler(noteServices)

	router := gin.Default()

	router.Use(middlewares.Request())

	router.GET("/", handlers.Home)

	apiRouter := router.Group("/api")

	todoRouter := apiRouter.Group("/todos")

	todoRouter.GET("", todoHandler.ListTodos)
	// todoRouter.GET("/:id", todoHandler.GetTodo)
	// todoRouter.POST("", todoHandler.CreateTodo)
	// todoRouter.PATCH("/:id", todoHandler.UpdateTodo)
	// todoRouter.DELETE("/:id", todoHandler.DeleteTodo)

	noteRouter := apiRouter.Group("/notes")

	noteRouter.Use(middlewares.RoleRequired())

	noteRouter.GET("", noteHandler.ListNotes)
	// noteRouter.GET("/:id", noteHandler.GetNote)
	// noteRouter.POST("", noteHandler.CreateNote)
	// noteRouter.PATCH("/:id", noteHandler.UpdateNote)
	// noteRouter.DELETE("/:id", noteHandler.DeleteNote)

	return router
}
