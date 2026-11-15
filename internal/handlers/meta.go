package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/BlackIQ/gin-todo/internal/responses"
)

func Home(c *gin.Context) {
	requestId := c.GetHeader("X-Request-ID")

	fmt.Println(requestId)

	responses.OK(
		c,
		http.StatusOK,
		requestId,
	)
}
