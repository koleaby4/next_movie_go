package main

import (
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/web/handlers"
	"log"
	"net/http"
)

func main() {
	appConfig, err := config.ReadFromFile()
	if err != nil {
		log.Fatalln("error reading config file", err)
	}

	h := handlers.New(appConfig)

	http.HandleFunc("/movies/most-popular", h.MostPopularMovies)
	http.HandleFunc("/movies/{movie_id}", h.MovieDetail)

	http.ListenAndServe(":8080", nil)
}
