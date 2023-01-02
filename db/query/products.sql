-- name: ReadCategoryProduct :many
SELECT * FROM v_products WHERE category_id = $1;