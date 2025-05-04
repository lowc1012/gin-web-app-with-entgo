package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
)

// Todo defines the interface for todo data operations.
type Todo interface {
	Create(ctx context.Context, todo *ent.Todo) (*ent.Todo, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Todo, error)
	GetAll(ctx context.Context) ([]*ent.Todo, error)
	Put(ctx context.Context, todo *ent.Todo) (*ent.Todo, error)
}
