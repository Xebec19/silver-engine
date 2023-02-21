// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: cart.sql

package db

import (
	"context"
	"database/sql"
)

const getCartID = `-- name: GetCartID :one
select cart_id from carts where user_id = $1
`

func (q *Queries) GetCartID(ctx context.Context, userID sql.NullInt32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getCartID, userID)
	var cart_id int32
	err := row.Scan(&cart_id)
	return cart_id, err
}

const insertCartItem = `-- name: InsertCartItem :exec
insert into cart_details (cart_id,product_id,product_price,quantity,delivery_price)
values($1,$2,(select price from products p where p.product_id = $2),$3,
(select delivery_price from products p where p.product_id = $2))
`

type InsertCartItemParams struct {
	CartID    sql.NullInt32 `json:"cart_id"`
	ProductID sql.NullInt32 `json:"product_id"`
	Quantity  sql.NullInt32 `json:"quantity"`
}

func (q *Queries) InsertCartItem(ctx context.Context, arg InsertCartItemParams) error {
	_, err := q.db.ExecContext(ctx, insertCartItem, arg.CartID, arg.ProductID, arg.Quantity)
	return err
}

const readProductQuantityInCard = `-- name: ReadProductQuantityInCard :one
select sum(quantity) as quantity from v_cart vc group by product_id having product_id = $1 and user_id = $2
`

type ReadProductQuantityInCardParams struct {
	ProductID sql.NullInt32 `json:"product_id"`
	UserID    sql.NullInt32 `json:"user_id"`
}

func (q *Queries) ReadProductQuantityInCard(ctx context.Context, arg ReadProductQuantityInCardParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, readProductQuantityInCard, arg.ProductID, arg.UserID)
	var quantity int64
	err := row.Scan(&quantity)
	return quantity, err
}

const updateProductQuantity = `-- name: UpdateProductQuantity :exec
update cart_details set quantity = $1 + quantity where product_id = $2 and cart_id = $3
`

type UpdateProductQuantityParams struct {
	Quantity  sql.NullInt32 `json:"quantity"`
	ProductID sql.NullInt32 `json:"product_id"`
	CartID    sql.NullInt32 `json:"cart_id"`
}

func (q *Queries) UpdateProductQuantity(ctx context.Context, arg UpdateProductQuantityParams) error {
	_, err := q.db.ExecContext(ctx, updateProductQuantity, arg.Quantity, arg.ProductID, arg.CartID)
	return err
}