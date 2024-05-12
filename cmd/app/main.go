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

func playWithMostRecentMovies(cfg tmdb.Config, minRating float64) {
	recentMovies, err := tmdb.GetRecentMovies(cfg, minRating, 1)
	enrichedMovies := tmdb.EnrichMoviesInfo(cfg, recentMovies)
	if err != nil {
		log.Fatalln("error getting newest recentMovies", err)
	}

	fmt.Println("number of recentMovies fetched:", len(recentMovies))

	for _, movie := range enrichedMovies {
		fmt.Println(movie.Title, movie.Rating, movie.PosterURL, movie.TrailerURL)
	}

}

func main() {
	tmdbConfig := tmdb.Config{
		BaseUrl: "https://api.themoviedb.org",
		ApiKey:  config.GetTmdbApiKey(),
	}

	appConfig := config.AppConfig{
		Tmdb: tmdbConfig,
	}
	fmt.Println("tmdbApiKey", appConfig.Tmdb.ApiKey)

	minRating := 7.0

	playWithMostRecentMovies(appConfig.Tmdb, minRating)
	fmt.Println("=====================================")
	playWithMostPopularMovies(appConfig.Tmdb, minRating)
}

func playWithMostPopularMovies(cfg tmdb.Config, minRating float64) {
	mostPopularMovies, err := tmdb.GetMostPopularMovies(cfg, minRating, 1)
	mostPopularMovies = tmdb.EnrichMoviesInfo(cfg, mostPopularMovies)

	if err != nil {
		log.Fatalln("error getting most popular movies", err)
	}

	fmt.Println("number of most popular movies fetched:", len(mostPopularMovies))
	for _, movie := range mostPopularMovies {
		fmt.Println(movie.Title, movie.Rating, movie.PosterURL, movie.TrailerURL)
	}
}
