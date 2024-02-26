package monitor

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// meter defines the meter used for api middleware metrics.
var meter = otel.Meter("api-middleware")

// metrics is a singleton type for the api middleware metrics.
var metrics *apiMetrics

type apiMetrics struct {
	requests metric.Int64Counter
	errors   metric.Int64Counter
}

// init initializes the metrics singleton.
func init() {
	metrics = &apiMetrics{}
	metrics.createMeasures()
}

// getMetrics returns a pointer to the metrics singleton.
func ApiMetrics() *apiMetrics {
	return metrics
}

// createMeasures creates all of the otel metrics for apiMetrics.
func (am *apiMetrics) createMeasures() {
	am.requests, _ = meter.Int64Counter("api_requests",
		metric.WithDescription("The number of api requests made"),
		metric.WithUnit("{request}"))

	am.errors, _ = meter.Int64Counter("api_errors",
		metric.WithDescription("The number of api errors made"),
		metric.WithUnit("{error}"))
}

func (am *apiMetrics) AddRequests(ctx context.Context, incr int64) {
	am.requests.Add(ctx, incr)
}

func (am *apiMetrics) AddErrors(ctx context.Context, incr int64) {
	am.errors.Add(ctx, incr)
}
