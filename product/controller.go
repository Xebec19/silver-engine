package product

import (
	"context"
	"strconv"

	db "github.com/Xebec19/silver-engine/db/sqlc"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber/v2"
)

// @Summary	Read all categories
// @Param		request	query		string	true	"query params"
// @Success	201
// @Router		/category/v1/list [get]
func readCategories(c *fiber.Ctx) error {
	category, err := db.DBQuery.ReadAllCategories(context.Background())

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	return c.Status(fiber.StatusCreated).JSON(util.SuccessResponse(category, "Category fetched successfully"))
}

/*
/products/items [get]
read products sends array of products in paginated form
*/
//	@Summary	Read all categories
//	@Param		request	query		string	true	"query params"
//	@Success	201
//	@Router		/category/v1/list [get]
func readProducts(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	size, err := strconv.Atoi(c.Query("size", "0"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	args := db.ReadAllProductsParams{
		Limit:  int32(size),
		Offset: int32(page * size),
	}

	products, err := db.DBQuery.ReadAllProducts(context.Background(), args)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(util.SuccessResponse(products, "Products fetched successfully"))
}

/*
/products/category/:category-id [get]
get products in a category
*/
func readCategoryWiseProducts(c *fiber.Ctx) error {
	categoryId, err := strconv.Atoi(c.Query("category-id", "0"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(util.ErrorResponse(err))
	}

	category, err := db.DBQuery.ReadOneCategory(context.Background(), int32(categoryId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(util.ErrorResponse(err))
	}

	return c.Status(fiber.StatusOK).JSON(util.SuccessResponse(category, "Category fetched successfully"))
}
