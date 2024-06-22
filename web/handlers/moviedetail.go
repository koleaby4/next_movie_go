package handlers

import (
	"errors"
	"github.com/jackc/pgx/v5"
	db2 "github.com/koleaby4/next_movie_go/db"
	"github.com/koleaby4/next_movie_go/tmdb"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// MovieDetail handles the movie detail request
func (h *Handlers) MovieDetail(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	idStr := parts[len(parts)-1]
	movieID, err := strconv.Atoi(idStr)

	if err != nil {
		log.Printf("error parsing movie id=%v; err=%v\n", idStr, err)
	}

	conn, ctx := db2.NewConnection(h.AppConfig.DbDsn)
	defer conn.Close(ctx)

	queries := db2.New(conn)

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

	session, err := cookieStore.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		IsLoggedIn bool
		Movie      db2.Movie
	}{
		IsLoggedIn: session.Values["AuthToken"] != nil,
		Movie:      movie,
	}

	tmpl, err := template.ParseFiles("web/templates/movie_detail.html", "web/templates/_watched_info_form.html", "web/templates/_navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println("error executing template", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
