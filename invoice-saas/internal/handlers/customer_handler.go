package handlers

import (
	"invoice-saas/internal/repository"
	"invoice-saas/internal/service"
	"invoice-saas/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *service.CustomerService
	orgRepo *repository.OrganizationRepository
}

func NewCustomerHandler(s *service.CustomerService, orgRepo *repository.OrganizationRepository) *CustomerHandler {
	return &CustomerHandler{service: s, orgRepo: orgRepo}
}

type createCustomerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (h *CustomerHandler) isMember(orgID, userID uint) bool {
	_, ok := h.orgRepo.FindMemberRole(orgID, userID)
	return ok
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if !h.isMember(orgID, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	var req createCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer, err := h.service.CreateCustomer(orgID, req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create customer"})
		return
	}

	c.JSON(http.StatusCreated, customer)

}

func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if !h.isMember(orgID, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	customers, err := h.service.ListCustomers(orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch customers"})
		return
	}

	c.JSON(http.StatusOK, customers)
}
