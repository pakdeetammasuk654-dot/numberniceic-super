package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// --- Handlers สำหรับ Pages (เหมือนเดิม) ---

func serveHomePage(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("home", fiber.Map{
		"Title": "หน้าแรก - NumberNiceIC",
	}, "layouts/main")
}

func serveAnalyzeNamePage(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("analyze_name", fiber.Map{
		"Title": "วิเคราะห์ชื่อ - NumberNiceIC",
	}, "layouts/main")
}

func serveApiDocsPage(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)
	return c.Render("api_docs", fiber.Map{
		"Title": "API Docs - NumberNiceIC",
	}, "layouts/main")
}

// --- จบส่วน Handlers Pages ---

func SetupRoutes(app *fiber.App, db *sql.DB) {

	// --- Setup for Numbers (ของเดิม) ---
	numberRepo := repository.NewNumberRepository(db)
	numberService := services.NewNumberService(numberRepo)
	numberHandler := handlers.NewNumberHandler(numberService)

	// --- 🚀 [เปลี่ยน] Setup for Analysis ---
	satNumRepo := repository.NewSatNumRepository(db)
	shaNumRepo := repository.NewShaNumRepository(db)

	// 👈 [เปลี่ยน] เรียก Service ใหม่
	analysisService := services.NewAnalysisService(satNumRepo, shaNumRepo)

	// 👈 [เปลี่ยน] เรียก Handler ใหม่
	analysisHandler := handlers.NewAnalysisHandler(analysisService)
	// --- จบส่วน Analysis ---

	// --- API Group ---
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// API Routes (สำหรับ Numbers)
	v1.Get("/numbers", numberHandler.GetAllNumbers)
	v1.Get("/numbers/:number", numberHandler.GetNumberByPairNumber)

	// 🚀 [เปลี่ยน] API Routes (สำหรับ Analysis)
	v1.Get("/satnums", analysisHandler.GetAllSatNums)                 // 👈 [เปลี่ยน]
	v1.Post("/satnums/calculate", analysisHandler.CalculateAstrology) // 👈 [เปลี่ยน]

	// --- Page Routes (สำหรับหน้าเว็บ) ---
	app.Get("/", serveHomePage)
	app.Get("/analyze-name", serveAnalyzeNamePage)
	app.Get("/api-docs", serveApiDocsPage)

}
