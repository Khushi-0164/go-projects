package worker

import (
	"log/slog"

	"gorm.io/gorm"
)

// Job represents one unit of background work — here, marking an invoice as paid.
type Job struct {
	OrgID     uint
	InvoiceID uint
}

type Pool struct {
	jobs chan Job
	db   *gorm.DB
}

func NewPool(db *gorm.DB, numWorkers, queueSize int) *Pool {
	p := &Pool{
		jobs: make(chan Job, queueSize),
		db:   db,
	}

	for i := 0; i < numWorkers; i++ {
		go p.startWorker(i)
	}

	return p
}

func (p *Pool) Enqueue(job Job) {
	p.jobs <- job
}

func (p *Pool) startWorker(id int) {
	for job := range p.jobs {
		slog.Info("processing payment confirmation job", "worker_id", id, "invoice_id", job.InvoiceID)

		result := p.db.Table("invoices").
			Where("organization_id = ? AND id = ?", job.OrgID, job.InvoiceID).
			Update("status", "paid")

		if result.Error != nil {
			slog.Error("failed to mark invoice as paid", "error", result.Error, "invoice_id", job.InvoiceID)
			continue
		}

		slog.Info("invoice marked as paid", "worker_id", id, "invoice_id", job.InvoiceID)
	}
}
