package persistence

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/domain/interfaces"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent/todo"
)

var _ interfaces.Todo = (*entTodoRepository)(nil)

type entTodoRepository struct {
	client *ent.Client
}

func NewEntTodoRepository(client *ent.Client) *entTodoRepository {
	return &entTodoRepository{
		client: client,
	}
}

func (r *entTodoRepository) GetAll(ctx context.Context) ([]*ent.Todo, error) {
	return r.client.Todo.Query().Where(todo.DeletedAtIsNil()).All(ctx)
}

func (r *entTodoRepository) GetByID(ctx context.Context, id uuid.UUID) (*ent.Todo, error) {
	return r.client.Todo.Query().
		Where(todo.ID(id)).
		Where(todo.DeletedAtIsNil()).
		Only(ctx)
}

func (r *entTodoRepository) Create(ctx context.Context, todo *ent.Todo) (*ent.Todo, error) {
	now := time.Now()
	return r.client.Todo.
		Create().
		SetID(todo.ID).
		SetTitle(todo.Title).
		SetDescription(todo.Description).
		SetCreatedAt(now).
		SetUpdatedAt(now).
		Save(ctx)
}

func (r *entTodoRepository) Delete(ctx context.Context, id uuid.UUID) error {
	todo, err := r.client.Todo.Query().Where(todo.ID(id)).Only(ctx)
	if err != nil {
		return err
	}

	_, err = todo.Update().SetDeletedAt(time.Now()).Save(ctx)
	return err
}

func (r *entTodoRepository) Put(ctx context.Context, todo *ent.Todo) (*ent.Todo, error) {
	return r.client.Todo.
		UpdateOneID(todo.ID).
		SetTitle(todo.Title).
		SetDescription(todo.Description).
		SetUpdatedAt(time.Now()).
		Save(ctx)
}
