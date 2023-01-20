-- name: CreateCourse :one
INSERT INTO courses (
  code,
  name
) VALUES (
  $1, $2
) RETURNING *;


-- name: CountCourse :one
SELECT COUNT(*) FROM courses;
