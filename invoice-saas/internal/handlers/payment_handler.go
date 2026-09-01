package handlers

import (
	"net/http"

	"invoice-saas/internal/repository"
	"invoice-saas/internal/service"
	"invoice-saas/internal/utils"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *service.PaymentService
	orgRepo *repository.OrganizationRepository
}

func NewPaymentHandler(s *service.PaymentService, orgRepo *repository.OrganizationRepository) *PaymentHandler {
	return &PaymentHandler{service: s, orgRepo: orgRepo}
}

func (h *PaymentHandler) CreateCheckoutSession(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if _, ok := h.orgRepo.FindMemberRole(orgID, userID); !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	invoiceID := utils.ParseUint(c.Param("invoiceId"))

	url, err := h.service.CreateCheckoutSession(orgID, invoiceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create checkout session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"checkout_url": url})
}
