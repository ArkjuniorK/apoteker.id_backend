package handlers

import (
	"fmt"

	"github.com/ArkjuniorK/apoteker.id_backend/database"
	"github.com/ArkjuniorK/apoteker.id_backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PegawaiHandler struct {
	log *zap.Logger
	db  *gorm.DB
}

type ResultPegawai struct {
	ID         uint   `json:"id"`
	FullName   string `json:"full_name"`
	Username   string `json:"user_name"`
	ProfilePic string `json:"profile_picture"`
	ApotekName string `json:"apotek_name"`
}

func NewPegawaiHandler(l *zap.Logger, d *gorm.DB) *PegawaiHandler {
	return &PegawaiHandler{log: l, db: d}
}

func (a PegawaiHandler) GetPegawais(c *fiber.Ctx) error {
	var result []ResultPegawai
	database.DB.Raw("SELECT pegawais.id, pegawais.full_name, pegawais.username, pegawais.profile_pic, apoteks.name AS apotek_name FROM pegawais INNER JOIN apoteks ON pegawais.apotek_id = apoteks.id").Scan(&result)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get all pegawais",
		"data":    result,
	})
}

func (a PegawaiHandler) GetPegawai(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")
	var result Result

	database.DB.Raw(fmt.Sprintf("SELECT pegawais.id, pegawais.full_name, pegawais.username, pegawais.profile_pic, apoteks.name AS apotek_name FROM pegawais INNER JOIN apoteks ON pegawais.apotek_id = apoteks.id WHERE pegawais.id='%d'", id)).Scan(&result)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "pegawai",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get pegawai",
		"data":    result,
	})
}

func (a PegawaiHandler) CreatePegawai(c *fiber.Ctx) error {

	var pegawai model.Pegawai

	if err := c.BodyParser(&pegawai); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	create := database.DB.Create(&pegawai)

	if create.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": create.Error,
			"data":    "{}",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "apoteker created successfully",
		"data":    pegawai,
	})
}

func (a PegawaiHandler) UpdatePegawai(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")
	var pegawai model.Pegawai

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer", "data": "{}"})
	}
	if err := c.BodyParser(&pegawai); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	update := database.DB.Where("id = ?", id).Save(&pegawai)

	if update.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": update.Error,
			"data":    "{}",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "pegawai created successfully",
		"data":    pegawai,
	})
}

func (a PegawaiHandler) DeletePegawai(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")
	var pegawai model.Pegawai

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer", "data": "{}"})
	}

	delete := database.DB.Where("id = ?", id).Unscoped().Delete(&pegawai)

	if delete.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": delete.Error,
			"data":    "{}",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "pegawai deleted successfully",
		"data":    "{}",
	})
}

// Local function to find Apotek
// func findApoteker(id int, apotek *model.Apoteker) error {
// 	database.DB.Raw(&model.Apoteker{}).Preload("Apotekers").Find(&apotek, "id = ?", id)
// 	if apotek.ID == 0 {
// 		return errors.New("apotek doesnt exist")
// 	}
// 	return nil
// }
