package responses

import (
	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, status int, data any) {
	c.JSON(status, Response{
		Success: true,
		Data:    data,
	})
}

func Fail(c *gin.Context, status int, code, message string) {
	c.JSON(status, Response{
		Success: false,
		Error: &ErrorInfo{
			Code:    code,
			Message: message,
		},
	})
}
