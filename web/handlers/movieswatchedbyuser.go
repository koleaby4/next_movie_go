package handlers

import (
	"fmt"
	"github.com/koleaby4/next_movie_go/db"
	"net/http"
	"strconv"
	"strings"
)

// UpdateWatchedStatus updates the watched status of a movie for a user
func (h *Handlers) UpdateWatchedStatus(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	if len(pathSegments) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	movieID, err := strconv.Atoi(pathSegments[2]) // Assuming the ID is the second segment
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the user ID and opinion from the form data
	session, err := cookieStore.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the user ID from the session
	userID, ok := session.Values["UserID"].(int)
	if !ok {
		http.Error(w, "User ID not found in session", http.StatusBadRequest)
		fmt.Println("User needs to be logged in to update watched status")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	conn, ctx := db.NewConnection(h.AppConfig.DbDsn)
	defer conn.Close(ctx)

	queries := db.New(conn)

	// Use these values to update the watched status of the movie for the user in the database
	opitions := map[string]int{
		"liked":       3,
		"neutral":     2,
		"disliked":    1,
		"not-watched": 0,
	}

	opinion, ok := opitions[r.FormValue("opinion")]
	if !ok {
		http.Error(w, "Invalid opinion", http.StatusBadRequest)
		return
	}

	if opinion == 0 {
		err = queries.RemoveMovieWatchedByUser(ctx, userID, movieID)
	} else {
		err = queries.UpsertMovieWatchedByUser(ctx, userID, movieID, opinion)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/movies/"+strconv.Itoa(movieID), http.StatusSeeOther)
}
