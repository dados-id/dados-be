-- name: CreateCorrection :one
INSERT INTO correction_forms (
  problem,
  correct_info,
  email,
  user_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: ListCorrection :many
SELECT * FROM correction_forms
LIMIT $1
OFFSET $2;

-- name: UpdateCorrection :one
UPDATE correction_forms
SET
  status = @status
WHERE
  id = @id::bigint
RETURNING *;
