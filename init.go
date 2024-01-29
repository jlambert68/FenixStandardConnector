package main

import (
	"FenixStandardConnector/sharedCode"
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

// Extract environment variables used by this Connector-code
func init() {

	// API-key used in TestInstruction 'TestInstruction_Standard_IsDateAPublicHoliday' to call external service
	sharedCode.ApiKey = mustGetenv("ApiKey")

}
