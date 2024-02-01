package main

import (
	"FenixStandardConnector/executionOrchestrator"
)

func main() {

	// Initiate ExecutionOrchestratorEngine
	executionOrchestrator.InitiateExecutionOrchestratorEngine(allowedUsers)

	// Keep program running
	var waitChannel chan bool
	waitChannel = make(chan bool)
	<-waitChannel

}
