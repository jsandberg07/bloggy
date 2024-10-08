-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)

SELECT 
    inserted_feed_follow.*, 
    feeds.name AS feed_name,
    users.name AS user_name
    FROM inserted_feed_follow
    INNER JOIN users ON inserted_feed_follow.user_id = users.id
    INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, u.name, f.name
FROM feed_follows AS ff
INNER JOIN users AS u ON ff.user_id = u.id
INNER JOIN feeds AS f ON ff.feed_id = f.id
WHERE $1 = u.name;
    
-- name: UnfollowFeedForUser :exec
WITH delete_batch AS (
    SELECT feed_follows.id FROM feed_follows
    INNER JOIN feeds ON feed_follows.feed_id = feeds.id
    WHERE $1 = feeds.url
    AND $2 = feed_follows.user_id
)

DELETE FROM feed_follows
USING delete_batch
WHERE delete_batch.id = feed_follows.id;