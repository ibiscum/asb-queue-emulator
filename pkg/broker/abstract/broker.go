package abstract

type MQBroker interface {
	SendMessage(queueName string, message []byte) error
	PeekMessage(queueName string) (*Message, error)
	PopMessage(queueName string) (*Message, error)
	DeleteMessage(queueName string, messageId string) (*Message, error)
	CreateQueue(queueName string) error
}
