-- name: CreateProfessor :one
INSERT INTO professors (
  first_name,
  last_name,
  faculty_id,
  school_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetProfessor :one
SELECT * FROM professors
WHERE id = $1;

-- name: GetProfessorInfoAggregate :one
SELECT
  P.total_review,
  P.rating,
  P.would_take_again,
  P.level_of_difficulty,
  SUM(CASE PR.quality when 1 then 1 else 0 end)::int as awful,
  SUM(CASE PR.quality when 2 then 1 else 0 end)::int as ok,
  SUM(CASE PR.quality when 3 then 1 else 0 end)::int as good,
  SUM(CASE PR.quality when 4 then 1 else 0 end)::int as great,
  SUM(CASE PR.quality when 5 then 1 else 0 end)::int as awesome
FROM professors P
  JOIN professor_ratings PR ON P.id = PR.professor_id
WHERE P.id = $1
GROUP BY P.id;

-- name: ListTop5Tags :many
SELECT T.name as tag_names FROM tags T
  JOIN professor_rating_tags PRT ON PRT.tag_id = T.id
  JOIN professor_ratings PR ON PRT.professor_id = PR.id
WHERE
  PR.professor_id = $1
GROUP BY PR.professor_id
ORDER BY COUNT(*)
LIMIT 5;

-- name: ListProfessors :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  P.total_review,
  P.would_take_again,
  P.level_of_difficulty,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
LIMIT $1
OFFSET $2;

-- name: SearchProfessorsByName :many
SELECT * FROM professors
WHERE first_name LIKE $1 OR last_name LIKE $1 OR concat(first_name, ' ', last_name) LIKE $1
LIMIT 5;

-- name: ListProfessorsBySchool :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  P.total_review,
  P.would_take_again,
  P.level_of_difficulty,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE school_id = $1
LIMIT $2
OFFSET $3;

-- name: ListProfessorsByFaculty :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  P.total_review,
  P.would_take_again,
  P.level_of_difficulty,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE faculty_id = $1
LIMIT $2
OFFSET $3;

-- name: ListProfessorsByFacultyAndSchool :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  P.total_review,
  P.would_take_again,
  P.level_of_difficulty,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE faculty_id = $1 AND school_id = $2
LIMIT $3
OFFSET $4;

-- name: UpdateProfessorStatusRequest :one
UPDATE professors
SET
  status = @status::text
WHERE
  id = @id::bigint
RETURNING *;