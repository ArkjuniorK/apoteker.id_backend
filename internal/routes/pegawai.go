package routes

import (
	"github.com/ArkjuniorK/apoteker.id_backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupPegawaiRoutes(r fiber.Router, l *zap.Logger, d *gorm.DB) {
	pegawai := r.Group("/pegawai")
	handler := handlers.NewPegawaiHandler(l, d)

	// // Read all pegawais
	pegawai.Get("/", handler.GetPegawais)
	// Create a apoteker
	pegawai.Post("/", handler.CreatePegawai)
	// // Read one pegawai
	pegawai.Get("/:apotekId", handler.GetPegawai)
	// // Update one pegawai
	pegawai.Put("/:apotekId", handler.UpdatePegawai)
	// // Delete one pegawai
	pegawai.Delete("/:apotekId", handler.DeletePegawai)
}
