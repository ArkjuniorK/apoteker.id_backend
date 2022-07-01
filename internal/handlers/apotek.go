package handlers

import "github.com/gofiber/fiber/v2"

func GetApoteks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func CreateApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func GetApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func UpdateApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func DeleteApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}
