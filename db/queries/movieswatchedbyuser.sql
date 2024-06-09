-- name: GetMoviesWatchedByUser :many
select *
from movies_watched_by_user
where user_id = $1;

-- name: UpsertMovieWatchedByUser :exec
INSERT INTO movies_watched_by_user (user_id, movie_id, experience_stars)
VALUES ($1, $2, $3)
ON CONFLICT (user_id, movie_id)
    DO UPDATE SET experience_stars = EXCLUDED.experience_stars;

-- name: RemoveMovieWatchedByUser :exec
delete
from movies_watched_by_user
where user_id = $1
  and movie_id = $2;
