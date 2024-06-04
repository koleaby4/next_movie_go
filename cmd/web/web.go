package main

import (
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/web/handlers"
	"log"
	"net/http"
)

func main() {
	appConfig, err := config.ReadFromFile("config/.env.json")
	if err != nil {
		log.Fatalln("error reading config file", err)
	}

	h := handlers.New(appConfig)

	http.HandleFunc("/movies/popular", h.PopularMovies)
	http.HandleFunc("/movies/{movie_id}", h.MovieDetail)

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.Login(w, r)
		} else if r.Method == http.MethodPost {
			h.LoginPost(w, r)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
