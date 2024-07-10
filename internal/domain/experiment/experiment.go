package experiment

import (
	"time"

	"github.com/google/uuid"
)

// Experiment represents a set of related training jobs.
type Experiment struct {
	ID          uuid.UUID
	Name        string
	Description string
	Jobs        []uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
