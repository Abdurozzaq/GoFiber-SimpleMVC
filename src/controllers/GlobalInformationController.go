package controllers

import (
	"github.com/Abdurozzaq/fiber-sewa-otomotif/src/database"
	"github.com/Abdurozzaq/fiber-sewa-otomotif/src/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateGlobalInformation(c *fiber.Ctx) error {
	db := database.DB
	globalInformation := new(models.GlobalInformation)

	// Store the body in the globalInformation and return error if encountered
	err := c.BodyParser(globalInformation)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the globalInformation
	globalInformation.ID = uuid.New()

	// Create the Note and return error if encountered
	err = db.Create(&globalInformation).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create globalInformation", "data": err})
	}

	// Return the created globalInformation
	return c.JSON(fiber.Map{"status": "success", "message": "Created GlobalInformation", "data": globalInformation})
}
