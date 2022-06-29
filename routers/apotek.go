package router

import (
	"apoteker.id_backend/handlers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func SetupApotek(l *zap.Logger, r fiber.Router) {
	// initiate Apotek handlers
	handler := handlers.NewApotek(l)

	r.Get("/", handler.Home)
}
