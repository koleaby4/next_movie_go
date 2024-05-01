-- name: GetMovie :one
select * from movies where id = $1;

-- name: ListMovies :many
select *
from movies
order by id;

-- name: CreateMovie :execresult
insert into movies (id, title) values($1, $2);

-- name: DeleteMovie :exec
delete from movies where id = $1;

-- name: UpdateMovie :exec
update movies
set title = $1 where id = $1;

-- name: SearchMovies :many
select id, title from movies
where title like $1;