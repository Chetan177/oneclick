package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/chetan177/oneclick/rest"
)

func main() {
	server := rest.NewServer()
	server.Start()

	// Create a channel to receive OS signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-sigChan

	// Perform any cleanup or shutdown tasks here
	server.Stop()
}
