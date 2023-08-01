package eventbus

import "context"

type (
	Event interface {
		Name() string
	}

	EventBus interface {
		Subscribe(handler EventHandler, events ...Event)
		Notify(ctx context.Context, event Event) error
	}
)

type EventHandler func(ctx context.Context, event Event) error
