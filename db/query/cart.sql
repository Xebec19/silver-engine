-- name: IsProductInCart :one
select case when exists (select product_id from v_cart vc where product_id = $1 and user_id = $2)
then 1 else 0 end is_product_in_cart;

-- name: ReadCartItemQuantity :one
select quantity from v_cart vc where product_id = $1 and user_id = $2;

-- name: UpdateCartItemQuantity :exec
update cart_details set quantity = $1 + quantity where product_id = $2 and cart_id = $3;

-- name: GetCartID :one
select cart_id from carts where user_id = $1;

-- name: InsertCartItem :exec
insert into cart_details (cart_id,product_id,product_price,quantity,delivery_price)
values($1,$2,(select price from products p where p.product_id = $2),$3,
(select delivery_price from products p where p.product_id = $2));

-- name: AddItemToCart :exec
CALL add_to_cart($1,$2,$3);

-- name: RemoveItemFromCart :exec
CALL remove_item($1,$2,$3);