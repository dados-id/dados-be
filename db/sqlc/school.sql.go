// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: school.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createSchool = `-- name: CreateSchool :one
INSERT INTO schools (
  name,
  nick_name,
  city,
  province,
  website,
  email
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, name, nick_name, city, province, website, email, status, verified_date
`

type CreateSchoolParams struct {
	Name     string   `json:"name"`
	NickName []string `json:"nickName"`
	City     string   `json:"city"`
	Province string   `json:"province"`
	Website  string   `json:"website"`
	Email    string   `json:"email"`
}

func (q *Queries) CreateSchool(ctx context.Context, arg CreateSchoolParams) (School, error) {
	row := q.db.QueryRowContext(ctx, createSchool,
		arg.Name,
		pq.Array(arg.NickName),
		arg.City,
		arg.Province,
		arg.Website,
		arg.Email,
	)
	var i School
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.NickName),
		&i.City,
		&i.Province,
		&i.Website,
		&i.Email,
		&i.Status,
		&i.VerifiedDate,
	)
	return i, err
}

const getSchoolInfo = `-- name: GetSchoolInfo :one
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
GROUP BY S.id
`

type GetSchoolInfoRow struct {
	Name          string `json:"name"`
	Reputation    string `json:"reputation"`
	Location      string `json:"location"`
	Opportunities string `json:"opportunities"`
	Facilities    string `json:"facilities"`
	Internet      string `json:"internet"`
	Food          string `json:"food"`
	Clubs         string `json:"clubs"`
	Social        string `json:"social"`
	Happiness     string `json:"happiness"`
	Safety        string `json:"safety"`
	OverallRating string `json:"overallRating"`
}

func (q *Queries) GetSchoolInfo(ctx context.Context, id int64) (GetSchoolInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getSchoolInfo, id)
	var i GetSchoolInfoRow
	err := row.Scan(
		&i.Name,
		&i.Reputation,
		&i.Location,
		&i.Opportunities,
		&i.Facilities,
		&i.Internet,
		&i.Food,
		&i.Clubs,
		&i.Social,
		&i.Happiness,
		&i.Safety,
		&i.OverallRating,
	)
	return i, err
}

const listSchools = `-- name: ListSchools :many
SELECT
  S.id,
  S.name,
  S.city,
  S.province
FROM schools S
ORDER BY
  CASE
    WHEN $3::varchar = 'name' AND $4::varchar = 'asc' THEN LOWER(S.name)
    ELSE NULL
  END,
  CASE
    WHEN $3::varchar = 'name' AND $4::varchar = 'desc' THEN LOWER(S.name)
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2
`

type ListSchoolsParams struct {
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ListSchoolsRow struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Province string `json:"province"`
}

func (q *Queries) ListSchools(ctx context.Context, arg ListSchoolsParams) ([]ListSchoolsRow, error) {
	rows, err := q.db.QueryContext(ctx, listSchools,
		arg.Limit,
		arg.Offset,
		arg.SortBy,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListSchoolsRow{}
	for rows.Next() {
		var i ListSchoolsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.City,
			&i.Province,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSchoolsByName = `-- name: ListSchoolsByName :many
SELECT
  id,
  name,
  city,
  province
FROM schools
WHERE $3::text ILIKE ANY(nick_name) OR name ILIKE $4::text
ORDER BY
  CASE
    WHEN $5::varchar = 'name' AND $6::varchar = 'asc' THEN LOWER(name)
    ELSE NULL
  END,
  CASE
    WHEN $5::varchar = 'name' AND $6::varchar = 'desc' THEN LOWER(name)
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2
`

type ListSchoolsByNameParams struct {
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	NickName  string `json:"nickName"`
	Name      string `json:"name"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ListSchoolsByNameRow struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Province string `json:"province"`
}

func (q *Queries) ListSchoolsByName(ctx context.Context, arg ListSchoolsByNameParams) ([]ListSchoolsByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, listSchoolsByName,
		arg.Limit,
		arg.Offset,
		arg.NickName,
		arg.Name,
		arg.SortBy,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListSchoolsByNameRow{}
	for rows.Next() {
		var i ListSchoolsByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.City,
			&i.Province,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const randomSchoolID = `-- name: RandomSchoolID :one
SELECT id FROM schools
ORDER BY RANDOM()
LIMIT 1
`

func (q *Queries) RandomSchoolID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, randomSchoolID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateSchoolStatusRequest = `-- name: UpdateSchoolStatusRequest :one
UPDATE schools
SET
  status = $1
WHERE
  id = $2::bigint
RETURNING id, name, nick_name, city, province, website, email, status, verified_date
`

type UpdateSchoolStatusRequestParams struct {
	Status Statusrequest `json:"status"`
	ID     int64         `json:"id"`
}

func (q *Queries) UpdateSchoolStatusRequest(ctx context.Context, arg UpdateSchoolStatusRequestParams) (School, error) {
	row := q.db.QueryRowContext(ctx, updateSchoolStatusRequest, arg.Status, arg.ID)
	var i School
	err := row.Scan(
		&i.ID,
		&i.Name,
		pq.Array(&i.NickName),
		&i.City,
		&i.Province,
		&i.Website,
		&i.Email,
		&i.Status,
		&i.VerifiedDate,
	)
	return i, err
}
