package cart

import (
	"database/sql"
	"errors"

	db "github.com/Xebec19/silver-engine/db/sqlc"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber/v2"
)

type addProductSchema struct {
	ProductId int32 `json:"product_id" binding:"required"`
	Quantity  int32 `json:"quantity" binding:"required"`
}

// @Summary Add a product to users cart
// @Router /cart/add-product
func addProductIntoCart(c *fiber.Ctx) error {
	req := new(addProductSchema)
	userId := c.Locals("userid").(int64)
	// parse and validate request body
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	stock, err := db.DBQuery.ReadQuantity(c.Context(), req.ProductId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	if req.Quantity > int32(stock.Int32) {
		return c.Status(fiber.StatusNotAcceptable).JSON(util.ErrorResponse(errors.New("out of stock")))
	}

	args := db.IsProductInCartParams{
		ProductID: sql.NullInt32{Int32: req.ProductId, Valid: true},
		UserID:    sql.NullInt32{Int32: int32(userId), Valid: true},
	}

	isProductInCart, err := db.DBQuery.IsProductInCart(c.Context(), args)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	if isProductInCart.(int64) != int64(0) {
		args := db.ReadCartItemQuantityParams{
			ProductID: sql.NullInt32{Int32: req.ProductId, Valid: true},
			UserID:    sql.NullInt32{Int32: int32(userId), Valid: true},
		}

		prevQuantity, err := db.DBQuery.ReadCartItemQuantity(c.Context(), args)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
		}

		if prevQuantity.Int32+req.Quantity > stock.Int32 {
			return c.Status(fiber.StatusNotAcceptable).JSON(util.ErrorResponse(errors.New("out of stock")))
		}

		cartId, err := db.DBQuery.GetCartID(c.Context(), sql.NullInt32{Int32: int32(userId), Valid: true})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
		}

		args2 := db.UpdateCartItemQuantityParams{
			ProductID: sql.NullInt32{Int32: req.ProductId, Valid: true},
			CartID:    sql.NullInt32{Int32: cartId, Valid: true},
			Quantity:  sql.NullInt32{Int32: int32(req.Quantity), Valid: true},
		}

		db.DBQuery.UpdateCartItemQuantity(c.Context(), args2)

		return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(nil, "Quantity updated in cart"))
	}

	getCartIDParams := sql.NullInt32{Int32: int32(userId), Valid: true}

	cartId, err := db.DBQuery.GetCartID(c.Context(), getCartIDParams)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	insertCartItemParams := db.InsertCartItemParams{
		CartID:    sql.NullInt32{Int32: int32(cartId), Valid: true},
		ProductID: sql.NullInt32{Int32: int32(req.ProductId), Valid: true},
		Quantity:  sql.NullInt32{Int32: int32(req.Quantity), Valid: true},
	}

	db.DBQuery.InsertCartItem(c.Context(), insertCartItemParams)
	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(nil, "Quantity updated in cart"))
}
