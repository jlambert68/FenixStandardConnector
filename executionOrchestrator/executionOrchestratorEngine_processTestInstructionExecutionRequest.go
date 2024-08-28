package executionOrchestrator

import (
	"FenixStandardConnector/sharedCode"
	fenixExecutionWorkerGrpcApi "github.com/jlambert68/FenixGrpcApi/FenixExecutionServer/fenixExecutionWorkerGrpcApi/go_grpc_api"
	"github.com/jlambert68/FenixTestInstructionsAdminShared/TypeAndStructs"
	"github.com/sirupsen/logrus"
)

func processTestInstructionExecutionRequest(
	testInstructionExecutionPubSubRequest *fenixExecutionWorkerGrpcApi.ProcessTestInstructionExecutionPubSubRequest) (
	testInstructionExecutionResultMessage *fenixExecutionWorkerGrpcApi.FinalTestInstructionExecutionResultMessage,
	err error) {

	// Depending on TestInstruction then choose how to execution the TestInstruction
	switch TypeAndStructs.OriginalElementUUIDType(testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid) {

	// When there are TestInstruction to be processed, the code will be here

	default:
		sharedCode.Logger.WithFields(logrus.Fields{
			"id": "83c4d742-812e-4598-8a39-feabff216e11",
			"testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionUuid": testInstructionExecutionPubSubRequest.TestInstruction.TestInstructionOriginalUuid,
		}).Fatal("Unknown TestInstruction Uuid")
	}

	return testInstructionExecutionResultMessage, err
}
