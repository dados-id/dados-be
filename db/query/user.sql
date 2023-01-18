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

-- name: SaveProfessor :exec
INSERT INTO user_save_professors (
  professor_id,
  user_id
) VALUES (
  $1, $2
);

-- name: UnsaveProfessor :exec
DELETE FROM user_save_professors
WHERE
  professor_id = $1
AND
  user_id = $2;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: UserListProfessorRatings :many
SELECT
  PR.id,
  PR.quality,
  PR.difficult,
  PR.would_take_again,
  PR.taken_for_credit,
  PR.use_textbooks,
  PR.attendance_mandatory,
  PR.grade,
  PR.tags,
  PR.review,
  PR.created_at,
  P.first_name as professor_first_name,
  P.last_name as professor_last_name,
  S.name as school_name,
  C.name as course_name
FROM professor_ratings PR
  JOIN professors P ON PR.professor_id = P.id
  JOIN schools S ON P.school_id = S.id
  JOIN courses C ON C.code = PR.course_code
WHERE
  PR.user_id = $1
LIMIT $2
OFFSET $3;

-- name: UserListSchoolRatings :many
SELECT
  SR.id,
  SR.reputation,
  SR.location,
  SR.opportunities,
  SR.facilities,
  SR.internet,
  SR.food,
  SR.clubs,
  SR.social,
  SR.happiness,
  SR.safety,
  SR.review,
  SR.overall_rating,
  SR.created_at,
  S.name as school_name
FROM school_ratings SR
  JOIN schools S ON SR.school_id = S.id
WHERE
  SR.user_id = $1
LIMIT $2
OFFSET $3;

-- name: UserListSavedProfessors :many
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
FROM user_save_professors USP
  JOIN professors P ON USP.professor_id = P.id
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE
  USP.user_id = $1
LIMIT $2
OFFSET $3;

-- name: UpdateUser :one
UPDATE users
SET
  first_name = COALESCE(sqlc.narg(first_name), first_name),
  last_name = COALESCE(sqlc.narg(last_name), last_name),
  school = COALESCE(sqlc.narg(school), school),
  expected_year_of_graduation = COALESCE(sqlc.narg(expected_year_of_graduation), expected_year_of_graduation)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: CountUser :one
SELECT COUNT(*) FROM users;
