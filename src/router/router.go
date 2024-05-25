package router

import (
	"github.com/Abdurozzaq/GoFiber-SimpleMVC/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	/**
	HOME ROUTES
	*/
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"title": "Welcome to Sewa Otomotif!",
		})
	})

	/**
	GLOBAL INFORMATION ROUTES
	*/
	// Create a Global Information
	app.Post("/global-information", controllers.CreateGlobalInformation)

	// Read all Global Information
	// note.Get("/", noteHandler.GetNotes)

	// Read one Global Information
	// note.Get("/:noteId", noteHandler.GetNote)

	// Update one Global Information
	// note.Put("/:noteId", noteHandler.UpdateNote)

	// Delete one Global Information
	// note.Delete("/:noteId", noteHandler.DeleteNote)
}
