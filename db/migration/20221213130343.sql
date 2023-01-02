create or replace view v_products as
select bp.product_id, bp.product_name, bp.product_image, bp.quantity, bp.created_on, 
bp.price, bp.delivery_price, bp.product_desc, bp.gender, bc.category_id, bc.category_name, bc2.country_id, bc2.country_name  
from products bp
join categories bc ON bp.category_id = bc.category_id 
join countries bc2 on bc2.country_id = bp.country_id 
where bc.status = 'active' and bp.status = 'active';
