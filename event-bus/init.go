package eventbus

import "context"

type DomainEventBus struct {
	handlers map[string][]EventHandler
}

func (e DomainEventBus) Subscribe(handler EventHandler, events ...Event) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

func (e DomainEventBus) Notify(ctx context.Context, event Event) error {
	for _, handler := range e.handlers[event.Name()] {
		if err := handler(ctx, event); err != nil {
			return err
		}
	}
	return nil
}
