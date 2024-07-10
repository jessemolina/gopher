package evaluation

import (
	"time"

	"github.com/google/uuid"
)

// Evaluation represents the evaluation metrics for a model.
type Evaluation struct {
	ID          uuid.UUID
	ModelID     uuid.UUID
	JobID       uuid.UUID
	Metrics     map[string]float64
	EvaluatedAt time.Time
}
