-- name: Matches :one
select count(*) as matches
from authentication
where user_id = $1
    and code = $2
    and expires > now()
limit 1;

-- name: Upsert :exec
INSERT INTO authentication (user_id, code, expires)
VALUES ($1, $2, NOW() + INTERVAL '24 hours')
ON CONFLICT (user_id)
    DO UPDATE SET code = $2, expires = NOW() + INTERVAL '24 hours';
