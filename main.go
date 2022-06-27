package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Unable to load env. Error: %s", err)
	}

	// get envs
	port := os.Getenv("PORT")

	// create new fiber app
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"hello": "world",
		})
	})

	// listen to port
	app.Listen(":" + port)
}
