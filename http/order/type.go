package order

type (
	AddOrder struct {
		ID       string `json:"id" validate:"required"`
		UserID   string `json:"user_id" validate:"required"`
		TicketID string `json:"ticket_id" validate:"required"`
		Amount   int64  `json:"amount" validate:"required"`
	}
)
