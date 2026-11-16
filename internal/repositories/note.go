package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{db: db}
}

func (r *NoteRepository) FindAll() ([]models.Note, error) {
	var notes []models.Note

	err := r.db.Find(&notes).Error

	return notes, err
}

func (r *NoteRepository) FindByID(id uuid.UUID) (*models.Note, error) {
	var note models.Note

	err := r.db.First(&note, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *NoteRepository) Create(note *models.Note) error {
	return r.db.Create(note).Error
}

func (r *NoteRepository) Update(id uuid.UUID, updates map[string]any) error {
	return r.db.Model(&models.Note{}).Where("id = ?", id).Updates(updates).Error
}

func (r *NoteRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Note{}, "id = ?", id).Error
}
