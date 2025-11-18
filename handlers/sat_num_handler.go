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

// 🚀 [แก้ไข] เพิ่มการบังคับ Header UTF-8
func (h *SatNumHandler) CalculateAstrology(c *fiber.Ctx) error {
	// ... (โค้ด BodyParser และตรวจสอบ Name เหมือนเดิม) ...
	var requestBody models.AstrologyRequest
	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	if requestBody.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name field is required",
		})
	}

	result, err := h.Service.CalculateNameAstrology(requestBody.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to calculate astrology",
		})
	}

	// 🚀 [สำคัญ] บังคับ Header ให้เป็น UTF-8 *ก่อน* ส่ง JSON
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// 5. ส่งผลลัพธ์กลับไป
	return c.JSON(result)
}

// 🚀 [แก้ไข] เพิ่มการบังคับ Header UTF-8
func (h *SatNumHandler) GetAllSatNums(c *fiber.Ctx) error {
	results, err := h.Service.GetAllSatNums()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve sat_nums",
		})
	}

	// 🚀 [สำคัญ] บังคับ Header ให้เป็น UTF-8 *ก่อน* ส่ง JSON
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// ส่งผลลัพธ์กลับไปเป็น JSON
	return c.JSON(results)
}
