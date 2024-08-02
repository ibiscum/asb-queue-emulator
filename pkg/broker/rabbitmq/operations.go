package rabbitmq

import (
	"asb-queue-emulator/pkg/broker/abstract"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	uuid "github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (rmq *RabbitMQBroker) SendMessage(queueName string, message []byte) error {
	conn, err := amqp.Dial(rmq.connectionString)
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to connect to RabbitMQ", err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to open a channel", err)
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclarePassive(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to declare queue", err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	messageId := uuid.New()
	err = ch.PublishWithContext(ctx,
		"",         // exchange
		queue.Name, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   messageId.String(),
			Body:        message,
		})
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to publish a message", err)
		return err
	}
	log.Printf(" [x] Sent %s\n", string(message))
	return nil
}

func (rmq *RabbitMQBroker) read(queueName string, acknowledge bool) (*abstract.Message, error) {
	conn, err := amqp.Dial(rmq.connectionString)
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to dial to rabbitmq", err)
		return nil, err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to open a channel", err)
		return nil, err
	}
	defer ch.Close()
	queue, err := ch.QueueDeclarePassive(queueName, true, false, false, false, nil)
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to declare a queue", err)
		return nil, err
	}
	if queue.Messages > 0 {
		msgs, err := ch.Consume(
			queueName,
			"",
			false, // Manual ack
			false, // non-exclusive
			false, // No local
			false, // no-wait
			nil,
		)
		if err != nil {
			log.Printf("ERROR!: %s: %s\n", "Failed to open a consumer connection", err)
			return nil, err
		}
		for d := range msgs {
			if acknowledge {
				d.Ack(false)
			}
			return &abstract.Message{MessageId: d.MessageId, Body: d.Body}, nil
		}
	}
	return nil, errors.New("no message was on the queue")
}

func (rmq *RabbitMQBroker) PeekMessage(queueName string) (*abstract.Message, error) {
	return rmq.read(queueName, false)
}

func (rmq *RabbitMQBroker) PopMessage(queueName string) (*abstract.Message, error) {
	return rmq.read(queueName, true)
}

func (rmq *RabbitMQBroker) CreateQueue(queueName string) error {
	conn, err := amqp.Dial(rmq.connectionString)
	if err != nil {
		log.Fatalf("FATAL ERROR! %s: %s", "Failed to connect to RabbitMQ when creating the queue", err.Error())
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("FATAL ERROR! Failed to open a channel when creating a queue: %s", err.Error())
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("FATAL ERROR! Failed to declare a queue: %s", err.Error())
		return err
	}

	return err
}


func (rmq *RabbitMQBroker) DeleteMessage(queueName string, messageId string) (*abstract.Message, error) {
	conn, err := amqp.Dial(rmq.connectionString)
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to dial to rabbitmq", err)
		return nil, err
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to open a channel", err)
		return nil, err
	}
	defer ch.Close()
	queue, err := ch.QueueDeclarePassive(queueName, true, false, false, false, nil)
	if err != nil {
		log.Printf("ERROR!: %s: %s\n", "Failed to declare a queue", err)
		return nil, err
	}
	if queue.Messages > 0 {
		msgs, err := ch.Consume(
			queueName,
			"",
			false, // Manual ack
			false, // non-exclusive
			false, // No local
			false, // no-wait
			nil,
		)
		if err != nil {
			log.Printf("ERROR!: %s: %s\n", "Failed to open a consumer connection", err)
			return nil, err
		}
		for d := range msgs {
			if d.MessageId == messageId {
				d.Ack(false)
				return &abstract.Message{MessageId: d.MessageId, Body: d.Body}, nil
			}
		}
	}
	return nil, fmt.Errorf("message of id %s was not on the queue", messageId)
}