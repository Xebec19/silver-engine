package main

import (
	"log"

	"github.com/Xebec19/miniature-giggle/backend/auth"
	db "github.com/Xebec19/miniature-giggle/backend/db/sqlc"
	"github.com/Xebec19/miniature-giggle/backend/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Connect to database
	db.Connect()

	// set Routes
	auth.SetRoute(app)

	// load env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// start server
	log.Printf("Server listening on %v", config.ServerAddress)
	app.Listen(config.ServerAddress)
}
