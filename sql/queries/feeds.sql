-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id, last_fetched_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    NULL
)
RETURNING *;

-- name: GetFeeds :many
SELECT f.name, url, u.name
FROM feeds AS f
LEFT JOIN users AS u ON f.user_id = u.id;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE $1 = url;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at = $1
WHERE $2 = id;
UPDATE feeds SET updated_at = $3
WHERE $2 = id;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST;