package handlers

import (
	"context"
	"fmt"
	"github.com/koleaby4/next_movie_go/internal/db"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

func (h *Handlers) SendCode(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	code := rand.IntN(999_999) // generate a random 6-digit code

	ctx := context.Background()
	conn := db.NewConnection(h.AppConfig.DbDsn, ctx)
	defer conn.Close(ctx)

	queries := db.New(conn)
	user, err := queries.GetUser(ctx, email)
	if err != nil {
		log.Fatalln("error sending authentication code", err)
	}

	_, err := queries.Upsert().InsertCode(context.Background(), db.InsertCodeParams{
		Email:     email,
		Code:      code,
		CreatedAt: time.Now(),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send the code to the user's email
	err = sendEmail(email, code) // implement this function to send an email
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Code sent to email")
}
