package persistence

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/domain/interfaces"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent/task"
)

// Ensure entTaskRepository implements the interface
var _ interfaces.Task = (*entTaskRepository)(nil)

type entTaskRepository struct {
	client *ent.Client
}

// NewEntTaskRepository creates a new repository implementation using Ent.
func NewEntTaskRepository(client *ent.Client) interfaces.Task {
	return &entTaskRepository{client: client}
}

func (r *entTaskRepository) GetAll(ctx context.Context) ([]*ent.Task, error) {
	return r.client.Task.Query().Where(task.DeletedAtIsNil()).All(ctx)
}

func (r *entTaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Task, error) {
    return r.client.Task.Query().Where(task.ID(id)).Where(task.DeletedAtIsNil()).Only(ctx)
}

func (r *entTaskRepository) Create(ctx context.Context, task *ent.Task) (*ent.Task, error) {
    return r.client.Task.Create().
        SetID(task.ID).
        SetTitle(task.Title).
        SetDescription(task.Description).
        SetTodoID(task.TodoID).
        SetParentID(task.ParentID).
        SetPriority(task.Priority).
        SetStatus(task.Status).
        SetCreatedAt(task.CreatedAt).
        SetUpdatedAt(task.UpdatedAt).
        Save(ctx)
}

func (r *entTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
    task, err := r.client.Task.Query().Where(task.ID(id)).Only(ctx)
    if err != nil {
        return err
    }
    _, err = task.Update().SetDeletedAt(time.Now()).Save(ctx)
    return err
}

func (r *entTaskRepository) Put(ctx context.Context, task *ent.Task) (*ent.Task, error) {
    return r.client.Task.UpdateOneID(task.ID).
        SetTitle(task.Title).
        SetDescription(task.Description).
        SetPriority(task.Priority).
        SetParentID(task.ParentID).
        SetTodoID(task.TodoID).
        SetStatus(task.Status).
        SetUpdatedAt(time.Now()).
        Save(ctx)
}
