package log

import (
	"os"

	"golang.org/x/exp/slog"
)

// NewLogger creates a JSON structured logger that writes to Stdout.
func NewLogger(service string) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", service)
	return logger
}
