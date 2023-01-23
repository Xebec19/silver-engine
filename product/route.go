package product

import "github.com/gofiber/fiber/v2"

func SetRoute(app *fiber.App) {
	router := app.Group("/category")
	router.Get("/v1/list", readCategories)
	router.Get("/v1/detail/:category-id", readCategoryWiseProducts)
}
