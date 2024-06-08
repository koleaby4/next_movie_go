package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func NewConnection(dsn string, ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalln("error connecting to db:", err)
	}
	return conn
}
