package gue

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/vgarvardt/gue/v5"
	"github.com/vgarvardt/gue/v5/adapter/pgxv4"
)

func New(pool *pgxpool.Pool) *gue.Client {
	client, err := gue.NewClient(pgxv4.NewConnPool(pool))
	if err != nil {
		panic(err)
	}
	return client
}
