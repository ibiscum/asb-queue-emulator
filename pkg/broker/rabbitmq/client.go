package rabbitmq

type RabbitMQBroker struct {
	connectionString string
}

func NewRabbitMQBroker(connectionString string) *RabbitMQBroker {
	return &RabbitMQBroker{
		connectionString: connectionString,
	}
}
