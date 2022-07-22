package handlers

import (
	"errors"

	"github.com/ArkjuniorK/apoteker.id_backend/database"
	"github.com/ArkjuniorK/apoteker.id_backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ApotekHandler struct {
	log *zap.Logger
	db  *gorm.DB
}

func NewApotekHandler(l *zap.Logger, d *gorm.DB) *ApotekHandler {
	return &ApotekHandler{
		log: l,
		db:  d,
	}
}

func (a ApotekHandler) GetApoteks(c *fiber.Ctx) error {
	var apoteks []model.Apotek

	database.DB.Model(&model.Apotek{}).Preload("Pegawais").Find(&apoteks)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get all apoteks",
		"data":    apoteks,
	})
}

func (a ApotekHandler) CreateApotek(c *fiber.Ctx) error {
	var apotek model.Apotek

	if err := c.BodyParser(&apotek); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	database.DB.Create(&apotek)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Apotek created successfully",
		"data":    apotek,
	})
}

func (a ApotekHandler) GetApotek(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")
	var apotek model.Apotek

	// database.DB.Raw(fmt.Sprintf("SELECT apotekers.id, apotekers.full_name, apotekers.username, apotekers.profile_pic, apoteks.name AS apotek_name FROM apotekers INNER JOIN apoteks ON apotekers.apotek_id = apoteks.id WHERE apotekers.id='%d'", id)).Scan(&result)
	// database.DB.Raw(fmt.Sprintf("SELECT * FROM apoteks WHERE apoteks.id='%d'", id)).Preload("Pegawais").Scan(&apotek)
	database.DB.Where("id = ?", id).Preload("Pegawais").First(&apotek)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "apotek",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get all apoteks",
		"data":    apotek,
	})
}

func (a ApotekHandler) UpdateApotek(c *fiber.Ctx) error {
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

func (a ApotekHandler) DeleteApotek(c *fiber.Ctx) error {
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
func findApotek(id int, apotek *model.Apotek) error {
	database.DB.Model(&model.Apotek{}).Preload("Apotekers").Find(&apotek, "id = ?", id)
	if apotek.ID == 0 {
		return errors.New("apotek doesnt exist")
	}
	return nil
}
