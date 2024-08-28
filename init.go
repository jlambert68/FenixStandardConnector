package main

import (
	"github.com/jlambert68/FenixSyncShared/environmentVariables"
	"log"
	"os"
)

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func init() {

	// Initiate object that extract Environment Variables or Injected Environment Variables
	environmentVariables.InitiateInjectedVariablesMap(&injectedVariablesMap)
}
