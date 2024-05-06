-- name: GetMovie :one
select *
from movies
where id = $1;

-- name: ListMovies :many
select *
from movies
order by id;

-- name: CreateMovie :execresult
insert into movies (id, title, description, poster_url, trailer_url)
values ($1, $2, $3, $4, $5);

-- name: DeleteMovie :exec
delete
from movies
where id = $1;

-- name: UpdateMovieTitle :exec
update movies
set title = $2
where id = $1;

-- name: UpdateMovieDescription :exec
update movies
set description = $2
where id = $1;

-- name: UpdateMoviePosterUrl :exec
update movies
set poster_url = $2
where id = $1;

-- name: UpdateMovieTrailerUrl :exec
update movies
set trailer_url = $2
where id = $1;

-- name: SearchMovies :many
select *
from movies
where title like $1;