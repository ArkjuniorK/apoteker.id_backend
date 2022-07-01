package routes

import (
	apotekHandler "github.com/ArkjuniorK/apoteker.id_backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupApotekRoutes(r fiber.Router) {
	apotek := r.Group("/apotek")

	// Read all apoteks
	apotek.Get("/", apotekHandler.GetApoteks)
	// Create a apotek
	apotek.Post("/", apotekHandler.CreateApotek)
	// Read one apotek
	apotek.Get("/:apotekId", apotekHandler.GetApotek)
	// Update one apotek
	apotek.Put("/:apotekId", apotekHandler.UpdateApotek)
	// Delete one apotek
	apotek.Delete("/:apotekId", apotekHandler.DeleteApotek)
}
