package telemetry

import (
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	jaegerUrl = "http://jaeger-collector.o11y-system.svc.cluster.local:14268/api/traces"
)

// NewTraceProvider creates a trace provider based on Config.
func NewTraceProvider(c Config) (*sdktrace.TracerProvider, error) {
	var exp sdktrace.SpanExporter

	res, err := NewResource(c)
	if err != nil {
		return nil, err
	}

	switch c.ExporterType {
	case "stdout":
		exp, err = stdouttrace.New()

	case "jaeger":
		exp, err = jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(jaegerUrl)))
		if err != nil {
			return nil, err
		}

	}
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(res),
	)

	return tp, nil
}
