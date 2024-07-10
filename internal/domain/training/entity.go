package training

import (
	"time"

	"github.com/google/uuid"
)

type Status string

type Job struct {
	ID          uuid.UUID
	ModelID     uuid.UUID
	DatasetID   uuid.UUID
	Status      Status
	StartedAt   time.Time
	EndedAt     time.Time
	Performance float64
	Metrics     map[string]float64
}
