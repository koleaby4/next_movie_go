// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Movie struct {
	ID          int32
	Title       string
	ReleaseDate string
	Overview    string
	PosterUrl   string
	TrailerUrl  string
	Rating      float64
	RawData     string
	CreatedAt   pgtype.Timestamp
}

type MoviesWatchedByUser struct {
	UserID          int32
	MovieID         int32
	ExperienceStars int32
}

type User struct {
	ID        int32
	Email     string
	AuthToken string
	Expiry    pgtype.Timestamp
}
