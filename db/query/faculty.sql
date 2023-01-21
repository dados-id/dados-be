-- name: CreateFaculty :one
INSERT INTO faculties (
  name
) VALUES (
  $1
) RETURNING *;

-- name: CountFaculty :one
SELECT COUNT(*) FROM faculties;
