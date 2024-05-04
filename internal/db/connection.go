package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/koleaby4/next_movie_go/config"
	"log"
)

func NewConnection(dsn string, ctx context.Context) *pgx.Conn {
	var err error
	dsn, err = config.GetDsn()
	if err != nil {
		log.Fatalln("error getting dsn", err)
	}
	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalln("error connecting to db:", err)
	}
	return conn
}
