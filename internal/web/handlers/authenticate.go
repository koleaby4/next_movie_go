package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/koleaby4/next_movie_go/internal/db"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var cookieStore = sessions.NewCookieStore([]byte("DUMMY_SESSION_KEY"))

func hashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 10)
	return bytes, err
}

func (h *Handlers) LoginPost(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /login")
	email := r.FormValue("email")
	password := []byte(r.FormValue("password"))

	ctx := context.Background()
	conn := db.NewConnection(h.AppConfig.DbDsn, ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)
	user, err := queries.GetUser(ctx, email)

	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Println("Error hashing password", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if user.ID == 0 { // user does not exist
		log.Println("User does not exist", user)
		user, err = queries.UpsertUser(ctx, db.User{Email: email, AuthToken: string(hashedPassword)})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		log.Println("User exists", user)
		err = bcrypt.CompareHashAndPassword(hashedPassword, password)
		if err != nil { // user exists, but password does not match
			fmt.Println("Passwords do not match")
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		} else {
			fmt.Println("Passwords match")
		}
	}

	session, err := cookieStore.Get(r, strconv.Itoa(user.ID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["AuthToken"] = hashedPassword
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/most-popular-movies", http.StatusSeeOther)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /login")
	tmpl, err := template.ParseFiles("internal/web/templates/login.html", "internal/web/templates/_navbar.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := cookieStore.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]bool{
		"IsLoggedIn": session.Values["AuthToken"] != nil,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
