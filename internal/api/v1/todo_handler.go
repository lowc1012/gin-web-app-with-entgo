package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/api/v1/dto"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/api/v1/mapper"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/application/todo"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent" // For ent specific errors or types
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
)

type TodoHandler struct {
	todoService *todo.Service
}

func NewTodoHandler(service *todo.Service) *TodoHandler {
	return &TodoHandler{todoService: service}
}

func (h *TodoHandler) RegisterTodoRoutes(router *gin.RouterGroup) {
	todos := router.Group("/todos")
	{
		todos.GET("", h.getAllTodosHandler)
		todos.GET("/:id", h.getTodoHandler)
		todos.POST("", h.createTodoHandler)
		todos.PUT("/:id", h.putTodoHandler)
		todos.DELETE("/:id", h.deleteTodoHandler)
	}
}

func (h *TodoHandler) getAllTodosHandler(c *gin.Context) {
	todos, err := h.todoService.GetAllTodos(c.Request.Context())
	if err != nil {
		log.Errorw("Failed to get all todos from service", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve todos"})
		return
	}

	// Map ent.Todo entities to TodoResponse DTOs
	todoResponses := make([]dto.TodoResponse, 0, len(todos))
	for _, t := range todos {
		todoResponses = append(todoResponses, mapper.MapEntTodoToResponse(t))
	}

	c.JSON(http.StatusOK, todoResponses)
}

func (h *TodoHandler) getTodoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.FromBytes([]byte(idStr))
	if err != nil {
		log.Warnw("Invalid todo ID format received", "id", idStr, "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid todo ID format"})
		return
	}

	todoItem, err := h.todoService.GetByID(c.Request.Context(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Infow("Todo not found", "id", id)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		} else {
			log.Errorw("Failed to get todo by ID from service", "id", id, "error", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve todo"})
		}
		return
	}

	c.JSON(http.StatusOK, mapper.MapEntTodoToResponse(todoItem))
}

func (h *TodoHandler) createTodoHandler(c *gin.Context) {
	var req dto.CreateTodoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warnw("Failed to bind create todo request", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "details": err.Error()})
		return
	}

	todo, err := h.todoService.Create(c.Request.Context(), req.Title, req.Description) // Example service call
	if err != nil {
		// Handle potential validation errors from the service layer or other errors
		log.Errorw("Failed to create todo via service", "request", req, "error", err)
		// TODO: Map specific service errors (e.g., validation) to 400 Bad Request
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, mapper.MapEntTodoToResponse(todo))
}

func (h *TodoHandler) putTodoHandler(c *gin.Context) {
	// TODO: Implement PUT handler
	// 1. Get ID from c.Param("id"), validate it.
	// 2. Bind JSON body to an UpdateTodoRequest DTO (define this DTO).
	// 3. Call a service method like h.todoService.UpdateTodo(ctx, id, updateData).
	// 4. Handle errors (NotFound, Validation, Internal).
	// 5. Map the updated todo to a DTO and return 200 OK.
	c.JSON(http.StatusNotImplemented, gin.H{"message": "PUT not implemented"})
}

func (h *TodoHandler) deleteTodoHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.FromBytes([]byte(idStr))
	if err != nil {
		log.Warnw("Invalid todo ID format received for delete", "id", idStr, "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid todo ID format"})
		return
	}

	err = h.todoService.Delete(c.Request.Context(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Infow("Todo not found for deletion", "id", id)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Todo not found"}) // Or maybe 204 anyway? Depends on preference.
		} else {
			log.Errorw("Failed to delete todo via service", "id", id, "error", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete todo"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
