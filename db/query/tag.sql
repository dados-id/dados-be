-- name: CreateTag :one
INSERT INTO tags (
  name
) VALUES (
  $1
) RETURNING *;

-- name: RandomTag :one
SELECT name FROM tags
ORDER BY RANDOM()
LIMIT 1;
