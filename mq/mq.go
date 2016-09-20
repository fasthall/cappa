package mq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type MQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewMQ() (*MQ, error) {
	mq := new(MQ)
	fmt.Println("Connecting to RabbitMQ...")
	var err error
	mq.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	mq.ch, err = mq.conn.Channel()
	if err != nil {
		return nil, err
	}
	return mq, nil
}

func (mq *MQ) Listen() error {
	q, err := mq.ch.QueueDeclare(
		"cappa",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := mq.ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		fmt.Printf("Received a message: %s\n", d.Body)
		d.Ack(false)
	}

	return nil
}
