-- name: CreateUser :one
INSERT INTO users (
  first_name,
  last_name,
  school,
  expected_year_of_graduation,
  email
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: UpdateUser :one
UPDATE users
SET
  first_name = COALESCE(sqlc.narg(first_name), first_name),
  last_name = COALESCE(sqlc.narg(last_name), last_name),
  school = COALESCE(sqlc.narg(school), school),
  expected_year_of_graduation = COALESCE(sqlc.narg(expected_year_of_graduation), expected_year_of_graduation),
  email = COALESCE(sqlc.narg(email), email)
WHERE
  id = sqlc.arg(id)
RETURNING *;
