package main

import (
	"log"
	"os"

	"github.com/Xebec19/silver-engine/auth"
	"github.com/Xebec19/silver-engine/cart"
	db "github.com/Xebec19/silver-engine/db/sqlc"
	_ "github.com/Xebec19/silver-engine/docs"
	"github.com/Xebec19/silver-engine/product"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title			khushi-backend
// @version		1.0
// @description	This is a Go application having JWT authentication, Unit tests,etc using postgresql as database
// @host			localhost:8080
// @BasePath		/
// @schemes		http
// @contact.name	Rohan Kumar Thakur
// @contact.email	rohandeveloper106@gmail.com
// @license.name	GNU GENERAL PUBLIC LICENSE
func main() {

	file, err := os.OpenFile("./my_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	loggerConfig := logger.Config{
		Output: file, // add file to save output
	}

	// load env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// set up rate limiter
	if config.Env != "development" {
		app.Use(limiter.New())
	}

	// set up logger
	app.Use(logger.New(loggerConfig))
	app.Use(logger.New())

	// set up cache
	app.Use(cache.New())

	// set up cors
	app.Use(cors.New())

	// Connect to database
	db.Connect()

	// render API Docs

	// Public Routes
	auth.SetRoute(app)
	product.SetRoute(app)

	app.Use(func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		claims, err := util.VerifyToken(token)

		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(util.ErrorResponse(err))
		}

		ctx.Locals("user", claims)

		ctx.Next()
		return nil
	})

	// Private Routes
	cart.SetRoute(app)

	// start server
	log.Printf("Server listening on %v", config.ServerAddress)
	app.Listen(config.ServerAddress)
}
