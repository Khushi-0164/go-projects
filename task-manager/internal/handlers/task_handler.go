package handlers

import (
	"errors"
	"net/http"

	"task-manager/internal/models"
	"task-manager/internal/service"
	"task-manager/internal/utils"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(s *service.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

type createTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	AssigneeID  *uint  `json:"assignee_id"`
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	projectID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	var req createTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.CreateTask(projectID, userID, req.Title, req.Description, req.AssigneeID)
	if err != nil {
		respondServiceError(c, err)
		return
	}

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	projectID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	tasks, err := h.service.ListTasks(projectID, userID)
	if err != nil {
		respondServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, tasks)
}

type updateTaskStatusRequest struct {
	Status models.TaskStatus `json:"status" binding:"required"`
}

func (h *TaskHandler) UpdateTaskStatus(c *gin.Context) {
	taskID := utils.ParseUint(c.Param("taskId"))
	userID := c.GetUint("user_id")

	var req updateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := h.service.UpdateTaskStatus(taskID, userID, req.Status)
	if err != nil {
		respondServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, task)
}

// respondServiceError maps known service-layer errors to HTTP status codes.
// Shared across handlers so we don't repeat this switch in every function.
func respondServiceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, service.ErrNotAuthorized):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrTaskNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, service.ErrUserNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
	}
}
