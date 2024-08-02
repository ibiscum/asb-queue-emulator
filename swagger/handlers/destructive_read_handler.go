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

type DestructiveReadHandler struct {
	HandlerContext utils.HandlerContext
}

// NewDestructiveReadHandler creates a new DestructiveReadHandler
func NewDestructiveReadHandler(handlerContext utils.HandlerContext) operations.DestructiveReadHandler {
	destructiveReadHandler := new(DestructiveReadHandler)
	destructiveReadHandler.HandlerContext = handlerContext
	return destructiveReadHandler
}

func (handler *DestructiveReadHandler) Handle(params operations.DestructiveReadParams) middleware.Responder {
	var message *abstract.Message
	message, err := handler.HandlerContext.MQBroker.PopMessage(params.QueueName)
	if err != nil {
		log.Print(err)
		if strings.Contains(err.Error(), "no message was on the queue") {
			return operations.NewDestructiveReadNoContent()
		}
		return operations.NewDestructiveReadInternalServerError()
	}
	brokerProperties := utils.BrokerProperties{
		MessageId: message.MessageId,
	}
	jsonBrokerProperties, err := json.Marshal(brokerProperties)
	if err != nil {
		log.Println("json marshal error:", err)
		return operations.NewDestructiveReadInternalServerError()
	}
	return operations.NewDestructiveReadOK().WithPayload(string(message.Body)).WithBrokerProperties(string(jsonBrokerProperties))
}