package routes

import (
	handlers "github.com/ArkjuniorK/apoteker.id_backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupApotekRoutes(r fiber.Router, l *zap.Logger, db *gorm.DB) {
	apotek := r.Group("/apotek")
	handler := handlers.NewApotekHandler(l, db)

	// Read all apoteks
	apotek.Get("/list", handler.GetApoteks)
	// Create a apotek
	apotek.Post("/create", handler.CreateApotek)
	// Read one apotek
	apotek.Get("/detail/:uuid", handler.GetApotek)
	// Update one apotek
	apotek.Put("/update/:uuid", handler.UpdateApotek)
	// Delete one apotek
	apotek.Delete("/delete/:uuid", handler.DeleteApotek)
}
