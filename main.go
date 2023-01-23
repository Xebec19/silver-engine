package main

import (
	"log"

	"github.com/Xebec19/silver-engine/auth"
	db "github.com/Xebec19/silver-engine/db/sqlc"
	_ "github.com/Xebec19/silver-engine/docs"
	"github.com/Xebec19/silver-engine/product"
	"github.com/Xebec19/silver-engine/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

//	@title			khushi-backend
//	@version		1.0
//	@description	This is a Go application having JWT authentication, Unit tests,etc using postgresql as database
//	@host			localhost:8080
//	@BasePath		/
//	@schemes		http
//	@contact.name	Rohan Kumar Thakur
//	@contact.email	rohandeveloper106@gmail.com
//	@license.name	GNU GENERAL PUBLIC LICENSE
func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Connect to database
	db.Connect()

	// render API Docs

	// set Routes
	auth.SetRoute(app)
	product.SetRoute(app)

	// load env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// start server
	log.Printf("Server listening on %v", config.ServerAddress)
	app.Listen(config.ServerAddress)
}
