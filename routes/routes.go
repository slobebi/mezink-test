package routes

import (
	"restful_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// CRUD
	app.Get("/api", controllers.Welcome)
	app.Get("/api/records", controllers.GetAllRecords)
	app.Get("/api/records/:id", controllers.GetRecord)
	app.Post("/api/records", controllers.CreateRecord)

	// Filter with totalMarks and createdAt
	app.Get("/api/records-filter", controllers.GetAllRecordsByTotalMarksAndDate)
}
