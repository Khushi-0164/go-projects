package service

import (
	"testing"

	"invoice-saas/internal/models"
)

func TestCreateInvoice_ComputesCorrectTotal(t *testing.T) {
	repo := newFakeInvoiceRepo()
	svc := NewInvoiceService(repo)

	items := []LineItemInput{
		{Description: "Consulting - 5 hours", Quantity: 5, UnitPriceCents: 15000},
		{Description: "Setup fee", Quantity: 1, UnitPriceCents: 5000},
	}

	invoice, err := svc.CreateInvoice(1, 1, items)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	expectedTotal := int64(80000) // (5*15000) + (1*5000)
	if invoice.TotalCents != expectedTotal {
		t.Errorf("expected total_cents %d, got %d", expectedTotal, invoice.TotalCents)
	}

	if len(invoice.LineItems) != 2 {
		t.Errorf("expected 2 line items, got %d", len(invoice.LineItems))
	}
}

func TestCreateInvoice_DefaultsToDraftStatus(t *testing.T) {
	repo := newFakeInvoiceRepo()
	svc := NewInvoiceService(repo)

	items := []LineItemInput{
		{Description: "Item", Quantity: 1, UnitPriceCents: 1000},
	}

	invoice, err := svc.CreateInvoice(1, 1, items)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if invoice.Status != models.InvoiceStatusDraft {
		t.Errorf("expected status %q, got %q", models.InvoiceStatusDraft, invoice.Status)
	}
}

func TestGetInvoice_RejectsWrongOrg(t *testing.T) {
	repo := newFakeInvoiceRepo()
	svc := NewInvoiceService(repo)

	items := []LineItemInput{{Description: "Item", Quantity: 1, UnitPriceCents: 1000}}
	invoice, _ := svc.CreateInvoice(1, 1, items)

	_, err := svc.GetInvoice(2, invoice.ID)
	if err == nil {
		t.Fatalf("expected an error when a different org tries to access this invoice, got nil")
	}
}

func TestUpdateStatus_ChangesStatus(t *testing.T) {
	repo := newFakeInvoiceRepo()
	svc := NewInvoiceService(repo)

	items := []LineItemInput{{Description: "Item", Quantity: 1, UnitPriceCents: 1000}}
	invoice, _ := svc.CreateInvoice(1, 1, items)

	if err := svc.UpdateStatus(1, invoice.ID, models.InvoiceStatusPaid); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	updated, _ := svc.GetInvoice(1, invoice.ID)
	if updated.Status != models.InvoiceStatusPaid {
		t.Errorf("expected status %q, got %q", models.InvoiceStatusPaid, updated.Status)
	}
}
