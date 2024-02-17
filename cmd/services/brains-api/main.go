package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/jessemolina/gopher/cmd/services/brains-api/v1/groups/tests"
	"github.com/jessemolina/gopher/internal/api/debug"
	"github.com/jessemolina/gopher/internal/api/router"
	"github.com/jessemolina/gopher/pkg/config"
	"github.com/jessemolina/gopher/pkg/telemetry"
	"go.opentelemetry.io/otel"
)

var build = "develop"
var service = "Brains"

func main() {
	logger := telemetry.NewLogger("brains-api")

	err := run(logger)
	if err != nil {
		logger.Error("startup", "ERROR", err)
		os.Exit(1)
	}

}

// run starts the api service.
func run(log *slog.Logger) error {

	/* Define service configuration */

	cfg := struct {
		Server struct {
			APIPort         string        `config:"default:3000"`
			DebugPort       string        `config:"default:4000"`
			ReadTimeout     time.Duration `config:"default:5s"`
			WriteTimeout    time.Duration `config:"default:10s"`
			IdleTimeout     time.Duration `config:"default:120s"`
			ShutdownTimeout time.Duration `config:"default:20s"`
		}
		OTEL struct {
			MeterExport string `config:"default:stdout"`
		}
	}{}

	config.Parse(&cfg, service)

	log.Info("service startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	/* Start the debug service */

	log.Info("starting debug server", "port", cfg.Server.DebugPort)

	go func() {
		if err := http.ListenAndServe(":"+cfg.Server.DebugPort, debug.DefaultMux()); err != nil {
			log.Error("shutting down debug server", "status", "ERROR")
		}
	}()

	/* Enable Telemetry via OTEL */

	mp, err := telemetry.NewMeterProvider(telemetry.Config{
		ServiceName:  service,
		ExporterType: cfg.OTEL.MeterExport,
	})

	if err != nil {
		return err
	}

	otel.SetMeterProvider(mp)

	log.Info("set otel meter provider", "meter", cfg.OTEL.MeterExport)

	/* Start the api service */

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	rc := router.Config{
		Log: log,
	}

	api := &http.Server{
		Addr:         ":" + cfg.Server.APIPort,
		Handler:      router.Build(tests.Routes(), rc),
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info("starting api server", "port", cfg.Server.APIPort)
		serverErrors <- api.ListenAndServe()
	}()

	/* Shutdown the service */

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		log.Info("shutdown started", "signal", sig)
		defer log.Info("shutdown complete", "signal", sig)
	}

	return nil
}
