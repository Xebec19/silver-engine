package cart

import (
	"context"
	"database/sql"
	"errors"

	db "github.com/Xebec19/silver-engine/db/sqlc"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type addProductSchema struct {
	ProductId int32 `json:"product_id" binding:"required"`
	Quantity  int16 `json:"quantity" binding:"required"`
}

// @Summary Add a product to users cart
// @Router /cart/add-product
func addProductIntoCart(c *fiber.Ctx) error {
	req := new(addProductSchema)
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userid"].(int32)
	// parse and validate request body
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	args := db.ReadProductQuantityInCardParams{
		ProductID: sql.NullInt32{Int32: req.ProductId, Valid: true},
		UserID:    sql.NullInt32{Int32: userId, Valid: true},
	}

	quantity, err := db.DBQuery.ReadProductQuantityInCard(context.Background(), args)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	// fetch current stock of given product
	stock, err := db.DBQuery.ReadQuantity(context.Background(), int32(req.ProductId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	// if cart already has given item, we simply increment
	// the quantity
	if quantity > 0 {
		// check if merchant have enough stock to complete order
		if int16(quantity)+req.Quantity > int16(stock.Int32) {
			return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(errors.New("out of stock")))
		}

		cartId, err := db.DBQuery.GetCartID(context.Background(), sql.NullInt32{Int32: userId, Valid: true})

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
		}

		args := db.UpdateProductQuantityParams{
			Quantity:  sql.NullInt32{Int32: int32(req.Quantity), Valid: true},
			ProductID: sql.NullInt32{Int32: int32(req.ProductId), Valid: true},
			CartID:    sql.NullInt32{Int32: cartId, Valid: true},
		}

		db.DBQuery.UpdateProductQuantity(context.Background(), args)

		return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(nil, "Quantity updated in cart"))
	}

	/*
		given we have handled the case when
		product is already available in cart, we now handle the case when product is newly added to
		cart
	*/
	cartId, err := db.DBQuery.GetCartID(context.Background(), sql.NullInt32{Int32: userId, Valid: true})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	args2 := db.InsertCartItemParams{
		CartID:    sql.NullInt32{Int32: cartId, Valid: true},
		ProductID: sql.NullInt32{Int32: req.ProductId, Valid: true},
		Quantity:  sql.NullInt32{Int32: int32(quantity), Valid: true},
	}

	db.DBQuery.InsertCartItem(context.Background(), args2)

	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(nil, "Product added in cart"))
}
