package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

func NewConnection(ctx context.Context, dsn string) *pgx.Conn {
	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalln("error connecting to db:", err)
	}
	return conn
}
