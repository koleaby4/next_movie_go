package handlers

import (
	"github.com/koleaby4/next_movie_go/internal/models"
	"github.com/koleaby4/next_movie_go/internal/tmdb"
	"html/template"
	"net/http"
)

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

	session, err := cookieStore.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		IsLoggedIn bool
		Movies     []models.Movie
	}{
		IsLoggedIn: session.Values["AuthToken"] != nil,
		Movies:     movies,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
