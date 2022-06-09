package controllers

import (
	"fmt"
	"restful_api/database"
	"restful_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

// serializer

// records filter by total marks and date request payload
type RecordsFilterByTotalMarksAndDateRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	MinCount  int64  `json:"minCount"`
	MaxCount  int64  `json:"maxCount"`
}

// records with total marks
type RecordsWithTotalMarks struct {
	ID         uint      `json:"id"`
	TotalMarks int64     `json:"totalMarks"`
	CreatedAt  time.Time `json:"createdAt"`
}

// records response
type RecordsResponse struct {
	Code    uint             `json:"code"`
	Msg     string           `json:"msg"`
	Records []models.Records `json:"records"`
	Error   string           `json:"error"`
}

// records response with total marks
type RecordsWithTotalMarksResponse struct {
	Code    uint                    `json:"code"`
	Msg     string                  `json:"msg"`
	Records []RecordsWithTotalMarks `json:"records"`
	Error   string                  `json:"error"`
}

func Welcome(c *fiber.Ctx) error {
	return c.SendString("It's Work!")
}

func CreateRecord(c *fiber.Ctx) error {
	record := new(models.Records)

	if err := c.BodyParser(record); err != nil {
		response := RecordsResponse{
			Code:  1,
			Msg:   "Error",
			Error: err.Error(),
		}
		return c.Status(400).JSON(response)
	}

	database.Database.Db.Create(record)
	response := RecordsResponse{
		Code:    0,
		Msg:     "Success",
		Records: []models.Records{*record},
	}
	return c.Status(200).JSON(response)
}

func GetAllRecords(c *fiber.Ctx) error {
	records := []models.Records{}
	database.Database.Db.Find(&records)

	response := RecordsResponse{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}

	return c.Status(200).JSON(response)
}

func GetRecord(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		response := RecordsResponse{
			Code:  1,
			Msg:   "Error",
			Error: "Failed to read id from parameter",
		}
		return c.Status(400).JSON(response)
	}

	record := new(models.Records)
	result := database.Database.Db.First(record, id)
	if result.Error != nil {
		response := RecordsResponse{
			Code:  1,
			Msg:   "Error",
			Error: fmt.Sprintf("Failed to retrive data with id %d", id),
		}
		return c.Status(400).JSON(response)
	}

	response := RecordsResponse{
		Code:    0,
		Msg:     "Success",
		Records: []models.Records{*record},
	}
	return c.Status(200).JSON(response)

}

func GetAllRecordsByTotalMarksAndDate(c *fiber.Ctx) error {
	request := new(RecordsFilterByTotalMarksAndDateRequest)

	if err := c.BodyParser(&request); err != nil {
		response := RecordsWithTotalMarksResponse{
			Code:  1,
			Msg:   "Error",
			Error: err.Error(),
		}
		return c.Status(400).JSON(response)
	}

	records := []RecordsWithTotalMarks{}
	result := database.Database.Db.Raw("SELECT * FROM (SELECT id, (SELECT SUM(s) FROM UNNEST(marks) AS s) AS total_marks, created_at FROM records) AS temp WHERE total_marks BETWEEN ? AND ? AND created_at BETWEEN ? AND ?;", request.MinCount, request.MaxCount, request.StartDate, request.EndDate).Scan(&records)

	if result.Error != nil {
		response := RecordsWithTotalMarksResponse{
			Code:  1,
			Msg:   "Error",
			Error: result.Error.Error(),
		}
		return c.Status(400).JSON(response)
	}

	response := RecordsWithTotalMarksResponse{
		Code:    0,
		Msg:     "Success",
		Records: records,
	}
	return c.Status(200).JSON(response)
}
