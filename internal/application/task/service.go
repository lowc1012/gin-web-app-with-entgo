package task

import (
	"context"

	"github.com/lowc1012/gin-web-app-with-entgo/internal/domain/interfaces"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
)

// Service provides task-related use cases.
type Service struct {
	tasks interfaces.Task
}

// NewService creates a new task service.
func NewService(tasks interfaces.Task) *Service {
	return &Service{tasks: tasks}
}

// GetAllTasks retrieves all tasks.
func (s *Service) GetAllTasks(ctx context.Context) ([]*ent.Task, error) {
	// Add any application-specific logic here (validation, orchestration) if needed
	return s.tasks.GetAll(ctx)
}

func (s *Service) GetByID(ctx context.Context, id int) (*ent.Task, error) {
    return nil, nil
}

func (s *Service) Update() {

}

func (s *Service) Delete() {

}

func (s *Service) Create() {

}
