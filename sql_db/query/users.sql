-- name: CreateUser :one
INSERT INTO users (
    username, password
) VALUES (
             $1, $2
         )
    RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at
LIMIT $1
OFFSET $2
;


-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;