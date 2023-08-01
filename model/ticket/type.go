package ticket

import (
	"database/sql"
	"time"
)

type (
	TicketSeatCategory string

	Ticket struct {
		ID           string
		OrderID      sql.NullString
		Price        int64
		SeatCode     string
		SeatCategory TicketSeatCategory
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	AddTicket struct {
		ID           string
		Price        int64
		SeatCode     string
		SeatCategory TicketSeatCategory
	}

	UpdateTicket = AddTicket

	UpdateTicketOrder struct {
		ID      string
		OrderID string
	}
)
