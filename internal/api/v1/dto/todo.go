package dto

import (
	"time"

	"github.com/google/uuid"
)

// TodoResponse represents a todo item returned in API responses.
type TodoResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
    UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateTodoRequest represents the data needed to create a todo item.
type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required,min=1"`
	Description string `json:"description"`
}
