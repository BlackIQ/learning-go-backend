package handlers

import (
	"github.com/gin-gonic/gin"

	"gin-basics/internal/responses"
	"gin-basics/internal/services"
)

type NoteHandler struct {
	service *services.NoteService
}

func NewNoteHandler(service *services.NoteService) *NoteHandler {
	return &NoteHandler{service: service}
}

func (h *NoteHandler) ListNotes(c *gin.Context) {
	notes, err := h.service.List()

	if err != nil {
		responses.Fail(c, 500, "DATABASE_ERROR", "Failed to fetch notes")
		return
	}

	responses.OK(c, 200, notes)
}
