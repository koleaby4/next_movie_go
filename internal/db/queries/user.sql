-- name: GetUser :one
select *
from users
where email = $1;

-- name: CreateUser :execresult
insert into users (email)
values ($1) on conflict (email) do nothing;