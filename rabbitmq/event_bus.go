package rabbitmq

import (
	"context"
	"encoding/json"
	eventbus "github.com/aryahmph/skripsi-common/event-bus"
	tModel "github.com/aryahmph/skripsi-common/model/ticket"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type eventBusRabbitMQ struct {
	ch *amqp.Channel
}

var _ eventbus.EventBus = &eventBusRabbitMQ{}

func NewEventBus(ch *amqp.Channel) *eventBusRabbitMQ {
	return &eventBusRabbitMQ{
		ch: ch,
	}
}

// Notify implements eventbus.EventBus.
func (e *eventBusRabbitMQ) Notify(ctx context.Context, event eventbus.Event) (err error) {
	var args []byte

	// publish event map
	switch ev := event.(type) {
	case tModel.TicketCreatedEvent:
		ticket := ev.Ticket()
		if args, err = json.Marshal(Ticket{
			ID:           ticket.ID,
			Price:        ticket.Price,
			SeatCode:     ticket.SeatCode,
			SeatCategory: string(ticket.SeatCategory),
		}); err != nil {
			return err
		}
	}

	return e.ch.PublishWithContext(ctx, ExchangeName, event.Name(), false, false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			MessageId:    uuid.New().String(),
			Timestamp:    time.Now(),
			Body:         args,
		},
	)
}

// Subscribe implements eventbus.EventBus.
func (e *eventBusRabbitMQ) Subscribe(handler eventbus.EventHandler, events ...eventbus.Event) {
	for _, event := range events {
		e.addQueue(event)

		messages, err := e.ch.Consume(event.Name(), "", false, false, false, false, nil)
		if err != nil {
			panic(err)
		}

		go e.worker(messages, handler, event)
	}
}

func (e *eventBusRabbitMQ) worker(messages <-chan amqp.Delivery, handler eventbus.EventHandler, event eventbus.Event) {
	for message := range messages {
		var err error
		var ev eventbus.Event

		// consume event map
		switch event.(type) {

		}

		if err != nil {
			message.Reject(false)
			continue
		}

		err = handler(context.Background(), ev)
		if err != nil {
			message.Reject(true)
			continue
		}

		message.Ack(false)
	}
}

func (e *eventBusRabbitMQ) addQueue(event eventbus.Event) {
	e.ch.ExchangeDeclare(
		ExchangeName,
		amqp.ExchangeTopic,
		true,
		false,
		false,
		false,
		nil,
	)

	queue, err := e.ch.QueueDeclare(event.Name(), true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	if err = e.ch.QueueBind(
		queue.Name,
		event.Name(),
		ExchangeName,
		false,
		nil,
	); err != nil {
		panic(err)
	}
}
