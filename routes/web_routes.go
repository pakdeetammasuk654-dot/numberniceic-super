package routes

import (
	"github.com/gofiber/fiber/v2"
)

// --- Handlers สำหรับ Pages ---

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

// SetupWebRoutes กำหนด Routes ทั้งหมดสำหรับหน้าเว็บ (HTML)
func SetupWebRoutes(app *fiber.App) {
	// --- Page Routes (สำหรับหน้าเว็บ) ---
	app.Get("/", serveHomePage)
	app.Get("/analyze-name", serveAnalyzeNamePage)
	app.Get("/api-docs", serveApiDocsPage)
}
