package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jessemolina/gopher/pkg/log"
	"golang.org/x/exp/slog"
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
