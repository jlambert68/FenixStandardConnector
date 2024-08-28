package main

import (
	"FenixStandardConnector/executionOrchestrator"
	"FenixStandardConnector/sharedCode"
	"fmt"
	uuidGenerator "github.com/google/uuid"
)

func main() {

	// Create Unique Uuid for run time instance used as identification when communication with GuiExecutionServer
	sharedCode.ApplicationRunTimeUuid = uuidGenerator.New().String()
	fmt.Println("sharedCode.ApplicationRunTimeUuid: " + sharedCode.ApplicationRunTimeUuid)

	// Initiate ExecutionOrchestratorEngine
	executionOrchestrator.InitiateExecutionOrchestratorEngine(
		allowedUsers,
		templateUrlParameters,
		[][]byte{embeddedFile_SubCustody_MainTestDataArea,
			embeddedFile_SubCustody_ExtraTestDataArea,
			embeddedFile_CustodyCash_MainTestDataArea})

	// Keep program running
	var waitChannel chan bool
	waitChannel = make(chan bool)
	<-waitChannel

}
