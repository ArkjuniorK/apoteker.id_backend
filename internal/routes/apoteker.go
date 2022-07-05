package routes

import (
	handlers "github.com/ArkjuniorK/apoteker.id_backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func SetupApotekerRoutes(r fiber.Router, l *zap.Logger) {
	apoteker := r.Group("/apoteker")
	// handler := handlers.New(l)

	// Read all apotekser
	apoteker.Get("/", handlers.GetApotekers)
	// Create a apoteker
	apoteker.Post("/", handlers.CreateApoteker)
	// // Read one apoteker
	apoteker.Get("/:apotekId", handlers.GetApoteker)
	// // Update one apoteker
	apoteker.Put("/:apotekId", handlers.UpdateApoteker)
	// // Delete one apoteker
	apoteker.Delete("/:apotekId", handlers.DeleteApoteker)
}
