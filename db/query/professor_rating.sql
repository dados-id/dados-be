-- name: CreateProfessorRating :one
INSERT INTO professor_ratings (
  quality,
  difficult,
  would_take_again,
  taken_for_credit,
  use_textbooks,
  attendance_mandatory,
  grade,
  review,
  professor_id,
  course_code,
  user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
) RETURNING *;

-- name: CreateProfessorCourseAssociation :exec
INSERT INTO professor_course_associations (
  course_code,
  professor_id
) VALUES (
  $1, $2
);

-- name: CreateProfessorRatingTags :exec
INSERT INTO professor_rating_tags (
  tag_name,
  professor_rating_id
) VALUES (
  $1, $2
);

-- name: GetProfessorRating :one
SELECT
  PR.quality,
  PR.difficult,
  PR.would_take_again,
  PR.taken_for_credit,
  PR.use_textbooks,
  PR.attendance_mandatory,
  PR.grade,
  PR.review,
  P.first_name as professor_first_name,
  P.last_name as professor_last_name,
  S.name as school_name,
  F.name as faculty_name,
  array_agg(PRT.tag_name)::varchar[] tags
FROM professor_ratings PR
  JOIN professors P ON PR.professor_id = P.id
  JOIN schools S ON P.school_id = S.id
  JOIN faculties F ON P.faculty_id = F.id
  JOIN professor_rating_tags PRT ON PR.id = PRT.professor_rating_id
WHERE
  P.id = @professor_id::bigint AND PR.id = @professor_rating_id::bigint
GROUP BY
  P.id, PR.id, S.name, F.name;

-- name: ListProfessorRatings :many
SELECT
  PR.id,
  PR.quality,
  PR.difficult,
  PR.would_take_again,
  PR.taken_for_credit,
  PR.use_textbooks,
  PR.attendance_mandatory,
  PR.grade,
  PR.review,
  PR.up_vote ,
  PR.down_vote ,
  PR.created_at,
  array_agg(PRT.tag_name)::varchar[] tags
FROM professor_ratings PR
  JOIN professor_rating_tags PRT ON PR.id = PRT.professor_rating_id
WHERE
  PR.professor_id = $1
GROUP BY
  PR.id
LIMIT $2
OFFSET $3;

-- name: ListProfessorRatingsFilterByCourse :many
SELECT
  PR.id,
  PR.quality,
  PR.difficult,
  PR.would_take_again,
  PR.taken_for_credit,
  PR.use_textbooks,
  PR.attendance_mandatory,
  PR.grade,
  PR.review,
  PR.up_vote ,
  PR.down_vote ,
  PR.created_at,
  array_agg(PRT.tag_name)::varchar[] tags
FROM professor_ratings PR
  JOIN professor_rating_tags PRT ON PR.id = PRT.professor_rating_id
WHERE
  PR.professor_id = $1 AND PR.course_code = $2
GROUP BY
  PR.id
LIMIT $3
OFFSET $4;

-- name: ListProfessorRatingsFilterByRating :many
SELECT
  PR.id,
  PR.quality,
  PR.difficult,
  PR.would_take_again,
  PR.taken_for_credit,
  PR.use_textbooks,
  PR.attendance_mandatory,
  PR.grade,
  PR.review,
  PR.up_vote ,
  PR.down_vote ,
  PR.created_at,
  array_agg(PRT.tag_name)::varchar[] tags
FROM professor_ratings PR
  JOIN professor_rating_tags PRT ON PR.id = PRT.professor_rating_id
WHERE
  PR.professor_id = $1 AND PR.quality = $2
GROUP BY
  PR.id
LIMIT $3
OFFSET $4;

-- name: UpdateProfessorRating :one
UPDATE professor_ratings
SET
  quality = COALESCE(sqlc.narg(quality), quality),
  difficult = COALESCE(sqlc.narg(difficult), difficult),
  would_take_again = COALESCE(sqlc.narg(would_take_again), would_take_again),
  taken_for_credit = COALESCE(sqlc.narg(taken_for_credit), taken_for_credit),
  use_textbooks = COALESCE(sqlc.narg(use_textbooks), use_textbooks),
  attendance_mandatory = COALESCE(sqlc.narg(attendance_mandatory), attendance_mandatory),
  grade = COALESCE(sqlc.narg(grade), grade),
  review = COALESCE(sqlc.narg(review), review),
  up_vote = COALESCE(sqlc.narg(up_vote), up_vote),
  down_vote = COALESCE(sqlc.narg(down_vote), down_vote),
  course_code = COALESCE(sqlc.narg(course_code), course_code)
WHERE
  id = @professor_id
RETURNING *;
