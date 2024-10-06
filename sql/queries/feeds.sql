-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, url, u.name
FROM feeds AS f
LEFT JOIN users AS u ON f.user_id = u.id;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE $1 = url;