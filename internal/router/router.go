package router

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gin-basics/internal/handlers"
	"gin-basics/internal/middlewares"
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

	authRouter := apiRouter.Group("/auth")

	authRouter.POST("/signup", handlers.Signup)
	authRouter.POST("/signin", handlers.Signin)

	userRouter := apiRouter.Group("/users")
	userRouter.Use(middlewares.AuthMiddleware())

	userRouter.GET("/me", handlers.Me)

	todoRouter := apiRouter.Group("/todos")
	todoRouter.Use(middlewares.AuthMiddleware())

	todoRouter.GET("", handlers.AllTodos)
	todoRouter.GET("/:id", handlers.OneTodo)
	todoRouter.POST("", handlers.CreateTodo)
	todoRouter.PATCH("/:id", handlers.UpdateTodo)
	todoRouter.DELETE("/:id", handlers.DeleteTodo)

	return router
}
