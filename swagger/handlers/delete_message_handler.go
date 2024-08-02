package handlers

import (
	"asb-queue-emulator/swagger/gen/restapi/operations"
	"asb-queue-emulator/swagger/utils"
	"log"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type DeleteMessageHandler struct {
	HandlerContext utils.HandlerContext
}

// NewPeekMessageHandler creates a new PeekMessageHandler
func NewDeleteMessageHandler(handlerContext utils.HandlerContext) operations.DeleteMessageHandler {
	deleteMessageHandler := new(DeleteMessageHandler)
	deleteMessageHandler.HandlerContext = handlerContext
	return deleteMessageHandler
}

func (s *DeleteMessageHandler) Handle(params operations.DeleteMessageParams) middleware.Responder {
	_, err := s.HandlerContext.MQBroker.DeleteMessage(params.QueueName, params.MessageID)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "message id of the destructive read doesn't match") {
			return operations.NewDeleteMessageNotFound()
		}
		return operations.NewDeleteMessageInternalServerError()
	}
	return operations.NewDeleteMessageOK()
}
