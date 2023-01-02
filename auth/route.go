package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ValidateWrapper(s interface{}) func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		// Parse the request body into the provided struct
		err := json.Unmarshal(c.Body(), s)
		if err != nil {
			fmt.Println(err)
			// If the request body doesn't validate to the struct, return an error
			c.SendStatus(http.StatusBadRequest)
			return
		}
		// If the request body validates to the struct, continue processing the request
		c.Next()
	}
}

func SetRoute(app *fiber.App) {
	router := app.Group("/auth")
	router.Post("/register", register)
	router.Post("/login", login)
}
