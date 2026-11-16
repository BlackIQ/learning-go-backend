package services

import (
	"github.com/google/uuid"

	"gin-basics/internal/dto"
	"gin-basics/internal/models"
	"gin-basics/internal/repositories"
)

type NoteService struct {
	repo *repositories.NoteRepository
}

func NewNoteService(repo *repositories.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) List() ([]models.Note, error) {
	return s.repo.FindAll()
}

func (s *NoteService) GetByID(id uuid.UUID) (*models.Note, error) {
	return s.repo.FindByID(id)
}

func (s *NoteService) Create(req dto.CreateNoteRequest) (*models.Note, error) {
	note := &models.Note{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
	}

	if err := s.repo.Create(note); err != nil {
		return nil, err
	}

	return note, nil
}

// Update

// Delete
