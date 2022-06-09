package main

import (
	"log"
	"restful_api/database"
	"restful_api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
