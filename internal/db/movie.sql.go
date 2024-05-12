// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: movie.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

const createMovie = `-- name: CreateMovie :execresult
insert into movies (id, title, overview, poster_url, trailer_url, rating, raw_data)
values ($1, $2, $3, $4, $5, $6, $7)
on conflict (id) do nothing
`

type CreateMovieParams struct {
	ID         string
	Title      string
	Overview   string
	PosterUrl  string
	TrailerUrl string
	Rating     float64
	RawData    string
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, createMovie,
		arg.ID,
		arg.Title,
		arg.Overview,
		arg.PosterUrl,
		arg.TrailerUrl,
		arg.Rating,
		arg.RawData,
	)
}

const getMovie = `-- name: GetMovie :one
select id, title, overview, poster_url, trailer_url, rating
from movies
where id = $1
`

type GetMovieRow struct {
	ID         string
	Title      string
	Overview   string
	PosterUrl  string
	TrailerUrl string
	Rating     float64
}

func (q *Queries) GetMovie(ctx context.Context, id string) (GetMovieRow, error) {
	row := q.db.QueryRow(ctx, getMovie, id)
	var i GetMovieRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Overview,
		&i.PosterUrl,
		&i.TrailerUrl,
		&i.Rating,
	)
	return i, err
}

const listMovies = `-- name: ListMovies :many
select id, title, overview, poster_url, trailer_url, rating
from movies
order by id
`

type ListMoviesRow struct {
	ID         string
	Title      string
	Overview   string
	PosterUrl  string
	TrailerUrl string
	Rating     float64
}

func (q *Queries) ListMovies(ctx context.Context) ([]ListMoviesRow, error) {
	rows, err := q.db.Query(ctx, listMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListMoviesRow
	for rows.Next() {
		var i ListMoviesRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Overview,
			&i.PosterUrl,
			&i.TrailerUrl,
			&i.Rating,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :execresult
update movies
set overview    = $2,
    title       = $3,
    overview    = $4,
    poster_url  = $5,
    trailer_url = $6,
    rating      = $7,
    raw_data    = $8,
    created_at  = NOW()
where id = $1
`

type UpdateMovieParams struct {
	ID         string
	Overview   string
	Title      string
	Overview_2 string
	PosterUrl  string
	TrailerUrl string
	Rating     float64
	RawData    string
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, updateMovie,
		arg.ID,
		arg.Overview,
		arg.Title,
		arg.Overview_2,
		arg.PosterUrl,
		arg.TrailerUrl,
		arg.Rating,
		arg.RawData,
	)
}
