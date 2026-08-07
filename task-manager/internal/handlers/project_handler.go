package handlers

import (
	"errors"
	"net/http"

	"task-manager/internal/models"
	"task-manager/internal/service"
	"task-manager/internal/utils"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	service *service.ProjectService
}

func NewProjectHandler(s *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: s}
}

type createProjectRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	project, err := h.service.CreateProject(req.Name, req.Description, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, project)
}

func (h *ProjectHandler) ListMyProjects(c *gin.Context) {
	userID := c.GetUint("user_id")

	projects, err := h.service.ListMyProjects(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

type addMemberRequest struct {
	Email string             `json:"email" binding:"required,email"`
	Role  models.ProjectRole `json:"role" binding:"required"`
}

func (h *ProjectHandler) AddMember(c *gin.Context) {
	projectID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	var req addMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err := h.service.AddMember(projectID, userID, req.Email, req.Role)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotAuthorized):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case errors.Is(err, service.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusConflict, gin.H{"error": "user is already a member of this project"})
		}
		return
	}

	c.JSON(http.StatusCreated, member)
}
