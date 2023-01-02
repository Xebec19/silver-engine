package products

import (
	"context"

	db "github.com/Xebec19/miniature-giggle/backend/db/sqlc"
	"github.com/Xebec19/miniature-giggle/backend/util"
	"github.com/gofiber/fiber/v2"
)

// read categories sends category details
func readCategories(c *fiber.Ctx) error {
	category, err := db.DBQuery.ReadAllCategories(context.Background())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(category, "Category fetched successfully"))
}
