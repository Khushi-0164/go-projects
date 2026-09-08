package handlers

import (
	"net/http"

	"invoice-saas/internal/repository"
	"invoice-saas/internal/service"
	"invoice-saas/internal/utils"

	"github.com/gin-gonic/gin"
)

type SummaryHandler struct {
	service *service.SummaryService
	orgRepo *repository.OrganizationRepository
}

func NewSummaryHandler(s *service.SummaryService, orgRepo *repository.OrganizationRepository) *SummaryHandler {
	return &SummaryHandler{service: s, orgRepo: orgRepo}
}

func (h *SummaryHandler) GetSummary(c *gin.Context) {
	orgID := utils.ParseUint(c.Param("id"))
	userID := c.GetUint("user_id")

	if _, ok := h.orgRepo.FindMemberRole(orgID, userID); !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "you are not a member of this organization"})
		return
	}

	summary, err := h.service.GetSummary(c.Request.Context(), orgID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch summary"})
		return
	}

	c.JSON(http.StatusOK, summary)
}
