package category

import (
	"context"
	"strconv"

	db "github.com/Xebec19/silver-engine/db/sqlc"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber/v2"
)

func readCategories(c *fiber.Ctx) error {
	category, err := db.DBQuery.ReadAllCategories(context.Background())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(category, "Category fetched successfully"))
}

func readCategoryWiseProducts(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	category, err := db.DBQuery.ReadOneCategory(context.Background(), int32(categoryId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(util.SuccessResponse(category, "Category fetched successfully"))
}
