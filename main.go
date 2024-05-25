package main

import (
	"github.com/Abdurozzaq/fiber-sewa-otomotif/src/database"
	"github.com/Abdurozzaq/fiber-sewa-otomotif/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize Template Engine
	engine := html.New("./src/views", ".html")

	// Initialize Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Initialize Database Connection
	database.ConnectDB()

	// Initialize Router
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
