-- name: CreateReport :one
INSERT INTO report_forms (
  comment,
  professor_rating_id,
  user_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: ListReport :many
SELECT * FROM report_forms
LIMIT $1
OFFSET $2;

-- name: UpdateReport :one
UPDATE report_forms
SET status = @status::text
WHERE id = @id::bigint
RETURNING *;
