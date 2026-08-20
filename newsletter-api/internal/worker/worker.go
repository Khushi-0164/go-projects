package worker

import (
	"log/slog"
	"time"
)

// Job represents a unit of background work — here, just an email to "send".
type Job struct {
	Email string
}

// Pool manages a fixed number of worker goroutines pulling jobs off a shared queue.
type Pool struct {
	jobs chan Job
}

// NewPool creates a worker pool with the given queue capacity, and starts
// numWorkers goroutines immediately, ready to process jobs as they arrive.
func NewPool(numWorkers, queueSize int) *Pool {
	p := &Pool{
		jobs: make(chan Job, queueSize),
	}

	for i := 0; i < numWorkers; i++ {
		go p.startWorker(i)
	}

	return p
}

// Enqueue adds a job to the queue. If the queue is full, this blocks —
// in a real system you might instead make this non-blocking with a
// select/default, dropping or erroring instead of waiting.
func (p *Pool) Enqueue(job Job) {
	p.jobs <- job
}

func (p *Pool) startWorker(id int) {
	for job := range p.jobs {
		slog.Info("processing welcome email job", "worker_id", id, "email", job.Email)

		// Simulate a slow external API call (a real email provider).
		time.Sleep(2 * time.Second)

		slog.Info("welcome email sent", "worker_id", id, "email", job.Email)
	}
}
