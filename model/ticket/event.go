package ticket

import eventbus "github.com/aryahmph/skripsi-common/event-bus"

type TicketEvent interface {
	eventbus.Event
	Ticket() Ticket
}

type TicketCreatedEvent struct {
	ticket Ticket
}

func NewTicketCreatedEvent(ticket Ticket) TicketCreatedEvent {
	return TicketCreatedEvent{ticket: ticket}
}

func (e TicketCreatedEvent) Name() string {
	return "event.ticket.created"
}

func (e TicketCreatedEvent) Ticket() Ticket {
	return e.ticket
}
