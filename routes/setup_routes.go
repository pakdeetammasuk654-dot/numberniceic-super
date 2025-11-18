package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// --- 🚀 [ใหม่] Handlers สำหรับ Pages ---

// 1. Handler สำหรับหน้าแรก
func serveHomePage(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"Title": "หน้าแรก - NumberNiceIC",
	}, "layouts/main")
}

// 2. Handler สำหรับหน้าวิเคราะห์ชื่อ
func serveAnalyzeNamePage(c *fiber.Ctx) error {
	return c.Render("analyze_name", fiber.Map{
		"Title": "วิเคราะห์ชื่อ - NumberNiceIC",
	}, "layouts/main")
}

// 3. Handler สำหรับหน้า API Docs (นี่คือฟังก์ชันเดิมของคุณ)
func serveApiDocsPage(c *fiber.Ctx) error {
	// "api_docs" คือชื่อไฟล์ .gohtml ใหม่ที่เราเพิ่งเปลี่ยน
	return c.Render("api_docs", fiber.Map{
		"Title": "API Docs - NumberNiceIC",
	}, "layouts/main")
}

// --- จบส่วน Handlers ใหม่ ---

func SetupRoutes(app *fiber.App, db *sql.DB) {

	// --- Setup for Numbers (ของเดิม) ---
	numberRepo := repository.NewNumberRepository(db)
	numberService := services.NewNumberService(numberRepo)
	numberHandler := handlers.NewNumberHandler(numberService)

	// --- Setup for SatNums (ของเดิม) ---
	satNumRepo := repository.NewSatNumRepository(db)
	satNumService := services.NewSatNumService(satNumRepo)
	satNumHandler := handlers.NewSatNumHandler(satNumService)

	// --- API Group (ของเดิม) ---
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// API Routes (ของเดิม)
	v1.Get("/numbers", numberHandler.GetAllNumbers)
	v1.Get("/numbers/:number", numberHandler.GetNumberByPairNumber)
	v1.Get("/satnums", satNumHandler.GetAllSatNums)
	v1.Post("/satnums/calculate", satNumHandler.CalculateAstrology)

	// --- 🚀 [ใหม่] Page Routes (อัปเดตส่วนนี้) ---

	// "/" (หน้าแรก) จะไปที่ serveHomePage
	app.Get("/", serveHomePage)

	// "/analyze-name" จะไปที่ serveAnalyzeNamePage
	app.Get("/analyze-name", serveAnalyzeNamePage)

	// "/api-docs" จะไปที่ serveApiDocsPage (หน้า API เดิม)
	app.Get("/api-docs", serveApiDocsPage)

}
