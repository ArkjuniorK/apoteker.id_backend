package handlers

import (
	"fmt"

	"github.com/ArkjuniorK/apoteker.id_backend/database"
	"github.com/ArkjuniorK/apoteker.id_backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// type ApotekerHandler struct {
// 	log *zap.Logger
// }
type Result struct {
	ID         uint   `json:"id"`
	FullName   string `json:"full_name"`
	Username   string `json:"user_name"`
	ProfilePic string `json:"profile_picture"`
	ApotekName string `json:"apotek_name"`
}

func GetApotekers(c *fiber.Ctx) error {
	var result []Result
	database.DB.Raw("SELECT apotekers.id, apotekers.full_name, apotekers.username, apotekers.profile_pic, apoteks.name AS apotek_name FROM apotekers INNER JOIN apoteks ON apotekers.apotek_id = apoteks.id").Scan(&result)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get all apoteks",
		"data":    result,
	})
}

func GetApoteker(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")
	var result Result

	database.DB.Raw(fmt.Sprintf("SELECT apotekers.id, apotekers.full_name, apotekers.username, apotekers.profile_pic, apoteks.name AS apotek_name FROM apotekers INNER JOIN apoteks ON apotekers.apotek_id = apoteks.id WHERE apotekers.id='%d'", id)).Scan(&result)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "apoteker",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "get all apoteks",
		"data":    result,
	})
}

func CreateApoteker(c *fiber.Ctx) error {

	var apoteker model.Apoteker

	if err := c.BodyParser(&apoteker); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	create := database.DB.Create(&apoteker)

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
		"data":    apoteker,
	})
}

func UpdateApoteker(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")
	var apoteker model.Apoteker

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer", "data": "{}"})
	}
	if err := c.BodyParser(&apoteker); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	update := database.DB.Where("id = ?", id).Save(&apoteker)

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
		"data":    apoteker,
	})
}

func DeleteApoteker(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")
	var apoteker model.Apoteker

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer", "data": "{}"})
	}

	delete := database.DB.Where("id = ?", id).Unscoped().Delete(&apoteker)

	if delete.RowsAffected == 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": delete.Error,
			"data":    "{}",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "apoteker deleted successfully",
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
