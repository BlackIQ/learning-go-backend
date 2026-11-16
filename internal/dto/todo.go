package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTodoRequest struct {
	Title   string `json:"title" binding:"required"`
	Details string `json:"details" binding:"required"`
}

type UpdateTodoRequest struct {
	Title   *string `json:"title"`
	Details *string `json:"details"`
}

type TodoResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
