package handlers

import (
	"asb-queue-emulator/swagger/gen/restapi/operations"
	"asb-queue-emulator/swagger/utils"

	"github.com/go-openapi/runtime/middleware"
)

type SendMessageHandler struct {
	HandlerContext utils.HandlerContext
}

// NewPeekMessageHandler creates a new PeekMessageHandler
func NewSendMessageHandler(handlerContext utils.HandlerContext) operations.SendMessageHandler {
	sendMessageHandler := new(SendMessageHandler)
	sendMessageHandler.HandlerContext = handlerContext
	return sendMessageHandler
}

func (s *SendMessageHandler) Handle(params operations.SendMessageParams) middleware.Responder {
	byteMessage := []byte(params.MessageContent)
	s.HandlerContext.MQBroker.SendMessage(params.QueueName, byteMessage)
	return operations.NewSendMessageCreated()
}
