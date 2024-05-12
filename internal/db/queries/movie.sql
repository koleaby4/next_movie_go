-- name: GetMovie :one
select id, title, overview, poster_url, trailer_url, rating
from movies
where id = $1;

-- name: ListMovies :many
select id, title, overview, poster_url, trailer_url, rating
from movies
order by id;

-- name: CreateMovie :execresult
insert into movies (id, title, overview, poster_url, trailer_url, rating, raw_data)
values ($1, $2, $3, $4, $5, $6, $7)
on conflict (id) do nothing;


-- name: UpdateMovie :execresult
update movies
set overview    = $2,
    title       = $3,
    overview    = $4,
    poster_url  = $5,
    trailer_url = $6,
    rating      = $7,
    raw_data    = $8,
    created_at  = NOW()
where id = $1;