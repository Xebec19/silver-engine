package cart

import (
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
