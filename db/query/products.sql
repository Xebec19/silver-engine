-- name: ReadCategoryProduct :many
SELECT * FROM v_products WHERE category_id = $1;

-- name: ReadAllProducts :many
SELECT product_id, product_name, product_image, quantity, product_desc from v_products limit $1 offset $2;

-- name: ReadCategoryItems :many
SELECT product_id, product_name, product_image, quantity, product_desc from v_products where category_id = $1 limit $2 offset $3;