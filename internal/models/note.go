package models

import (
	"github.com/google/uuid"
)

type Note struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
}
