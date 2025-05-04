package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
)

// Task defines the interface for task data operations.
type Task interface {
	Create(ctx context.Context, task *ent.Task) (*ent.Task, error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*ent.Task, error)
	GetAll(ctx context.Context) ([]*ent.Task, error)
	Put(ctx context.Context, task *ent.Task) (*ent.Task, error)
}
