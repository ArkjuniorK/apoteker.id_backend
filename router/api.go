package router

import (
	routes "github.com/ArkjuniorK/apoteker.id_backend/internal/routes"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRouter(app *fiber.App, log *zap.Logger, db *gorm.DB) {
	// group the api router
	api := app.Group("/api")

	// router for apotek
	routes.SetupApotekRoutes(api, log, db)

	// router for apoteker
	routes.SetupApotekerRoutes(api, log)
}
