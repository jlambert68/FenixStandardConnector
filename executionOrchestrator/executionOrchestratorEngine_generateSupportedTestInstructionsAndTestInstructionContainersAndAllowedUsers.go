package executionOrchestrator

import (
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
)

// Generates the 'SupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers' that will be sent via gRPC to Worker
func generateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers() (
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers *TestInstructionAndTestInstuctionContainerTypes.
		TestInstructionsAndTestInstructionsContainersStruct) {

	// Generate the full structure for supported TestInstructions, TestInstructionContainers and Allowed Users
	TestInstructionsAndTesInstructionContainersAndAllowedUsers.
		GenerateTestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard(allowedUsers)

	// Get the full structure for supported TestInstructions, TestInstructionContainers and Allowed Users
	supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers =
		TestInstructionsAndTesInstructionContainersAndAllowedUsers.
			TestInstructionsAndTestInstructionContainersAndAllowedUsers_FenixStandard

	return supportedTestInstructionsAndTestInstructionContainersAndAllowedUsers

}
