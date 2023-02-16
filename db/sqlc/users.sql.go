// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: users.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO USERS (first_name, last_name, email, phone, password)
values ($1,$2,$3,$4,$5) RETURNING user_id
`

type CreateUserParams struct {
	FirstName string         `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Email     string         `json:"email"`
	Phone     sql.NullInt64  `json:"phone"`
	Password  string         `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.Phone,
		arg.Password,
	)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}

const findUserOne = `-- name: FindUserOne :one
SELECT user_id, email, CONCAT(first_name, ' ', last_name) AS user_name, password FROM USERS u 
WHERE u.EMAIL = $1
`

type FindUserOneRow struct {
	UserID   int32       `json:"user_id"`
	Email    string      `json:"email"`
	UserName interface{} `json:"user_name"`
	Password string      `json:"password"`
}

func (q *Queries) FindUserOne(ctx context.Context, email string) (FindUserOneRow, error) {
	row := q.db.QueryRowContext(ctx, findUserOne, email)
	var i FindUserOneRow
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.UserName,
		&i.Password,
	)
	return i, err
}
