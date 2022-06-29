package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Apotek struct {
	log *zap.Logger
}

// initiate new apotek structure
func NewApotek(logger *zap.Logger) *Apotek {
	return &Apotek{
		log: logger,
	}
}

// Home handler that encapsulated to Apotek struct
func (a *Apotek) Home(c *fiber.Ctx) error {
	a.log.Info("here at home")
	return c.JSON(&fiber.Map{
		"home": "hello apotek!",
	})
}
