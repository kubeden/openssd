package main

import (
	"log"

	"github.com/kubeden/openssd/go/api"
	"github.com/kubeden/openssd/go/client"
	"github.com/kubeden/openssd/types"
)

func main() {
	config := types.Config{
		GithubUsername: "kubeden",
		GithubRepo:     "kubeden",
		ReadmeFile:     "README.md",
		InfoFile:       "INFO.md",
		XUserFullName:  "Kuberdenis",
		XUserName:      "kubeden",
		TemplateChoice: "default",
	}

	// Start API server
	go func() {
		if err := api.StartServer(config); err != nil {
			log.Fatalf("API server failed to start: %v", err)
		}
	}()

	// Start client server
	if err := client.StartServer(config); err != nil {
		log.Fatalf("Client server failed to start: %v", err)
	}
}
