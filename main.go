package main

import (
	"log"

	"github.com/kubeden/openssd/go/api"
	"github.com/kubeden/openssd/go/client"
)

func main() {
	// Start API server
	go func() {
		log.Println("Starting API server on :8081")
		if err := api.StartServer(); err != nil {
			log.Fatalf("API server failed to start: %v", err)
		}
	}()

	// Start client server
	log.Println("Starting client server on :8080")
	if err := client.StartServer(); err != nil {
		log.Fatalf("Client server failed to start: %v", err)
	}
}
