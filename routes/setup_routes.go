package routes

import (
	"database/sql"
	"numberniceic/handlers"
	"numberniceic/repository"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// 👈 [1. แก้ไข] อัปเดตฟังก์ชันนี้
func serveLandingPage(c *fiber.Ctx) error {
	// เปลี่ยนจาก c.SendString() มาเป็น c.Render()
	// "index" คือชื่อไฟล์ index.gohtml (ไม่ต้องใส่นามสกุล)
	// "layouts/main" คือไฟล์ Layout ที่เราต้องการใช้ห่อหุ้ม
	// fiber.Map{...} คือข้อมูลที่เราจะส่งเข้าไปใน Template (เช่น .Title)
	return c.Render("index", fiber.Map{
		"Title": "API Landing Page - NumberNiceIC",
	}, "layouts/main") // 👈 ระบุ Layout ที่จะใช้
}

func SetupRoutes(app *fiber.App, db *sql.DB) {

	// --- ส่วนของ API Routes (เหมือนเดิม) ---
	numberRepo := repository.NewNumberRepository(db)
	numberService := services.NewNumberService(numberRepo)
	numberHandler := handlers.NewNumberHandler(numberService)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/numbers", numberHandler.GetAllNumbers)
	v1.Get("/numbers/:number", numberHandler.GetNumberByPairNumber)

	// --- ส่วนของ Page Route (อัปเดตแล้ว) ---
	// Route "/" (หน้าแรก) จะเรียกใช้ฟังก์ชัน serveLandingPage
	// ที่เราเพิ่งแก้ไขให้ Render Template
	app.Get("/", serveLandingPage)
}
