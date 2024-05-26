package main

import (
	"github.com/Abdurozzaq/GoFiber-SimpleMVC/src/database"
	"github.com/Abdurozzaq/GoFiber-SimpleMVC/src/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize Template Engine
	engine := html.New("./src/views", ".html")

	// Initialize Fiber
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Initialize default config
	app.Use(logger.New())

	// Initialize public folder for static files
	app.Static("/", "./src/public")

	// Initialize Database Connection
	database.ConnectDB()

	// Initialize Router
	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
