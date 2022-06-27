package main

import (
	"apoteker.id_backend/config"
	"apoteker.id_backend/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// create new fiber app
	app := fiber.New()

	router.SetupRoutes(app)

	// get envs
	port := config.Config("PORT")

	// listen to port
	app.Listen(":" + port)
}
