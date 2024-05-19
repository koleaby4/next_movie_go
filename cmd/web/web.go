package main

import (
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"html/template"
	"net/http"
)

func mostPopularMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := tmdb.GetMostPopularMovies(tmdbConfig, 8)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("internal/web/templates/movies.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var tmdbConfig = tmdb.Config{
	BaseUrl: "https://api.themoviedb.org/3",
	ApiKey:  config.GetTmdbApiKey(),
}

func main() {
	http.HandleFunc("/movies/most-popular", mostPopularMovies)
	http.ListenAndServe(":8080", nil)
}
