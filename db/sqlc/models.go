// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"database/sql"
)

type Cart struct {
	CartID        int32          `json:"cart_id"`
	UserID        sql.NullInt32  `json:"user_id"`
	Price         sql.NullString `json:"price"`
	DeliveryPrice sql.NullString `json:"delivery_price"`
	CreatedOn     sql.NullTime   `json:"created_on"`
	UpdatedOn     sql.NullTime   `json:"updated_on"`
	Total         sql.NullString `json:"total"`
}

type CartDetail struct {
	CdID          int32          `json:"cd_id"`
	CartID        sql.NullInt32  `json:"cart_id"`
	ProductID     sql.NullInt32  `json:"product_id"`
	ProductPrice  sql.NullString `json:"product_price"`
	Quantity      sql.NullInt32  `json:"quantity"`
	DeliveryPrice sql.NullString `json:"delivery_price"`
}

type Category struct {
	CategoryID   int32          `json:"category_id"`
	CategoryName string         `json:"category_name"`
	CreatedOn    sql.NullTime   `json:"created_on"`
	ImageUrl     sql.NullString `json:"image_url"`
	Status       sql.NullString `json:"status"`
}

type Country struct {
	CountryID      int32          `json:"country_id"`
	CountryName    sql.NullString `json:"country_name"`
	Currency       sql.NullString `json:"currency"`
	CurrencySymbol sql.NullString `json:"currency_symbol"`
}

type Order struct {
	OrderID       string       `json:"order_id"`
	UserID        int32        `json:"user_id"`
	Price         string       `json:"price"`
	DeliveryPrice string       `json:"delivery_price"`
	Total         string       `json:"total"`
	CreatedOn     sql.NullTime `json:"created_on"`
	Email         string       `json:"email"`
	Address       string       `json:"address"`
}

type OrderDetail struct {
	OdID          int32          `json:"od_id"`
	OrderID       sql.NullString `json:"order_id"`
	ProductID     sql.NullInt32  `json:"product_id"`
	ProductPrice  sql.NullString `json:"product_price"`
	Quantity      sql.NullInt32  `json:"quantity"`
	DeliveryPrice sql.NullString `json:"delivery_price"`
}

type Product struct {
	ProductID     int32          `json:"product_id"`
	CategoryID    sql.NullInt32  `json:"category_id"`
	ProductName   string         `json:"product_name"`
	ProductImage  sql.NullString `json:"product_image"`
	Quantity      sql.NullInt32  `json:"quantity"`
	CreatedOn     sql.NullTime   `json:"created_on"`
	UpdatedOn     sql.NullTime   `json:"updated_on"`
	Status        sql.NullString `json:"status"`
	Price         string         `json:"price"`
	DeliveryPrice sql.NullString `json:"delivery_price"`
	ProductDesc   sql.NullString `json:"product_desc"`
	Gender        sql.NullString `json:"gender"`
	CountryID     sql.NullInt32  `json:"country_id"`
}

type Token struct {
	TokenID    int32         `json:"token_id"`
	UserID     sql.NullInt32 `json:"user_id"`
	Token      string        `json:"token"`
	CreatedOn  sql.NullTime  `json:"created_on"`
	LastAccess sql.NullTime  `json:"last_access"`
}

type User struct {
	UserID    int32          `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Email     string         `json:"email"`
	Phone     sql.NullInt64  `json:"phone"`
	Password  string         `json:"password"`
	CreatedOn sql.NullTime   `json:"created_on"`
	UpdatedOn sql.NullTime   `json:"updated_on"`
	Status    sql.NullString `json:"status"`
	Access    sql.NullString `json:"access"`
}

type VCart struct {
	UserID        sql.NullInt32  `json:"user_id"`
	CartID        int32          `json:"cart_id"`
	CardItemID    int32          `json:"card_item_id"`
	ProductID     sql.NullInt32  `json:"product_id"`
	Price         string         `json:"price"`
	Quantity      sql.NullInt32  `json:"quantity"`
	DeliveryPrice sql.NullString `json:"delivery_price"`
}

type VProduct struct {
	ProductID     int32          `json:"product_id"`
	ProductName   string         `json:"product_name"`
	ProductImage  sql.NullString `json:"product_image"`
	Quantity      sql.NullInt32  `json:"quantity"`
	CreatedOn     sql.NullTime   `json:"created_on"`
	Price         string         `json:"price"`
	DeliveryPrice sql.NullString `json:"delivery_price"`
	ProductDesc   sql.NullString `json:"product_desc"`
	Gender        sql.NullString `json:"gender"`
	CategoryID    int32          `json:"category_id"`
	CategoryName  string         `json:"category_name"`
	CountryID     int32          `json:"country_id"`
	CountryName   sql.NullString `json:"country_name"`
}
