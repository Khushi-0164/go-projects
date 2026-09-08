package handlers

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"invoice-saas/config"
	"invoice-saas/internal/worker"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/webhook"
)

type WebhookHandler struct {
	pool *worker.Pool
}

func NewWebhookHandler(pool *worker.Pool) *WebhookHandler {
	return &WebhookHandler{pool: pool}
}

func (h *WebhookHandler) StripeWebhook(c *gin.Context) {
	payload, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
		return
	}

	signature := c.GetHeader("Stripe-Signature")

	// Verify this request genuinely came from Stripe, using the shared
	// webhook secret. Without this check, anyone could POST a fake
	// "payment succeeded" event and mark any invoice as paid for free.
	event, err := webhook.ConstructEventWithOptions(payload, signature, config.StripeWebhookSecret(), webhook.ConstructEventOptions{
		IgnoreAPIVersionMismatch: true,
	})
	if err != nil {
		slog.Warn("webhook signature verification failed", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid signature"})
		return
	}

	if event.Type == "checkout.session.completed" {
		var session stripe.CheckoutSession
		if err := parseEventData(event, &session); err != nil {
			slog.Error("failed to parse checkout session", "error", err)
			c.JSON(http.StatusOK, gin.H{"received": true})
			return
		}

		invoiceID, _ := strconv.ParseUint(session.Metadata["invoice_id"], 10, 64)
		orgID, _ := strconv.ParseUint(session.Metadata["organization_id"], 10, 64)

		if invoiceID > 0 && orgID > 0 {
			h.pool.Enqueue(worker.Job{
				OrgID:     uint(orgID),
				InvoiceID: uint(invoiceID),
			})
		}
	}

	// Always respond 200 quickly — Stripe retries webhooks that don't
	// get a fast 2xx response, which could cause duplicate processing.
	c.JSON(http.StatusOK, gin.H{"received": true})
}

func parseEventData(event stripe.Event, target interface{}) error {
	return json.Unmarshal(event.Data.Raw, target)
}
