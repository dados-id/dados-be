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
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FacultyID int64  `json:"faculty_id"`
	SchoolID  int64  `json:"school_id"`
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

const getProfessor = `-- name: GetProfessor :one
SELECT id, first_name, last_name, rating, total_review, would_take_again, level_of_difficulty, created_at, status, verified_date, faculty_id, school_id FROM professors
WHERE id = $1
`

func (q *Queries) GetProfessor(ctx context.Context, id int64) (Professor, error) {
	row := q.db.QueryRowContext(ctx, getProfessor, id)
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

const getProfessorInfoAggregate = `-- name: GetProfessorInfoAggregate :one
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
GROUP BY P.id
`

type GetProfessorInfoAggregateRow struct {
	TotalReview       int32  `json:"total_review"`
	Rating            int16  `json:"rating"`
	WouldTakeAgain    int16  `json:"would_take_again"`
	LevelOfDifficulty string `json:"level_of_difficulty"`
	Awful             int32  `json:"awful"`
	Ok                int32  `json:"ok"`
	Good              int32  `json:"good"`
	Great             int32  `json:"great"`
	Awesome           int32  `json:"awesome"`
}

func (q *Queries) GetProfessorInfoAggregate(ctx context.Context, id int64) (GetProfessorInfoAggregateRow, error) {
	row := q.db.QueryRowContext(ctx, getProfessorInfoAggregate, id)
	var i GetProfessorInfoAggregateRow
	err := row.Scan(
		&i.TotalReview,
		&i.Rating,
		&i.WouldTakeAgain,
		&i.LevelOfDifficulty,
		&i.Awful,
		&i.Ok,
		&i.Good,
		&i.Great,
		&i.Awesome,
	)
	return i, err
}

const listProfessors = `-- name: ListProfessors :many
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
OFFSET $2
`

type ListProfessorsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ListProfessorsRow struct {
	ID                int64  `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Rating            int16  `json:"rating"`
	TotalReview       int32  `json:"total_review"`
	WouldTakeAgain    int16  `json:"would_take_again"`
	LevelOfDifficulty string `json:"level_of_difficulty"`
	FacultyName       string `json:"faculty_name"`
	SchoolName        string `json:"school_name"`
}

func (q *Queries) ListProfessors(ctx context.Context, arg ListProfessorsParams) ([]ListProfessorsRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessors, arg.Limit, arg.Offset)
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
			&i.TotalReview,
			&i.WouldTakeAgain,
			&i.LevelOfDifficulty,
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
OFFSET $3
`

type ListProfessorsByFacultyParams struct {
	FacultyID int64 `json:"faculty_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

type ListProfessorsByFacultyRow struct {
	ID                int64  `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Rating            int16  `json:"rating"`
	TotalReview       int32  `json:"total_review"`
	WouldTakeAgain    int16  `json:"would_take_again"`
	LevelOfDifficulty string `json:"level_of_difficulty"`
	FacultyName       string `json:"faculty_name"`
	SchoolName        string `json:"school_name"`
}

func (q *Queries) ListProfessorsByFaculty(ctx context.Context, arg ListProfessorsByFacultyParams) ([]ListProfessorsByFacultyRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorsByFaculty, arg.FacultyID, arg.Limit, arg.Offset)
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
			&i.TotalReview,
			&i.WouldTakeAgain,
			&i.LevelOfDifficulty,
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
OFFSET $4
`

type ListProfessorsByFacultyAndSchoolParams struct {
	FacultyID int64 `json:"faculty_id"`
	SchoolID  int64 `json:"school_id"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"offset"`
}

type ListProfessorsByFacultyAndSchoolRow struct {
	ID                int64  `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Rating            int16  `json:"rating"`
	TotalReview       int32  `json:"total_review"`
	WouldTakeAgain    int16  `json:"would_take_again"`
	LevelOfDifficulty string `json:"level_of_difficulty"`
	FacultyName       string `json:"faculty_name"`
	SchoolName        string `json:"school_name"`
}

func (q *Queries) ListProfessorsByFacultyAndSchool(ctx context.Context, arg ListProfessorsByFacultyAndSchoolParams) ([]ListProfessorsByFacultyAndSchoolRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorsByFacultyAndSchool,
		arg.FacultyID,
		arg.SchoolID,
		arg.Limit,
		arg.Offset,
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
			&i.TotalReview,
			&i.WouldTakeAgain,
			&i.LevelOfDifficulty,
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
OFFSET $3
`

type ListProfessorsBySchoolParams struct {
	SchoolID int64 `json:"school_id"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

type ListProfessorsBySchoolRow struct {
	ID                int64  `json:"id"`
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Rating            int16  `json:"rating"`
	TotalReview       int32  `json:"total_review"`
	WouldTakeAgain    int16  `json:"would_take_again"`
	LevelOfDifficulty string `json:"level_of_difficulty"`
	FacultyName       string `json:"faculty_name"`
	SchoolName        string `json:"school_name"`
}

func (q *Queries) ListProfessorsBySchool(ctx context.Context, arg ListProfessorsBySchoolParams) ([]ListProfessorsBySchoolRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorsBySchool, arg.SchoolID, arg.Limit, arg.Offset)
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
			&i.TotalReview,
			&i.WouldTakeAgain,
			&i.LevelOfDifficulty,
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

const listTop5Tags = `-- name: ListTop5Tags :many
SELECT T.name as tag_name FROM tags T
  JOIN professor_rating_tags PRT ON PRT.tag_id = T.id
  JOIN professor_ratings PR ON PRT.professor_id = PR.id
WHERE
  PR.professor_id = $1
GROUP BY PR.professor_id
ORDER BY COUNT(*)
LIMIT 5
`

func (q *Queries) ListTop5Tags(ctx context.Context, professorID int64) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listTop5Tags, professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var tag_name string
		if err := rows.Scan(&tag_name); err != nil {
			return nil, err
		}
		items = append(items, tag_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const searchProfessorsByName = `-- name: SearchProfessorsByName :many
SELECT id, first_name, last_name, rating, total_review, would_take_again, level_of_difficulty, created_at, status, verified_date, faculty_id, school_id FROM professors
WHERE first_name LIKE $1 OR last_name LIKE $1 OR concat(first_name, ' ', last_name) LIKE $1
LIMIT 10
`

func (q *Queries) SearchProfessorsByName(ctx context.Context, firstName string) ([]Professor, error) {
	rows, err := q.db.QueryContext(ctx, searchProfessorsByName, firstName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Professor{}
	for rows.Next() {
		var i Professor
		if err := rows.Scan(
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

const updateProfessorStatusRequest = `-- name: UpdateProfessorStatusRequest :one
UPDATE professors
SET
  status = $1::text
WHERE
  id = $2::bigint
RETURNING id, first_name, last_name, rating, total_review, would_take_again, level_of_difficulty, created_at, status, verified_date, faculty_id, school_id
`

type UpdateProfessorStatusRequestParams struct {
	Status string `json:"status"`
	ID     int64  `json:"id"`
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
