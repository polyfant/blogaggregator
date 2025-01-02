-- name: CreatePost :one
INSERT INTO posts (
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPostsForUser :many
SELECT p.* FROM posts p
JOIN feeds f ON p.feed_id = f.id
JOIN feed_follows ff ON f.id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY p.published_at DESC
LIMIT $2;
