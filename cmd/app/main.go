package main

import (
	"context"
	"fmt"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/db"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"log"
)

func playWithUsersTable() {
	fmt.Println("starting playWithUsersTable...")

	ctx := context.Background()
	conn := db.NewConnection("", ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)

	alexResult, err := queries.CreateUser(ctx, "alexResult@gmail.com")
	if err != nil {
		log.Fatalln("error creating user alexResult:", err)
	}
	log.Println("user alexResult:", alexResult)

	fmt.Println("finished playWithUsersTable")

}

func playWithMoviesTable() {
	fmt.Println("starting playWithMoviesTable...")
	ctx := context.Background()
	conn := db.NewConnection("", ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)
	matrixResult, err := queries.CreateMovie(ctx, db.CreateMovieParams{ID: "abc", Title: "Matrix"})
	if err != nil {
		return
	}
	log.Println("movie matrixResult:", matrixResult)

	fmt.Println("finished playWithMoviesTable")

}

func main() {
	tmdbApiKey := config.GetTmdbApiKey()
	fmt.Println("tmdbApiKey", tmdbApiKey)

	recentMovies, err := tmdb.GetRecentMovies(tmdbApiKey, 7)
	enrichedMovies := tmdb.EnrichMovies(tmdbApiKey, recentMovies)
	if err != nil {
		log.Fatalln("error getting newest recentMovies", err)
	}

	fmt.Println("number of recentMovies fetched:", len(recentMovies))

	for _, movie := range enrichedMovies {
		fmt.Println(movie.Title, movie.Rating, movie.TrailerURL)
	}
}
