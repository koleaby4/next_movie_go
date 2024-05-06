// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: movie.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

const createMovie = `-- name: CreateMovie :execresult
insert into movies (id, title, description, poster_url, trailer_url)
values ($1, $2, $3, $4, $5)
`

type CreateMovieParams struct {
	ID          string
	Title       string
	Description string
	PosterUrl   pgtype.Text
	TrailerUrl  pgtype.Text
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, createMovie,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.PosterUrl,
		arg.TrailerUrl,
	)
}

const deleteMovie = `-- name: DeleteMovie :exec
delete
from movies
where id = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteMovie, id)
	return err
}

const getMovie = `-- name: GetMovie :one
select id, title, description, poster_url, trailer_url
from movies
where id = $1
`

func (q *Queries) GetMovie(ctx context.Context, id string) (Movie, error) {
	row := q.db.QueryRow(ctx, getMovie, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.PosterUrl,
		&i.TrailerUrl,
	)
	return i, err
}

const listMovies = `-- name: ListMovies :many
select id, title, description, poster_url, trailer_url
from movies
order by id
`

func (q *Queries) ListMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.Query(ctx, listMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.PosterUrl,
			&i.TrailerUrl,
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

const searchMovies = `-- name: SearchMovies :many
select id, title, description, poster_url, trailer_url
from movies
where title like $1
`

func (q *Queries) SearchMovies(ctx context.Context, title string) ([]Movie, error) {
	rows, err := q.db.Query(ctx, searchMovies, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.PosterUrl,
			&i.TrailerUrl,
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

const updateMovieDescription = `-- name: UpdateMovieDescription :exec
update movies
set description = $2
where id = $1
`

type UpdateMovieDescriptionParams struct {
	ID          string
	Description string
}

func (q *Queries) UpdateMovieDescription(ctx context.Context, arg UpdateMovieDescriptionParams) error {
	_, err := q.db.Exec(ctx, updateMovieDescription, arg.ID, arg.Description)
	return err
}

const updateMoviePosterUrl = `-- name: UpdateMoviePosterUrl :exec
update movies
set poster_url = $2
where id = $1
`

type UpdateMoviePosterUrlParams struct {
	ID        string
	PosterUrl pgtype.Text
}

func (q *Queries) UpdateMoviePosterUrl(ctx context.Context, arg UpdateMoviePosterUrlParams) error {
	_, err := q.db.Exec(ctx, updateMoviePosterUrl, arg.ID, arg.PosterUrl)
	return err
}

const updateMovieTitle = `-- name: UpdateMovieTitle :exec
update movies
set title = $2
where id = $1
`

type UpdateMovieTitleParams struct {
	ID    string
	Title string
}

func (q *Queries) UpdateMovieTitle(ctx context.Context, arg UpdateMovieTitleParams) error {
	_, err := q.db.Exec(ctx, updateMovieTitle, arg.ID, arg.Title)
	return err
}

const updateMovieTrailerUrl = `-- name: UpdateMovieTrailerUrl :exec
update movies
set trailer_url = $2
where id = $1
`

type UpdateMovieTrailerUrlParams struct {
	ID         string
	TrailerUrl pgtype.Text
}

func (q *Queries) UpdateMovieTrailerUrl(ctx context.Context, arg UpdateMovieTrailerUrlParams) error {
	_, err := q.db.Exec(ctx, updateMovieTrailerUrl, arg.ID, arg.TrailerUrl)
	return err
}
