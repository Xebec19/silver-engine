package cart

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/Xebec19/silver-engine/db/sqlc"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber/v2"
)

type updateCartSchema struct {
	ProductId int32 `json:"product_id" binding:"required"`
	Quantity  int32 `json:"quantity" binding:"required"`
}

// @Summary Add a product to user's cart
// @Router /cart/add-product
func addProductIntoCart(c *fiber.Ctx) error {
	req := new(updateCartSchema)
	userId := c.Locals("userid").(int64)
	// parse and validate request body
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	args := db.AddItemToCartParams{
		ProductID: req.ProductId,
		UserID:    int32(userId),
		Quantity:  req.Quantity,
	}

	/**
	* todo get available quantity of product
	* todo available quantity should not be greater than given quantity of the product
	* todo check if product exists in cart details table
	* todo if exists and given quantity + existing quantity < available quantity update cart details quantity
	* todo else create new entry in cart details
	 */
	quantity, err := db.DBQuery.ReadQuantity(c.Context(), req.ProductId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	if req.Quantity > quantity.Int32 {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(errors.New("out of stock")))
	}

	cartId, err := db.DBQuery.GetCartID(context.Background(), sql.NullInt32{Int32: int32(userId), Valid: true})

	cartDetailArgs := db.CheckCartDetailParams{
		CartID:    sql.NullInt32{Int32: int32(cartId), Valid: true},
		ProductID: sql.NullInt32{Int32: int32(req.ProductId), Valid: true},
	}

	cartDetail, err := db.DBQuery.CheckCartDetail(c.Context(), cartDetailArgs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	if cartDetail == 0 {
		insertCartItemArgv := db.InsertCartItemParams{
			CartID:    sql.NullInt32{Int32: int32(cartId), Valid: true},
			ProductID: sql.NullInt32{Int32: int32(req.ProductId), Valid: true},
			Quantity:  sql.NullInt32{Int32: int32(req.Quantity), Valid: true},
		}

		// todo insert new item in cart details
	}

	// todo update existing item in cart details

	// call a procedure in db, to add product in cart
	db.DBQuery.AddItemToCart(c.Context(), args)

	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(nil, "Product Added Successfully"))
}

// @Summary Remove a product from user's cart
// @Router /cart/remove-product
func removeProductFromCart(c *fiber.Ctx) error {
	req := new(updateCartSchema)
	userId := c.Locals("userid").(int64)
	// parse and validate request body
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	args := db.RemoveItemFromCartParams{
		ProductID: req.ProductId,
		UserID:    int32(userId),
		Quantity:  req.Quantity,
	}

	// call a procedure in db, to remove products from cart
	db.DBQuery.RemoveItemFromCart(c.Context(), args)

	return c.Status(fiber.StatusAccepted).JSON(util.SuccessResponse(nil, "Product Removed Successfully"))
}
