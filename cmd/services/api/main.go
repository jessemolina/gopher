package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/jessemolina/gopher/cmd/services/api/handlers"
)

var build = "dev"

type config struct {
	port int
}

func main() {

	// ================================================================
	// RUN

	if err := run(); err != nil {
		log.Println("run error")
		os.Exit(1)
	}
}

func run() error {

	// ================================================================
	// STARTUP

	// Log service start up.
	log.Println("starting service", build)
	defer log.Println("service ended")

	// ================================================================
	// DEBUG API

	// TODO make config dynamic for cli input
	cfg := config{
		port: 4000,
	}

	// ================================================================
	// DEBUG API

	debugMux := handlers.DebugMux()
	addr := ":" + strconv.Itoa(cfg.port)

	if err := http.ListenAndServe(addr, debugMux); err != nil {
		log.Println("error starting debug api")
	}

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
