package service

import (
	"errors"

	"invoice-saas/internal/models"
)

var errNotFound = errors.New("not found")

type memberKey struct {
	OrgID  uint
	UserID uint
}

type fakeOrganizationRepository struct {
	orgs      map[uint]*models.Organization
	members   map[memberKey]models.OrgRole
	users     map[string]*models.User
	nextOrgID uint
}

func newFakeOrgRepo() *fakeOrganizationRepository {
	return &fakeOrganizationRepository{
		orgs:      make(map[uint]*models.Organization),
		members:   make(map[memberKey]models.OrgRole),
		users:     make(map[string]*models.User),
		nextOrgID: 1,
	}
}

func (f *fakeOrganizationRepository) CreateWithOwner(org *models.Organization, ownerID uint) error {
	org.ID = f.nextOrgID
	f.nextOrgID++
	f.orgs[org.ID] = org
	f.members[memberKey{OrgID: org.ID, UserID: ownerID}] = models.OrgRoleOwner
	return nil
}

func (f *fakeOrganizationRepository) FindOrgsForUser(userID uint) ([]models.Organization, error) {
	var result []models.Organization
	for key := range f.members {
		if key.UserID == userID {
			if org, ok := f.orgs[key.OrgID]; ok {
				result = append(result, *org)
			}
		}
	}
	return result, nil
}

func (f *fakeOrganizationRepository) FindMemberRole(orgID, userID uint) (models.OrgRole, bool) {
	role, ok := f.members[memberKey{OrgID: orgID, UserID: userID}]
	return role, ok
}

func (f *fakeOrganizationRepository) FindUserByEmail(email string) (*models.User, error) {
	user, ok := f.users[email]
	if !ok {
		return nil, errNotFound
	}
	return user, nil
}

func (f *fakeOrganizationRepository) AddMember(member *models.OrgMember) error {
	f.members[memberKey{OrgID: member.OrganizationID, UserID: member.UserID}] = member.Role
	return nil
}

func (f *fakeOrganizationRepository) addUser(user *models.User) {
	f.users[user.Email] = user
}

type fakeCustomerRepository struct {
	customers map[uint]*models.Customer
	nextID    uint
}

func newFakeCustomerRepo() *fakeCustomerRepository {
	return &fakeCustomerRepository{
		customers: make(map[uint]*models.Customer),
		nextID:    1,
	}
}

func (f *fakeCustomerRepository) Create(customer *models.Customer) error {
	customer.ID = f.nextID
	f.nextID++
	f.customers[customer.ID] = customer
	return nil
}

func (f *fakeCustomerRepository) FindAll(orgID uint) ([]models.Customer, error) {
	var result []models.Customer
	for _, c := range f.customers {
		if c.OrganizationID == orgID {
			result = append(result, *c)
		}
	}
	return result, nil
}

func (f *fakeCustomerRepository) FindByID(orgID, customerID uint) (*models.Customer, error) {
	c, ok := f.customers[customerID]
	if !ok || c.OrganizationID != orgID {
		return nil, errNotFound
	}
	return c, nil
}

type fakeInvoiceRepository struct {
	invoices map[uint]*models.Invoice
	nextID   uint
}

func newFakeInvoiceRepo() *fakeInvoiceRepository {
	return &fakeInvoiceRepository{
		invoices: make(map[uint]*models.Invoice),
		nextID:   1,
	}
}

func (f *fakeInvoiceRepository) CreateWithLineItems(invoice *models.Invoice, lineItems []models.InvoiceLineItem) error {
	var total int64
	for _, item := range lineItems {
		total += int64(item.Quantity) * item.UnitPriceCents
	}
	invoice.TotalCents = total
	invoice.ID = f.nextID
	f.nextID++

	for i := range lineItems {
		lineItems[i].InvoiceID = invoice.ID
	}
	invoice.LineItems = lineItems

	f.invoices[invoice.ID] = invoice
	return nil
}

func (f *fakeInvoiceRepository) FindAll(orgID uint) ([]models.Invoice, error) {
	var result []models.Invoice
	for _, inv := range f.invoices {
		if inv.OrganizationID == orgID {
			result = append(result, *inv)
		}
	}
	return result, nil
}

func (f *fakeInvoiceRepository) FindByID(orgID, invoiceID uint) (*models.Invoice, error) {
	inv, ok := f.invoices[invoiceID]
	if !ok || inv.OrganizationID != orgID {
		return nil, errNotFound
	}
	return inv, nil
}

func (f *fakeInvoiceRepository) UpdateStatus(orgID, invoiceID uint, status models.InvoiceStatus) error {
	inv, ok := f.invoices[invoiceID]
	if !ok || inv.OrganizationID != orgID {
		return errNotFound
	}
	inv.Status = status
	return nil
}
