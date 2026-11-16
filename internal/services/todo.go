package services

import (
	"github.com/google/uuid"

	"gin-basics/internal/dto"
	"gin-basics/internal/models"
	"gin-basics/internal/repositories"
)

type TodoService struct {
	repo *repositories.TodoRepository
}

func NewTodoService(repo *repositories.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) List() ([]models.Todo, error) {
	return s.repo.FindAll()
}

func (s *TodoService) GetByID(id uuid.UUID) (*models.Todo, error) {
	return s.repo.FindByID(id)
}

func (s *TodoService) Create(req dto.CreateTodoRequest) (*models.Todo, error) {
	todo := &models.Todo{
		ID:      uuid.New(),
		Title:   req.Title,
		Details: req.Details,
	}

	if err := s.repo.Create(todo); err != nil {
		return nil, err
	}

	return todo, nil
}

// Update

// Delete
