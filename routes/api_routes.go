package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// SetupApiRoutes กำหนด Routes ทั้งหมดสำหรับ API (JSON)
func SetupApiRoutes(app *fiber.App, db *sql.DB) {

	// --- Setup for Numbers ---
	numberRepo := repository.NewNumberRepository(db)
	numberService := services.NewNumberService(numberRepo)
	numberHandler := handlers.NewNumberHandler(numberService)

	// --- Setup for Analysis ---
	satNumRepo := repository.NewSatNumRepository(db)
	shaNumRepo := repository.NewShaNumRepository(db)
	kakisDayRepo := repository.NewKakisDayRepository(db)

	analysisService := services.NewAnalysisService(satNumRepo, shaNumRepo, kakisDayRepo)
	analysisHandler := handlers.NewAnalysisHandler(analysisService)

	// --- API Group ---
	// นี่คือส่วนที่ Mobile App จะเรียกใช้
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// API Routes (สำหรับ Numbers)
	v1.Get("/numbers", numberHandler.GetAllNumbers)
	v1.Get("/numbers/:number", numberHandler.GetNumberByPairNumber)

	// API Routes (สำหรับ Analysis)
	v1.Get("/satnums", analysisHandler.GetAllSatNums)
	v1.Post("/satnums/calculate", analysisHandler.CalculateAstrology)
}
