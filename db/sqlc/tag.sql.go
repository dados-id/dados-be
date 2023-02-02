// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: tag.sql

package db

import (
	"context"
)

const createTag = `-- name: CreateTag :one
INSERT INTO tags (
  name
) VALUES (
  $1
) RETURNING name
`

func (q *Queries) CreateTag(ctx context.Context, name string) (string, error) {
	row := q.db.QueryRowContext(ctx, createTag, name)
	err := row.Scan(&name)
	return name, err
}

const listRandomTag = `-- name: ListRandomTag :many
SELECT name FROM tags
ORDER BY RANDOM()
LIMIT 3
`

func (q *Queries) ListRandomTag(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listRandomTag)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []string{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
