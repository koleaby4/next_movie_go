package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/db"
	"log"
)

func main() {
	fmt.Println("starting...")

	dsn, err := config.GetDsn()
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)

	if err != nil {
		log.Fatalln("error connecting to db:", err)
	}
	defer conn.Close(ctx)

	queries := db.New(conn)
	if err != nil {
		log.Fatalln("error fetching dsn:", err)
	}

	alexResult, err := queries.CreateUser(ctx, "alexResult@gmail.com")
	if err != nil {
		log.Fatalln("error creating user alexResult:", err)
	}
	log.Println("user alexResult:", alexResult)

	matrixResult, err := queries.CreateMovie(ctx, db.CreateMovieParams{ID: "abc", Title: "Matrix"})
	if err != nil {
		return
	}
	log.Println("movie matrixResult:", matrixResult)

	fmt.Println("Done.")
}
