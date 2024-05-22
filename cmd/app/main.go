package main

import (
	"context"
	"fmt"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/db"
	"github.com/koleaby4/next_movie_go/internal/models"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"log"
	"time"
)

func playWithUsersTable(dsn string) {
	fmt.Println("starting playWithUsersTable...")

	ctx := context.Background()
	conn := db.NewConnection(dsn, ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)

	alexResult, err := queries.CreateUser(ctx, "alexResult@gmail.com")
	if err != nil {
		log.Fatalln("error creating user alexResult:", err)
	}
	log.Println("user alexResult:", alexResult)

	fmt.Println("finished playWithUsersTable")

}

func playWithMoviesTable(dsn string) {
	fmt.Println("starting playWithMoviesTable...")
	ctx := context.Background()
	conn := db.NewConnection(dsn, ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)
	matrixResult, err := queries.InsertMovie(ctx, models.Movie{ID: 123, Title: "Matrix"})
	if err != nil {
		return
	}
	log.Println("movie matrixResult:", matrixResult)

	fmt.Println("finished playWithMoviesTable")
}

func LoadGoodMovies(queries *db.Queries, cfg config.TmdbConfig, ctx context.Context) {
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

			_, err := queries.InsertMovie(ctx, movie)
			if err != nil {
				log.Printf("error persisting movie=%v. err=%v\n", movie, err)
			}
			counter++
		}

	}
}

func main() {

	appConfig, err := config.ReadFromFile()

	if err != nil {
		log.Fatalln("error reading config file", err)
	}

	ctx := context.Background()
	conn := db.NewConnection(appConfig.DbDsn, ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)

	LoadGoodMovies(queries, appConfig.TmdbConfig, ctx)
}

func playWithMostPopularMovies(cfg config.TmdbConfig, minRating float64) {
	mostPopularMovies, err := tmdb.GetMostPopularMovies(cfg, minRating)

	if err != nil {
		log.Fatalln("error getting most popular movies", err)
	}

	fmt.Println("number of most popular movies fetched:", len(mostPopularMovies))
	for _, movie := range mostPopularMovies {
		fmt.Println(movie.Title, movie.Rating, movie.PosterUrl, movie.TrailerUrl)
	}
}
