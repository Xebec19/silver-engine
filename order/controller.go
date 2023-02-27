package order

import (
	db "github.com/Xebec19/silver-engine/db/sqlc"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber"
)

// @Summary create an order
// @Router /order/create-order
func createOrder(c *fiber.Ctx) error {
	userId := c.Locals("userid").(int64)

	db.DBQuery.CreateOrder(c.Context(), int32(userId))

	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(nil, "Order Created"))
}

// @Summary get list of orders in paginated way
func listOrders(c *fiber.Ctx) error {
	userId := c.Locals("userid").(int64)

	orders, err := db.DBQuery.ListOrder(c.Context(), int32(userId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.SuccessResponse(nil, "Order Creation Failed"))
	}

	return c.Status(fiber.StatusFound).JSON(util.SuccessResponse(orders, "Orders Fetched"))
}

// @Summary It sends list of items included in an order
// @Router /order/order-details
func listOrderItems(c *fiber.Ctx) error {
	return c.Status(fiber.StatusFound).JSON(util.SuccessResponse(nil, "Order Items Fetched"))
}
