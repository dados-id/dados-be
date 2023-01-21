-- name: CreateCourse :one
INSERT INTO courses (
  code,
  name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetCourseByProfessor :many
SELECT C.* FROM professor_course_associations PCA
  JOIN courses C ON PCA.course_code = C.code
WHERE PCA.professor_id = $1;

-- name: RandomCourseCode :one
SELECT code FROM courses
ORDER BY RANDOM()
LIMIT 1;
