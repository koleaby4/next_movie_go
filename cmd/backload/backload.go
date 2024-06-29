package main

import (
	"context"
	"fmt"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/db"
	"github.com/koleaby4/next_movie_go/tmdb"
	"github.com/koleaby4/next_movie_go/web/handlers"
	"log"
	"os"
	"regexp"
	"time"
)

// LoadGoodMovies loads good movies
func LoadGoodMovies(ctx context.Context, queries *db.Queries, cfg config.TmdbConfig) (time.Time, error) {
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

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln("error reading config file", err)
	}

	conn, ctx := db.NewConnection(cfg.DbDsn)
	defer conn.Close(ctx)

	queries := db.New(conn)

	h := handlers.New(cfg, queries)

	watermarkDate, err := LoadGoodMovies(ctx, h.Queries, cfg.TmdbConfig)
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
