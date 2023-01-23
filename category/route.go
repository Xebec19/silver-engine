package category

import "github.com/gofiber/fiber/v2"

func SetRoute(app *fiber.App) {
	router := app.Group("/category")
	router.Get("/v1/list", readCategories)
	router.Get("/v1/details/:id", readCategoryWiseProducts)
}
