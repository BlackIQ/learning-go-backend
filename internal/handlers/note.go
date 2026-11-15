package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/BlackIQ/gin-todo/internal/models"
)

var notes = []*models.Note{}

func ListNotes(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		notes,
	)
}

func GetNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			c.JSON(
				http.StatusOK,
				notes[i],
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Note not found",
		},
	)
}

func CreateNote(c *gin.Context) {
	var newNote *models.Note

	if err := c.ShouldBindBodyWithJSON(&newNote); err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	newNote.ID = uuid.New()

	notes = append(notes, newNote)

	c.JSON(
		http.StatusCreated,
		newNote,
	)
}

func UpdateNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	var updateNote *models.Note

	if err := c.ShouldBindBodyWithJSON(&updateNote); err != nil {
		c.JSON(
			http.StatusConflict,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			updateNote.ID = noteId
			notes[i] = updateNote

			c.JSON(
				http.StatusOK,
				notes[i],
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Note not found",
		},
	)
}

func DeleteNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Invalid ID",
			},
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			notes = append(notes[:i], notes[i+1:]...)

			c.Status(
				http.StatusNoContent,
			)
			return
		}
	}

	c.JSON(
		http.StatusNotFound,
		gin.H{
			"error": "Todo not found",
		},
	)
}
