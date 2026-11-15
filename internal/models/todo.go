package models

import (
	"github.com/google/uuid"
)

type Todo struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title" binding:"required"`
	Details string    `json:"details" binding:"required"`
}
