package main

import (
	"log"

	"github.com/kubeden/openssd/go/api"
	"github.com/kubeden/openssd/go/client"
)

func main() {
	// Start API server
	go func() {
		if err := api.StartServer(); err != nil {
			log.Fatalf("API server failed to start: %v", err)
		}
	}()

	// Start client server
	if err := client.StartServer(); err != nil {
		log.Fatalf("Client server failed to start: %v", err)
	}
}
