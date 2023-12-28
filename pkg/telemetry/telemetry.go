package telemetry

import (
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// Config defines the instrumentation configuration.
type Config struct {
	ServiceName  string
	ExporterType string
}

// NewResource defines an OTEL resource based off Config.
func NewResource(c Config) (*sdkresource.Resource, error) {
	r, err := sdkresource.Merge(
		sdkresource.Default(),
		sdkresource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(c.ServiceName),
		),
	)

	if err != nil {
		return nil, err
	}

	return r, nil
}

