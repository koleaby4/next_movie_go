package main

import (
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/db"
	"github.com/koleaby4/next_movie_go/web/handlers"
	"log"
	"net/http"
)

func main() {
	appConfig, err := config.GetConfig()
	if err != nil {
		log.Fatalln("error reading config file", err)
	}

	conn, ctx := db.NewConnection(appConfig.DbDsn)
	defer conn.Close(ctx)

	queries := db.New(conn)

	h := handlers.New(appConfig, queries)

	http.HandleFunc("/most-popular-movies", h.MostPopularMovies)
	http.HandleFunc("/movies/{movie_id}", h.MovieDetail)
	http.HandleFunc("/movies/{movie_id}/watched_status", h.UpdateWatchedStatus)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.Login(w, r)
		} else if r.Method == http.MethodPost {
			h.LoginPost(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/logout", h.Logout)

	http.ListenAndServe(":8080", nil)
}
