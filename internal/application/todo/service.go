package todo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/domain/interfaces"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
)

type Service struct {
	todos interfaces.Todo
}

func NewService(todos interfaces.Todo) *Service {
	return &Service{
		todos: todos,
	}
}

// GetAllTodos retrieves all todos.
func (s *Service) GetAllTodos(ctx context.Context) ([]*ent.Todo, error) {
	todos, err := s.todos.GetAll(ctx)
	// TODO: Add any validation or default value logic

	return todos, err // Return the fetched todos and error
}

// GetByID retrieves a single todo by its ID
func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*ent.Todo, error) {
	// TODO: Add any validation or default value logic
	return s.todos.GetByID(ctx, id)
}

func (s *Service) Create(ctx context.Context, title string, description string) (*ent.Todo, error) {
	// TODO: Add any validation or default value logic
	id := uuid.New()

	return s.todos.Create(ctx, &ent.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
}

// Delete removes a todo by its ID
func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	// Add authorization checks or other logic here if needed
	return s.todos.Delete(ctx, id)
}

// Update updates an existing todo
func (s *Service) Update(ctx context.Context, id uuid.UUID, title string, description string) (*ent.Todo, error) {
	// Add validation or business logic here

	return s.todos.Put(ctx, &ent.Todo{
		ID:          id,
		Title:       title,
		Description: description,
	})
}
