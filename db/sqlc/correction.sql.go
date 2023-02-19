// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: correction.sql

package db

import (
	"context"
)

const createCorrection = `-- name: CreateCorrection :one
INSERT INTO correction_forms (
  problem,
  correct_info,
  email,
  user_id
) VALUES (
  $1, $2, $3, $4
) RETURNING id, problem, correct_info, email, status, request_date, verified_date, user_id
`

type CreateCorrectionParams struct {
	Problem     string `json:"problem"`
	CorrectInfo string `json:"correctInfo"`
	Email       string `json:"email"`
	UserID      string `json:"userID"`
}

func (q *Queries) CreateCorrection(ctx context.Context, arg CreateCorrectionParams) (CorrectionForm, error) {
	row := q.db.QueryRowContext(ctx, createCorrection,
		arg.Problem,
		arg.CorrectInfo,
		arg.Email,
		arg.UserID,
	)
	var i CorrectionForm
	err := row.Scan(
		&i.ID,
		&i.Problem,
		&i.CorrectInfo,
		&i.Email,
		&i.Status,
		&i.RequestDate,
		&i.VerifiedDate,
		&i.UserID,
	)
	return i, err
}

const listCorrection = `-- name: ListCorrection :many
SELECT id, problem, correct_info, email, status, request_date, verified_date, user_id FROM correction_forms
LIMIT $1
OFFSET $2
`

type ListCorrectionParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCorrection(ctx context.Context, arg ListCorrectionParams) ([]CorrectionForm, error) {
	rows, err := q.db.QueryContext(ctx, listCorrection, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CorrectionForm{}
	for rows.Next() {
		var i CorrectionForm
		if err := rows.Scan(
			&i.ID,
			&i.Problem,
			&i.CorrectInfo,
			&i.Email,
			&i.Status,
			&i.RequestDate,
			&i.VerifiedDate,
			&i.UserID,
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

const updateCorrection = `-- name: UpdateCorrection :one
UPDATE correction_forms
SET
  status = $1
WHERE
  id = $2::int
RETURNING id, problem, correct_info, email, status, request_date, verified_date, user_id
`

type UpdateCorrectionParams struct {
	Status Statusrequest `json:"status"`
	ID     int32         `json:"id"`
}

func (q *Queries) UpdateCorrection(ctx context.Context, arg UpdateCorrectionParams) (CorrectionForm, error) {
	row := q.db.QueryRowContext(ctx, updateCorrection, arg.Status, arg.ID)
	var i CorrectionForm
	err := row.Scan(
		&i.ID,
		&i.Problem,
		&i.CorrectInfo,
		&i.Email,
		&i.Status,
		&i.RequestDate,
		&i.VerifiedDate,
		&i.UserID,
	)
	return i, err
}
