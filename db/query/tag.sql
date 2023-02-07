-- name: CreateTag :one
INSERT INTO tags (
  name
) VALUES (
  $1
) RETURNING *;

-- name: ListRandomTag :many
SELECT name FROM tags
ORDER BY RANDOM()
LIMIT 3;
