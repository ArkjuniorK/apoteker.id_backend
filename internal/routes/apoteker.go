package routes

import (
	"github.com/ArkjuniorK/apoteker.id_backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupApotekerRoutes(r fiber.Router, l *zap.Logger, d *gorm.DB) {
	apoteker := r.Group("/apoteker")
	handler := handlers.NewApotekerHandler(l, d)

	// Read all apotekser
	apoteker.Get("/", handler.GetApotekers)
	// Create a apoteker
	apoteker.Post("/", handler.CreateApoteker)
	// // Read one apoteker
	apoteker.Get("/:apotekId", handler.GetApoteker)
	// // Update one apoteker
	apoteker.Put("/:apotekId", handler.UpdateApoteker)
	// // Delete one apoteker
	apoteker.Delete("/:apotekId", handler.DeleteApoteker)
}
