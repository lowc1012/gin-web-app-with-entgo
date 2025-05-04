package mapper

import (
	"github.com/lowc1012/gin-web-app-with-entgo/internal/api/v1/dto"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent"
)

// MapEntTaskToResponse converts an ent.Task entity to a TaskResponse DTO.
func MapEntTaskToResponse(task *ent.Task) dto.TaskResponse {
	// Add nil check for safety, although Ent usually returns non-nil objects or errors
	if task == nil {
		return dto.TaskResponse{} // Or handle as appropriate
	}
	return dto.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		Status:    task.Status,
		Priority:  task.Priority,
	}
}

// MapEntTodoToResponse converts an ent.Todo entity to a TodoResponse DTO.
func MapEntTodoToResponse(todo *ent.Todo) dto.TodoResponse {
	if todo == nil {
		return dto.TodoResponse{}
	}
	return dto.TodoResponse{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		CreatedAt:   todo.CreatedAt,
	}
}
