package handlers

import (
	"asb-queue-emulator/pkg/broker/abstract"
	"asb-queue-emulator/swagger/gen/restapi/operations"
	"asb-queue-emulator/swagger/utils"
	"encoding/json"
	"log"
	"strings"

	"github.com/go-openapi/runtime/middleware"
)

type PeekMessageHandler struct {
	HandlerContext utils.HandlerContext
}

// NewPeekMessageHandler creates a new PeekMessageHandler
func NewPeekMessageHandler(handlerContext utils.HandlerContext) operations.PeekMessageHandler {
	peekMessageHandler := new(PeekMessageHandler)
	peekMessageHandler.HandlerContext = handlerContext
	return peekMessageHandler
}

func (s *PeekMessageHandler) Handle(params operations.PeekMessageParams) middleware.Responder {
	var message *abstract.Message
	message, err := s.HandlerContext.MQBroker.PeekMessage(params.QueueName)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "no message was on the queue") {
			return operations.NewPeekMessageNoContent()
		}
		return operations.NewPeekMessageInternalServerError()
	}

	brokerProperties := utils.BrokerProperties{
		MessageId: message.MessageId,
	}
	jsonBrokerProperties, err := json.Marshal(brokerProperties)
	if err != nil {
		log.Println("json marshal error:", err)
		return operations.NewPeekMessageInternalServerError()
	}

	return operations.NewPeekMessageCreated().WithPayload(string(message.Body)).WithBrokerProperties(string(jsonBrokerProperties))
}
