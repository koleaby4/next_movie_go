-- name: GetMovie :one
select id,
       title,
       release_date,
       overview,
       rating,
       poster_url,
       trailer_url,
       raw_data
from movies
where id = $1;

-- name: ListMovies :many
select id,
       title,
       release_date,
       overview,
       rating,
       poster_url,
       trailer_url,
       raw_data
from movies
order by release_date desc;

-- name: InsertMovie :execresult
insert into movies (id, title, release_date, overview, rating, poster_url, trailer_url, raw_data)
values ($1, $2, $3, $4, $5, $6, $7, $8)
on conflict (id) do nothing;

