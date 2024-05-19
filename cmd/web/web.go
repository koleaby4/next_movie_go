package main

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/db"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"html/template"
	"net/http"
	"strings"
)

func movieDetail(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	movieID := parts[len(parts)-1]

	ctx := context.Background()
	conn := db.NewConnection("", ctx)
	defer conn.Close(ctx)
	queries := db.New(conn)

	movie, err := queries.GetMovie(ctx, movieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			movie, err = tmdb.GetMovie(tmdbConfig, movieID)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = queries.InsertMovie(ctx, movie)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	tmpl, err := template.ParseFiles("internal/web/templates/movie_detail.html", "internal/web/templates/_watched_info_form.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

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
	http.HandleFunc("/movies/{movie_id}", movieDetail)
	http.ListenAndServe(":8080", nil)
}
