package amqp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	amqp "github.com/Azure/go-amqp"
)

// TODO Get definition from Swagger
var sendMessageMethod = http.MethodPut
var sendMessageUrl = "/test"

/*
** Converts an AMQP Message to an HTTP Request as expected by ServiceBus API
 */
func AmqpToHttp(amqpMessage *amqp.Message) (*http.Request, error) {

	// Amqp Properties and Header go into BrokerProperties Header
	// Example:
	// BrokerProperties: {"Label":"M1","State":"Active","TimeToLive":10}
	// Priority: High
	// Customer: 12345,ABC

	headerJson, err := json.Marshal(amqpMessage.Header)
	propJson, err := json.Marshal(amqpMessage.Properties)

	// Both Header and Properties need to be combined into one HTTP Header
	out := map[string]interface{}{}
	json.Unmarshal([]byte(headerJson), &out)
	json.Unmarshal([]byte(propJson), &out)
	BrokerPropertiesJson, _ := json.Marshal(out)

	// Message Data / Body
	ampqData := amqpMessage.GetData()
	body := bytes.NewBuffer(ampqData)

	method := sendMessageMethod
	url := sendMessageUrl

	// Build HTTP Request
	httpMessage, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	httpMessage.Header.Add("BrokerProperties", string(BrokerPropertiesJson))

	return httpMessage, nil
}

// If a batch of messages is sent, these properties are part of the JSON-encoded HTTP body. For more information, see Send Message and Send Message Batch.
// Reference: https://learn.microsoft.com/en-us/rest/api/servicebus/message-headers-and-properties
func AmqpToHttpBatch(amqpMessage *amqp.Message) (*http.Request, error) {
	// TODO Difference in Batch HTTP Request
	return nil, nil // Not implemented
}

/*
** Converts a HTTP Response to an AMPQ Message
 */
func HttpToAmqp(httpResponse *http.Response) (*amqp.Message, error) {
	// TODO implemtation - Headers / Properties
	responseBody := httpResponse.Body
	body, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}

	ampqMessage := amqp.NewMessage(body)
	return ampqMessage, nil
}
