package handlers

import (
	"errors"

	"github.com/ArkjuniorK/apoteker.id_backend/database"
	"github.com/ArkjuniorK/apoteker.id_backend/internal/model"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ApotekHandler struct {
	log *zap.Logger
}

type ApotekSerialize struct {
	ID      uint   `json:"id"`
	Logo    string `json:"logo"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func New(l *zap.Logger) *ApotekHandler {
	return &ApotekHandler{
		log: l,
	}
}

func CreateResponseApotek(apotekModel model.Apotek) ApotekSerialize {
	return ApotekSerialize{ID: apotekModel.ID, Logo: apotekModel.Logo, Name: apotekModel.Name, Address: apotekModel.Address}
}

func (a ApotekHandler) GetApoteks(c *fiber.Ctx) error {
	apoteks := []model.Apotek{}

	database.DB.Find(&apoteks)
	responses := []ApotekSerialize{}

	for _, apotek := range apoteks {
		response := CreateResponseApotek(apotek)
		responses = append(responses, response)
	}
	return c.Status(200).JSON(responses)
}

func (a ApotekHandler) CreateApotek(c *fiber.Ctx) error {
	var apotek model.Apotek

	if err := c.BodyParser(&apotek); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&apotek)
	response := CreateResponseApotek(apotek)

	return c.Status(200).JSON(response)
}

func findApotek(id int, apotek *model.Apotek) error {
	database.DB.Find(&apotek, "id = ?", id)
	if apotek.ID == 0 {
		return errors.New("apotek doesnt exist")
	}
	return nil
}

func (a ApotekHandler) GetApotek(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")

	var apotek model.Apotek

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer"})
	}

	if err := findApotek(id, &apotek); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	response := CreateResponseApotek(apotek)
	return c.Status(200).JSON(response)
}

func (a ApotekHandler) UpdateApotek(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")

	var apotek model.Apotek

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer"})
	}

	if err := findApotek(id, &apotek); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateApotek struct {
		Logo    string `json:"logo"`
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	var updateData UpdateApotek

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	apotek.Logo = updateData.Logo
	apotek.Name = updateData.Name
	apotek.Address = updateData.Address

	database.DB.Save(&apotek)

	response := CreateResponseApotek(apotek)
	return c.Status(200).JSON(response)
}

func (a ApotekHandler) DeleteApotek(c *fiber.Ctx) error {
	id, err := c.ParamsInt("apotekId")

	var apotek model.Apotek

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "message": "Please ensure that id is an integer"})
	}

	if err := findApotek(id, &apotek); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&apotek).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).JSON(fiber.Map{"status": true, "message": "apotek delete successfully"})
}
