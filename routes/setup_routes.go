package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// --- Handlers สำหรับ Pages ---

// 1. Handler สำหรับหน้าแรก
func serveHomePage(c *fiber.Ctx) error {
	// 🚀 [เพิ่ม] บังคับ Header ให้เป็น UTF-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	return c.Render("home", fiber.Map{
		"Title": "หน้าแรก - NumberNiceIC",
	}, "layouts/main")
}

// 2. Handler สำหรับหน้าวิเคราะห์ชื่อ
func serveAnalyzeNamePage(c *fiber.Ctx) error {
	// 🚀 [เพิ่ม] บังคับ Header ให้เป็น UTF-8
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTMLCharsetUTF8)

	return c.Render("analyze_name", fiber.Map{
		"Title": "วิเคราะห์ชื่อ - NumberNiceIC",
	}, "layouts/main")
}

// 3. Handler สำหรับหน้า API Docs
func serveApiDocsPage(c *fiber.Ctx) error {
	// 🚀 [เพิ่ม] บังคับ Header ให้เป็น UTF-8
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

	// --- Setup for Analysis (เลขศาสตร์ + พลังเงา) (ของเดิม) ---
	satNumRepo := repository.NewSatNumRepository(db)
	shaNumRepo := repository.NewShaNumRepository(db)
	satNumService := services.NewSatNumService(satNumRepo, shaNumRepo)
	satNumHandler := handlers.NewSatNumHandler(satNumService)

	// --- API Group (ของเดิม) ---
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// API Routes (ของเดิม)
	v1.Get("/numbers", numberHandler.GetAllNumbers)
	v1.Get("/numbers/:number", numberHandler.GetNumberByPairNumber)
	v1.Get("/satnums", satNumHandler.GetAllSatNums)
	v1.Post("/satnums/calculate", satNumHandler.CalculateAstrology)

	// --- Page Routes (ของเดิม) ---
	app.Get("/", serveHomePage)
	app.Get("/analyze-name", serveAnalyzeNamePage)
	app.Get("/api-docs", serveApiDocsPage)
}
