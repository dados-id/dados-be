-- name: CreateProfessor :one
INSERT INTO professors (
  first_name,
  last_name,
  faculty_id,
  school_id
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetProfessorInfo :one
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.total_review,
  P.rating,
  P.would_take_again,
  P.level_of_difficulty,
  F.name as faculty_name,
  S.name as school_name,
  SUM(CASE PR.quality when 1 then 1 else 0 end)::int as terrible,
  SUM(CASE PR.quality when 2 then 1 else 0 end)::int as poor,
  SUM(CASE PR.quality when 3 then 1 else 0 end)::int as fair,
  SUM(CASE PR.quality when 4 then 1 else 0 end)::int as good,
  SUM(CASE PR.quality when 5 then 1 else 0 end)::int as excellent
FROM professors P
  LEFT JOIN professor_ratings PR ON P.id = PR.professor_id
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE
  P.id = $1
GROUP BY P.id, F.id, S.id;

-- name: ListTopTags :many
SELECT PRT.tag_name as tag_names FROM professor_rating_tags PRT
  JOIN professor_ratings PR ON PRT.professor_rating_id = PR.id
WHERE
  PR.professor_id = $1
GROUP BY PRT.tag_name
ORDER BY COUNT(PRT.tag_name) DESC
LIMIT 5;

-- name: ListTopCoursesTaught :many
SELECT PR.course_code FROM professor_ratings PR
  JOIN courses C on C.code = PR.course_code
WHERE
  PR.professor_id = $1
GROUP BY
  PR.course_code
ORDER BY COUNT(*)::int DESC
LIMIT 3;

-- name: ListProfessors :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
ORDER BY
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2;

-- name: CountListProfessors :one
SELECT COUNT(*)::int FROM professors;

-- name: ListProfessorsByName :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE LOWER(P.first_name) LIKE LOWER(@name::varchar)
  OR LOWER(P.last_name) LIKE LOWER(@name::varchar)
  OR LOWER(concat(P.first_name, ' ', P.last_name)) LIKE LOWER(@name::varchar)
ORDER BY
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2;

-- name: CountListProfessorsByName :one
SELECT COUNT(*)::int FROM professors P
 WHERE LOWER(P.first_name) LIKE LOWER(@name::varchar)
 OR LOWER(P.last_name) LIKE LOWER(@name::varchar)
 OR LOWER(concat(P.first_name, ' ', P.last_name)) LIKE LOWER(@name::varchar);

-- name: ListProfessorsBySchool :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE P.school_id = $1
ORDER BY
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $2
OFFSET $3;

-- name: CountListProfessorsBySchool :one
SELECT COUNT(*)::int FROM professors
  WHERE school_id = $1;

-- name: ListProfessorsByFaculty :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE P.faculty_id = $1
ORDER BY
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $2
OFFSET $3;

-- name: CountListProfessorsByFaculty :one
SELECT COUNT(*)::int FROM professors
  WHERE faculty_id = $1;

-- name: ListProfessorsByFacultyAndSchool :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  F.name as faculty_name,
  S.name as school_name
FROM professors P
  LEFT JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE P.faculty_id = $1 AND P.school_id = $2
ORDER BY
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'rating' AND @sort_order::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $3
OFFSET $4;

-- name: UpdateProfessorStatusRequest :one
UPDATE professors
SET
  status = @status
WHERE
  id = @id::int
RETURNING *;

-- name: RandomProfessorID :one
SELECT id FROM professors
ORDER BY RANDOM()
LIMIT 1;
