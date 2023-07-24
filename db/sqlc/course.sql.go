// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: course.sql

package db

import (
	"context"
)

const createCourse = `-- name: CreateCourse :one
INSERT INTO courses (
  code,
  name
) VALUES (
  $1, $2
) RETURNING code, name
`

type CreateCourseParams struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) (Course, error) {
	row := q.db.QueryRowContext(ctx, createCourse, arg.Code, arg.Name)
	var i Course
	err := row.Scan(&i.Code, &i.Name)
	return i, err
}

const listCoursesByProfessorId = `-- name: ListCoursesByProfessorId :many
SELECT C.code FROM professor_course_associations PCA
  JOIN courses C ON PCA.course_code = C.code
WHERE PCA.professor_id = $1
`

func (q *Queries) ListCoursesByProfessorId(ctx context.Context, professorID int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listCoursesByProfessorId, professorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err != nil {
			return nil, err
		}
		items = append(items, code)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRandomCourseCode = `-- name: ListRandomCourseCode :many
SELECT code FROM courses
ORDER BY RANDOM()
LIMIT 3
`

func (q *Queries) ListRandomCourseCode(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listRandomCourseCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var code string
		if err := rows.Scan(&code); err != nil {
			return nil, err
		}
		items = append(items, code)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
