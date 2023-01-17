// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: professor_rating.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createProfessorCourseAssociation = `-- name: CreateProfessorCourseAssociation :exec
INSERT INTO professor_course_associations (
  course_code,
  professor_id
) VALUES (
  $1, $2
)
`

type CreateProfessorCourseAssociationParams struct {
	CourseCode  string `json:"course_code"`
	ProfessorID int64  `json:"professor_id"`
}

func (q *Queries) CreateProfessorCourseAssociation(ctx context.Context, arg CreateProfessorCourseAssociationParams) error {
	_, err := q.db.ExecContext(ctx, createProfessorCourseAssociation, arg.CourseCode, arg.ProfessorID)
	return err
}

const createProfessorRating = `-- name: CreateProfessorRating :one
INSERT INTO professor_ratings (
  quality,
  difficult,
  would_take_again,
  taken_for_credit,
  use_textbooks,
  attendance_mandatory,
  grade,
  tags,
  review,
  professor_id,
  course_code,
  user_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
) RETURNING id, quality, difficult, would_take_again, taken_for_credit, use_textbooks, attendance_mandatory, grade, tags, review, up_vote, down_vote, created_at, edited_at, professor_id, course_code, user_id, verified
`

type CreateProfessorRatingParams struct {
	Quality             string         `json:"quality"`
	Difficult           string         `json:"difficult"`
	WouldTakeAgain      int16          `json:"would_take_again"`
	TakenForCredit      sql.NullBool   `json:"taken_for_credit"`
	UseTextbooks        sql.NullBool   `json:"use_textbooks"`
	AttendanceMandatory int16          `json:"attendance_mandatory"`
	Grade               sql.NullString `json:"grade"`
	Tags                []string       `json:"tags"`
	Review              string         `json:"review"`
	ProfessorID         int64          `json:"professor_id"`
	CourseCode          string         `json:"course_code"`
	UserID              int64          `json:"user_id"`
}

func (q *Queries) CreateProfessorRating(ctx context.Context, arg CreateProfessorRatingParams) (ProfessorRating, error) {
	row := q.db.QueryRowContext(ctx, createProfessorRating,
		arg.Quality,
		arg.Difficult,
		arg.WouldTakeAgain,
		arg.TakenForCredit,
		arg.UseTextbooks,
		arg.AttendanceMandatory,
		arg.Grade,
		pq.Array(arg.Tags),
		arg.Review,
		arg.ProfessorID,
		arg.CourseCode,
		arg.UserID,
	)
	var i ProfessorRating
	err := row.Scan(
		&i.ID,
		&i.Quality,
		&i.Difficult,
		&i.WouldTakeAgain,
		&i.TakenForCredit,
		&i.UseTextbooks,
		&i.AttendanceMandatory,
		&i.Grade,
		pq.Array(&i.Tags),
		&i.Review,
		&i.UpVote,
		&i.DownVote,
		&i.CreatedAt,
		&i.EditedAt,
		&i.ProfessorID,
		&i.CourseCode,
		&i.UserID,
		&i.Verified,
	)
	return i, err
}

const getProfessorRating = `-- name: GetProfessorRating :one
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
  P.first_name as professor_first_name,
  P.last_name as professor_last_name,
  S.name as school_name
FROM professor_ratings PR
  JOIN professors P ON PR.professor_id = P.id
  JOIN schools S ON P.school_id = S.id
WHERE P.id = $1::bigint AND PR.id = $2::bigint
`

type GetProfessorRatingParams struct {
	ProfessorID       int64 `json:"professor_id"`
	ProfessorRatingID int64 `json:"professor_rating_id"`
}

type GetProfessorRatingRow struct {
	ID                  int64          `json:"id"`
	Quality             string         `json:"quality"`
	Difficult           string         `json:"difficult"`
	WouldTakeAgain      int16          `json:"would_take_again"`
	TakenForCredit      sql.NullBool   `json:"taken_for_credit"`
	UseTextbooks        sql.NullBool   `json:"use_textbooks"`
	AttendanceMandatory int16          `json:"attendance_mandatory"`
	Grade               sql.NullString `json:"grade"`
	Tags                []string       `json:"tags"`
	Review              string         `json:"review"`
	ProfessorFirstName  string         `json:"professor_first_name"`
	ProfessorLastName   string         `json:"professor_last_name"`
	SchoolName          string         `json:"school_name"`
}

func (q *Queries) GetProfessorRating(ctx context.Context, arg GetProfessorRatingParams) (GetProfessorRatingRow, error) {
	row := q.db.QueryRowContext(ctx, getProfessorRating, arg.ProfessorID, arg.ProfessorRatingID)
	var i GetProfessorRatingRow
	err := row.Scan(
		&i.ID,
		&i.Quality,
		&i.Difficult,
		&i.WouldTakeAgain,
		&i.TakenForCredit,
		&i.UseTextbooks,
		&i.AttendanceMandatory,
		&i.Grade,
		pq.Array(&i.Tags),
		&i.Review,
		&i.ProfessorFirstName,
		&i.ProfessorLastName,
		&i.SchoolName,
	)
	return i, err
}

const listProfessorRatings = `-- name: ListProfessorRatings :many
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
  PR.up_vote ,
  PR.down_vote ,
  PR.created_at
FROM professor_ratings PR
WHERE PR.professor_id = $1
LIMIT $2
OFFSET $3
`

type ListProfessorRatingsParams struct {
	ProfessorID int64 `json:"professor_id"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

type ListProfessorRatingsRow struct {
	ID                  int64          `json:"id"`
	Quality             string         `json:"quality"`
	Difficult           string         `json:"difficult"`
	WouldTakeAgain      int16          `json:"would_take_again"`
	TakenForCredit      sql.NullBool   `json:"taken_for_credit"`
	UseTextbooks        sql.NullBool   `json:"use_textbooks"`
	AttendanceMandatory int16          `json:"attendance_mandatory"`
	Grade               sql.NullString `json:"grade"`
	Tags                []string       `json:"tags"`
	Review              string         `json:"review"`
	UpVote              int32          `json:"up_vote"`
	DownVote            int32          `json:"down_vote"`
	CreatedAt           time.Time      `json:"created_at"`
}

func (q *Queries) ListProfessorRatings(ctx context.Context, arg ListProfessorRatingsParams) ([]ListProfessorRatingsRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorRatings, arg.ProfessorID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorRatingsRow{}
	for rows.Next() {
		var i ListProfessorRatingsRow
		if err := rows.Scan(
			&i.ID,
			&i.Quality,
			&i.Difficult,
			&i.WouldTakeAgain,
			&i.TakenForCredit,
			&i.UseTextbooks,
			&i.AttendanceMandatory,
			&i.Grade,
			pq.Array(&i.Tags),
			&i.Review,
			&i.UpVote,
			&i.DownVote,
			&i.CreatedAt,
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

const listProfessorRatingsJoinProfessorFilterByCourse = `-- name: ListProfessorRatingsJoinProfessorFilterByCourse :many
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
  PR.up_vote ,
  PR.down_vote ,
  PR.created_at
FROM professor_ratings PR
WHERE PR.professor_id = $1 AND PR.course_code = $2
LIMIT $3
OFFSET $4
`

type ListProfessorRatingsJoinProfessorFilterByCourseParams struct {
	ProfessorID int64  `json:"professor_id"`
	CourseCode  string `json:"course_code"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

type ListProfessorRatingsJoinProfessorFilterByCourseRow struct {
	ID                  int64          `json:"id"`
	Quality             string         `json:"quality"`
	Difficult           string         `json:"difficult"`
	WouldTakeAgain      int16          `json:"would_take_again"`
	TakenForCredit      sql.NullBool   `json:"taken_for_credit"`
	UseTextbooks        sql.NullBool   `json:"use_textbooks"`
	AttendanceMandatory int16          `json:"attendance_mandatory"`
	Grade               sql.NullString `json:"grade"`
	Tags                []string       `json:"tags"`
	Review              string         `json:"review"`
	UpVote              int32          `json:"up_vote"`
	DownVote            int32          `json:"down_vote"`
	CreatedAt           time.Time      `json:"created_at"`
}

func (q *Queries) ListProfessorRatingsJoinProfessorFilterByCourse(ctx context.Context, arg ListProfessorRatingsJoinProfessorFilterByCourseParams) ([]ListProfessorRatingsJoinProfessorFilterByCourseRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorRatingsJoinProfessorFilterByCourse,
		arg.ProfessorID,
		arg.CourseCode,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorRatingsJoinProfessorFilterByCourseRow{}
	for rows.Next() {
		var i ListProfessorRatingsJoinProfessorFilterByCourseRow
		if err := rows.Scan(
			&i.ID,
			&i.Quality,
			&i.Difficult,
			&i.WouldTakeAgain,
			&i.TakenForCredit,
			&i.UseTextbooks,
			&i.AttendanceMandatory,
			&i.Grade,
			pq.Array(&i.Tags),
			&i.Review,
			&i.UpVote,
			&i.DownVote,
			&i.CreatedAt,
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

const listProfessorRatingsJoinProfessorFilterByRating = `-- name: ListProfessorRatingsJoinProfessorFilterByRating :many
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
  PR.up_vote ,
  PR.down_vote ,
  PR.created_at
FROM professor_ratings PR
WHERE PR.professor_id = $1 AND PR.quality = $2
LIMIT $3
OFFSET $4
`

type ListProfessorRatingsJoinProfessorFilterByRatingParams struct {
	ProfessorID int64  `json:"professor_id"`
	Quality     string `json:"quality"`
	Limit       int32  `json:"limit"`
	Offset      int32  `json:"offset"`
}

type ListProfessorRatingsJoinProfessorFilterByRatingRow struct {
	ID                  int64          `json:"id"`
	Quality             string         `json:"quality"`
	Difficult           string         `json:"difficult"`
	WouldTakeAgain      int16          `json:"would_take_again"`
	TakenForCredit      sql.NullBool   `json:"taken_for_credit"`
	UseTextbooks        sql.NullBool   `json:"use_textbooks"`
	AttendanceMandatory int16          `json:"attendance_mandatory"`
	Grade               sql.NullString `json:"grade"`
	Tags                []string       `json:"tags"`
	Review              string         `json:"review"`
	UpVote              int32          `json:"up_vote"`
	DownVote            int32          `json:"down_vote"`
	CreatedAt           time.Time      `json:"created_at"`
}

func (q *Queries) ListProfessorRatingsJoinProfessorFilterByRating(ctx context.Context, arg ListProfessorRatingsJoinProfessorFilterByRatingParams) ([]ListProfessorRatingsJoinProfessorFilterByRatingRow, error) {
	rows, err := q.db.QueryContext(ctx, listProfessorRatingsJoinProfessorFilterByRating,
		arg.ProfessorID,
		arg.Quality,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProfessorRatingsJoinProfessorFilterByRatingRow{}
	for rows.Next() {
		var i ListProfessorRatingsJoinProfessorFilterByRatingRow
		if err := rows.Scan(
			&i.ID,
			&i.Quality,
			&i.Difficult,
			&i.WouldTakeAgain,
			&i.TakenForCredit,
			&i.UseTextbooks,
			&i.AttendanceMandatory,
			&i.Grade,
			pq.Array(&i.Tags),
			&i.Review,
			&i.UpVote,
			&i.DownVote,
			&i.CreatedAt,
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

const updateProfessorRating = `-- name: UpdateProfessorRating :one
UPDATE professor_ratings
SET
  quality = COALESCE($1, quality),
  difficult = COALESCE($2, difficult),
  would_take_again = COALESCE($3, would_take_again),
  taken_for_credit = COALESCE($4, taken_for_credit),
  use_textbooks = COALESCE($5, use_textbooks),
  attendance_mandatory = COALESCE($6, attendance_mandatory),
  grade = COALESCE($7, grade),
  tags = COALESCE($8, tags),
  review = COALESCE($9, review),
  up_vote = COALESCE($10, up_vote),
  down_vote = COALESCE($11, down_vote),
  course_code = COALESCE($12, course_code)
WHERE
  id = $13
RETURNING id, quality, difficult, would_take_again, taken_for_credit, use_textbooks, attendance_mandatory, grade, tags, review, up_vote, down_vote, created_at, edited_at, professor_id, course_code, user_id, verified
`

type UpdateProfessorRatingParams struct {
	Quality             sql.NullString `json:"quality"`
	Difficult           sql.NullString `json:"difficult"`
	WouldTakeAgain      sql.NullInt16  `json:"would_take_again"`
	TakenForCredit      sql.NullBool   `json:"taken_for_credit"`
	UseTextbooks        sql.NullBool   `json:"use_textbooks"`
	AttendanceMandatory sql.NullInt16  `json:"attendance_mandatory"`
	Grade               sql.NullString `json:"grade"`
	Tags                []string       `json:"tags"`
	Review              sql.NullString `json:"review"`
	UpVote              sql.NullInt32  `json:"up_vote"`
	DownVote            sql.NullInt32  `json:"down_vote"`
	CourseCode          sql.NullString `json:"course_code"`
	ID                  int64          `json:"id"`
}

func (q *Queries) UpdateProfessorRating(ctx context.Context, arg UpdateProfessorRatingParams) (ProfessorRating, error) {
	row := q.db.QueryRowContext(ctx, updateProfessorRating,
		arg.Quality,
		arg.Difficult,
		arg.WouldTakeAgain,
		arg.TakenForCredit,
		arg.UseTextbooks,
		arg.AttendanceMandatory,
		arg.Grade,
		pq.Array(arg.Tags),
		arg.Review,
		arg.UpVote,
		arg.DownVote,
		arg.CourseCode,
		arg.ID,
	)
	var i ProfessorRating
	err := row.Scan(
		&i.ID,
		&i.Quality,
		&i.Difficult,
		&i.WouldTakeAgain,
		&i.TakenForCredit,
		&i.UseTextbooks,
		&i.AttendanceMandatory,
		&i.Grade,
		pq.Array(&i.Tags),
		&i.Review,
		&i.UpVote,
		&i.DownVote,
		&i.CreatedAt,
		&i.EditedAt,
		&i.ProfessorID,
		&i.CourseCode,
		&i.UserID,
		&i.Verified,
	)
	return i, err
}