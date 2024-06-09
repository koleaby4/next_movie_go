-- name: GetMoviesWatchedByUser :many
select *
from movies_watched_by_user
where user_id = $1;

-- name: AddMovieWatchedByUser :execresult
insert into movies_watched_by_user (user_id, movie_id, experience_stars)
values ($1, $2, $3);

-- name: RemoveMovieWatchedByUser :exec
delete
from movies_watched_by_user
where user_id = $1
  and movie_id = $2;

-- name: UpdateMovieWatchedByUser :exec
update movies_watched_by_user
set experience_stars = $1
where user_id = $2
  and movie_id = $3;

