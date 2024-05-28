package handlers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/koleaby4/next_movie_go/config"
	"github.com/koleaby4/next_movie_go/internal/db"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Handlers struct {
	AppConfig config.AppConfig
}

func New(cfg config.AppConfig) *Handlers {
	return &Handlers{
		AppConfig: cfg,
	}
}

func (h *Handlers) MovieDetail(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	movieIdStr := parts[len(parts)-1]
	movieID, err := strconv.Atoi(movieIdStr)
	if err != nil {
		log.Printf("error parsing movie id=%v; err=%v\n", movieIdStr, err)
	}

	ctx := context.Background()
	conn := db.NewConnection(h.AppConfig.DbDsn, ctx)
	defer conn.Close(ctx)
	queries := db.New(conn)

	movie, err := queries.GetMovie(ctx, movieID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			movie, err = tmdb.GetMovie(h.AppConfig.TmdbConfig, movieID)
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

	tmpl, err := template.ParseFiles("internal/web/templates/movie_detail.html", "internal/web/templates/_watched_info_form.html", "internal/web/templates/_navbar.html")
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

func (h *Handlers) MostPopularMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := tmdb.GetMostPopularMovies(h.AppConfig.TmdbConfig, 8)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("internal/web/templates/movies.html", "internal/web/templates/_navbar.html")
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
