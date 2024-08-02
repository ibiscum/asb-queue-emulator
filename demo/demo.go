// Go demo showing off an insertion and read of Azure Service Bus Queue through the Azure SDK for Go.
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

// We'll want the functions InsertMessage and ReadMessage to be able to access the queueName and connectionString and Main will be the entry point for the demo.

// InsertMessage inserts a message into the queue.
func InsertMessage(client *azservicebus.Client, queueName string) {

	// Create a Service Bus sender
	fmt.Println("Creating Service Bus sender")
	sender, err := client.NewSender(queueName, nil)
	if err != nil {
		fmt.Println("Failed to create Service Bus sender")
		os.Exit(1)
	}

	// Create a message
	message := &azservicebus.Message{
		Body: []byte("Hello, World!"),
	}

	// Send the message
	fmt.Println("Sending message")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err = sender.SendMessage(ctx, message, nil)
	if err != nil {
		fmt.Printf("Failed to send message: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Sent message")
}

// ReadMessage reads a message from the queue.
func ReadMessage(client *azservicebus.Client, queueName string) {
	// Create a Service Bus receiver
	fmt.Println("Creating Service Bus receiver")
	receiver, err := client.NewReceiverForQueue(queueName, nil)
	if err != nil {
		fmt.Println("Failed to create Service Bus receiver")
		os.Exit(1)
	}

	// Receive a message
	fmt.Println("Receiving message")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	message, err := receiver.PeekMessages(ctx, 1, nil)
	if err != nil {
		fmt.Printf("Failed to receive message: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Received message")
	// Convert the RawAMQPMessage to a Human Readable format
	if valueBytes, ok := message[0].RawAMQPMessage.Body.Value.([]byte); ok {
		fmt.Printf("Received message: %s\n", string(valueBytes))
	} else {
		fmt.Println("Failed to convert message to human readable format")
	}
}

// NewDefaultAzureCredential creates a new default Azure credential.
func NewDefaultAzureCredential() (*azidentity.DefaultAzureCredential, error) {
	return azidentity.NewDefaultAzureCredential(nil)
}

// Main is the entry point for the demo.
func main() {
	queueName := "test1"
	if queueName == "" {
		fmt.Println("SERVICE_BUS_QUEUE_NAME must be set")
		os.Exit(1)
	}

	connectionString := "Endpoint=sb://localhost;SharedAccessKeyName=YourKeyName;SharedAccessKey=YourAccessKey" //os.Getenv("SERVICE_BUS_CONNECTION_STRING")

	// Create a client options struct with disabled SSL certificate validation
	clientOptions := &azservicebus.ClientOptions{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client, err := azservicebus.NewClientFromConnectionString(connectionString, clientOptions)
	if err != nil {
		fmt.Printf("Failed to create Service Bus client: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Inserting message")
	InsertMessage(client, queueName)

	fmt.Println("Reading message")
	ReadMessage(client, queueName)
}
