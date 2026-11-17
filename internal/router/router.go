package router

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gin-basics/internal/handlers"
)

func SetupRouter(diskFS *embed.FS) *gin.Engine {
	router := gin.Default()

	router.Use(cors.Default())

	// Frontend
	frontendFS, err := fs.Sub(*diskFS, "frontend/dist")
	if err != nil {
		panic(err)
	}

	assetsFS, err := fs.Sub(frontendFS, "assets")
	if err != nil {
		panic(err)
	}

	router.StaticFS("/assets", http.FS(assetsFS))

	router.GET("/", func(c *gin.Context) {
		index, err := fs.ReadFile(frontendFS, "index.html")

		if err != nil {
			c.AbortWithStatus(
				http.StatusInternalServerError,
			)

			return
		}

		c.Data(http.StatusOK, "text/html; charset=utf-8", index)
	})

	// Backend
	apiRouter := router.Group("/api")

	todoRouter := apiRouter.Group("/todos")

	todoRouter.GET("", handlers.AllTodos)
	todoRouter.GET("/:id", handlers.OneTodo)
	todoRouter.POST("", handlers.CreateTodo)
	todoRouter.PATCH("/:id", handlers.UpdateTodo)
	todoRouter.DELETE("/:id", handlers.DeleteTodo)

	return router
}
