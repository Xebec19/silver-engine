// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: products.sql

package db

import (
	"context"
	"database/sql"
)

const readAllProducts = `-- name: ReadAllProducts :many
SELECT product_id, product_name, product_image, quantity, product_desc from v_products limit $1 offset $2
`

type ReadAllProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type ReadAllProductsRow struct {
	ProductID    int32          `json:"product_id"`
	ProductName  string         `json:"product_name"`
	ProductImage sql.NullString `json:"product_image"`
	Quantity     sql.NullInt32  `json:"quantity"`
	ProductDesc  sql.NullString `json:"product_desc"`
}

func (q *Queries) ReadAllProducts(ctx context.Context, arg ReadAllProductsParams) ([]ReadAllProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, readAllProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadAllProductsRow
	for rows.Next() {
		var i ReadAllProductsRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.ProductImage,
			&i.Quantity,
			&i.ProductDesc,
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

const readCategoryItems = `-- name: ReadCategoryItems :many
SELECT product_id, product_name, product_image, quantity, product_desc from v_products where category_id = $1 limit $2 offset $3
`

type ReadCategoryItemsParams struct {
	CategoryID int32 `json:"category_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

type ReadCategoryItemsRow struct {
	ProductID    int32          `json:"product_id"`
	ProductName  string         `json:"product_name"`
	ProductImage sql.NullString `json:"product_image"`
	Quantity     sql.NullInt32  `json:"quantity"`
	ProductDesc  sql.NullString `json:"product_desc"`
}

func (q *Queries) ReadCategoryItems(ctx context.Context, arg ReadCategoryItemsParams) ([]ReadCategoryItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, readCategoryItems, arg.CategoryID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ReadCategoryItemsRow
	for rows.Next() {
		var i ReadCategoryItemsRow
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.ProductImage,
			&i.Quantity,
			&i.ProductDesc,
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

const readCategoryProduct = `-- name: ReadCategoryProduct :many
SELECT product_id, product_name, product_image, quantity, created_on, price, delivery_price, product_desc, gender, category_id, category_name, country_id, country_name FROM v_products WHERE category_id = $1
`

func (q *Queries) ReadCategoryProduct(ctx context.Context, categoryID int32) ([]VProduct, error) {
	rows, err := q.db.QueryContext(ctx, readCategoryProduct, categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VProduct
	for rows.Next() {
		var i VProduct
		if err := rows.Scan(
			&i.ProductID,
			&i.ProductName,
			&i.ProductImage,
			&i.Quantity,
			&i.CreatedOn,
			&i.Price,
			&i.DeliveryPrice,
			&i.ProductDesc,
			&i.Gender,
			&i.CategoryID,
			&i.CategoryName,
			&i.CountryID,
			&i.CountryName,
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

const readOneProduct = `-- name: ReadOneProduct :one
SELECT product_id, product_name, product_image, product_desc, price, delivery_price, category_id, category_name from v_products where product_id = $1
`

type ReadOneProductRow struct {
	ProductID     int32          `json:"product_id"`
	ProductName   string         `json:"product_name"`
	ProductImage  sql.NullString `json:"product_image"`
	ProductDesc   sql.NullString `json:"product_desc"`
	Price         string         `json:"price"`
	DeliveryPrice sql.NullString `json:"delivery_price"`
	CategoryID    int32          `json:"category_id"`
	CategoryName  string         `json:"category_name"`
}

func (q *Queries) ReadOneProduct(ctx context.Context, productID int32) (ReadOneProductRow, error) {
	row := q.db.QueryRowContext(ctx, readOneProduct, productID)
	var i ReadOneProductRow
	err := row.Scan(
		&i.ProductID,
		&i.ProductName,
		&i.ProductImage,
		&i.ProductDesc,
		&i.Price,
		&i.DeliveryPrice,
		&i.CategoryID,
		&i.CategoryName,
	)
	return i, err
}

const readQuantity = `-- name: ReadQuantity :one
SELECT quantity from v_products where product_id = $1
`

func (q *Queries) ReadQuantity(ctx context.Context, productID int32) (sql.NullInt32, error) {
	row := q.db.QueryRowContext(ctx, readQuantity, productID)
	var quantity sql.NullInt32
	err := row.Scan(&quantity)
	return quantity, err
}
