package handlers

import (
	"numberniceic/models"
	"numberniceic/services"

	"github.com/gofiber/fiber/v2"
)

// SatNumHandler handles the HTTP requests for SatNum.
type SatNumHandler struct {
	Service services.SatNumService
}

// NewSatNumHandler creates a new SatNumHandler.
func NewSatNumHandler(service services.SatNumService) *SatNumHandler {
	return &SatNumHandler{Service: service}
}

func (h *SatNumHandler) CalculateAstrology(c *fiber.Ctx) error {
	// 1. สร้างตัวแปรรับ JSON body
	var requestBody models.AstrologyRequest

	// 2. Parse JSON body เข้า struct
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// 3. ตรวจสอบว่า "name" ไม่ได้ว่าง
	if requestBody.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name field is required",
		})
	}

	// 4. เรียก Service เพื่อคำนวณ
	result, err := h.Service.CalculateNameAstrology(requestBody.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to calculate astrology",
		})
	}

	// 5. ส่งผลลัพธ์กลับไป
	return c.JSON(result)
}

// GetAllSatNums is the handler for GET /api/v1/satnums
func (h *SatNumHandler) GetAllSatNums(c *fiber.Ctx) error {
	results, err := h.Service.GetAllSatNums()
	if err != nil {
		// ถ้ามี error, ส่ง 500 Internal Server Error กลับไป
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve sat_nums",
		})
	}

	// ส่งผลลัพธ์กลับไปเป็น JSON
	return c.JSON(results)
}
