package amqp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/Azure/go-amqp"
	"github.com/stretchr/testify/require"
)

// TODO Validate against real raw example
var exampleBrokerProperties = "{\"AbsoluteExpiryTime\":null,\"ContentEncoding\":null,\"ContentType\":null,\"CorrelationID\":\"0000\",\"CreationTime\":null,\"DeliveryCount\":5,\"Durable\":false,\"FirstAcquirer\":false,\"GroupID\":null,\"GroupSequence\":null,\"MessageID\":\"12345\",\"Priority\":1,\"ReplyTo\":null,\"ReplyToGroupID\":null,\"Subject\":null,\"TTL\":10000000000,\"To\":null,\"UserID\":null}"

func TestAmqpToHttp(t *testing.T) {
	// Set up object coming from Azure SDK
	amqpBody := "This is a test body"
	data, err := json.Marshal(amqpBody)
	require.Empty(t, err)
	incomingAmpqMessage := amqp.NewMessage(data)
	incomingAmpqMessage.Header = &amqp.MessageHeader{
		TTL:           10 * time.Second,
		DeliveryCount: 5,
		Priority:      1,
	}
	incomingAmpqMessage.Properties = &amqp.MessageProperties{
		MessageID:     "12345",
		CorrelationID: "0000",
	}

	// Set up expected object
	// TODO defined in swagger
	method := http.MethodPut
	url := "/test"
	httpBody := bytes.NewBuffer(data)
	expectedHttpMessage, err := http.NewRequest(method, url, httpBody)
	expectedHttpMessage.Header.Add("BrokerProperties", exampleBrokerProperties)
	require.Empty(t, err)

	// Make request and validate
	outputHttpMessage, err := AmqpToHttp(incomingAmpqMessage)
	require.Empty(t, err)

	require.Equal(t, expectedHttpMessage.Method, outputHttpMessage.Method, "Method not expected")
	require.Equal(t, expectedHttpMessage.URL, outputHttpMessage.URL, "URL not expected")
	require.Equal(t, expectedHttpMessage.Body, outputHttpMessage.Body, "Body not expected")
	require.Equal(t, expectedHttpMessage.Header, outputHttpMessage.Header, "Header not expected")
}

func TestHttpToAmqp(t *testing.T) {
	// TODO Testing of headers
	// Set up object response
	amqpBody := "This is a test body"
	data, err := json.Marshal(amqpBody)
	require.Empty(t, err)

	httpBody := io.NopCloser(bytes.NewBuffer(data))

	incomingHttpResponse := &http.Response{}
	incomingHttpResponse.Body = httpBody

	// Set up expected object
	expectedAmpqMessage := amqp.NewMessage(data)

	// Make request and validate
	outputAmpqMessage, err := HttpToAmqp(incomingHttpResponse)
	require.Empty(t, err)

	require.Equal(t, expectedAmpqMessage.Header, outputAmpqMessage.Header, "Header not expected")
	require.Equal(t, expectedAmpqMessage.Data, outputAmpqMessage.Data, "Data not expected")
}
