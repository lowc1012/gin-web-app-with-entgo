package dto

import (
	"time"

	"github.com/google/uuid"
)

// TaskResponse represents a task returned in API responses.
type TaskResponse struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	Status    string    `json:"status"`
	Priority  int       `json:"priority"`
}

// CreateTaskRequest represents the data needed to create a task.
type CreateTaskRequest struct {
	Title    string `json:"title" binding:"required"` // Add validation tags
	Priority *int   `json:"priority"`                 // Use pointer for optional fields
}
