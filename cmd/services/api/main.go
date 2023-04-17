package main

import (
	"log"
)

var build = "dev"

func main() {

	// TODO start a basic logger
	// ================================================================
	// LOGGER
	log.Println("starting service", build)
	defer log.Println("service ended")

	// TODO create a service shutdown channel
	// ================================================================
	// SHUTDOWN
}

// TODO build a run function to load handlers and start server mux
