package handlers

import (
	"github.com/koleaby4/next_movie_go/db"
	"github.com/koleaby4/next_movie_go/tmdb"
	"html/template"
	"net/http"
)

// MostPopularMovies handles the most popular movies request
func (h *Handlers) MostPopularMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := tmdb.GetMostPopularMovies(h.Config.TmdbConfig, 8)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("web/templates/movies.html", "web/templates/_navbar.html")
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
		Movies     []db.Movie
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
