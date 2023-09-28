package main

import (
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jessemolina/gopher/pkg/log"
	"github.com/jessemolina/gopher/pkg/config"
)

var build = "develop"

func main() {
	logger := log.NewLogger("models-api")

	err := run(logger)
	if err != nil {
		logger.Error("startup", "ERROR", err)
		os.Exit(1)
	}

}

// run starts the api service.
func run(log *slog.Logger) error {

	// Configuration

	cfg := struct {
		APIPort   string `config:"default:3000"`
		DebugPort string `config:"default:4000"`
	}{}

	config.Parse(&cfg)

	log.Info("config", "API Port", cfg.APIPort, "Debug Port", cfg.DebugPort)

	// Start the service.

	log.Info("startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// Shutdown the service.

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	sig := <-shutdown

	log.Info("shutdown started", "signal", sig)
	defer log.Info("shutdown complete", "signal", sig)

	return nil
}
