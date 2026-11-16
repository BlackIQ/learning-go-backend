package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"gin-basics/internal/models"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo

	err := r.db.Find(&todos).Error

	return todos, err
}

func (r *TodoRepository) FindByID(id uuid.UUID) (*models.Todo, error) {
	var todo models.Todo

	err := r.db.First(&todo, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *TodoRepository) Create(todo *models.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) Update(id uuid.UUID, updates map[string]any) error {
	return r.db.Model(&models.Todo{}).Where("id = ?", id).Updates(updates).Error
}

func (r *TodoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Todo{}, "id = ?", id).Error
}
