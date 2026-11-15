package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-basics/internal/responses"
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
