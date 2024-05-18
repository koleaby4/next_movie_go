package main

import (
	"context"
	"fmt"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/db"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"log"
	"strconv"
	"time"
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
	matrixResult, err := queries.InsertMovie(ctx, db.InsertMovieParams{ID: "abc", Title: "Matrix"})
	if err != nil {
		return
	}
	log.Println("movie matrixResult:", matrixResult)

	fmt.Println("finished playWithMoviesTable")
}

func LoadGoodMovies(queries *db.Queries, cfg tmdb.Config, ctx context.Context) {
	var from time.Time
	latestLoadedReleaseDate, err := queries.GetLastKnownReleaseDate(context.Background())
	if err != nil || latestLoadedReleaseDate == nil {
		log.Println("error getting latest inserted date", err)
		from = time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	} else {
		releaseDate, ok := latestLoadedReleaseDate.(string)
		if !ok {
			log.Println("error casting latest inserted date to string", latestLoadedReleaseDate)
		}
		from, err = time.Parse("2006-01-02", releaseDate)
		if err != nil {
			fmt.Println("error parsing date=", err)
		}
	}

	log.Println("latestLoadedReleaseDate", from)

	const minRating = 7.0

	to := from.AddDate(0, 0, 30)

	var counter int
	for counter < 300 {
		enrichedMovies, err := tmdb.GetMoviesReleasedBetween(cfg, from, to, minRating)
		if err != nil {
			log.Fatalln("error getting newest recentMovies", err)
		}

		from = to
		to = from.AddDate(0, 1, 0)

		for _, movie := range enrichedMovies {
			data := db.InsertMovieParams{
				ID:          strconv.Itoa(movie.Id),
				Title:       movie.Title,
				ReleaseDate: movie.ReleaseDate,
				Overview:    movie.Overview,
				Rating:      movie.Rating,
				PosterUrl:   movie.PosterURL,
				TrailerUrl:  movie.TrailerURL,
				RawData:     movie.RawData,
			}

			_, err := queries.InsertMovie(ctx, data)
			if err != nil {
				log.Printf("error persisting movie=%v. err=%v\n", data, err)
			}
			counter++
		}

	}
}

func main() {
	tmdbConfig := tmdb.Config{
		BaseUrl: "https://api.themoviedb.org/3",
		ApiKey:  config.GetTmdbApiKey(),
	}

	appConfig := config.AppConfig{
		Tmdb: tmdbConfig,
	}

	ctx := context.Background()
	conn := db.NewConnection("", ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)

	LoadGoodMovies(queries, appConfig.Tmdb, ctx)
}

func playWithMostPopularMovies(cfg tmdb.Config, minRating float64) {
	mostPopularMovies, err := tmdb.GetMostPopularMovies(cfg, minRating, 1)

	if err != nil {
		log.Fatalln("error getting most popular movies", err)
	}

	fmt.Println("number of most popular movies fetched:", len(mostPopularMovies))
	for _, movie := range mostPopularMovies {
		fmt.Println(movie.Title, movie.Rating, movie.PosterURL, movie.TrailerURL)
	}
}
