package brokerutils

import (
	"asb-queue-emulator/pkg/broker/abstract"
	"asb-queue-emulator/pkg/broker/rabbitmq"
	"errors"

	"github.com/spf13/viper"
)

func GetBroker() (abstract.MQBroker, error) {
	switch abstract.BrokerType(viper.GetString("broker.type")) {
	case abstract.RabbitMQ:
		return rabbitmq.NewRabbitMQBroker(
			viper.GetString("broker.connectionString"),
		), nil
	default:
		return nil, errors.New("failed to create MQBroker")
	}
}

func CreateQueues(broker abstract.MQBroker) error {
	queues := viper.GetStringSlice("broker.queues")

	for i := range queues {
		err := broker.CreateQueue(queues[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func GetServerPort() int {
	return viper.GetInt("serverPort")
}
