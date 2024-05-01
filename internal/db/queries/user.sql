-- name: GetUser :one
select * from users where email = $1;

-- name: ListUsers :many
select *
from users
order by email;

-- name: CreateUser :execresult
insert into users (email) values($1);

-- name: DeleteUser :exec
delete from users where email = $1;

-- name: UpdateUser :exec
update users
set email = $1 where email = $1;

-- name: SearchUsers :many
select id, email from users
where email like $1;