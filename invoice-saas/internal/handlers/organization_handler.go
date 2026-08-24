package handlers

import (
	"errors"
	"invoice-saas/internal/models"
	"invoice-saas/internal/service"
	"invoice-saas/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	service *service.OrganizationService
}

func NewOrganizationHandler(s *service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{service: s}
}

type createOrgRequest struct {
	Name string `json:"name" binding:"required"`
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var req createOrgRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("user_id")

	org, err := h.service.CreateOrganization(req.Name, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create organization"})
		return
	}

	c.JSON(http.StatusCreated, org)
}

func (h *OrganizationHandler) ListMyOrganizations(c *gin.Context) {
	userID := c.GetUint("user_id")

	orgs, err := h.service.ListMyOrganizations(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch organizations"})
		return
	}

	c.JSON(http.StatusOK, orgs)
}

type addMemberRequest struct {
	Email string         `json:"email" binding:"required,email"`
	Role  models.OrgRole `json:"role" binding:"required"`
}

func (h *OrganizationHandler) AddMember(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")
	var req addMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	member, err := h.service.AddMember(orgID, userID, req.Email, req.Role)
	if err != nil {
		switch {
		case errors.Is(err, service.ErrNotAuthorized):
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		case errors.Is(err, service.ErrUserNotFound):
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusConflict, gin.H{"error": "user is already a member of this organization"})
		}
		return
	}
	c.JSON(http.StatusCreated, member)
}
