-- name: CreateSchool :one
INSERT INTO schools (
  name,
  nick_name,
  city,
  province,
  website,
  email
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetSchoolInfoAggregate :one
SELECT
  S.name,
  COALESCE(ROUND(AVG(SR.reputation), 1), 0.0)::text as reputation,
  COALESCE(ROUND(AVG(SR.location), 1), 0.0)::text as location,
  COALESCE(ROUND(AVG(SR.opportunities), 1), 0.0)::text as opportunities,
  COALESCE(ROUND(AVG(SR.facilities), 1), 0.0)::text as facilities,
  COALESCE(ROUND(AVG(SR.internet), 1), 0.0)::text as internet,
  COALESCE(ROUND(AVG(SR.food), 1), 0.0)::text as food,
  COALESCE(ROUND(AVG(SR.clubs), 1), 0.0)::text as clubs,
  COALESCE(ROUND(AVG(SR.social), 1), 0.0)::text as social,
  COALESCE(ROUND(AVG(SR.happiness), 1), 0.0)::text as happiness,
  COALESCE(ROUND(AVG(SR.safety), 1), 0.0)::text as safety,
  COALESCE(ROUND(AVG(SR.overall_rating), 1), 0.0)::text as overall_rating
FROM schools S
  LEFT JOIN school_ratings SR ON S.id = SR.school_id
WHERE
  S.id = $1
GROUP BY S.id;

-- name: ListSchools :many
SELECT
  S.id,
  S.name
FROM schools S
LIMIT $1
OFFSET $2;

-- name: SearchSchoolsByNameOrNickName :many
SELECT
  id,
  name,
  city,
  province
FROM schools
WHERE @name_arr::text ILIKE ANY(nick_name) OR name ILIKE @name::text
ORDER BY id ASC
LIMIT 5;

-- name: UpdateSchoolStatusRequest :one
UPDATE schools
SET
  status = @status
WHERE
  id = @id::bigint
RETURNING *;

-- name: CountSchool :one
SELECT COUNT(*) FROM schools;
