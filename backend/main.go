package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	log.Fatal(app.Listen(":8080"))
}
