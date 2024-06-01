-- name: GetUser :one
select *
from users
where email = $1;


-- name: UpsertUser :one
INSERT INTO users (email, auth_token, expiry)
VALUES ($1, $2, NOW() + INTERVAL '24 hours') ON CONFLICT (user_id)
DO
UPDATE SET code = $2, expires = NOW() + INTERVAL '24 hours'
    RETURNING id, email, auth_token, expiry;