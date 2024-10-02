-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;
-- dollar signs are params that are passed into
-- :one at the end tells SQLC that we expect to get a single row back

-- name: GetUser :one
SELECT * FROM users
WHERE $1 = name;

-- name: ResetUser :exec
DELETE FROM users *;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY id;