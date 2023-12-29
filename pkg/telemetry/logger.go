package telemetry

import (
	"log/slog"
	"os"
)

// NewLogger returns a structured logger that writes to stdout.
func NewLogger(c Config) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil)).With("service", c.ServiceName)
	return logger
}
