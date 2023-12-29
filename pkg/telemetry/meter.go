package telemetry

import (
	"time"

	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

func NewMeterProvider(c Config) (*sdkmetric.MeterProvider, error) {
	var expOption sdkmetric.Option

	res, err := NewResource(c)
	if err != nil {
		return nil, err
	}

	switch c.ExporterType {
	case "stdout":
		exp, err := stdoutmetric.New()
		if err != nil {
			return nil, err
		}
		expOption = sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exp,
			// Default is 1m. Set to 3s for demonstrative purposes.
			sdkmetric.WithInterval(3*time.Second)))
	case "prometheus":
		exp, err := prometheus.New()
		if err != nil {
			return nil, err
		}
		expOption = sdkmetric.WithReader(exp)

	}
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		expOption,
	)

	return mp, nil
}
