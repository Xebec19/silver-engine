table "bazaar_cart_details" {
  schema = schema.public
  column "cd_id" {
    null = false
    type = serial
  }
  column "cart_id" {
    null = true
    type = integer
  }
  column "product_id" {
    null = true
    type = integer
  }
  column "product_price" {
    null = true
    type = numeric
  }
  column "quantity" {
    null = true
    type = integer
  }
  column "delivery_price" {
    null = true
    type = numeric
  }
  primary_key {
    columns = [column.cd_id]
  }
  foreign_key "fk_cart" {
    columns     = [column.cart_id]
    ref_columns = [table.bazaar_carts.column.cart_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "fk_product" {
    columns     = [column.product_id]
    ref_columns = [table.bazaar_products.column.product_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "bazaar_carts" {
  schema = schema.public
  column "cart_id" {
    null = false
    type = serial
  }
  column "user_id" {
    null = true
    type = integer
  }
  column "price" {
    null = true
    type = numeric
  }
  column "delivery_price" {
    null = true
    type = numeric
  }
  column "created_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "updated_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "total" {
    null = true
    type = numeric
  }
  primary_key {
    columns = [column.cart_id]
  }
  foreign_key "fk_users" {
    columns     = [column.user_id]
    ref_columns = [table.bazaar_users.column.user_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "bazaar_categories" {
  schema = schema.public
  column "category_id" {
    null = false
    type = serial
  }
  column "category_name" {
    null = false
    type = text
  }
  column "created_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "status" {
    null = true
    type = text
  }
  primary_key {
    columns = [column.category_id]
  }
}
table "bazaar_countries" {
  schema = schema.public
  column "country_id" {
    null = false
    type = serial
  }
  column "country_name" {
    null = true
    type = character_varying(20)
  }
  column "currency" {
    null = true
    type = character_varying(10)
  }
  column "currency_symbol" {
    null = true
    type = character_varying(5)
  }
  primary_key {
    columns = [column.country_id]
  }
}
table "bazaar_order" {
  schema = schema.public
  column "order_id" {
    null = false
    type = character_varying(20)
  }
  column "user_id" {
    null = false
    type = integer
  }
  column "price" {
    null = false
    type = numeric
  }
  column "delivery_price" {
    null = false
    type = numeric
  }
  column "total" {
    null = false
    type = numeric
  }
  column "created_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "email" {
    null = false
    type = character_varying(50)
  }
  column "address" {
    null = false
    type = text
  }
  primary_key {
    columns = [column.order_id]
  }
  foreign_key "fk_user" {
    columns     = [column.user_id]
    ref_columns = [table.bazaar_users.column.user_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "bazaar_order_details" {
  schema = schema.public
  column "od_id" {
    null = false
    type = serial
  }
  column "order_id" {
    null = true
    type = character_varying(20)
  }
  column "product_id" {
    null = true
    type = integer
  }
  column "product_price" {
    null = true
    type = numeric
  }
  column "quantity" {
    null = true
    type = integer
  }
  column "delivery_price" {
    null = true
    type = numeric
  }
  primary_key {
    columns = [column.od_id]
  }
  foreign_key "fk_order" {
    columns     = [column.order_id]
    ref_columns = [table.bazaar_order.column.order_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "fk_product" {
    columns     = [column.product_id]
    ref_columns = [table.bazaar_products.column.product_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "bazaar_products" {
  schema = schema.public
  column "product_id" {
    null = false
    type = serial
  }
  column "category_id" {
    null = true
    type = integer
  }
  column "product_name" {
    null = false
    type = character_varying
  }
  column "product_image" {
    null = true
    type = character_varying
  }
  column "quantity" {
    null = true
    type = integer
  }
  column "created_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "updated_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "status" {
    null    = true
    type    = text
    default = "active"
  }
  column "price" {
    null = false
    type = numeric
  }
  column "delivery_price" {
    null    = true
    type    = numeric
    default = 0
  }
  column "product_desc" {
    null    = true
    type    = character_varying(500)
    default = "No description"
  }
  column "gender" {
    null    = true
    type    = character_varying(10)
    default = "Not specified"
  }
  column "country_id" {
    null    = true
    type    = integer
    default = 1
  }
  primary_key {
    columns = [column.product_id]
  }
  foreign_key "fk_category" {
    columns     = [column.category_id]
    ref_columns = [table.bazaar_categories.column.category_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  foreign_key "fk_country" {
    columns     = [column.country_id]
    ref_columns = [table.bazaar_countries.column.country_id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}
table "bazaar_tokens" {
  schema = schema.public
  column "token_id" {
    null = false
    type = serial
  }
  column "user_id" {
    null = true
    type = integer
  }
  column "token" {
    null = false
    type = character_varying
  }
  column "created_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "last_access" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  primary_key {
    columns = [column.token_id]
  }
  foreign_key "fk_usertoken" {
    columns     = [column.user_id]
    ref_columns = [table.bazaar_users.column.user_id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
}
table "bazaar_users" {
  schema = schema.public
  column "user_id" {
    null = false
    type = serial
  }
  column "first_name" {
    null = false
    type = text
  }
  column "last_name" {
    null = true
    type = text
  }
  column "email" {
    null = false
    type = character_varying
  }
  column "phone" {
    null = true
    type = bigint
  }
  column "password" {
    null = false
    type = character_varying
  }
  column "created_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "updated_on" {
    null    = true
    type    = timestamptz
    default = sql("now()")
  }
  column "status" {
    null    = true
    type    = character_varying(10)
    default = "active"
  }
  column "access" {
    null    = true
    type    = character_varying(20)
    default = "user"
  }
  primary_key {
    columns = [column.user_id]
  }
  index "bazaar_users_email_key" {
    unique  = true
    columns = [column.email]
  }
}
schema "public" {
}
