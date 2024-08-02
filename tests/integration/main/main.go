package main

import (
	"asb-queue-emulator/pkg/broker/rabbitmq"
	"asb-queue-emulator/tests/integration/base"
	rabbitmqtest "asb-queue-emulator/tests/integration/rabbitmq"
	"log"
)

const rabbitmqConnectionString = "amqp://guest:guest@localhost:5672/"

var (
	rmqQueues = []string{"testQueue", "secondTestQueue"}
	tests     = []base.TestSuite{
		&rabbitmqtest.RabbitMQBrokerTests{ConnectionString: rabbitmqConnectionString, QueueNames: rmqQueues, Broker: rabbitmq.NewRabbitMQBroker(rabbitmqConnectionString)},
	}
)

func RunTest(testSuite base.TestSuite) {
	defer testSuite.AfterSuite()
	err := testSuite.BeforeSuite()
	if err != nil {
		log.Panicf("Error running the before suite %s", err.Error())
		return
	}
	testSuite.RunSuite()
}

func main() {
	for _, test := range tests {
		RunTest(test)
	}
}
