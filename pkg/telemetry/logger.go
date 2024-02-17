package telemetry

import (
	"log/slog"
	"os"
)

// NewLogger returns a structured logger that writes to stdout.
func NewLogger(service string) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", service)
	return logger
}
