package v1

import (
	"github.com/lowc1012/gin-web-app-with-entgo/internal/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/api/v1/dto" // Import DTOs
	"github.com/lowc1012/gin-web-app-with-entgo/internal/api/v1/mapper"
	taskApp "github.com/lowc1012/gin-web-app-with-entgo/internal/application/task"
	"github.com/lowc1012/gin-web-app-with-entgo/internal/ent" // Keep for ent specific errors or types if needed
)

// TaskHandler holds dependencies for task handlers.
type TaskHandler struct {
	taskService *taskApp.Service
}

// NewTaskHandler creates a new TaskHandler.
func NewTaskHandler(taskService *taskApp.Service) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// RegisterTaskRoutes registers task routes with the Gin engine.
func (h *TaskHandler) RegisterTaskRoutes(router *gin.RouterGroup) {
	tasks := router.Group("/tasks")
	{
		tasks.GET("", h.getAllTasksHandler)
		tasks.GET("/:id", h.getTaskHandler)
		tasks.POST("", h.createTaskHandler)
		tasks.PUT("/:id", h.putTaskHandler)
		tasks.DELETE("/:id", h.deleteTaskHandler)
	}
}

// getAllTasksHandler handles GET requests for all tasks.
func (h *TaskHandler) getAllTasksHandler(c *gin.Context) {
	// Use the injected service and pass the request context
	tasks, err := h.taskService.GetAllTasks(c.Request.Context())
	if err != nil {
		log.Errorw("Failed to get all tasks from service", "error", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve tasks"})
		return
	}

	// Map ent.Task entities to TaskResponse DTOs
	taskResponses := make([]dto.TaskResponse, 0, len(tasks))
	for _, t := range tasks {
		taskResponses = append(taskResponses, mapper.MapEntTaskToResponse(t))
	}

	// Return the result
	c.JSON(http.StatusOK, taskResponses)
}

func (h *TaskHandler) getTaskHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Warnw("Invalid task ID format received", "id", idStr, "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid task ID format"})
		return
	}

	task, err := h.taskService.GetByID(c.Request.Context(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			log.Infow("Task not found", "id", id)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		} else {
			log.Errorw("Failed to get task by ID from service", "id", id, "error", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve task"})
		}
		return
	}

	c.JSON(http.StatusOK, mapper.MapEntTaskToResponse(task))
}

func (h *TaskHandler) createTaskHandler(c *gin.Context) {
	var req dto.CreateTaskRequest

	// Bind and validate request body
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Errorw("Failed to bind create task request", "error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body", "details": err.Error()})
		return
	}

	// TODO:
	// 1. Call h.taskService.Create
	// 2. Handle potential service errors (validation errors, internal errors)
	// 3. Map the created task to a DTO
	// 4. Return HTTP StatusCreated (201) with the created task DTO

	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (h *TaskHandler) putTaskHandler(c *gin.Context) {
	// TODO: Implement using h.taskService, get ID from c.Param("id"), bind request body
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}

func (h *TaskHandler) deleteTaskHandler(c *gin.Context) {
	// TODO: Implement using h.taskService, get ID from c.Param("id")
	c.JSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
}
