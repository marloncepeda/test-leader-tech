package queues

import (
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func (rmq *RabbitMQ) Connect() error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	rmq.conn = conn

	return nil
}

func (rmq *RabbitMQ) SendMessage(queueName string, message []byte) error {
	channel, err := rmq.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	_, err = channel.QueueDeclare(
		queueName, // Nombre de la cola
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = channel.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}

	return nil
}
