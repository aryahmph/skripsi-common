package pg

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func New(dbURL string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), dbURL)
}
