package category

import "github.com/gofiber/fiber/v2"

func SetRoute(app *fiber.App) {
	router := app.Group("/category")
	router.Get("/list", readCategories)
	router.Get("/detail/:category-id", readCategoryWiseProducts)
}
