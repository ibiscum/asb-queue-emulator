package rabbitmq

import (
	"asb-queue-emulator/pkg/broker/abstract"
	"asb-queue-emulator/tests/integration/base"
	"asb-queue-emulator/tests/integration/utils"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBrokerTests struct {
	ConnectionString string
	QueueNames       []string
	Broker           abstract.MQBroker
}

var _ base.TestSuite = &RabbitMQBrokerTests{}

func (rmqt *RabbitMQBrokerTests) RunSuite() {
	rmqt.testNonExistantQueue()
	rmqt.testBrokerOperationsSingleQueue()
	rmqt.testBrokerOperationsMultipleQueues()
}

func (rmqt *RabbitMQBrokerTests) BeforeSuite() error {
	conn, err := amqp.Dial(rmqt.ConnectionString)
	if err != nil {
		log.Panicln("error when establishing communication with rabbitmq container, make sure you have it running")
		return err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Panicln("error when creating a communication channel to rabbitmq")
		return err
	}
	for _, queueName := range rmqt.QueueNames {
		_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
		if err != nil {
			log.Panicf("error trying to create queue %s: %s", queueName, err.Error())
		}
	}

	return nil
}

func (rmqt *RabbitMQBrokerTests) AfterSuite() {
	conn, err := amqp.Dial(rmqt.ConnectionString)
	if err != nil {
		log.Panicln("error when establishing communication with rabbitmq container, make sure you have it running")
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Panicln("error when creating a communication channel to rabbitmq")
		return
	}
	for _, queueName := range rmqt.QueueNames {
		_, err = ch.QueueDelete(queueName, false, false, false)
		if err != nil {
			log.Panicf("error trying to delete queue %s: %s", queueName, err.Error())
		}
	}
}

func (rmqt *RabbitMQBrokerTests) testNonExistantQueue() {
	log.Println("Start of non existant queue tests")
	nonExistantQueueName := "the queue that doesn't exists"

	// Put a message in non existant queue
	err := rmqt.Broker.SendMessage(nonExistantQueueName, []byte{})
	utils.AssertIsNotNil(rmqt, err)
	utils.AssertString(rmqt, "Exception (404) Reason: \"NOT_FOUND - no queue 'the queue that doesn't exists' in vhost '/'\"", err.Error())

	// Peek a message in non existant queue
	msg, err := rmqt.Broker.PeekMessage(nonExistantQueueName)
	utils.AssertIsNil(rmqt, msg)
	utils.AssertIsNotNil(rmqt, err)
	utils.AssertString(rmqt, "Exception (404) Reason: \"NOT_FOUND - no queue 'the queue that doesn't exists' in vhost '/'\"", err.Error())

	// Pop a message in non existant queue
	msg, err = rmqt.Broker.PopMessage(nonExistantQueueName)
	utils.AssertIsNil(rmqt, msg)
	utils.AssertIsNotNil(rmqt, err)
	utils.AssertString(rmqt, "Exception (404) Reason: \"NOT_FOUND - no queue 'the queue that doesn't exists' in vhost '/'\"", err.Error())
	log.Println("[PASSED] End of non existant queue tests")
}

func (rmqt *RabbitMQBrokerTests) testBrokerOperationsSingleQueue() {
	log.Println("Start of Normal Operations on single queue test")
	queueName := rmqt.QueueNames[0]
	firstMessage := []byte("Hello")
	secondMessage := []byte("World!")
	// Sending the messages to the queue should result on non errors
	err := rmqt.Broker.SendMessage(queueName, firstMessage)
	utils.AssertIsNil(rmqt, err)
	err = rmqt.Broker.SendMessage(queueName, secondMessage)
	utils.AssertIsNil(rmqt, err)

	// Peeking twice the messages should return the same message twice
	msg, err := rmqt.Broker.PeekMessage(queueName)
	utils.AssertIsNil(rmqt, err)
	utils.AssertEqual(rmqt, msg.Body, firstMessage)
	secondPeek, err := rmqt.Broker.PeekMessage(queueName)
	utils.AssertIsNil(rmqt, err)
	utils.AssertEqual(rmqt, secondPeek.MessageId, msg.MessageId, "Peeking two times without pop should return the same message twice")
	utils.AssertEqual(rmqt, secondPeek.Body, msg.Body)

	// Pop in a message will return it and delete it from the queue
	msg, err = rmqt.Broker.PopMessage(queueName)
	utils.AssertIsNil(rmqt, err)
	utils.AssertEqual(rmqt, msg.Body, secondPeek.Body)

	// Peeking the next message should return the second message
	secondMsg, err := rmqt.Broker.PeekMessage(queueName)
	utils.AssertIsNil(rmqt, err)
	utils.AssertEqual(rmqt, secondMsg.Body, secondMessage)

	// Removing the second message from the queue to empty it
	msg, err = rmqt.Broker.DeleteMessage(queueName, secondMsg.MessageId)
	utils.AssertIsNil(rmqt, err)
	utils.AssertEqual(rmqt, msg.Body, secondMsg.Body)

	// Peeking and pop of the empty queue should return nothing and respective errors
	msg, err = rmqt.Broker.PeekMessage(queueName)
	utils.AssertIsNil(rmqt, msg)
	utils.AssertIsNotNil(rmqt, err)
	utils.AssertString(rmqt, "no message was on the queue", err.Error())
	msg, err = rmqt.Broker.PopMessage(queueName)
	utils.AssertIsNil(rmqt, msg)
	utils.AssertIsNotNil(rmqt, err)
	utils.AssertString(rmqt, "no message was on the queue", err.Error())
	log.Println("[PASSED] End of single queue operations")
}

func (rmqt *RabbitMQBrokerTests) testBrokerOperationsMultipleQueues() {
	log.Println("Start of multiple queue tests")
	firstQueue := rmqt.QueueNames[0]
	secondQueue := rmqt.QueueNames[1]
	firstMessage := []byte("Hello")
	secondMessage := []byte("World!")

	// Sending the messages to the different queues
	err := rmqt.Broker.SendMessage(firstQueue, firstMessage)
	utils.AssertIsNil(rmqt, err)
	err = rmqt.Broker.SendMessage(secondQueue, secondMessage)
	utils.AssertIsNil(rmqt, err)

	// Peeking at different queues should return different messages

	firstQueueMsg, err := rmqt.Broker.PeekMessage(firstQueue)
	utils.AssertIsNil(rmqt, err)
	utils.AssertIsNotNil(rmqt, firstQueueMsg)
	secondQueueMsg, err := rmqt.Broker.PeekMessage(secondQueue)
	utils.AssertIsNil(rmqt, err)
	utils.AssertIsNotNil(rmqt, secondQueueMsg)

	utils.AssertNotEqual(rmqt, firstQueueMsg.Body, secondQueueMsg.Body)

	// Pop at different queues should only affect one of them
	secondQueueMsg, err = rmqt.Broker.PopMessage(secondQueue)
	utils.AssertIsNil(rmqt, err)
	utils.AssertIsNotNil(rmqt, secondQueueMsg)
	firstQueueMsg, err = rmqt.Broker.PeekMessage(firstQueue)
	utils.AssertIsNil(rmqt, err)
	utils.AssertIsNotNil(rmqt, firstQueueMsg)

	// Empty the first queue
	rmqt.Broker.PopMessage(firstQueue)

	log.Println("[PASSED] End of multiple queue tests")
}
