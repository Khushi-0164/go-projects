package service

import "testing"

func TestCreateCustomer_ScopedToOrg(t *testing.T) {
	repo := newFakeCustomerRepo()
	svc := NewCustomerService(repo)

	customer, err := svc.CreateCustomer(1, "Acme Client", "client@acme.com")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if customer.OrganizationID != 1 {
		t.Errorf("expected OrganizationID 1, got %d", customer.OrganizationID)
	}
}

func TestGetCustomer_RejectsWrongOrg(t *testing.T) {
	repo := newFakeCustomerRepo()
	svc := NewCustomerService(repo)

	customer, _ := svc.CreateCustomer(1, "Acme Client", "client@acme.com")

	// Org 1 can see it.
	_, err := svc.GetCustomer(1, customer.ID)
	if err != nil {
		t.Fatalf("expected org 1 to access its own customer, got error: %v", err)
	}

	// Org 2 (a different tenant) should NOT be able to see it.
	_, err = svc.GetCustomer(2, customer.ID)
	if err == nil {
		t.Fatalf("expected an error when a different org tries to access this customer, got nil")
	}
}

func TestListCustomers_OnlyReturnsOwnOrg(t *testing.T) {
	repo := newFakeCustomerRepo()
	svc := NewCustomerService(repo)

	svc.CreateCustomer(1, "Org 1's Client", "a@example.com")
	svc.CreateCustomer(2, "Org 2's Client", "b@example.com")

	customers, err := svc.ListCustomers(1)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if len(customers) != 1 {
		t.Fatalf("expected exactly 1 customer for org 1, got %d", len(customers))
	}
	if customers[0].Name != "Org 1's Client" {
		t.Errorf("expected %q, got %q", "Org 1's Client", customers[0].Name)
	}
}
