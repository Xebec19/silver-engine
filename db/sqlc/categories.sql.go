// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: categories.sql

package db

import (
	"context"
	"database/sql"
)

const readAllCategories = `-- name: ReadAllCategories :many
SELECT category_id, category_name, image_url FROM categories
WHERE status = 'active'
`

type ReadAllCategoriesRow struct {
	CategoryID   int32          `json:"category_id"`
	CategoryName string         `json:"category_name"`
	ImageUrl     sql.NullString `json:"image_url"`
}

func (q *Queries) ReadAllCategories(ctx context.Context) ([]ReadAllCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, readAllCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadAllCategoriesRow
	for rows.Next() {
		var i ReadAllCategoriesRow
		if err := rows.Scan(&i.CategoryID, &i.CategoryName, &i.ImageUrl); err != nil {
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

const readOneCategory = `-- name: ReadOneCategory :one
SELECT category_id, category_name, image_url FROM categories 
WHERE status = 'active' AND category_id = $1
`

type ReadOneCategoryRow struct {
	CategoryID   int32          `json:"category_id"`
	CategoryName string         `json:"category_name"`
	ImageUrl     sql.NullString `json:"image_url"`
}

func (q *Queries) ReadOneCategory(ctx context.Context, categoryID int32) (ReadOneCategoryRow, error) {
	row := q.db.QueryRowContext(ctx, readOneCategory, categoryID)
	var i ReadOneCategoryRow
	err := row.Scan(&i.CategoryID, &i.CategoryName, &i.ImageUrl)
	return i, err
}
