// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type MoviesWatchedByUser struct {
	UserID          int
	MovieID         int
	ExperienceStars int
}

type User struct {
	ID        int
	Email     string
	AuthToken string
	Expiry    pgtype.Timestamp
}

type Movie struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	Overview    string  `json:"overview"`
	Rating      float64 `json:"vote_average"`
	PosterUrl   string  `json:"poster_path"`
	TrailerUrl  string  `json:"trailer_path"`
	RawData     string  `json:"-"`
}