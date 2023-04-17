package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var build = "dev"

func main() {

	// ================================================================
	// RUN

	if err := run(); err != nil {
		log.Println("run error")
		os.Exit(1)
	}

}

// TODO build a run function to load handlers and start server mux
func run() error {

	// ================================================================
	// STARTUP

	// Log service start up.
	log.Println("starting service", build)
	defer log.Println("service ended")

	// ================================================================
	// SHUTDOWN

	// make a channel with 1 buffer for an os.Signal
	// block on the channel until it receives either SIGINT or SIGTERM
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)
	<-shutdown
	log.Println("shuting down service")

	return nil
}
