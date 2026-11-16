package handlers

import (
	"github.com/gin-gonic/gin"

	"gin-basics/internal/responses"
	"gin-basics/internal/services"
)

type TodoHandler struct {
	service *services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) ListTodos(c *gin.Context) {
	todos, err := h.service.List()

	if err != nil {
		responses.Fail(c, 500, "DATABASE_ERROR", "Failed to fetch todos")
		return
	}

	responses.OK(c, 200, todos)
}
