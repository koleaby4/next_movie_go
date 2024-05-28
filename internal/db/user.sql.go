// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgconn"
)

const createUser = `-- name: CreateUser :execresult
insert into users (email)
values ($1) on conflict (email) do nothing
`

func (q *Queries) CreateUser(ctx context.Context, email string) (pgconn.CommandTag, error) {
	return q.db.Exec(ctx, createUser, email)
}

const getUser = `-- name: GetUser :one
select id, email
from users
where email = $1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(&i.ID, &i.Email)
	return i, err
}
