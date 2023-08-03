package order

import "time"

type (
	OrderStatus string

	Order struct {
		ID          string
		UserID      string
		TicketID    string
		Amount      int64
		OrderStatus OrderStatus
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	AddOrder struct {
		ID       string
		UserID   string
		TicketID string
		Amount   int64
	}

	UpdateOrderStatus struct {
		ID          string
		OrderStatus OrderStatus
	}
)
