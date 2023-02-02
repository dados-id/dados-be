// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  id,
  first_name,
  last_name,
  expected_year_of_graduation,
  email,
  school_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, first_name, last_name, expected_year_of_graduation, email, created_at, school_id
`

type CreateUserParams struct {
	ID                       string        `json:"id"`
	FirstName                string        `json:"firstName"`
	LastName                 string        `json:"lastName"`
	ExpectedYearOfGraduation sql.NullInt16 `json:"expectedYearOfGraduation"`
	Email                    string        `json:"email"`
	SchoolID                 sql.NullInt32 `json:"schoolID"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.ExpectedYearOfGraduation,
		arg.Email,
		arg.SchoolID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.ExpectedYearOfGraduation,
		&i.Email,
		&i.CreatedAt,
		&i.SchoolID,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT
  U.id,
  U.first_name,
  U.last_name,
  U.expected_year_of_graduation,
  U.email,
  S.name
FROM users U
  JOIN schools S ON S.id = U.school_id
WHERE U.id = $1
`

type GetUserRow struct {
	ID                       string        `json:"id"`
	FirstName                string        `json:"firstName"`
	LastName                 string        `json:"lastName"`
	ExpectedYearOfGraduation sql.NullInt16 `json:"expectedYearOfGraduation"`
	Email                    string        `json:"email"`
	Name                     string        `json:"name"`
}

func (q *Queries) GetUser(ctx context.Context, id string) (GetUserRow, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i GetUserRow
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.ExpectedYearOfGraduation,
		&i.Email,
		&i.Name,
	)
	return i, err
}

const listRandomUserID = `-- name: ListRandomUserID :many
SELECT id FROM users
ORDER BY RANDOM()
LIMIT 3
`

func (q *Queries) ListRandomUserID(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listRandomUserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const saveProfessor = `-- name: SaveProfessor :exec
INSERT INTO user_save_professors (
  professor_id,
  user_id
) VALUES (
  $1, $2
)
`

type SaveProfessorParams struct {
	ProfessorID int32  `json:"professorID"`
	UserID      string `json:"userID"`
}

func (q *Queries) SaveProfessor(ctx context.Context, arg SaveProfessorParams) error {
	_, err := q.db.ExecContext(ctx, saveProfessor, arg.ProfessorID, arg.UserID)
	return err
}

const unsaveProfessor = `-- name: UnsaveProfessor :exec
DELETE FROM user_save_professors
WHERE
  professor_id = $1
AND
  user_id = $2
`

type UnsaveProfessorParams struct {
	ProfessorID int32  `json:"professorID"`
	UserID      string `json:"userID"`
}

func (q *Queries) UnsaveProfessor(ctx context.Context, arg UnsaveProfessorParams) error {
	_, err := q.db.ExecContext(ctx, unsaveProfessor, arg.ProfessorID, arg.UserID)
	return err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
  first_name = COALESCE($1, first_name),
  last_name = COALESCE($2, last_name),
  school_id = COALESCE($3, school_id),
  expected_year_of_graduation = COALESCE($4, expected_year_of_graduation)
WHERE
  id = $5
RETURNING id, first_name, last_name, expected_year_of_graduation, email, created_at, school_id
`

type UpdateUserParams struct {
	FirstName                sql.NullString `json:"firstName"`
	LastName                 sql.NullString `json:"lastName"`
	SchoolID                 sql.NullInt32  `json:"schoolID"`
	ExpectedYearOfGraduation sql.NullInt16  `json:"expectedYearOfGraduation"`
	ID                       string         `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.SchoolID,
		arg.ExpectedYearOfGraduation,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.ExpectedYearOfGraduation,
		&i.Email,
		&i.CreatedAt,
		&i.SchoolID,
	)
	return i, err
}

const userListProfessorRatings = `-- name: UserListProfessorRatings :many
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
  PR.created_at,
  P.first_name as professor_first_name,
  P.last_name as professor_last_name,
  S.name as school_name,
  C.name as course_name,
  array_agg(PRT.tag_name)::varchar[] tags
FROM professor_ratings PR
  JOIN professors P ON PR.professor_id = P.id
  JOIN schools S ON P.school_id = S.id
  JOIN courses C ON C.code = PR.course_code
  JOIN professor_rating_tags PRT ON PR.id = PRT.professor_rating_id
WHERE
  PR.user_id = $1
GROUP BY
  PR.id, P.id, S.id, C.code
LIMIT $2
OFFSET $3
`

type UserListProfessorRatingsParams struct {
	UserID string `json:"userID"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type UserListProfessorRatingsRow struct {
	ID                  int32     `json:"id"`
	Quality             string    `json:"quality"`
	Difficult           string    `json:"difficult"`
	WouldTakeAgain      int16     `json:"wouldTakeAgain"`
	TakenForCredit      int16     `json:"takenForCredit"`
	UseTextbooks        int16     `json:"useTextbooks"`
	AttendanceMandatory int16     `json:"attendanceMandatory"`
	Grade               string    `json:"grade"`
	Review              string    `json:"review"`
	CreatedAt           time.Time `json:"createdAt"`
	ProfessorFirstName  string    `json:"professorFirstName"`
	ProfessorLastName   string    `json:"professorLastName"`
	SchoolName          string    `json:"schoolName"`
	CourseName          string    `json:"courseName"`
	Tags                []string  `json:"tags"`
}

func (q *Queries) UserListProfessorRatings(ctx context.Context, arg UserListProfessorRatingsParams) ([]UserListProfessorRatingsRow, error) {
	rows, err := q.db.QueryContext(ctx, userListProfessorRatings, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserListProfessorRatingsRow{}
	for rows.Next() {
		var i UserListProfessorRatingsRow
		if err := rows.Scan(
			&i.ID,
			&i.Quality,
			&i.Difficult,
			&i.WouldTakeAgain,
			&i.TakenForCredit,
			&i.UseTextbooks,
			&i.AttendanceMandatory,
			&i.Grade,
			&i.Review,
			&i.CreatedAt,
			&i.ProfessorFirstName,
			&i.ProfessorLastName,
			&i.SchoolName,
			&i.CourseName,
			pq.Array(&i.Tags),
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

const userListSavedProfessors = `-- name: UserListSavedProfessors :many
SELECT
  P.id,
  P.first_name,
  P.last_name,
  P.rating,
  F.name as faculty_name,
  S.name as school_name
FROM user_save_professors USP
  JOIN professors P ON USP.professor_id = P.id
  JOIN faculties F ON P.faculty_id = F.id
  JOIN schools S ON P.school_id = S.id
WHERE
  USP.user_id = $1
LIMIT $2
OFFSET $3
`

type UserListSavedProfessorsParams struct {
	UserID string `json:"userID"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type UserListSavedProfessorsRow struct {
	ID          int32  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Rating      string `json:"rating"`
	FacultyName string `json:"facultyName"`
	SchoolName  string `json:"schoolName"`
}

func (q *Queries) UserListSavedProfessors(ctx context.Context, arg UserListSavedProfessorsParams) ([]UserListSavedProfessorsRow, error) {
	rows, err := q.db.QueryContext(ctx, userListSavedProfessors, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserListSavedProfessorsRow{}
	for rows.Next() {
		var i UserListSavedProfessorsRow
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

const userListSchoolRatings = `-- name: UserListSchoolRatings :many
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
OFFSET $3
`

type UserListSchoolRatingsParams struct {
	UserID string `json:"userID"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type UserListSchoolRatingsRow struct {
	ID            int32     `json:"id"`
	Reputation    int16     `json:"reputation"`
	Location      int16     `json:"location"`
	Opportunities int16     `json:"opportunities"`
	Facilities    int16     `json:"facilities"`
	Internet      int16     `json:"internet"`
	Food          int16     `json:"food"`
	Clubs         int16     `json:"clubs"`
	Social        int16     `json:"social"`
	Happiness     int16     `json:"happiness"`
	Safety        int16     `json:"safety"`
	Review        string    `json:"review"`
	OverallRating string    `json:"overallRating"`
	CreatedAt     time.Time `json:"createdAt"`
	SchoolName    string    `json:"schoolName"`
}

func (q *Queries) UserListSchoolRatings(ctx context.Context, arg UserListSchoolRatingsParams) ([]UserListSchoolRatingsRow, error) {
	rows, err := q.db.QueryContext(ctx, userListSchoolRatings, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []UserListSchoolRatingsRow{}
	for rows.Next() {
		var i UserListSchoolRatingsRow
		if err := rows.Scan(
			&i.ID,
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
			&i.Review,
			&i.OverallRating,
			&i.CreatedAt,
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
