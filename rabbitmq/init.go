package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	routes = map[string]string{
		"event.order.created": "create_orders",
	}
)

func New(rabbitmqURL string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(rabbitmqURL)
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	err = initExchange(ch)
	if err != nil {
		panic(err)
	}

	err = initQueue(ch)
	if err != nil {
		panic(err)
	}

	return conn, ch
}

func initExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		ExchangeName,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)
}

func initQueue(ch *amqp.Channel) error {
	for key, val := range routes {
		queue, err := ch.QueueDeclare(val, true, false, false, false, nil)
		if err != nil {
			return err
		}

		err = ch.QueueBind(
			queue.Name,
			key,
			ExchangeName,
			false,
			nil,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
