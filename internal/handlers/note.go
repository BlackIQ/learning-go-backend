package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gin-basics/internal/dto"
	"gin-basics/internal/inits"
	"gin-basics/internal/models"
	"gin-basics/internal/responses"
)

func ListNotes(c *gin.Context) {
	var notes []models.Note

	result := inits.DB.Find(&notes)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	responses.OK(c, http.StatusOK, notes)
}

func GetNote(c *gin.Context) {
	noteID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_ID",
			"Invalid note ID",
		)
		return
	}

	var note models.Note

	result := inits.DB.First(&note, "id = ?", noteID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		responses.Fail(
			c,
			http.StatusNotFound,
			"NOTE_NOT_FOUND",
			"Note not found",
		)
		return
	}

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	responses.OK(c, http.StatusOK, note)
}

func CreateNote(c *gin.Context) {
	var request dto.CreateNoteRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_SCHEMA",
			err.Error(),
		)
		return
	}

	note := models.Note{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
	}

	result := inits.DB.Create(&note)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	responses.OK(c, http.StatusCreated, note)
}

func UpdateNote(c *gin.Context) {
	noteID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_ID",
			"Invalid note ID",
		)
		return
	}

	var request dto.UpdateNoteRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_SCHEMA",
			err.Error(),
		)
		return
	}

	updates := map[string]any{}

	if request.Title != nil {
		updates["title"] = *request.Title
	}

	if request.Description != nil {
		updates["description"] = *request.Description
	}

	if len(updates) == 0 {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"EMPTY_UPDATE",
			"No fields provided for update",
		)
		return
	}

	result := inits.DB.Model(&models.Note{}).
		Where("id = ?", noteID).
		Updates(updates)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	if result.RowsAffected == 0 {
		responses.Fail(
			c,
			http.StatusNotFound,
			"NOTE_NOT_FOUND",
			"Note not found",
		)
		return
	}

	var updatedNote models.Note

	if err := inits.DB.First(&updatedNote, "id = ?", noteID).Error; err != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			err.Error(),
		)
		return
	}

	responses.OK(c, http.StatusOK, updatedNote)
}

func DeleteNote(c *gin.Context) {
	noteID, err := uuid.Parse(c.Param("id"))

	if err != nil {
		responses.Fail(
			c,
			http.StatusBadRequest,
			"INVALID_ID",
			"Invalid note ID",
		)
		return
	}

	result := inits.DB.Delete(&models.Note{}, "id = ?", noteID)

	if result.Error != nil {
		responses.Fail(
			c,
			http.StatusInternalServerError,
			"DATABASE_ERROR",
			result.Error.Error(),
		)
		return
	}

	if result.RowsAffected == 0 {
		responses.Fail(
			c,
			http.StatusNotFound,
			"NOTE_NOT_FOUND",
			"Note not found",
		)
		return
	}

	c.Status(http.StatusNoContent)
}
