package order

import eventbus "github.com/aryahmph/skripsi-common/event-bus"

type OrderEvent interface {
	eventbus.Event
	Order() Order
}

type OrderCreatedEvent struct {
	order Order
}

func NewOrderCreatedEvent(order Order) OrderCreatedEvent {
	return OrderCreatedEvent{order: order}
}

func (e OrderCreatedEvent) Name() string {
	return "event.order.created"
}

func (e OrderCreatedEvent) Order() Order {
	return e.order
}

type OrderCancelledEvent struct {
	order Order
}

func NewOrderCancelledEvent(order Order) OrderCancelledEvent {
	return OrderCancelledEvent{order: order}
}

func (e OrderCancelledEvent) Name() string {
	return "event.order.cancelled"
}

func (e OrderCancelledEvent) Order() Order {
	return e.order
}
