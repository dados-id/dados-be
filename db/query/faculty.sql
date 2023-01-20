-- name: CreateFaculty :one
INSERT INTO faculties (
  name
) VALUES (
  $1
) RETURNING *;
