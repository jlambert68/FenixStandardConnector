package executionOrchestrator

import (
	"FenixStandardConnector/sharedCode"
	"fmt"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain"
	TestInstruction_SendTestDataToThisDomain_version1_0 "github.com/jlambert68/FenixStandardTestInstructionAdmin/TestInstructionsAndTesInstructionContainersAndAllowedUsers/TestInstructions/TestInstruction_SendTestDataToThisDomain/version_1_0"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func getMaxExpectedFinishedTimeStamp(testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.
	ProcessTestInstructionExecutionPubSubRequest) (
	maxExpectedFinishedTimeStamp time.Time,
	err error) {

	// expectedExecutionDurationInSeconds is extracted from TestInstruction-data
	var expectedExecutionDurationInSeconds int64

	// Create Version for TestInstruction
	var version string
	version = fmt.Sprintf("v%s.%s",
		strconv.Itoa(int(testInstructionExecutionPubSubRequest.TestInstruction.GetMajorVersionNumber())),
		strconv.Itoa(int(testInstructionExecutionPubSubRequest.TestInstruction.GetMinorVersionNumber())))

	// Depending on TestInstruction calculate or set 'MaxExpectedFinishedTimeStamp'
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {

	case TestInstruction_Standard_SendTestDataToThisDomain.TestInstructionUUID_FenixSentToUsersDomain_SendTestDataToThisDomain:

		// Extract execution duration based on TestInstruction version
		switch version {
		case "v1.0":

			// Extract duration
			expectedExecutionDurationInSeconds = TestInstruction_SendTestDataToThisDomain_version1_0.
				ExpectedMaxTestInstructionExecutionDurationInSeconds

			// Create Max Finished TimeStamp
			maxExpectedFinishedTimeStamp = time.Now().Add(time.Duration(expectedExecutionDurationInSeconds) * time.Second)

		default:
			sharedCode.Logger.WithFields(logrus.Fields{
				"id": "94843bac-ae44-482a-9dff-981319534f28",
				"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionOriginalUuid,
				"TestInstructionName": testInstructionExecutionPubSubRequest.TestInstruction.
					TestInstructionName,
				"version": version,
			}).Fatalln("Unhandled version, will exit")

		}

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "4d073695-8990-42c0-9347-9a66ef4133af",
			"TestInstructionOriginalUuid": testInstructionExecutionPubSubRequest.
				TestInstruction.TestInstructionOriginalUuid,
			"TestInstructionName": testInstructionExecutionPubSubRequest.
				TestInstruction.TestInstructionName,
		}).Fatalln("Unknown TestInstruction Uuid, will Exit")

	}

	return maxExpectedFinishedTimeStamp, err
}
