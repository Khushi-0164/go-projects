package service

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

type PaymentService struct {
	invoiceRepo InvoiceRepository
}

func NewPaymentService(invoiceRepo InvoiceRepository) *PaymentService {
	return &PaymentService{invoiceRepo: invoiceRepo}
}

func (s *PaymentService) CreateCheckoutSession(orgID, invoiceID uint) (string, error) {
	invoice, err := s.invoiceRepo.FindByID(orgID, invoiceID)
	if err != nil {
		return "", err
	}
	params := &stripe.CheckoutSessionParams{
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency:   stripe.String("usd"),
					UnitAmount: stripe.Int64(invoice.TotalCents),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Invoice #" + itoa(invoice.ID)),
					},
				},
				Quantity: stripe.Int64(1),
			},
		},
		Metadata: map[string]string{
			"invoice_id":      itoa(invoice.ID),
			"organization_id": itoa(invoice.OrganizationID),
		},
		SuccessURL: stripe.String("http://localhost:8080/payment-success"),
		CancelURL:  stripe.String("http://localhost:8080/payment-cancelled"),
	}

	sess, err := session.New(params)
	if err != nil {
		return "", err
	}

	return sess.URL, nil
}

func itoa(n uint) string {
	if n == 0 {
		return "0"
	}
	digits := []byte{}
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}
