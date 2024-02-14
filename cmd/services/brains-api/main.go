package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jessemolina/gopher/cmd/services/brains-api/handlers"
	"github.com/jessemolina/gopher/internal/api/debug"
	"github.com/jessemolina/gopher/pkg/config"
	"github.com/jessemolina/gopher/pkg/log"
)

var build = "develop"

func main() {
	// TODO Replace the logger with telemetry pkg logger.
	logger := log.NewLogger("brains-api")

	err := run(logger)
	if err != nil {
		logger.Error("startup", "ERROR", err)
		os.Exit(1)
	}

}

// run starts the api service.
func run(log *slog.Logger) error {

	// ================================================================
	// Configuration

	cfg := struct {
		Service struct {
			APIPort   string `config:"default:3000"`
			DebugPort string `config:"default:4000"`
		}
	}{}

	config.Parse(&cfg, "Brains")

	log.Info("service startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// ================================================================
	// Start the debug service.

	log.Info("starting debug server", "port", cfg.Service.DebugPort)

	go func() {
		if err := http.ListenAndServe(":"+cfg.Service.DebugPort, debug.DefaultMux()); err != nil {
			log.Error("shutting down debug server", "status", "ERROR")
		}
	}()

	// ================================================================
	// Start the api service.

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	apiMux := handlers.APIMux(handlers.APIMuxConfig{
		Log: log,
	})

	server := &http.Server{
		Addr:    ":" + cfg.Service.APIPort,
		Handler: apiMux,
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.Info("starting api server", "port", cfg.Service.APIPort)
		serverErrors <- server.ListenAndServe()
	}()

	// ================================================================
	// Shutdown the service.

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)
	case sig := <-shutdown:
		log.Info("shutdown started", "signal", sig)
		defer log.Info("shutdown complete", "signal", sig)
	}

	return nil
}
