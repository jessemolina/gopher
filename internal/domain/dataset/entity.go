package dataset

import (
	"time"

	"github.com/google/uuid"
)

type Dataset struct {
	ID          uuid.UUID
	Name        string
	Description string
	Location    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
