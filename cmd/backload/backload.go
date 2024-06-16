package main

import (
	"context"
	"fmt"
	"github.com/koleaby4/next_movie_go"
	db2 "github.com/koleaby4/next_movie_go/db"
	"github.com/koleaby4/next_movie_go/tmdb"
	"log"
	"os"
	"regexp"
	"time"
)

//func playWithUsersTable(dsn string) {
//	fmt.Println("starting playWithUsersTable...")
//
//	ctx := context.Background()
//	conn := db.NewConnection(dsn, ctx)
//	defer conn.Close(ctx)
//
//	queries := db.New(conn)
//
//	alexResult, err := queries.CreateUser(ctx, "alexResult@gmail.com")
//	if err != nil {
//		log.Fatalln("error creating user alexResult:", err)
//	}
//	log.Println("user alexResult:", alexResult)
//
//	fmt.Println("finished playWithUsersTable")
//
//}

//func playWithMoviesTable(dsn string) {
//	fmt.Println("starting playWithMoviesTable...")
//	ctx := context.Background()
//	conn := db2.NewConnection(dsn, ctx)
//	defer conn.Close(ctx)
//
//	queries := db2.New(conn)
//	matrixResult, err := queries.InsertMovie(ctx, db2.Movie{ID: 123, Title: "Matrix"})
//	if err != nil {
//		return
//	}
//	log.Println("movie matrixResult:", matrixResult)
//
//	fmt.Println("finished playWithMoviesTable")
//}

// LoadGoodMovies loads good movies
func LoadGoodMovies(ctx context.Context, queries *db2.Queries, cfg next_movie_go.TmdbConfig) (time.Time, error) {
	from, err := time.Parse("2006-01-02", cfg.BackloadHighWatermarkDate)
	if err != nil {
		return time.Time{}, err
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

	return to, nil
}

func main() {

	appConfig, err := next_movie_go.GetAppConfig()

	if err != nil {
		log.Fatalln("error reading config file", err)
	}

	conn, ctx := db2.NewConnection(appConfig.DbDsn)
	defer conn.Close(ctx)

	queries := db2.New(conn)

	watermarkDate, err := LoadGoodMovies(ctx, queries, appConfig.TmdbConfig)
	if err != nil {
		log.Fatalln("error in LoadGoodMovies", err)
	}
	err = updateHighWatermark(watermarkDate, ".env")
	if err != nil {
		log.Fatalln("error updating high watermark", err)
	}

	fmt.Println("finished backload")

}

func updateHighWatermark(newWatermark time.Time, configPath string) error {
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	newWatermarkStr := "TMDB_BACKLOAD_HIGH_WATERMARK_DATE=" + newWatermark.Format("2006-01-02")

	re := regexp.MustCompile(`TMDB_BACKLOAD_HIGH_WATERMARK_DATE=\d{4}-\d{2}-\d{2}`)
	updatedContent := re.ReplaceAllString(string(data), newWatermarkStr)

	err = os.WriteFile(configPath, []byte(updatedContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

//func playWithMostPopularMovies(cfg config.TmdbConfig, minRating float64) {
//	mostPopularMovies, err := tmdb.GetMostPopularMovies(cfg, minRating)
//
//	if err != nil {
//		log.Fatalln("error getting most popular movies", err)
//	}
//
//	fmt.Println("number of most popular movies fetched:", len(mostPopularMovies))
//	for _, movie := range mostPopularMovies {
//		fmt.Println(movie.Title, movie.Rating, movie.PosterUrl, movie.TrailerUrl)
//	}
//}
