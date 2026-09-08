package repository

import "gorm.io/gorm"

type OrgSummary struct {
	TotalInvoices     int64 `json:"total_invoices"`
	TotalRevenueCents int64 `json:"total_revenue_cents"`
	OutstandingCents  int64 `json:"outstanding_cents"`
}

type SummaryRepository struct {
	DB *gorm.DB
}

func NewSummaryRepository(db *gorm.DB) *SummaryRepository {
	return &SummaryRepository{DB: db}
}

// GetSummary computes aggregate stats for an org. This is deliberately
// the kind of query worth caching — it scans and sums across every
// invoice for the org, which gets more expensive as data grows.
func (r *SummaryRepository) GetSummary(orgID uint) (*OrgSummary, error) {
	var summary OrgSummary

	if err := r.DB.Model(&struct{}{}).Table("invoices").
		Where("organization_id = ?", orgID).
		Count(&summary.TotalInvoices).Error; err != nil {
		return nil, err
	}

	if err := r.DB.Table("invoices").
		Where("organization_id = ? AND status = ?", orgID, "paid").
		Select("COALESCE(SUM(total_cents), 0)").
		Row().Scan(&summary.TotalRevenueCents); err != nil {
		return nil, err
	}

	if err := r.DB.Table("invoices").
		Where("organization_id = ? AND status IN ?", orgID, []string{"sent", "overdue"}).
		Select("COALESCE(SUM(total_cents), 0)").
		Row().Scan(&summary.OutstandingCents); err != nil {
		return nil, err
	}

	return &summary, nil
}
