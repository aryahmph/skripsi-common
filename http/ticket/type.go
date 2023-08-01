package ticket

import (
	"time"
)

type (
	Ticket struct {
		ID           string    `json:"id"`
		OrderID      *string   `json:"order_id"`
		Price        int64     `json:"price"`
		SeatCode     string    `json:"seat_code"`
		SeatCategory string    `json:"seat_category"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
