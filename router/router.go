package router

import (
	apotekRoutes "apoteker.id_backend/internal/routes/apotek"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func SetupRouter(app *fiber.App, log *zap.Logger) {
	// group the api router
	api := app.Group("/api")

	// router for apotek
	apotekRoutes.SetupApotekRoutes(api)

}
