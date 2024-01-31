package executionOrchestrator

import (
	"FenixStandardConnector/sharedCode"
	"errors"
	"fmt"
	fenixConnectorAdminShared_sharedCode "github.com/jlambert68/FenixConnectorAdminShared/common_config"
	"github.com/jlambert68/FenixConnectorAdminShared/fenixConnectorAdminShared"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TestInstructionAndTestInstuctionContainerTypes"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"time"
)

// Initiate Call-Back-struct and initiate
var connectorFunctionsToDoCallBackOn *fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct

var allowedUsers []byte

func InitiateExecutionOrchestratorEngine(tempAllowedUsers []byte) {

	allowedUsers = tempAllowedUsers

	connectorFunctionsToDoCallBackOn = &fenixConnectorAdminShared_sharedCode.ConnectorCallBackFunctionsStruct{
		GetMaxExpectedFinishedTimeStamp:        getMaxExpectedFinishedTimeStamp,
		ProcessTestInstructionExecutionRequest: processTestInstructionExecutionRequest,
		InitiateLogger:                         initiateLogger,
		GenerateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers: generateSupportedTestInstructionsAndTestInstructionContainersAndAllowedUsers,
	}
	fenixConnectorAdminShared.InitiateFenixConnectorAdminShared(connectorFunctionsToDoCallBackOn)

}

func getMaxExpectedFinishedTimeStamp(testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	maxExpectedFinishedTimeStamp time.Time,
	err error) {

	var expectedExecutionDuration time.Duration

	// Depending on TestInstruction calculate or set 'MaxExpectedFinishedTimeStamp'
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid) {

	// When there are TestInstruction to be processed, the code will be here

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4192bcf6-09f7-4ee7-ad3d-4640bca4b2ba",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid,
		}).Error("Unknown TestInstruction Uuid")

		err = errors.New(fmt.Sprintf("Unknown TestInstruction Uuid: %s",
			testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid))

		expectedExecutionDuration = 0 * time.Minute
		maxExpectedFinishedTimeStamp = time.Now().Add(expectedExecutionDuration)
	}

	return maxExpectedFinishedTimeStamp, err
}

// Initiate logger with same logger as Shared Connector code uses
func initiateLogger(logger *logrus.Logger) {
	sharedCode.Logger = logger
}

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

func processTestInstructionExecutionRequest(
	testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	testInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage,
	err error) {

	// Depending on TestInstruction then choose how to execution the TestInstruction
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid) {

	// When there are TestInstruction to be processed, the code will be here

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "83c4d742-812e-4598-8a39-feabff216e11",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid,
		}).Fatal("Unknown TestInstruction Uuid")
	}

	return testInstructionExecutionResultMessage, err
}
