package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// serveLandingPage function (เหมือนเดิม)
func serveLandingPage(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "API Landing Page - NumberNiceIC",
	}, "layouts/main")
}

func SetupRoutes(app *fiber.App, db *sql.DB) {

	// --- Setup for Numbers (ของเดิม) ---
	numberRepo := repository.NewNumberRepository(db)
	numberService := services.NewNumberService(numberRepo)
	numberHandler := handlers.NewNumberHandler(numberService)

	// --- 🚀 [ใหม่] Setup for SatNums ---
	// 1. สร้าง Repo
	satNumRepo := repository.NewSatNumRepository(db)
	// 2. สร้าง Service โดยฉีด Repo เข้าไป
	satNumService := services.NewSatNumService(satNumRepo)
	// 3. สร้าง Handler โดยฉีด Service เข้าไป
	satNumHandler := handlers.NewSatNumHandler(satNumService)
	// --- จบส่วนใหม่ ---

	// --- API Group (ของเดิม) ---
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Number routes (ของเดิม)
	v1.Get("/numbers", numberHandler.GetAllNumbers)
	v1.Get("/numbers/:number", numberHandler.GetNumberByPairNumber)

	// --- 🚀 [ใหม่] SatNum route ---
	// 4. เพิ่ม Route ใหม่
	v1.Get("/satnums", satNumHandler.GetAllSatNums)
	// --- จบส่วนใหม่ ---

	// --- Page Route (ของเดิม) ---
	app.Get("/", serveLandingPage)
}
