-- name: ReadProductQuantityInCard :one
select sum(quantity) as quantity from v_cart vc group by product_id having product_id = $1 and user_id = $2 ;

-- name: UpdateProductQuantity :exec
update cart_details set quantity = $1 + quantity where product_id = $2 and cart_id = $3;

-- name: GetCartID :one
select cart_id from carts where user_id = $1;

-- name: InsertCartItem :exec
insert into cart_details (cart_id,product_id,product_price,quantity,delivery_price)
values($1,$2,(select price from products p where p.product_id = $2),$3,
(select delivery_price from products p where p.product_id = $2));