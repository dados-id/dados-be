// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: faculty.sql

package db

import (
	"context"
)

const createFaculty = `-- name: CreateFaculty :one
INSERT INTO faculties (
  name
) VALUES (
  $1
) RETURNING id, name
`

func (q *Queries) CreateFaculty(ctx context.Context, name string) (Faculty, error) {
	row := q.db.QueryRowContext(ctx, createFaculty, name)
	var i Faculty
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listFacultyBySchool = `-- name: ListFacultyBySchool :many
SELECT f.id, f.name FROM faculties F
JOIN school_faculty_associations SFA ON SFA.faculty_id = F.id
WHERE SFA.school_id = $1
`

func (q *Queries) ListFacultyBySchool(ctx context.Context, schoolID int32) ([]Faculty, error) {
	rows, err := q.db.QueryContext(ctx, listFacultyBySchool, schoolID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Faculty{}
	for rows.Next() {
		var i Faculty
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const listRandomFacultyID = `-- name: ListRandomFacultyID :many
SELECT id FROM faculties
ORDER BY RANDOM()
LIMIT 3
`

func (q *Queries) ListRandomFacultyID(ctx context.Context) ([]int32, error) {
	rows, err := q.db.QueryContext(ctx, listRandomFacultyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var id int32
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

const randomFacultyID = `-- name: RandomFacultyID :one
SELECT id FROM faculties
ORDER BY RANDOM()
LIMIT 1
`

func (q *Queries) RandomFacultyID(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, randomFacultyID)
	var id int32
	err := row.Scan(&id)
	return id, err
}
