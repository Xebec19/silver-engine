package cart

import (
	"context"

	db "github.com/Xebec19/silver-engine/db/sqlc"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber"
)

type addProductSchema struct {
	ProductId int64 `json:"product_id"binding:"required"`
	Quantity  int16 `json:"quantity"binding:"required"`
}

func addProductIntoCart(c *fiber.Ctx) error {
	req := new(addProductSchema)
	// parse and validate request body
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	// check if required quantity is in stock
	quantity, err := db.DBQuery.ReadQuantity(context.Background(), int32(req.ProductId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	if req.Quantity > int16(quantity.Int32) {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	/*
		todo if product exists in user cart
		todo update quantity else create new entry in cart
	*/

	return nil
}
