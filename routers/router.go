package routers

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	api := app.Group("api")

	apotek := api.Group("apotek")

	apotek.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "wolrd",
		})
	})

}
