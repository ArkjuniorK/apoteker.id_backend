package routes

import (
	"github.com/ArkjuniorK/apoteker.id_backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func SetupApotekRoutes(r fiber.Router, l *zap.Logger) {
	apotek := r.Group("/apotek")
	handler := handlers.New(l)

	// Read all apoteks
	apotek.Get("/", handler.GetApoteks)
	// Create a apotek
	apotek.Post("/", handler.CreateApotek)
	// Read one apotek
	apotek.Get("/:apotekId", handler.GetApotek)
	// Update one apotek
	apotek.Put("/:apotekId", handler.UpdateApotek)
	// Delete one apotek
	apotek.Delete("/:apotekId", handler.DeleteApotek)
}
