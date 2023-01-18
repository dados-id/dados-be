-- name: CreateSchool :one
INSERT INTO schools (
  name,
  nick_name,
  country,
  province,
  website,
  email
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetSchool :one
SELECT * FROM schools
WHERE id = $1;

-- name: GetSchoolInfoAggregate :one
SELECT
  AVG(SR.reputation)::smallint as reputation,
  AVG(SR.location)::smallint as location,
  AVG(SR.opportunities)::smallint as opportunities,
  AVG(SR.facilities)::smallint as facilities,
  AVG(SR.internet)::smallint as internet,
  AVG(SR.food)::smallint as food,
  AVG(SR.clubs)::smallint as clubs,
  AVG(SR.social)::smallint as social,
  AVG(SR.happiness)::smallint as happiness,
  AVG(SR.safety)::smallint as safety,
  AVG(SR.overall_rating)::smallint as overall_rating
FROM schools S
  JOIN school_ratings SR ON S.id = SR.school_id
WHERE
  S.id = $1
GROUP BY S.id;

-- name: ListSchools :many
SELECT * FROM schools
LIMIT $1
OFFSET $2;

-- name: SearchSchoolsByNameOrNickName :many
SELECT * FROM schools
WHERE name LIKE $1 OR $1 LIKE ANY(nick_name)
LIMIT 5;

-- name: UpdateSchoolStatusRequest :one
UPDATE schools
SET
  status = @status::text
WHERE
  id = @id::bigint
RETURNING *;

-- name: CountSchool :one
SELECT COUNT(*) FROM schools;
