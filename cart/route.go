package cart

import "github.com/gofiber/fiber"

func SetRoute(app *fiber.App) {
	router := app.Group("/cart")
	router.Post("/add-product")
}
