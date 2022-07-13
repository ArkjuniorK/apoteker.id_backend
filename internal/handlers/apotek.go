package handlers

import (
	"errors"

	"github.com/ArkjuniorK/apoteker.id_backend/database"
	"github.com/ArkjuniorK/apoteker.id_backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ApotekHandler struct {
	log *zap.Logger
	db  *gorm.DB
}

// Function to generate new ApotekHandler
//
// To add new feature to handler, pass the feature as parameter
func NewApotekHandler(l *zap.Logger, d *gorm.DB) *ApotekHandler {
	return &ApotekHandler{
		log: l,
		db:  d,
	}
}

// List all apotek, it would be use by Admin only
func (a *ApotekHandler) GetApoteks(c *fiber.Ctx) error {
	var apoteks []model.Apotek

	// find all the apotek and send to client
	a.db.Model(&model.Apotek{}).Preload("Apotekers").Find(&apoteks)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get all apoteks",
		"data":    apoteks,
	})
}

// Create new apotek by given body
func (a *ApotekHandler) CreateApotek(c *fiber.Ctx) error {
	var apotek model.Apotek

	// parse body to apotek struct and check the error
	if err := c.BodyParser(&apotek); err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}

	a.log.Sugar().Info(apotek)

	// add apotek data into database
	// and send the data to response
	a.db.Create(&apotek)
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "Apotek created successfully",
		"data":    apotek,
	})
}

// Get detail apotek by given UUID
func (a *ApotekHandler) GetApotek(c *fiber.Ctx) error {
	var apotek model.Apotek

	// get the params and parse to UUID
	// and check the error
	apotek_id := c.Params("apotekId")
	uid, err := uuid.Parse(apotek_id)
	if err != nil {
		return c.Status(500).Send([]byte(err.Error()))
	}

	// find apotek by given uid
	// and verify the error
	if err := a.findApotek(uid, &apotek); err != nil {
		return c.Status(400).Send([]byte(err.Error()))
	}

	// if apotek exists then send to client
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "get apotek",
		"data":    apotek,
	})
}

func (a *ApotekHandler) UpdateApotek(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")

	type Apotek struct {
		gorm.Model `json:"-"`
		ID         uint   `json:"id"`
		Logo       string `json:"logo"`
		Name       string `json:"name"`
		Address    string `json:"address"`
	}

	var apotek Apotek

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer", "data": "{}"})
	}
	if err := c.BodyParser(&apotek); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	update := database.DB.Where("id = ?", id).Save(&apotek)

	if update.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": update.Error,
			"data":    "{}",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "apoteker created successfully",
		"data":    apotek,
	})
}

func (a *ApotekHandler) DeleteApotek(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")

	var apotek model.Apotek
	var apoteker model.Apoteker

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer"})
	}

	deleteApoteker := database.DB.Where("apotek_id = ?", id).Unscoped().Delete(&apoteker)
	deleteApotek := database.DB.Where("id = ?", id).Unscoped().Delete(&apotek)

	if deleteApotek.RowsAffected == 0 && deleteApoteker.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "deleteApotek.Error",
			"data":    "{}",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "apotek deleted successfully",
		"data":    "{}",
	})
}

// Local function to find Apotek
func (a *ApotekHandler) findApotek(uid uuid.UUID, apotek *model.Apotek) error {
	var u uuid.NullUUID

	a.db.Model(&model.Apotek{}).Preload("Apotekers").Find(&apotek, "uuid = ?", uid).Scan(&u)
	if !u.Valid {
		return errors.New("apotek doesn't exists")
	}

	return nil
}
