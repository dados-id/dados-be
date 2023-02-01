// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: school_rating.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createSchoolFacultyAssociation = `-- name: CreateSchoolFacultyAssociation :exec
INSERT INTO school_faculty_associations (
  faculty_id,
  school_id
) VALUES (
  $1, $2
)
`

type CreateSchoolFacultyAssociationParams struct {
	FacultyID int32 `json:"facultyID"`
	SchoolID  int32 `json:"schoolID"`
}

func (q *Queries) CreateSchoolFacultyAssociation(ctx context.Context, arg CreateSchoolFacultyAssociationParams) error {
	_, err := q.db.ExecContext(ctx, createSchoolFacultyAssociation, arg.FacultyID, arg.SchoolID)
	return err
}

const createSchoolRating = `-- name: CreateSchoolRating :one
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
) RETURNING id, reputation, location, opportunities, facilities, internet, food, clubs, social, happiness, safety, review, up_vote, down_vote, overall_rating, created_at, edited_at, status, verified_date, user_id, school_id
`

type CreateSchoolRatingParams struct {
	UserID        string `json:"userID"`
	SchoolID      int32  `json:"schoolID"`
	Reputation    int16  `json:"reputation"`
	Location      int16  `json:"location"`
	Opportunities int16  `json:"opportunities"`
	Facilities    int16  `json:"facilities"`
	Internet      int16  `json:"internet"`
	Food          int16  `json:"food"`
	Clubs         int16  `json:"clubs"`
	Social        int16  `json:"social"`
	Happiness     int16  `json:"happiness"`
	Safety        int16  `json:"safety"`
	Review        string `json:"review"`
}

func (q *Queries) CreateSchoolRating(ctx context.Context, arg CreateSchoolRatingParams) (SchoolRating, error) {
	row := q.db.QueryRowContext(ctx, createSchoolRating,
		arg.UserID,
		arg.SchoolID,
		arg.Reputation,
		arg.Location,
		arg.Opportunities,
		arg.Facilities,
		arg.Internet,
		arg.Food,
		arg.Clubs,
		arg.Social,
		arg.Happiness,
		arg.Safety,
		arg.Review,
	)
	var i SchoolRating
	err := row.Scan(
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
		&i.UpVote,
		&i.DownVote,
		&i.OverallRating,
		&i.CreatedAt,
		&i.EditedAt,
		&i.Status,
		&i.VerifiedDate,
		&i.UserID,
		&i.SchoolID,
	)
	return i, err
}

const getSchoolRating = `-- name: GetSchoolRating :one
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
WHERE S.id = $1::int AND SR.id = $2::int
`

type GetSchoolRatingParams struct {
	SchoolID       int32 `json:"schoolID"`
	SchoolRatingID int32 `json:"schoolRatingID"`
}

type GetSchoolRatingRow struct {
	ID            int32  `json:"id"`
	Reputation    int16  `json:"reputation"`
	Location      int16  `json:"location"`
	Opportunities int16  `json:"opportunities"`
	Facilities    int16  `json:"facilities"`
	Internet      int16  `json:"internet"`
	Food          int16  `json:"food"`
	Clubs         int16  `json:"clubs"`
	Social        int16  `json:"social"`
	Happiness     int16  `json:"happiness"`
	Safety        int16  `json:"safety"`
	Review        string `json:"review"`
	SchoolName    string `json:"schoolName"`
}

func (q *Queries) GetSchoolRating(ctx context.Context, arg GetSchoolRatingParams) (GetSchoolRatingRow, error) {
	row := q.db.QueryRowContext(ctx, getSchoolRating, arg.SchoolID, arg.SchoolRatingID)
	var i GetSchoolRatingRow
	err := row.Scan(
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
		&i.SchoolName,
	)
	return i, err
}

const listSchoolRatings = `-- name: ListSchoolRatings :many
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
OFFSET $3
`

type ListSchoolRatingsParams struct {
	SchoolID int32 `json:"schoolID"`
	Limit    int32 `json:"limit"`
	Offset   int32 `json:"offset"`
}

type ListSchoolRatingsRow struct {
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
	UpVote        int32     `json:"upVote"`
	DownVote      int32     `json:"downVote"`
	OverallRating string    `json:"overallRating"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (q *Queries) ListSchoolRatings(ctx context.Context, arg ListSchoolRatingsParams) ([]ListSchoolRatingsRow, error) {
	rows, err := q.db.QueryContext(ctx, listSchoolRatings, arg.SchoolID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListSchoolRatingsRow{}
	for rows.Next() {
		var i ListSchoolRatingsRow
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
			&i.UpVote,
			&i.DownVote,
			&i.OverallRating,
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

const updateSchoolRating = `-- name: UpdateSchoolRating :one
UPDATE school_ratings
SET
  reputation = COALESCE($1, reputation),
  location = COALESCE($2, location),
  opportunities = COALESCE($3, opportunities),
  facilities = COALESCE($4, facilities),
  internet = COALESCE($5, internet),
  food = COALESCE($6, food),
  clubs = COALESCE($7, clubs),
  social = COALESCE($8, social),
  happiness = COALESCE($9, happiness),
  safety = COALESCE($10, safety),
  review = COALESCE($11, review),
  up_vote = COALESCE($12, up_vote),
  down_vote = COALESCE($13, down_vote)
WHERE
  id = $14 AND school_id = $15
RETURNING id, reputation, location, opportunities, facilities, internet, food, clubs, social, happiness, safety, review, up_vote, down_vote, overall_rating, created_at, edited_at, status, verified_date, user_id, school_id
`

type UpdateSchoolRatingParams struct {
	Reputation     sql.NullInt16  `json:"reputation"`
	Location       sql.NullInt16  `json:"location"`
	Opportunities  sql.NullInt16  `json:"opportunities"`
	Facilities     sql.NullInt16  `json:"facilities"`
	Internet       sql.NullInt16  `json:"internet"`
	Food           sql.NullInt16  `json:"food"`
	Clubs          sql.NullInt16  `json:"clubs"`
	Social         sql.NullInt16  `json:"social"`
	Happiness      sql.NullInt16  `json:"happiness"`
	Safety         sql.NullInt16  `json:"safety"`
	Review         sql.NullString `json:"review"`
	UpVote         sql.NullInt32  `json:"upVote"`
	DownVote       sql.NullInt32  `json:"downVote"`
	SchoolRatingID int32          `json:"schoolRatingID"`
	SchoolID       int32          `json:"schoolID"`
}

func (q *Queries) UpdateSchoolRating(ctx context.Context, arg UpdateSchoolRatingParams) (SchoolRating, error) {
	row := q.db.QueryRowContext(ctx, updateSchoolRating,
		arg.Reputation,
		arg.Location,
		arg.Opportunities,
		arg.Facilities,
		arg.Internet,
		arg.Food,
		arg.Clubs,
		arg.Social,
		arg.Happiness,
		arg.Safety,
		arg.Review,
		arg.UpVote,
		arg.DownVote,
		arg.SchoolRatingID,
		arg.SchoolID,
	)
	var i SchoolRating
	err := row.Scan(
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
		&i.UpVote,
		&i.DownVote,
		&i.OverallRating,
		&i.CreatedAt,
		&i.EditedAt,
		&i.Status,
		&i.VerifiedDate,
		&i.UserID,
		&i.SchoolID,
	)
	return i, err
}
