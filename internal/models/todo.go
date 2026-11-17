package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Details   string         `gorm:"not null" json:"details"`
	IsDone    bool           `gorm:"default:false" json:"isDone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
