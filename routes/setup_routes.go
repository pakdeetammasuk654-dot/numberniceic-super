package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes คือฟังก์ชันหลักที่จะถูกเรียกโดย main.go
// ทำหน้าที่ตั้งค่า Routes ทั้งหมดในระบบ
func SetupRoutes(app *fiber.App, db *sql.DB) {

	// 1. ตั้งค่า Routes สำหรับหน้าเว็บ (HTML)
	SetupWebRoutes(app)

	// 2. ตั้งค่า Routes สำหรับ API (JSON)
	// (นี่คือส่วนที่ Mobile App จะเรียกใช้)
	SetupApiRoutes(app, db)

}
