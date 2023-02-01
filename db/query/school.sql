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

-- name: GetSchoolInfo :one
SELECT
  S.name,
  COALESCE(ROUND(AVG(SR.reputation), 1), 0.0)::varchar as reputation,
  COALESCE(ROUND(AVG(SR.location), 1), 0.0)::varchar as location,
  COALESCE(ROUND(AVG(SR.opportunities), 1), 0.0)::varchar as opportunities,
  COALESCE(ROUND(AVG(SR.facilities), 1), 0.0)::varchar as facilities,
  COALESCE(ROUND(AVG(SR.internet), 1), 0.0)::varchar as internet,
  COALESCE(ROUND(AVG(SR.food), 1), 0.0)::varchar as food,
  COALESCE(ROUND(AVG(SR.clubs), 1), 0.0)::varchar as clubs,
  COALESCE(ROUND(AVG(SR.social), 1), 0.0)::varchar as social,
  COALESCE(ROUND(AVG(SR.happiness), 1), 0.0)::varchar as happiness,
  COALESCE(ROUND(AVG(SR.safety), 1), 0.0)::varchar as safety,
  COALESCE(ROUND(AVG(SR.overall_rating), 1), 0.0)::varchar as overall_rating
FROM schools S
  LEFT JOIN school_ratings SR ON S.id = SR.school_id
WHERE
  S.id = $1
GROUP BY S.id;

-- name: ListSchools :many
SELECT
  S.id,
  S.name,
  S.city,
  S.province
FROM schools S
ORDER BY
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'asc' THEN LOWER(S.name)
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'desc' THEN LOWER(S.name)
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2;

-- name: CountListSchools :one
SELECT COUNT(*) FROM schools;

-- name: ListSchoolsByName :many
SELECT
  S.id,
  S.name,
  S.city,
  S.province
FROM schools S
WHERE @nick_name::varchar ILIKE ANY(S.nick_name) OR S.name ILIKE @name::varchar
ORDER BY
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'asc' THEN LOWER(S.name)
    ELSE NULL
  END,
  CASE
    WHEN @sort_by::varchar = 'name' AND @sort_order::varchar = 'desc' THEN LOWER(S.name)
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2;

-- name: CountListSchoolsByName :one
SELECT COUNT(*) FROM schools
  WHERE @nick_name::varchar ILIKE ANY(S.nick_name)
  OR S.name ILIKE @name::varchar;

-- name: UpdateSchoolStatusRequest :one
UPDATE schools
SET
  status = @status
WHERE
  id = @id::bigint
RETURNING *;

-- name: RandomSchoolID :one
SELECT id FROM schools
ORDER BY RANDOM()
LIMIT 1;
