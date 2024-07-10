package model

import (
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID          uuid.UUID
	Name        string
	Description string
	Version     int
	HyperParams HyperParameters
	Parameters  Parameters
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Parameters struct {
	Weights []float64
	Biases  []float64
}

type HyperParameters struct {
	LearningRate float64
	Epochs       int
	BatchSize    int
}
