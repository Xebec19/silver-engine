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

// readCategoryDetails goDoc
//
//	@Summary	read one category
//	@Param		id	path	string	true	"Category ID"
//	@Success	201	{ data:interface{}, message:string, status:string}
//	@Router		/category/v1/details/:id [get]
func readCategoryDetails(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Params("id", "0"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	categoryDetails, err := db.DBQuery.ReadOneCategory(context.Background(), int32(categoryId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(categoryDetails, "Category details fetched"))
}
