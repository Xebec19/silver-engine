package util

import "github.com/gofiber/fiber"

// func Validate(schema interface{}) func(*fiber.Ctx) error {
// 	return func(c *fiber.Ctx) error {
// 		// Parse the request body into a map
// 		var body map[string]interface{}
// 		err := c.BodyParser(&body)
// 		if err != nil {
// 			return err
// 		}

// 		// Validate the request body against the given schema
// 		err = govalidator.New(govalidator.Options{
// 			Request: c.Request(),
// 			Data:    body,
// 			Rules:   schema,
// 		}).ValidateJSON()
// 		if err != nil {
// 			return err
// 		}

// 		// If the request body is valid, continue with the next middleware
// 		c.Next()
// 		return nil
// 	}
// }

func Validate(schema interface{}) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Next()
		return nil
	}
}
