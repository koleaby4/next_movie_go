package handlers

import (
	"fmt"
	"github.com/gorilla/sessions"
	db2 "github.com/koleaby4/next_movie_go/db"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

var cookieStore = sessions.NewCookieStore([]byte("DUMMY_SESSION_KEY"))

func hashPassword(password []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(password, 10)
	return bytes, err
}

// LoginPost handles the login POST request
func (h *Handlers) LoginPost(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /login")
	email := r.FormValue("email")
	password := []byte(r.FormValue("password"))

	conn, ctx := db2.NewConnection(h.AppConfig.DbDsn)
	defer conn.Close(ctx)

	queries := db2.New(conn)
	user, err := queries.GetUser(ctx, email)
	if err != nil {
		log.Printf("error (%v) fetching user with email=%v\n", err, email)
	}

	if user.ID == 0 { // user does not exist
		log.Println("User does not exist", user)
		hashedPassword, err := hashPassword(password)
		if err != nil {
			log.Println("Error hashing password", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		user, err = queries.UpsertUser(ctx, email, string(hashedPassword))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	log.Println("User exists", user)
	err = bcrypt.CompareHashAndPassword([]byte(user.AuthToken), password)
	if err != nil { // user exists, but password does not match
		fmt.Println("Passwords do not match")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	fmt.Println("Passwords match")

	session, err := cookieStore.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["UserID"] = user.ID
	session.Values["AuthToken"] = user.AuthToken

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/most-popular-movies", http.StatusSeeOther)
}

func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /login")
	tmpl, err := template.ParseFiles("web/templates/login.html", "web/templates/_navbar.html")
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

func (h *Handlers) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := cookieStore.Get(r, "user-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	delete(session.Values, "AuthToken")

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
