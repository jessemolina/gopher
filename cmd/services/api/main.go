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
	apiPort   int
	debugPort int
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
	// CONFIGUARTION

	// TODO make config dynamic for cli input
	cfg := config{
		apiPort:   4000,
		debugPort: 3000,
	}

	// ================================================================
	// DEBUG API

	debugMux := handlers.DebugMux()
	debugAddr := ":" + strconv.Itoa(cfg.debugPort)

	// run debug on different goroutine
	go func() {
		if err := http.ListenAndServe(debugAddr, debugMux); err != nil {
			log.Println("error starting debug api", err)
		}

	}()

	// ================================================================
	// SERVICE API

	// TODO listen and server service api
	apiMux := handlers.APIMux()
	apiAddr := ":" + strconv.Itoa(cfg.apiPort)

	// run service api on different goroutine
	go func() {
		if err := http.ListenAndServe(apiAddr, apiMux); err != nil {
			log.Println("error starting service api", err)
		}

	}()

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
