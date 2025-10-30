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
-- name: GetAllFeeds :many
SELECT *
FROM feeds;
-- name: FollowFeed :one
INSERT INTO feed_following(id, created_at, updated_at, user_id, feed_id)
VALUES($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetAllFollowingFeeds :many
SELECT *
FROM feed_following
WHERE user_id = $1;
-- name: UnfollowFeed :exec
DELETE FROM feed_following
WHERE user_id = $1
    AND id = $2;
-- name: GetNextFeedsToFetch :many
SELECT *
FROM feeds
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT $1;
-- name: MarkFeedAsFetched :one
UPDATE feeds
SET updated_at = NOW(),
    last_fetched_at = NOW()
WHERE id = $1
RETURNING *;