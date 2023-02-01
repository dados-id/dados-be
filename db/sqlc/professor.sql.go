// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: professor.sql

package db

import (
	"context"
)

const createProfessor = `-- name: CreateProfessor :one
INSERT INTO professors (
  first_name,
  last_name,
  faculty_id,
  school_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, first_name, last_name, rating, total_review, would_take_again, level_of_difficulty, created_at, status, verified_date, faculty_id, school_id
`

type CreateProfessorParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	FacultyID int64  `json:"facultyID"`
	SchoolID  int64  `json:"schoolID"`
}

func (q *Queries) CreateProfessor(ctx context.Context, arg CreateProfessorParams) (Professor, error) {
	row := q.db.QueryRowContext(ctx, createProfessor,
		arg.FirstName,
		arg.LastName,
		arg.FacultyID,
		arg.SchoolID,
	)
	var i Professor
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Rating,
		&i.TotalReview,
		&i.WouldTakeAgain,
		&i.LevelOfDifficulty,
		&i.CreatedAt,
		&i.Status,
		&i.VerifiedDate,
		&i.FacultyID,
		&i.SchoolID,
	)
	return i, err
}

const getProfessorInfo = `-- name: GetProfessorInfo :one
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
GROUP BY P.id, F.id, S.id
`

type GetProfessorInfoRow struct {
	ID                int64  `json:"id"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	TotalReview       int32  `json:"totalReview"`
	Rating            string `json:"rating"`
	WouldTakeAgain    int16  `json:"wouldTakeAgain"`
	LevelOfDifficulty string `json:"levelOfDifficulty"`
	FacultyName       string `json:"facultyName"`
	SchoolName        string `json:"schoolName"`
	Terrible          int32  `json:"terrible"`
	Poor              int32  `json:"poor"`
	Fair              int32  `json:"fair"`
	Good              int32  `json:"good"`
	Excellent         int32  `json:"excellent"`
}

func (q *Queries) GetProfessorInfo(ctx context.Context, id int64) (GetProfessorInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getProfessorInfo, id)
	var i GetProfessorInfoRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.TotalReview,
		&i.Rating,
		&i.WouldTakeAgain,
		&i.LevelOfDifficulty,
		&i.FacultyName,
		&i.SchoolName,
		&i.Terrible,
		&i.Poor,
		&i.Fair,
		&i.Good,
		&i.Excellent,
	)
	return i, err
}

const listProfessors = `-- name: ListProfessors :many
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
    WHEN $3::varchar = 'name' AND $4::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN $3::varchar = 'name' AND $4::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN $3::varchar = 'rating' AND $4::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN $3::varchar = 'rating' AND $4::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2
`

type ListProfessorsParams struct {
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ListProfessorsRow struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Rating      string `json:"rating"`
	FacultyName string `json:"facultyName"`
	SchoolName  string `json:"schoolName"`
}

func (q *Queries) ListProfessors(ctx context.Context, arg ListProfessorsParams) ([]ListProfessorsRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessors,
		arg.Limit,
		arg.Offset,
		arg.SortBy,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorsRow{}
	for rows.Next() {
		var i ListProfessorsRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Rating,
			&i.FacultyName,
			&i.SchoolName,
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

const listProfessorsByFaculty = `-- name: ListProfessorsByFaculty :many
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
    WHEN $4::varchar = 'name' AND $5::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN $4::varchar = 'name' AND $5::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN $4::varchar = 'rating' AND $5::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN $4::varchar = 'rating' AND $5::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $2
OFFSET $3
`

type ListProfessorsByFacultyParams struct {
	FacultyID int64  `json:"facultyID"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ListProfessorsByFacultyRow struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Rating      string `json:"rating"`
	FacultyName string `json:"facultyName"`
	SchoolName  string `json:"schoolName"`
}

func (q *Queries) ListProfessorsByFaculty(ctx context.Context, arg ListProfessorsByFacultyParams) ([]ListProfessorsByFacultyRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorsByFaculty,
		arg.FacultyID,
		arg.Limit,
		arg.Offset,
		arg.SortBy,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorsByFacultyRow{}
	for rows.Next() {
		var i ListProfessorsByFacultyRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Rating,
			&i.FacultyName,
			&i.SchoolName,
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

const listProfessorsByFacultyAndSchool = `-- name: ListProfessorsByFacultyAndSchool :many
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
    WHEN $5::varchar = 'name' AND $6::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN $5::varchar = 'name' AND $6::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN $5::varchar = 'rating' AND $6::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN $5::varchar = 'rating' AND $6::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $3
OFFSET $4
`

type ListProfessorsByFacultyAndSchoolParams struct {
	FacultyID int64  `json:"facultyID"`
	SchoolID  int64  `json:"schoolID"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ListProfessorsByFacultyAndSchoolRow struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Rating      string `json:"rating"`
	FacultyName string `json:"facultyName"`
	SchoolName  string `json:"schoolName"`
}

func (q *Queries) ListProfessorsByFacultyAndSchool(ctx context.Context, arg ListProfessorsByFacultyAndSchoolParams) ([]ListProfessorsByFacultyAndSchoolRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorsByFacultyAndSchool,
		arg.FacultyID,
		arg.SchoolID,
		arg.Limit,
		arg.Offset,
		arg.SortBy,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorsByFacultyAndSchoolRow{}
	for rows.Next() {
		var i ListProfessorsByFacultyAndSchoolRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Rating,
			&i.FacultyName,
			&i.SchoolName,
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

const listProfessorsByName = `-- name: ListProfessorsByName :many
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
WHERE LOWER(P.first_name) LIKE LOWER($3::varchar) OR LOWER(P.last_name) LIKE LOWER($3::varchar) OR LOWER(concat(P.first_name, ' ', P.last_name)) LIKE LOWER($3::varchar)
ORDER BY
  CASE
    WHEN $4::varchar = 'name' AND $5::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN $4::varchar = 'name' AND $5::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN $4::varchar = 'rating' AND $5::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN $4::varchar = 'rating' AND $5::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $1
OFFSET $2
`

type ListProfessorsByNameParams struct {
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	Name      string `json:"name"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ListProfessorsByNameRow struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Rating      string `json:"rating"`
	FacultyName string `json:"facultyName"`
	SchoolName  string `json:"schoolName"`
}

func (q *Queries) ListProfessorsByName(ctx context.Context, arg ListProfessorsByNameParams) ([]ListProfessorsByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorsByName,
		arg.Limit,
		arg.Offset,
		arg.Name,
		arg.SortBy,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorsByNameRow{}
	for rows.Next() {
		var i ListProfessorsByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Rating,
			&i.FacultyName,
			&i.SchoolName,
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

const listProfessorsBySchool = `-- name: ListProfessorsBySchool :many
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
    WHEN $4::varchar = 'name' AND $5::varchar = 'asc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END,
  CASE
    WHEN $4::varchar = 'name' AND $5::varchar = 'desc' THEN LOWER(concat(P.first_name, ' ', P.last_name))
    ELSE NULL
  END DESC,
  CASE
    WHEN $4::varchar = 'rating' AND $5::varchar = 'asc' THEN P.rating
    ELSE NULL
  END,
  CASE
    WHEN $4::varchar = 'rating' AND $5::varchar = 'desc' THEN P.rating
    ELSE NULL
  END DESC
LIMIT $2
OFFSET $3
`

type ListProfessorsBySchoolParams struct {
	SchoolID  int64  `json:"schoolID"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

type ListProfessorsBySchoolRow struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Rating      string `json:"rating"`
	FacultyName string `json:"facultyName"`
	SchoolName  string `json:"schoolName"`
}

func (q *Queries) ListProfessorsBySchool(ctx context.Context, arg ListProfessorsBySchoolParams) ([]ListProfessorsBySchoolRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorsBySchool,
		arg.SchoolID,
		arg.Limit,
		arg.Offset,
		arg.SortBy,
		arg.SortOrder,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorsBySchoolRow{}
	for rows.Next() {
		var i ListProfessorsBySchoolRow
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Rating,
			&i.FacultyName,
			&i.SchoolName,
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

const listTopCoursesTaught = `-- name: ListTopCoursesTaught :many
SELECT PR.course_code FROM professor_ratings PR
  JOIN courses C on C.code = PR.course_code
WHERE
  PR.professor_id = $1
GROUP BY
  PR.course_code
ORDER BY COUNT(*) DESC
LIMIT 3
`

func (q *Queries) ListTopCoursesTaught(ctx context.Context, professorID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listTopCoursesTaught, professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var course_code string
		if err := rows.Scan(&course_code); err != nil {
			return nil, err
		}
		items = append(items, course_code)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTopTags = `-- name: ListTopTags :many
SELECT PRT.tag_name as tag_names FROM professor_rating_tags PRT
  JOIN professor_ratings PR ON PRT.professor_rating_id = PR.id
WHERE
  PR.professor_id = $1
GROUP BY PRT.tag_name
ORDER BY COUNT(PRT.tag_name) DESC
LIMIT 5
`

func (q *Queries) ListTopTags(ctx context.Context, professorID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listTopTags, professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var tag_names string
		if err := rows.Scan(&tag_names); err != nil {
			return nil, err
		}
		items = append(items, tag_names)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const randomProfessorID = `-- name: RandomProfessorID :one
SELECT id FROM professors
ORDER BY RANDOM()
LIMIT 1
`

func (q *Queries) RandomProfessorID(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, randomProfessorID)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateProfessorStatusRequest = `-- name: UpdateProfessorStatusRequest :one
UPDATE professors
SET
  status = $1
WHERE
  id = $2::bigint
RETURNING id, first_name, last_name, rating, total_review, would_take_again, level_of_difficulty, created_at, status, verified_date, faculty_id, school_id
`

type UpdateProfessorStatusRequestParams struct {
	Status Statusrequest `json:"status"`
	ID     int64         `json:"id"`
}

func (q *Queries) UpdateProfessorStatusRequest(ctx context.Context, arg UpdateProfessorStatusRequestParams) (Professor, error) {
	row := q.db.QueryRowContext(ctx, updateProfessorStatusRequest, arg.Status, arg.ID)
	var i Professor
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Rating,
		&i.TotalReview,
		&i.WouldTakeAgain,
		&i.LevelOfDifficulty,
		&i.CreatedAt,
		&i.Status,
		&i.VerifiedDate,
		&i.FacultyID,
		&i.SchoolID,
	)
	return i, err
}
