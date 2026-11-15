package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-basics/internal/models"
	"gin-basics/internal/responses"
)

var notes = []*models.Note{}

func ListNotes(c *gin.Context) {
	responses.OK(
		c,
		http.StatusOK,
		notes,
	)
}

func GetNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"BadRequest",
			"Invalid ID provided",
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			responses.OK(
				c,
				http.StatusOK,
				notes[i],
			)
			return
		}
	}

	responses.Fail(
		c,
		http.StatusNotFound,
		"NotFound",
		"Note is not found",
	)
}

func CreateNote(c *gin.Context) {
	var newNote *models.Note

	if err := c.ShouldBindBodyWithJSON(&newNote); err != nil {
		responses.Fail(
			c,
			http.StatusConflict,
			"Conflict",
			"Invalid schema",
		)
		return
	}

	newNote.ID = uuid.New()

	notes = append(notes, newNote)

	responses.OK(
		c,
		http.StatusCreated,
		newNote,
	)
}

func UpdateNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"BadRequest",
			"Invalid ID provided",
		)
		return
	}

	var updateNote *models.Note

	if err := c.ShouldBindBodyWithJSON(&updateNote); err != nil {
		responses.Fail(
			c,
			http.StatusConflict,
			"Conflict",
			"Invalid schema",
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			updateNote.ID = noteId
			notes[i] = updateNote

			responses.OK(
				c,
				http.StatusOK,
				notes[i],
			)
			return
		}
	}

	responses.Fail(
		c,
		http.StatusNotFound,
		"NotFound",
		"Note is not found",
	)
}

func DeleteNote(c *gin.Context) {
	noteId, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"BadRequest",
			"Invalid ID provided",
		)
		return
	}

	for i, note := range notes {
		if noteId == note.ID {
			notes = append(notes[:i], notes[i+1:]...)

			responses.OK(
				c,
				http.StatusNoContent,
				nil,
			)
			return
		}
	}

	responses.Fail(
		c,
		http.StatusNotFound,
		"NotFound",
		"Note is not found",
	)
}
