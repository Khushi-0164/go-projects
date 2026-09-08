package handlers

import (
	"net/http"
	"strconv"

	"invoice-saas/internal/models"
	"invoice-saas/internal/repository"
	"invoice-saas/internal/service"
	"invoice-saas/internal/utils"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	service *service.InvoiceService
	orgRepo *repository.OrganizationRepository
}

func NewInvoiceHandler(s *service.InvoiceService, orgRepo *repository.OrganizationRepository) *InvoiceHandler {
	return &InvoiceHandler{service: s, orgRepo: orgRepo}
}

func (h *InvoiceHandler) isMember(orgID, userID uint) bool {
	_, ok := h.orgRepo.FindMemberRole(orgID, userID)
	return ok
}

type lineItemRequest struct {
	Description    string `json:"description" binding:"required"`
	Quantity       int    `json:"quantity" binding:"required,gt=0"`
	UnitPriceCents int64  `json:"unit_price_cents" binding:"required,gt=0"`
}

type createInvoiceRequest struct {
	CustomerID uint              `json:"customer_id" binding:"required"`
	LineItems  []lineItemRequest `json:"line_items" binding:"required,min=1,dive"`
}

func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if !h.isMember(orgID, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	var req createInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items := make([]service.LineItemInput, len(req.LineItems))
	for i, li := range req.LineItems {
		items[i] = service.LineItemInput{
			Description:    li.Description,
			Quantity:       li.Quantity,
			UnitPriceCents: li.UnitPriceCents,
		}
	}

	invoice, err := h.service.CreateInvoice(orgID, req.CustomerID, items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create invoice"})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

func (h *InvoiceHandler) ListInvoices(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if !h.isMember(orgID, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}
	status := c.Query("status")

	invoices, total, err := h.service.ListInvoices(orgID, page, limit, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch invoices"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        invoices,
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": (total + int64(limit) - 1) / int64(limit),
	})
}

func (h *InvoiceHandler) GetInvoice(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if !h.isMember(orgID, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	invoiceID := utils.ParseUint(c.Param("invoiceId"))
	invoice, err := h.service.GetInvoice(orgID, invoiceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

type updateStatusRequest struct {
	Status models.InvoiceStatus `json:"status" binding:"required"`
}

func (h *InvoiceHandler) UpdateStatus(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if !h.isMember(orgID, userID) {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	var req updateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoiceID := utils.ParseUint(c.Param("invoiceId"))
	if err := h.service.UpdateStatus(orgID, invoiceID, req.Status); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invoice status updated"})
}
