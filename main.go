package main

import (
	"log"
	"os"

	"github.com/kubeden/openssd/go/api"
	"github.com/kubeden/openssd/go/client"
	"github.com/kubeden/openssd/types"
)

func main() {
	config := types.Config{
		GithubUsername: getEnv("GITHUB_USERNAME", "kubeden"),
		GithubRepo:     getEnv("GITHUB_REPO", "kubeden"),
		ReadmeFile:     getEnv("README_FILE", "README.md"),
		InfoFile:       getEnv("INFO_FILE", "INFO.md"),
		XUserFullName:  getEnv("X_USER_FULL_NAME", "Kuberdenis"),
		XUserName:      getEnv("X_USERNAME", "kubeden"),
		TemplateChoice: getEnv("TEMPLATE_CHOICE", "ssi"),
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
