package products

import "github.com/gofiber/fiber/v2"

func SetRoute(app *fiber.App) {
	router := app.Group("/products")
	router.Get("/categories", readCategories)
	router.Get("/items", readProducts)
}
