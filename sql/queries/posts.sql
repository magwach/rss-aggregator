-- name: CreatePost :one
INSERT INTO posts (
        id,
        created_at,
        updated_at,
        title,
        description,
        published_at,
        url,
        feed_id
    )
VALUES (
        $1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
    )
RETURNING *;
-- name: GetAllPostsFromFollowingFeeds :many
SELECT posts.*
FROM posts
    JOIN feed_following ON posts.feed_id = feed_following.feed_id
WHERE feed_following.user_id = $1
ORDER BY posts.published_at DESC;