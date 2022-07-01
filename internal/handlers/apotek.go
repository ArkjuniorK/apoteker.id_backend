package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ApotekHandler struct {
	log *zap.Logger
}

func New(l *zap.Logger) *ApotekHandler {
	return &ApotekHandler{
		log: l,
	}
}

func (a ApotekHandler) GetApoteks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func (a ApotekHandler) CreateApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func (a ApotekHandler) GetApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func (a ApotekHandler) UpdateApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}

func (a ApotekHandler) DeleteApotek(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": ""})
}
