-- name: CreateSchoolRating :one
INSERT INTO school_ratings (
  user_id,
  school_id,
  reputation,
  location,
  opportunities,
  facilities,
  internet,
  food,
  clubs,
  social,
  happiness,
  safety,
  review
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
) RETURNING *;

-- name: GetSchoolRating :one
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
  S.name as school_name
FROM school_ratings SR
  JOIN schools S ON SR.school_id = S.id
WHERE S.id = @school_id::bigint AND SR.id = @school_rating_id::bigint;

-- name: CreateSchoolFacultyAssociation :exec
INSERT INTO school_faculty_associations (
  faculty_id,
  school_id
) VALUES (
  $1, $2
);

-- name: ListSchoolRatings :many
SELECT
  id,
  reputation,
  location,
  opportunities,
  facilities,
  internet,
  food,
  clubs,
  social,
  happiness,
  safety,
  review,
  up_vote,
  down_vote,
  overall_rating,
  created_at
FROM school_ratings
WHERE school_id = $1
LIMIT $2
OFFSET $3;

-- name: UpdateSchoolRating :one
UPDATE school_ratings
SET
  reputation = COALESCE(sqlc.narg(reputation), reputation),
  location = COALESCE(sqlc.narg(location), location),
  opportunities = COALESCE(sqlc.narg(opportunities), opportunities),
  facilities = COALESCE(sqlc.narg(facilities), facilities),
  internet = COALESCE(sqlc.narg(internet), internet),
  food = COALESCE(sqlc.narg(food), food),
  clubs = COALESCE(sqlc.narg(clubs), clubs),
  social = COALESCE(sqlc.narg(social), social),
  happiness = COALESCE(sqlc.narg(happiness), happiness),
  safety = COALESCE(sqlc.narg(safety), safety),
  review = COALESCE(sqlc.narg(review), review),
  up_vote = COALESCE(sqlc.narg(up_vote), up_vote),
  down_vote = COALESCE(sqlc.narg(down_vote), down_vote)
WHERE
  id = sqlc.arg(id)
RETURNING *;
