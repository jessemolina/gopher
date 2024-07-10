package prediction

import (
	"time"

	"github.com/google/uuid"
)

// Prediction represents the result of using a trained model to make predictions on new data.
type Prediction struct {
	ID        uuid.UUID
	ModelID   uuid.UUID
	InputData string // Input data in JSON or another format
	Result    string // Prediction result in JSON or another format
	CreatedAt time.Time
}
